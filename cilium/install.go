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
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/google/go-github/github"
	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/status"
	"github.com/layer5io/meshery-cilium/internal/config"
	mesherykube "github.com/layer5io/meshkit/utils/kubernetes"
)

const (
	platform = runtime.GOOS
	arch = runtime.GOARCH
)

func (h *Handler) installCilium(del bool, version, ns string) (string, error) {
	h.Log.Debug(fmt.Sprintf("Requested install of version: %s", version))
	h.Log.Debug(fmt.Sprintf("Requested action is delete: %v", del))
	h.Log.Debug(fmt.Sprintf("Requested action is in namespace: %s", ns))

	st := status.Installing
	if del {
		st = status.Removing
	}

	err := h.Config.GetObject(adapter.MeshSpecKey, h)
	if err != nil {
		return st, ErrMeshConfig(err)
	}

	h.Log.Info("Installing...")
	err = h.applyHelmChart(del, version, ns)
	if err != nil {
		h.Log.Error(ErrInstallCilium((err)))
		
		err = h.runCiliumCliCmd(ns, del)
		if err != nil {
			return st, ErrInstallCilium(err)
		}
	}

	st = status.Installed
	if del {
		st = status.Removed
	}

	return st, nil
}

func (h *Handler) applyHelmChart(del bool, version, namespace string) error {
	kClient := h.MesheryKubeclient

	repo := "https://helm.cilium.io/"
	chart := "cilium"
	var act mesherykube.HelmChartAction
	if del {
		act = mesherykube.UNINSTALL
	} else {
		act = mesherykube.INSTALL
	}
	return kClient.ApplyHelmChart(mesherykube.ApplyHelmChartConfig{
		ChartLocation: mesherykube.HelmChartLocation{
			Repository: repo,
			Chart:      chart,
			Version:    version,
		},
		Namespace:       "kube-system",
		Action:          act,
		CreateNamespace: true,
		ReleaseName:     chart,
	})
}

func (h *Handler) runCiliumCliCmd(namespace string, isDeleteOp bool) error {
	var (
		out bytes.Buffer
		er  bytes.Buffer
	)

	version, err := getReleaseTag()
	if (err != nil) {
		return ErrGettingRelease(err)
	}

	Executable, err := h.getExecutable(version)
	if err != nil {
		return ErrDownloadBinary(err)
	}
	execCmd := []string{"install"}
	if isDeleteOp {
		execCmd = []string{"uninstall"}
	}

	// We need a variable executable here hence using nosec
	// #nosec
	command := exec.Command(Executable, execCmd...)
	command.Stdout = &out
	command.Stderr = &er
	err = command.Run()
	if err != nil {        
		return ErrRunExecutable(err)
	}

	return nil
}

// getExecutable looks for the executable in
// 1. $PATH
// 2. Root config path
//
// If it doesn't find the executable in the path then it proceeds
// to download the binary from github releases and installs it
// in the root config path
func (h *Handler) getExecutable(release string) (string, error) {
	const binaryName = "cilium"
	alternateBinaryName := generatePlatformSpecificBinaryName("cilium-" , platform)

	// Look for the executable in the path
	h.Log.Info("Looking for cilium in the path...")
	executable, err := exec.LookPath(binaryName)
	if err == nil {
		return executable, nil
	}
	executable, err = exec.LookPath(alternateBinaryName)
	if err == nil {
		return executable, nil
	}

	// Look for config in the root path
	binPath := path.Join(config.RootPath(), "bin")
	h.Log.Info("Looking for cilium in", binPath, "...")
	executable = path.Join(binPath, binaryName)
	if _, err := os.Stat(executable); err == nil {
		return executable, nil
	}

	// Proceed to download the binary in the config root path
	h.Log.Info("cilium not found in the path, downloading...")
	res, err := downloadTar(release) // realease now hard-coded, resolve this

	if err != nil {
		return "", ErrDownloadingTar(err)
	}
	err = extractTar(res, binPath)
	
	// Install the binary
	h.Log.Info("Installing...")
	
	// Move binary to the right location
	// err = os.Rename(path.Join(downloadLocation, binaryName), path.Join(binPath, "cilium"))
	if err != nil {
		return "", ErrInstallCilium(err)
	}
	if err != nil {
		return "", err
	}

	h.Log.Info("Done")
	return path.Join(binPath, binaryName), nil
}

func downloadTar(release string) (*http.Response, error) {
	var url = "https://github.com/cilium/cilium-cli/releases/download"
	switch platform {
	case "windows":
		url = fmt.Sprintf("%s/%s/cilium-%s-%s.tar.gz", url, release, platform, arch)
	case "darwin":
		fallthrough
	case "linux":
		url = fmt.Sprintf("%s/%s/cilium-%s-%s.tar.gz", url, release, platform, arch)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, ErrDownloadingTar(err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrDownloadingTar(fmt.Errorf("bad status: %s", resp.Status))
	}
	

	return resp, nil
}

func getReleaseTag() (string, error) {
	client := github.NewClient(nil)

	tags, _, err := client.Repositories.ListTags(context.Background(), "cilium", "cilium-cli", nil)
	
	if err != nil {
		return "", err
	}
	return *tags[0].Name, nil
}

func extractTar(res *http.Response, location string) error {
	// Close the response body
	defer func() {
		if err := res.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	switch platform {
	case "darwin":
		fallthrough
	case "linux":
		if err := tarxzf(location, res.Body); err != nil {
			//ErrExtracingFromTar
			return ErrUnpackingTar(err)
		}
	case "windows":
		if err := unzip(location, res.Body); err != nil {
			return ErrUnpackingTar(err)
		}
	}
	

	return nil
}

func tarxzf(location string, stream io.Reader) error {
	uncompressedStream, err := gzip.NewReader(stream)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return ErrTarXZF(err)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			// File traversal is required to store the extracted manifests at the right place
			// #nosec
			if err := os.MkdirAll(path.Join(location, header.Name), 0750); err != nil {
				return ErrTarXZF(err)
			}
		case tar.TypeReg:
			// File traversal is required to store the extracted manifests at the right place
			// #nosec
			outFile, err := os.Create(path.Join(location, header.Name))
			if err != nil {
				return ErrTarXZF(err)
			}
			// Trust cilium tar
			// #nosec
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return ErrTarXZF(err)
			}
			if err = outFile.Close(); err != nil {
				return ErrTarXZF(err)
			}

			if header.FileInfo().Name() == "cilium" {
				// cilium binary needs to be executable
				// #nosec
				if err = os.Chmod(outFile.Name(), 0750); err != nil {
					return ErrInstallBinary(err)
				}
			}
			

		default:
			return ErrTarXZF(err)
		}
	}

	return nil
}


func unzip(location string, zippedContent io.Reader) error {
	// Keep file in memory: Approx size ~ 50MB
	// TODO: Find a better approach
	zipped, err := ioutil.ReadAll(zippedContent)
	if err != nil {
		return ErrUnzipFile(err)
	}

	zReader, err := zip.NewReader(bytes.NewReader(zipped), int64(len(zipped)))
	if err != nil {
		return ErrUnzipFile(err)
	}

	for _, file := range zReader.File {
		zippedFile, err := file.Open()
		if err != nil {
			return ErrUnzipFile(err)
		}
		defer func() {
			if err := zippedFile.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		// need file traversal to place the extracted files at the right place, hence
		// #nosec
		extractedFilePath := path.Join(location, file.Name)

		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(extractedFilePath, file.Mode()); err != nil {
				return ErrUnzipFile(err)
			}
		} else {
			// we need a variable path hence,
			// #nosec
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				return ErrUnzipFile(err)
			}

			/* #nosec G307 */

			defer func() {
				if err := outputFile.Close(); err != nil {
					fmt.Println(err)
				}
			}()

			// Trust cilium zip hence,
			// #nosec
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				return ErrUnzipFile(err)
			}
		}
	}

	return nil
}


func generatePlatformSpecificBinaryName(binName, platform string) string {
	if platform == "windows" && !strings.HasSuffix(binName, ".exe") {
		return binName + platform + ".exe"
	}

	return binName + platform
}
