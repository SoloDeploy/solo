package build

import (
	"os"

	"github.com/spf13/cobra"
)

// NewCmdBuild Go away linter
func NewCmdBuild() *cobra.Command {
	c := &cobra.Command{
		Use:   "build",
		Short: "Build all manner of magical things",
		Long:  `Build something`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	return c
}
