package cilium

import (
	"fmt"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/status"
	mesherykube "github.com/layer5io/meshkit/utils/kubernetes"
)

// noneNamespace indicates unset namespace
const noneNamespace = ""

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

// createNS handles the creatin as well as deletion of namespaces
func createNS(h *Handler, ns string, del bool) error {
	manifest := fmt.Sprintf(`
apiVersion: v1
kind: Namespace
metadata:
  name: %s
`,
		ns,
	)

	if err := h.applyManifest([]byte(manifest), del, noneNamespace); err != nil {
		return err
	}

	return nil
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
