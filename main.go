package main

import (
	"fmt"
	"os"

	"github.com/SoloDeploy/solo/cmd"
	"github.com/SoloDeploy/solo/core/configuration"
	"github.com/SoloDeploy/solo/core/output"
)

func main() {
	configuration, err := configuration.LoadConfiguration()
	if err != nil {
		output.PrintError(err)
		os.Exit(1)
	}
	// TODO: register dependency in inject object graph instead of passing it down the execution tree

	fmt.Printf("Name: %v\n", configuration.Project.Name)
	fmt.Printf("RootFolder: %v\n", configuration.Project.RootFolder)
	command := cmd.NewCmdSolo(configuration)

	if err := command.Execute(); err != nil {
		output.PrintError(err)
		os.Exit(1)
	}
}
