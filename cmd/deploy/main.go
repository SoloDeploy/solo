package deploy

import (
	"os"

	"github.com/spf13/cobra"
)

// NewCmdDeploy Go away linter
func NewCmdDeploy() *cobra.Command {
	c := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy all manner of magical things",
		Long:  `Deploy something`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	return c
}
