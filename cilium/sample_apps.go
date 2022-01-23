package cilium

import (
	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/status"
	mesherykube "github.com/layer5io/meshkit/utils/kubernetes"
)

func (h *Handler) installSampleApp(del bool, namespace string, templates []adapter.Template) (string, error) {
	st := status.Installing
	if del {
		st = status.Removing
	}
	for _, template := range templates {
		err := h.applyManifest([]byte(template.String()), del, namespace)
		if err != nil {
			return st, ErrSampleApp(err)
		}
	}
	return status.Installed, nil
}

func (h *Handler) applyManifest(contents []byte, isDel bool, namespace string) error {
	kclient := h.MesheryKubeclient
	if kclient == nil {
		return ErrNilClient
	}

	err := kclient.ApplyManifest(contents, mesherykube.ApplyOptions{
		Namespace: namespace,
		Update:    true,
		Delete:    isDel,
	})

	if err != nil {
		return err
	}

	return nil
}
