package init

import (
	"log"
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
			err := handler()
			if err != nil {
				log.Printf("Error initialising SoloDeploy project environment: %v", err)
				os.Exit(1)
			}
			os.Exit(0)
		},
	}

	return c
}
