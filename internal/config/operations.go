package config

import (
	"strings"

	smp "github.com/layer5io/service-mesh-performance/spec"
)

var (
	CiliumOperation = strings.ToLower(smp.ServiceMesh_CILIUM_SERVICE_MESH.Enum().String())
	ServiceName     = "service_name"
)
