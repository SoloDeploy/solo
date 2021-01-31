package destroy

import (
	"os"

	"github.com/spf13/cobra"
)

// NewCmdDestroy Go away linter
func NewCmdDestroy() *cobra.Command {
	c := &cobra.Command{
		Use:   "destroy",
		Short: "Destroy all manner of magical things",
		Long:  `Destroy something`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	return c
}
