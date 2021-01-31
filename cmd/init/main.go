package init

import (
	"os"

	"github.com/spf13/cobra"
)

// NewCmdInit Go away linter
func NewCmdInit() *cobra.Command {
	c := &cobra.Command{
		Use:   "init",
		Short: "Initialize my local environment with all manner of magical things",
		Long:  `Initialize local environment`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	return c
}
