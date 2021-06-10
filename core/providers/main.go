package providers

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"

	"github.com/SoloDeploy/solo/core/configuration"
	"github.com/SoloDeploy/solo/core/filesystem"
	"github.com/SoloDeploy/solo/core/network"
	"github.com/SoloDeploy/solo/core/output"
)

func isUrl(url string) bool {
	r, _ := regexp.Compile(`^https{0:1}:\/\/`)
	return r.MatchString(url)
}

// ConstructProvidersPath takes in the provider folder path and the provider name and
// returns the combined file path with a `.exe` extension if on Windows.
func ConstructProvidersPath(providersFolderPath string, providerName string) string {
	suffix := ""
	if runtime.GOOS == "windows" {
		suffix = ".exe"
	}
	return filepath.Join(providersFolderPath, fmt.Sprintf("%v%v", providerName, suffix))
}

// GetProviderPath returns the path to the executable for a specific named provider
func GetProviderPath(providerName string) (providerPath string, err error) {
	if providersFolderPath, err := GetProvidersFolder(); err == nil {
		providerPath = ConstructProvidersPath(providersFolderPath, providerName)
	}
	return
}

type providerInit struct {
	ProviderSource string
	ProviderDest   string
}

func initialiseProvider(p *providerInit) (err error) {
	if len(p.ProviderSource) > 0 {
		// TODO: check if it starts with `solo-providers-` and replace with the correct GitHub release artifact
		if isUrl(p.ProviderSource) {
			output.PrintlnfLog("Downloading Provider from %v", p.ProviderSource)
			err = network.DownloadFile(p.ProviderSource, p.ProviderDest)
			return
		}
		// TODO: make this work with relative paths
		if exists, err := filesystem.FileExists(p.ProviderSource); exists && err == nil {
			output.PrintlnfLog("Copying Provider from absolute path %v", p.ProviderSource)
			err = filesystem.CopyFile(p.ProviderSource, p.ProviderDest)
			if err != nil {
				output.PrintlnError(err)
			}
			return err
		}
		output.PrintlnfError("Provider not a URL or a local file: %v", p.ProviderSource)
	}
	return
}

func addProviderInit(providersList []*providerInit, providerConfig *configuration.ProviderConfiguration, destinationPath string) []*providerInit {
	if providerConfig != nil {
		return append(providersList, &providerInit{providerConfig.Provider, destinationPath})
	}
	return providersList
}

// InitialiseProviders downloads or copies all the Solo Providers found in the configuration
// to the Providers directory.
func InitialiseProviders(config *configuration.Configuration) (err error) {
	providerPath, err := GetProvidersFolder()

	if err != nil {
		return
	}

	providerInitList := make([]*providerInit, 0)
	providerInitList = addProviderInit(providerInitList, &config.Providers.Git, ConstructProvidersPath(providerPath, "git"))
	providerInitList = addProviderInit(providerInitList, &config.Providers.DeployerArtifacts, ConstructProvidersPath(providerPath, "deployer_artifacts"))
	providerInitList = addProviderInit(providerInitList, &config.Providers.DeployerRuntime, ConstructProvidersPath(providerPath, "deployer_runtime"))
	providerInitList = addProviderInit(providerInitList, &config.Providers.Configuration, ConstructProvidersPath(providerPath, "configuration"))
	providerInitList = addProviderInit(providerInitList, &config.Providers.Secrets, ConstructProvidersPath(providerPath, "secrets"))

	errorsOut := make(chan error, len(providerInitList))
	wgDone := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(len(providerInitList))
	for _, p := range providerInitList {
		go func(p *providerInit) {
			defer wg.Done()
			initErr := initialiseProvider(p)
			if initErr != nil {
				errorsOut <- initErr
			}
		}(p)
	}

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		break
	case err = <-errorsOut:
		close(errorsOut)
		return
	}

	return
}

// GetProviderPath returns the path to the folder where the provider executables have
// been initialised. In project folders it would be `${project}/.solo/providers` and in
// Git repositories it would be in `./.solo/providers`.
func GetProvidersFolder() (providerPath string, err error) {
	// TODO: find providers folder in project folder
	providerPath, err = filepath.Abs(filepath.Join(".solo", "providers"))
	if _, err := os.Stat(providerPath); os.IsNotExist(err) {
		output.PrintlnfInfo("Creating new Solo Providers local directory at %v", providerPath)
		err = os.MkdirAll(providerPath, fs.ModePerm)
	}
	return
}
