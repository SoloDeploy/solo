package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/SoloDeploy/solo/cmd/build"
	"github.com/SoloDeploy/solo/cmd/create"
	"github.com/SoloDeploy/solo/cmd/deploy"
	"github.com/SoloDeploy/solo/cmd/destroy"
	initcmd "github.com/SoloDeploy/solo/cmd/init"
	"github.com/SoloDeploy/solo/cmd/install"
	"github.com/SoloDeploy/solo/cmd/promote"
	"github.com/SoloDeploy/solo/cmd/publish"
	"github.com/SoloDeploy/solo/cmd/verify"
	"github.com/SoloDeploy/solo/core/configuration"
)

var configFile string

// NewCmdSolo Go away linter
func NewCmdSolo(configuration *configuration.Configuration) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:                   "solo",
		DisableFlagsInUseLine: true,
		Short:                 "solo is the tool for Solo Deploy",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}

	rootCmd.AddCommand(build.NewCmdBuild())
	rootCmd.AddCommand(create.NewCmdCreate())
	rootCmd.AddCommand(deploy.NewCmdDeploy())
	rootCmd.AddCommand(destroy.NewCmdDestroy())
	rootCmd.AddCommand(initcmd.NewCmdInit(configuration))
	rootCmd.AddCommand(install.NewCmdInstall())
	rootCmd.AddCommand(promote.NewCmdPromote())
	rootCmd.AddCommand(publish.NewCmdPublish())
	rootCmd.AddCommand(verify.NewCmdVerify())

	return rootCmd
}
