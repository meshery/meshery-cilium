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
	"github.com/layer5io/meshery-adapter-library/status"
	"github.com/layer5io/meshery-cilium/cilium/oam"
	internalconfig "github.com/layer5io/meshery-cilium/internal/config"
	meshkitCfg "github.com/layer5io/meshkit/config"
	"github.com/layer5io/meshkit/logger"
	"github.com/layer5io/meshkit/models/oam/core/v1alpha1"
)

// Handler instance for this adapter
type Handler struct {
	adapter.Adapter
}

// New initializes a new handler instance
func New(config meshkitCfg.Handler, log logger.Handler, kc meshkitCfg.Handler) adapter.Handler {
	return &Handler{
		Adapter: adapter.Adapter{
			Config:            config,
			Log:               log,
			KubeconfigHandler: kc,
		},
	}
}

// ApplyOperation function contains the operation handlers
func (h *Handler) ApplyOperation(ctx context.Context, request adapter.OperationRequest) error {
	operations := make(adapter.Operations)
	err := h.Config.GetObject(adapter.OperationsKey, &operations)
	if err != nil {
		return err
	}

	e := &adapter.Event{
		Operationid: request.OperationID,
		Summary:     status.Deploying,
		Details:     "Operation is not supported",
	}

	//deployment
	switch request.OperationName {
	case internalconfig.CiliumOperation:
		go func(hh *Handler, ee *adapter.Event) {
			version := string(operations[request.OperationName].Versions[len(operations[request.OperationName].Versions)-1])
			fmt.Println("version: ", version)
			stat, err := hh.installCilium(request.IsDeleteOperation, version, request.Namespace)
			if err != nil {
				e.Summary = fmt.Sprintf("Error while %s Cilium service mesh", stat)
				e.Details = err.Error()
				hh.StreamErr(e, err)
				return
			}
			ee.Summary = fmt.Sprintf("Cilium service mesh %s successfully", stat)
			ee.Details = fmt.Sprintf("Cilium service mesh is now %s.", stat)
			hh.StreamInfo(e)
		}(h, e)
	case
		common.BookInfoOperation,
		common.HTTPBinOperation,
		common.ImageHubOperation,
		common.EmojiVotoOperation:
		go func(hh *Handler, ee *adapter.Event) {
			appName := operations[request.OperationName].AdditionalProperties[common.ServiceName]
			stat, err := hh.installSampleApp(request.IsDeleteOperation, request.Namespace, operations[request.OperationName].Templates)
			if err != nil {
				e.Summary = fmt.Sprintf("Error while %s %s application", stat, appName)
				e.Details = err.Error()
				hh.StreamErr(e, err)
				return
			}
			ee.Summary = fmt.Sprintf("%s application %s successfully", appName, stat)
			ee.Details = fmt.Sprintf("The %s application is now %s.", appName, stat)
			hh.StreamInfo(e)
		}(h, e)
	case common.SmiConformanceOperation:
		go func(hh *Handler, ee *adapter.Event) {
			name := operations[request.OperationName].Description
			_, err := hh.RunSMITest(adapter.SMITestOptions{
				Ctx:         context.TODO(),
				OperationID: ee.Operationid,
				Manifest:    string(operations[request.OperationName].Templates[0]),
				Namespace:   "meshery",
				Labels: map[string]string{
					"cilium.io/monitored-by": "cilium",
				},
				Annotations: make(map[string]string),
			})
			if err != nil {
				e.Summary = fmt.Sprintf("Error while %s %s test", status.Running, name)
				e.Details = err.Error()
				hh.StreamErr(e, err)
				return
			}
			ee.Summary = fmt.Sprintf("%s test %s successfully", name, status.Completed)
			ee.Details = ""
			hh.StreamInfo(e)
		}(h, e)
	default:
		h.StreamErr(e, ErrOpInvalid)
	}
	return nil
}

// ProcessOAM will handles the grpc invocation for handling OAM objects
func (h *Handler) ProcessOAM(ctx context.Context, oamReq adapter.OAMRequest) (string, error) {
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
		msg2, err := h.HandleApplicationConfiguration(config, oamReq.DeleteOp)
		if err != nil {
			return msg2, ErrProcessOAM(err)
		}

		// Process components
		msg1, err := h.HandleComponents(comps, oamReq.DeleteOp)
		if err != nil {
			return msg1 + "\n" + msg2, ErrProcessOAM(err)
		}

		return msg1 + "\n" + msg2, nil
	}

	// Process components
	msg1, err := h.HandleComponents(comps, oamReq.DeleteOp)
	if err != nil {
		return msg1, ErrProcessOAM(err)
	}

	// Process configuration
	msg2, err := h.HandleApplicationConfiguration(config, oamReq.DeleteOp)
	if err != nil {
		return msg1 + "\n" + msg2, ErrProcessOAM(err)
	}

	return msg1 + "\n" + msg2, nil
}
