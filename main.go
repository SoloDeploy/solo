package main

import (
	"fmt"
	"os"

	solo "github.com/SoloDeploy/solo/cmd"
)

func main() {
	command := solo.NewCmdSolo()

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
