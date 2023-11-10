// Copyright 2020 Layer5, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cilium

import (
	"context"
	"fmt"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/common"
	"github.com/layer5io/meshery-adapter-library/meshes"
	"github.com/layer5io/meshery-adapter-library/status"
	"github.com/layer5io/meshery-cilium/cilium/oam"
	internalconfig "github.com/layer5io/meshery-cilium/internal/config"
	meshkitCfg "github.com/layer5io/meshkit/config"
	"github.com/layer5io/meshkit/errors"
	"github.com/layer5io/meshkit/logger"
	"github.com/layer5io/meshkit/models"
	"github.com/layer5io/meshkit/utils"
	"github.com/layer5io/meshkit/models/oam/core/v1alpha1"
	"github.com/layer5io/meshkit/utils/events"
	"gopkg.in/yaml.v2"
)

// Cilium  represents the Cilium adapter and embeds adapter.Adapter
type Cilium struct {
	adapter.Adapter
}

// New initializes a new handler instance
func New(c meshkitCfg.Handler, l logger.Handler, kc meshkitCfg.Handler, ev *events.EventStreamer) adapter.Handler {
	return &Cilium{
		Adapter: adapter.Adapter{
			Config:            c,
			Log:               l,
			KubeconfigHandler: kc,
			EventStreamer:     ev,
		},
	}
}

// CreateKubeconfigs creates and writes passed kubeconfig onto the filesystem
func (h *Cilium) CreateKubeconfigs(kubeconfigs []string) error {
	var errs = make([]error, 0)
	for _, kubeconfig := range kubeconfigs {
		kconfig := models.Kubeconfig{}
		err := yaml.Unmarshal([]byte(kubeconfig), &kconfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		// To have control over what exactly to take in on kubeconfig
		h.KubeconfigHandler.SetKey("kind", kconfig.Kind)
		h.KubeconfigHandler.SetKey("apiVersion", kconfig.APIVersion)
		h.KubeconfigHandler.SetKey("current-context", kconfig.CurrentContext)
		err = h.KubeconfigHandler.SetObject("preferences", kconfig.Preferences)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = h.KubeconfigHandler.SetObject("clusters", kconfig.Clusters)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = h.KubeconfigHandler.SetObject("users", kconfig.Users)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = h.KubeconfigHandler.SetObject("contexts", kconfig.Contexts)
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return mergeErrors(errs)
}

// ApplyOperation function contains the operation handlers
func (h *Cilium) ApplyOperation(ctx context.Context, opReq adapter.OperationRequest) error {
	err := h.CreateKubeconfigs(opReq.K8sConfigs)
	if err != nil {
		return err
	}
	kubeconfigs := opReq.K8sConfigs
	operations := make(adapter.Operations)
	requestedVersion := adapter.Version(opReq.Version)
	err = h.Config.GetObject(adapter.OperationsKey, &operations)
	if err != nil {
		return err
	}

	e := &meshes.EventsResponse{
		OperationId:   opReq.OperationID,
		Summary:       status.Deploying,
		Details:       "Operation is not supported",
		Component:     internalconfig.ServerDefaults["type"],
		ComponentName: internalconfig.ServerDefaults["name"],
	}

	//deployment
	switch opReq.OperationName {
	case internalconfig.CiliumOperation:
		go func(hh *Cilium, ee *meshes.EventsResponse) {
			var err error
			var stat, version string
			fmt.Println("dd", operations[opReq.OperationName].Versions)
			if len(operations[opReq.OperationName].Versions) == 0 {
				err = ErrFetchIstioVersions
			} else {
				version = string(operations[opReq.OperationName].Versions[len(operations[opReq.OperationName].Versions)-1])
				if utils.Contains[[]adapter.Version, adapter.Version](operations[opReq.OperationName].Versions, requestedVersion) {
					version = requestedVersion.String()
				}
				stat, err = hh.installCilium(opReq.IsDeleteOperation, version, opReq.Namespace, kubeconfigs)
			}
			if err != nil {
				ee.Summary = fmt.Sprintf("Error while %s Cilium service mesh with version %s", stat, version)
				ee.Details = err.Error()
				ee.ErrorCode = errors.GetCode(err)
				ee.ProbableCause = errors.GetCause(err)
				ee.SuggestedRemediation = errors.GetRemedy(err)
				hh.StreamErr(ee, err)
				return
			}
			ee.Summary = fmt.Sprintf("Cilium service mesh %s successfully", stat)
			ee.Details = fmt.Sprintf("Cilium service mesh is now %s.", stat)
			hh.StreamInfo(ee)
		}(h, e)
	case
		common.BookInfoOperation,
		common.HTTPBinOperation,
		common.ImageHubOperation,
		common.EmojiVotoOperation:
		go func(hh *Cilium, ee *meshes.EventsResponse) {
			appName := operations[opReq.OperationName].AdditionalProperties[common.ServiceName]
			stat, err := hh.installSampleApp(opReq.IsDeleteOperation, opReq.Namespace, operations[opReq.OperationName].Templates, kubeconfigs)
			if err != nil {
				summary := fmt.Sprintf("Error while %s %s application", stat, appName)
				hh.streamErr(summary, e, err)
				return
			}
			ee.Summary = fmt.Sprintf("%s application %s successfully", appName, stat)
			ee.Details = fmt.Sprintf("The %s application is now %s.", appName, stat)
			hh.StreamInfo(e)
		}(h, e)
	case common.SmiConformanceOperation:
		go func(hh *Cilium, ee *meshes.EventsResponse) {
			name := operations[opReq.OperationName].Description
			_, err := hh.RunSMITest(adapter.SMITestOptions{
				Ctx:         context.TODO(),
				OperationID: ee.OperationId,
				Manifest:    string(operations[opReq.OperationName].Templates[0]),
				Namespace:   "meshery",
				Labels: map[string]string{
					"cilium.io/monitored-by": "cilium",
				},
				Kubeconfigs: kubeconfigs,
				Annotations: make(map[string]string),
			})
			if err != nil {
				summary := fmt.Sprintf("Error while %s %s test", status.Running, name)
				hh.streamErr(summary, e, err)
				return
			}
			ee.Summary = fmt.Sprintf("%s test %s successfully", name, status.Completed)
			ee.Details = ""
			hh.StreamInfo(e)
		}(h, e)
	default:
		h.streamErr("Invalid Operation", e, ErrOpInvalid)
	}
	return nil
}

// ProcessOAM will handles the grpc invocation for handling OAM objects
func (h *Cilium) ProcessOAM(ctx context.Context, oamReq adapter.OAMRequest) (string, error) {
	err := h.CreateKubeconfigs(oamReq.K8sConfigs)
	if err != nil {
		return "", err
	}
	kubeconfigs := oamReq.K8sConfigs
	var comps []v1alpha1.Component
	for _, acomp := range oamReq.OamComps {
		comp, err := oam.ParseApplicationComponent(acomp)
		if err != nil {
			h.Log.Error(ErrParseOAMComponent)
			continue
		}

		comps = append(comps, comp)
	}

	config, err := oam.ParseApplicationConfiguration(oamReq.OamConfig)
	if err != nil {
		h.Log.Error(ErrParseOAMConfig)
	}

	// If operation is delete then first HandleConfiguration and then handle the deployment
	if oamReq.DeleteOp {
		// Process configuration
		msg2, err := h.HandleApplicationConfiguration(config, oamReq.DeleteOp, kubeconfigs)
		if err != nil {
			return msg2, ErrProcessOAM(err)
		}

		// Process components
		msg1, err := h.HandleComponents(comps, oamReq.DeleteOp, kubeconfigs)
		if err != nil {
			return msg1 + "\n" + msg2, ErrProcessOAM(err)
		}

		return msg1 + "\n" + msg2, nil
	}

	// Process components
	msg1, err := h.HandleComponents(comps, oamReq.DeleteOp, kubeconfigs)
	if err != nil {
		return msg1, ErrProcessOAM(err)
	}

	// Process configuration
	msg2, err := h.HandleApplicationConfiguration(config, oamReq.DeleteOp, kubeconfigs)
	if err != nil {
		return msg1 + "\n" + msg2, ErrProcessOAM(err)
	}

	return msg1 + "\n" + msg2, nil
}

func (h *Cilium) streamErr(summary string, e *meshes.EventsResponse, err error) {
	e.Summary = summary
	e.Details = err.Error()
	e.ErrorCode = errors.GetCode(err)
	e.ProbableCause = errors.GetCause(err)
	e.SuggestedRemediation = errors.GetRemedy(err)
	h.StreamErr(e, err)
}
