package init

import (
	"log"

	"github.com/SoloDeploy/solo/core/configuration"
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
	log.Printf("Repositories: %v", names)
	return
}
