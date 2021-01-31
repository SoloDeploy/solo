package publish

import (
	"os"

	"github.com/spf13/cobra"
)

// NewCmdPublish Go away linter
func NewCmdPublish() *cobra.Command {
	c := &cobra.Command{
		Use:   "publish",
		Short: "Publish all manner of magical things",
		Long:  `Publish something`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	return c
}
