package solo

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/SoloDeploy/solo/cmd/build"
	"github.com/SoloDeploy/solo/cmd/create"
	"github.com/SoloDeploy/solo/cmd/deploy"
	"github.com/SoloDeploy/solo/cmd/destroy"
	initcmd "github.com/SoloDeploy/solo/cmd/init"
	"github.com/SoloDeploy/solo/cmd/install"
	"github.com/SoloDeploy/solo/cmd/promote"
	"github.com/SoloDeploy/solo/cmd/publish"
	"github.com/SoloDeploy/solo/cmd/verify"
)

var configFile string

// NewCmdSolo Go away linter
func NewCmdSolo() *cobra.Command {
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
	rootCmd.AddCommand(initcmd.NewCmdInit())
	rootCmd.AddCommand(install.NewCmdInstall())
	rootCmd.AddCommand(promote.NewCmdPromote())
	rootCmd.AddCommand(publish.NewCmdPublish())
	rootCmd.AddCommand(verify.NewCmdVerify())

	return rootCmd
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	configFile = ".solo.yml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("SOLO")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}
}
