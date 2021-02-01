package install

import (
	"os"

	"github.com/spf13/cobra"
)

// NewCmdInstall Go away linter
func NewCmdInstall() *cobra.Command {
	c := &cobra.Command{
		Use:   "install",
		Short: "Install all manner of magical things",
		Long:  `Install something`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	return c
}
