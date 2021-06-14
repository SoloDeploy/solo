package main

import (
	"fmt"
	"os"

	solo "github.com/SoloDeploy/solo/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	command := solo.NewCmdSolo(version, commit, date)

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
