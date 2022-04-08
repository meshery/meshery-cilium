// Package cilium - Error codes for the adapter
package cilium

import (
	"github.com/layer5io/meshkit/errors"
)

var (
	// Error code for failed service mesh installation

	// ErrInstallCiliumCode represents the errors which are generated
	// during Cilium service mesh install process
	ErrInstallCiliumCode = "1000"

	// ErrTarXZFCode represents the errors which are generated
	// during decompressing and extracting tar.gz file
	ErrTarXZFCode = "1001"

	// ErrMeshConfigCode represents the errors which are generated
	// when an invalid mesh config is found
	ErrMeshConfigCode = "1002"

	// ErrRunCiliumCmdCode represents the errors which are generated
	// during fetch manifest process
	ErrRunCiliumCmdCode = "1003"

	// ErrDownloadBinaryCode represents the errors which are generated
	// during binary download process
	ErrDownloadBinaryCode = "1004"

	// ErrInstallBinaryCode represents the errors which are generated
	// during binary installation process
	ErrInstallBinaryCode = "1005"

	// ErrSampleAppCode represents the errors which are generated
	// duing sample app installation
	ErrSampleAppCode = "1006"

	// ErrCustomOperationCode represents the errors which are generated
	// when an invalid addon operation is requested
	ErrCustomOperationCode = "1007"

	// ErrCreatingNSCode represents the errors which are generated
	// during the process of creating a namespace
	ErrCreatingNSCode = "1008"

	// ErrRunExecutableCode represents the errors which are generated
	// during the running a executable
	ErrRunExecutableCode = "1009"

	// ErrSidecarInjectionCode represents the errors which are generated
	// during the process of enabling/disabling sidecar injection
	ErrSidecarInjectionCode = "1010"

	// ErrOpInvalidCode represents the error which is generated when
	// there is an invalid operation
	ErrOpInvalidCode = "1011"

	// ErrApplyHelmChartCode represents the error which are generated
	// during the process of applying helm chart
	ErrApplyHelmChartCode = "1012"

	// ErrNilClientCode represents the error code which is
	// generated when kubernetes client is nil
	ErrNilClientCode = "1013"

	// ErrInvalidOAMComponentTypeCode represents the error code which is
	// generated when an invalid oam component is requested
	ErrInvalidOAMComponentTypeCode = "1014"

	// ErrCiliumCoreComponentFailCode represents the error code which is
	// generated when an Cilium core operations fails
	ErrCiliumCoreComponentFailCode = "1015"
	// ErrProcessOAMCode represents the error code which is
	// generated when an OAM operations fails
	ErrProcessOAMCode = "1016"
	// ErrParseCiliumCoreComponentCode represents the error code which is
	// generated when Cilium core component manifest parsing fails
	ErrParseCiliumCoreComponentCode = "1017"
	// ErrParseOAMComponentCode represents the error code which is
	// generated during the OAM component parsing
	ErrParseOAMComponentCode = "1018"
	// ErrParseOAMConfigCode represents the error code which is
	// generated during the OAM configuration parsing
	ErrParseOAMConfigCode = "1019"

	// ErrOpInvalid represents the errors which are generated
	// when an invalid operation is requested
	ErrOpInvalid = errors.New(ErrOpInvalidCode, errors.Alert, []string{"Invalid operation"}, []string{"Cilium adapter recived an invalid operation from the meshey server"}, []string{"The operation is not supported by the adapter", "Invalid operation name"}, []string{"Check if the operation name is valid and supported by the adapter"})

	// ErrNilClient represents the error which is
	// generated when kubernetes client is nil
	ErrNilClient = errors.New(ErrNilClientCode, errors.Alert, []string{"kubernetes client not initialized"}, []string{"Kubernetes client is nil"}, []string{"kubernetes client not initialized"}, []string{"Reconnect the adaptor to Meshery server"})

	// ErrParseOAMComponent represents the error which is
	// generated during the OAM component parsing
	ErrParseOAMComponent = errors.New(ErrParseOAMComponentCode, errors.Alert, []string{"error parsing the component"}, []string{"Error occured while prasing application component in the OAM request made"}, []string{"Invalid OAM component passed in OAM request"}, []string{"Check if your request has vaild OAM components"})

	// ErrParseOAMConfig represents the error which is
	// generated during the OAM configuration parsing
	ErrParseOAMConfig = errors.New(ErrParseOAMConfigCode, errors.Alert, []string{"error parsing the configuration"}, []string{"Error occured while prasing component config in the OAM request made"}, []string{"Invalid OAM config passed in OAM request"}, []string{"Check if your request has vaild OAM config"})

	// ErrGetLatestReleaseCode represents the error which is
	// generated when the latest stable version could not
	// be fetched during runtime component registeration
	ErrGetLatestReleaseCode = "1020"

	// ErrMakingBinExecutableCode implies error while makng cilium cli executable
	ErrMakingBinExecutableCode = "1025"

	//ErrLoadNamespaceCode occur during the process of applying namespace
	ErrLoadNamespaceCode = "1026"

	// ErrUnpackingTarCode implies error while unpacking cilium release
	// bundle tar
	ErrUnpackingTarCode = "1027"

	// ErrUnzipFileCode represents the errors which are generated
	// during unzip process
	ErrUnzipFileCode = "1028"

	// ErrDownloadingTarCode implies error while downloading cilium tar
	ErrDownloadingTarCode = "1029"

	// ErrGettingReleaseCode implies error while fetching latest release for cilium cli
	ErrGettingReleaseCode = "1030"
)

// ErrInstallCilium is the error for install mesh
func ErrInstallCilium(err error) error {
	return errors.New(ErrInstallCiliumCode, errors.Alert, []string{"Error with Cilium operation"}, []string{"Error occured while installing Cilium mesh through Cilium", err.Error()}, []string{}, []string{})
}

// ErrTarXZF is the error for unzipping the file
func ErrTarXZF(err error) error {
	return errors.New(ErrTarXZFCode, errors.Alert, []string{"Error while extracting file"}, []string{err.Error()}, []string{"The gzip might be corrupt"}, []string{"Retry the operation"})
}

// ErrMeshConfig is the error for mesh config
func ErrMeshConfig(err error) error {
	return errors.New(ErrMeshConfigCode, errors.Alert, []string{"Error configuration mesh"}, []string{err.Error(), "Error getting MeshSpecKey config from in-memory configuration"}, []string{}, []string{"Reconnect the adaptor to the meshkit server"})
}

// ErrRunCiliumCmd is the error for mesh port forward
func ErrRunCiliumCmd(err error, des string) error {
	return errors.New(ErrRunCiliumCmdCode, errors.Alert, []string{"Error running cilium command"}, []string{err.Error()}, []string{"Corrupted cilium binary", "Command might be invalid"}, []string{})
}

// ErrDownloadBinary is the error while downloading Cilium binary
func ErrDownloadBinary(err error) error {
	return errors.New(ErrDownloadBinaryCode, errors.Alert, []string{"Error downloading Cilium binary"}, []string{err.Error(), "Error occured while download Cilium binary from its github release"}, []string{"Checkout https://docs.github.com/en/rest/reference/repos#releases for more details"}, []string{})
}

// ErrInstallBinary is the error while downloading Cilium binary
func ErrInstallBinary(err error) error {
	return errors.New(ErrInstallBinaryCode, errors.Alert, []string{"Error installing Cilium binary"}, []string{err.Error()}, []string{"Corrupted Cilium release binary", "Invalid installation location"}, []string{})
}

// ErrSampleApp is the error for streaming event
func ErrSampleApp(err error) error {
	return errors.New(ErrSampleAppCode, errors.Alert, []string{"Error with sample app operation"}, []string{err.Error(), "Error occured while trying to install a sample application using manifests"}, []string{"Invalid kubeclient config", "Invalid manifest"}, []string{"Reconnect your adapter to Meshery Server to refresh the kubeclient"})
}

// ErrCustomOperation is the error for streaming event
func ErrCustomOperation(err error) error {
	return errors.New(ErrCustomOperationCode, errors.Alert, []string{"Error with custom operation"}, []string{"Error occured while applying custom manifest to the cluster", err.Error()}, []string{"Invalid kubeclient config", "Invalid manifest"}, []string{"Upload the kubconfig in the Meshery Server and reconnect the adapter"})
}

// ErrCreatingNS is the error while creating the namespace
func ErrCreatingNS(err error) error {
	return errors.New(ErrCreatingNSCode, errors.Alert, []string{"Error creating namespace"}, []string{"Error occured while applying manifest to create a namespace", err.Error()}, []string{"Invalid kubeclient config", "Invalid manifest"}, []string{"Upload the kubeconfig in the Meshery Server and reconnect the adapter"})
}

// ErrRunExecutable is the error while running an executable
func ErrRunExecutable(err error) error {
	return errors.New(ErrRunExecutableCode, errors.Alert, []string{"Error running executable"}, []string{err.Error()}, []string{"Corrupted binary", "Invalid operation"}, []string{"Check if the adaptor is executing a deprecated command"})
}

// ErrApplyHelmChart is the error for applying helm chart
func ErrApplyHelmChart(err error) error {
	return errors.New(ErrApplyHelmChartCode, errors.Alert, []string{"Error occured while applying Helm Chart"}, []string{err.Error()}, []string{}, []string{})
}

// ErrParseCiliumCoreComponent is the error when Cilium core component manifest parsing fails
func ErrParseCiliumCoreComponent(err error) error {
	return errors.New(ErrParseCiliumCoreComponentCode, errors.Alert, []string{"Cilium core component manifest parsing failing"}, []string{err.Error()}, []string{}, []string{})
}

// ErrInvalidOAMComponentType is the error when the OAM component name is not valid
func ErrInvalidOAMComponentType(compName string) error {
	return errors.New(ErrInvalidOAMComponentTypeCode, errors.Alert, []string{"invalid OAM component name: ", compName}, []string{}, []string{}, []string{})
}

// ErrCiliumCoreComponentFail is the error when core Cilium component processing fails
func ErrCiliumCoreComponentFail(err error) error {
	return errors.New(ErrCiliumCoreComponentFailCode, errors.Alert, []string{"error in Cilium core component"}, []string{err.Error()}, []string{}, []string{})
}

// ErrProcessOAM is a generic error which is thrown when an OAM operations fails
func ErrProcessOAM(err error) error {
	return errors.New(ErrProcessOAMCode, errors.Alert, []string{"error performing OAM operations"}, []string{err.Error()}, []string{}, []string{})
}

// ErrGetLatestRelease is the error for get latest versions
func ErrGetLatestRelease(err error) error {
	return errors.New(ErrGetLatestReleaseCode, errors.Alert, []string{"Could not get latest version"}, []string{err.Error()}, []string{"Latest version could not be found at the specified url"}, []string{})
}

// ErrLoadNamespace is the occurend while applying namespace
func ErrLoadNamespace(err error, s string) error {
	return errors.New(ErrLoadNamespaceCode, errors.Alert, []string{"Error occured while applying namespace "}, []string{err.Error()}, []string{"Trying to access a namespace which is not available"}, []string{"Verify presence of namespace. Confirm Meshery ServiceAccount permissions"})

}

// ErrMakingBinExecutable occurs when cilium cli binary couldn't be made
// executable
func ErrMakingBinExecutable(err error) error {
	return errors.New(ErrMakingBinExecutableCode, errors.Alert, []string{"Error while making cilium cli an executable"}, []string{err.Error()}, []string{"Download might be corrupted."}, []string{"Please retry operation."})
}

// ErrUnpackingTar is the error when tar unpack fails
func ErrUnpackingTar(err error) error {
	return errors.New(ErrUnpackingTarCode, errors.Alert, []string{"Error occured while unpacking tar"}, []string{err.Error()}, []string{"The gzip might be corrupt"}, []string{"Please retry operation."})
}

// ErrUnzipFile is the error for unzipping the file
func ErrUnzipFile(err error) error {
	return errors.New(ErrUnzipFileCode, errors.Alert, []string{"Error while unzipping"}, []string{err.Error()}, []string{"File might be corrupt"}, []string{"Please retry operation."})
}

// ErrDownloadingTar is the error when tar download fails
func ErrDownloadingTar(err error) error {
	return errors.New(ErrDownloadingTarCode, errors.Alert, []string{"Error occured while downloading Cilium tar"}, []string{err.Error()}, []string{"Error occured while download cilium tar from its release url"}, []string{"Checkout https://github.com/cilium/cilium-cli/releases/download/<release>/cilium-<platform>-<arch>.tar.gz for more details"})
}

// ErrGettingRelease is the error when getting release tag fails
func ErrGettingRelease(err error) error {
	return errors.New(ErrGettingReleaseCode, errors.Alert, []string{"Could not get latest version"}, []string{err.Error()}, []string{"Latest version could not be found at the specified url"}, []string{})
}

