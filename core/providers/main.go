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

func ConstructProvidersPath(providersFolderPath string, providerName string) string {
	suffix := ""
	if runtime.GOOS == "windows" {
		suffix = ".exe"
	}
	return filepath.Join(providersFolderPath, fmt.Sprintf("%v%v", providerName, suffix))
}

func GetProviderPath(providerName string) (providerPath string, err error) {
	if providersFolderPath, err := GetProvidersFolder(); err == nil {
		providerPath = ConstructProvidersPath(providersFolderPath, providerName)
	}
	return
}

type ProviderInit struct {
	ProviderSource string
	ProviderDest   string
}

func initialiseProvider(providerInit *ProviderInit) (err error) {
	if len(providerInit.ProviderSource) > 0 {
		if isUrl(providerInit.ProviderSource) {
			output.FPrintlnLog("Downloading Provider from %v", providerInit.ProviderSource)
			err = network.DownloadFile(providerInit.ProviderSource, providerInit.ProviderDest)
			return
		}
		if exists, err := filesystem.FileExists(providerInit.ProviderSource); exists && err == nil {
			output.FPrintlnLog("Copying Provider from absolute path %v", providerInit.ProviderSource)
			err = filesystem.CopyFile(providerInit.ProviderSource, providerInit.ProviderDest)
			if err != nil {
				output.PrintlnError(err)
			}
			return err
		}
		output.FPrintlnError("Provider not a URL or a local file: %v", providerInit.ProviderSource)
	}
	return
}

func addProviderInit(providersList []*ProviderInit, providerConfig *configuration.ProviderConfiguration, destinationPath string) []*ProviderInit {
	if providerConfig != nil {
		return append(providersList, &ProviderInit{providerConfig.Provider, destinationPath})
	}
	return providersList
}

func InitialiseProviders(config *configuration.Configuration) (err error) {
	providerPath, err := GetProvidersFolder()

	if err != nil {
		return
	}

	providerInitList := make([]*ProviderInit, 0)
	providerInitList = addProviderInit(providerInitList, &config.Providers.Git, ConstructProvidersPath(providerPath, "git"))
	providerInitList = addProviderInit(providerInitList, &config.Providers.ContainerArtifacts, ConstructProvidersPath(providerPath, "container_artifacts"))
	providerInitList = addProviderInit(providerInitList, &config.Providers.ContainerRuntime, ConstructProvidersPath(providerPath, "container_runtime"))
	providerInitList = addProviderInit(providerInitList, &config.Providers.Configuration, ConstructProvidersPath(providerPath, "configuration"))
	providerInitList = addProviderInit(providerInitList, &config.Providers.Secrets, ConstructProvidersPath(providerPath, "secrets"))

	errorsOut := make(chan error, len(providerInitList))
	wgDone := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(len(providerInitList))
	for _, providerInit := range providerInitList {
		go func(providerInit *ProviderInit) {
			defer wg.Done()
			initErr := initialiseProvider(providerInit)
			if initErr != nil {
				errorsOut <- initErr
			}
		}(providerInit)
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

func GetProvidersFolder() (providerPath string, err error) {
	// TODO: find providers folder in project folder
	providerPath, err = filepath.Abs(filepath.Join(".solo", "providers"))
	if _, err := os.Stat(providerPath); os.IsNotExist(err) {
		output.FPrintlnInfo("Creating new Solo Providers local directory at %v", providerPath)
		err = os.MkdirAll(providerPath, fs.ModePerm)
	}
	return
}
