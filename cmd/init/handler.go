package init

import (
	"log"

	"github.com/SoloDeploy/solo/providers"
)

func handler() (err error) {
	gitProvider, err := providers.NewGitProvider()
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
