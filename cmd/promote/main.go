package promote

import (
	"os"

	"github.com/spf13/cobra"
)

// NewCmdPromote Go away linter
func NewCmdPromote() *cobra.Command {
	c := &cobra.Command{
		Use:   "promote",
		Short: "Promote all manner of magical things",
		Long:  `Promote something`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	return c
}
