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

func initialiseProvider(provider *configuration.ProviderConfiguration, downloadPath string, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	if provider != nil && len(provider.Provider) > 0 {
		if isUrl(provider.Provider) {
			output.FPrintlnLog("Downloading Provider from %v", provider.Provider)
			err = network.DownloadFile(provider.Provider, downloadPath)
			return
		}
		if exists, err := filesystem.FileExists(provider.Provider); exists && err == nil {
			output.FPrintlnLog("Copying Provider from absolute path %v", provider.Provider)
			err = filesystem.CopyFile(provider.Provider, downloadPath)
			if err != nil {
				output.PrintlnError(err)
			}
			return err
		}
		output.FPrintlnError("Provider not a URL or a local file: %v", provider.Provider)
	}
	return
}

func InitialiseProviders(config *configuration.Configuration) (err error) {
	providerPath, err := GetProvidersFolder()

	if err != nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(5)
	go initialiseProvider(&config.Providers.Git, ConstructProvidersPath(providerPath, "git"), &wg)
	go initialiseProvider(&config.Providers.ContainerArtifacts, ConstructProvidersPath(providerPath, "container_artifacts"), &wg)
	go initialiseProvider(&config.Providers.ContainerRuntime, ConstructProvidersPath(providerPath, "container_runtime"), &wg)
	go initialiseProvider(&config.Providers.Configuration, ConstructProvidersPath(providerPath, "configuration"), &wg)
	go initialiseProvider(&config.Providers.Secrets, ConstructProvidersPath(providerPath, "secrets"), &wg)
	wg.Wait()
	// TODO: handle errors from above async operations
	return nil
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
