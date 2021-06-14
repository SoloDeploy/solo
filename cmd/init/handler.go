package init

import (
	"github.com/SoloDeploy/solo/core/configuration"
	"github.com/SoloDeploy/solo/core/output"
	"github.com/SoloDeploy/solo/core/providers"
)

func handler(configuration *configuration.Configuration) (err error) {
	providers.InitialiseProviders(configuration)
	gitProvider, err := providers.NewGitProvider(configuration)
	if err != nil {
		return
	}
	defer gitProvider.Close()
	names, err := gitProvider.GetRepositoryNames()
	if err != nil {
		return
	}
	output.PrintlnfInfo("Repositories: %v", names)
	return
}
