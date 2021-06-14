package init

import (
	"errors"

	"github.com/SoloDeploy/solo/core/configuration"
)

var noGitProviderError = errors.New(`No Git Provider configured.
The init command clones all the existing project repositories. It therefore needs to interact with Git using the Git Provider.
Configure the provider in the repository configuration file (./solo.yml) or in the root of the project folder (.solo/config.yml).`)

func validate(config *configuration.Configuration) error {
	if config.Providers.Git.Provider == "" {
		return noGitProviderError
	}
	return nil
}
