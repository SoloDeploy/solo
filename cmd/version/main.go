package version

import (
	"fmt"

	"github.com/spf13/cobra"
	goVersion "go.hein.dev/go-version"
)

var output = "json"

// NewCmdVersion creates a command to output the current version of SoloDeploy
func NewCmdVersion(version string, commit string, date string) *cobra.Command {
	c := &cobra.Command{
		Use:   "version",
		Short: "Version will output the current build information",
		Long:  "Prints the version, Git commit ID and commit date in JSON or YAML format using the go.hein.dev/go-version package.",
		Run: func(_ *cobra.Command, _ []string) {
			resp := goVersion.FuncWithOutput(false, version, commit, date, output)
			fmt.Print(resp)
		},
	}

	c.Flags().StringVarP(&output, "output", "o", "json", "Output format. One of 'yaml' or 'json'.")

	return c
}
