package init

import (
	"log"
	"os"

	"github.com/SoloDeploy/solo/core/configuration"
	"github.com/SoloDeploy/solo/core/output"
	"github.com/spf13/cobra"
)

const long_description = `Run this command to initialise the local environment with a new or existing project.

If no config file (.solo/config.yml) is found, a new default project is initialised
using the default settings. Settings can be overiden using commandline flags.

If an existing config file is found, no flags are required and the CLI assumes that
the project already exists. It then finds all the controller, manifest, and deployable
Git repositories and clones them into a logical sub-directory structure. If any of the
Git repositories are already cloned locally, the CLI prints a warning, but does not
clone it.

In either case the CLI downloads the correct Git Provider implementation, starts it,
and tests that the connection to the remote Git server is authorised.`

// NewCmdInit Go away linter
func NewCmdInit(configuration *configuration.Configuration) *cobra.Command {
	c := &cobra.Command{
		Use:   "init",
		Short: "Initialise the local environment",
		Long:  long_description,
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("Initialising local project folder")
			err := validate(configuration)
			if err != nil {
				output.PrintfError("Validation failed: %v", err)
				os.Exit(1)
			}
			err = handler(configuration)
			if err != nil {
				output.PrintfError("Error initialising SoloDeploy project environment: %v", err)
				os.Exit(1)
			}
			os.Exit(0)
		},
	}

	return c
}
