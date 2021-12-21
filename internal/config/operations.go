package config

import (
	"strings"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/meshes"
	smp "github.com/layer5io/service-mesh-performance/spec"
)

var (
	CiliumOperation = strings.ToLower(smp.ServiceMesh_CILIUM_SERVICE_MESH.Enum().String())
	ServiceName     = "service_name"
)

func getOperations(op adapter.Operations) adapter.Operations {
	versions, _ := getLatestReleaseNames(3)

	op[CiliumOperation] = &adapter.Operation{
		Type:        int32(meshes.OpCategory_INSTALL),
		Description: "Cilium Service Mesh",
		Versions:    versions,
		Templates:   []adapter.Template{},
	}

	return op
}
