package create

import (
	"os"

	"github.com/spf13/cobra"
)

// NewCmdCreate Go away linter
func NewCmdCreate() *cobra.Command {
	c := &cobra.Command{
		Use:   "create",
		Short: "Create all manner of magical things",
		Long:  `Create something`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	return c
}
