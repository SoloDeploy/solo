package verify

import (
	"os"

	"github.com/spf13/cobra"
)

// NewCmdVerify Go away linter
func NewCmdVerify() *cobra.Command {
	c := &cobra.Command{
		Use:   "verify",
		Short: "Verify all manner of magical things",
		Long:  `Verify something`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	return c
}
