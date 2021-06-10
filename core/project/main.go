package project

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/SoloDeploy/solo/core/output"
)

// FindRootFolder traverses the current folder hierarchy to look for a project
// configuration file (`./.solo/config.yml`). If found it returns the path to
// the project folder, otherwise it returns an empty string.
func FindRootFolder() (string, error) {
	output.PrintlnLog("Looking for project directory")
	directoryUnderTest, err := os.Getwd()

	for {
		configPath := filepath.Join(directoryUnderTest, ".solo/config.yml")

		_, err = ioutil.ReadFile(configPath)
		if err == nil {
			output.PrintlnLog("Project directory found")
			return directoryUnderTest, nil
		}

		parentPath := filepath.Dir(directoryUnderTest)
		if parentPath == directoryUnderTest {
			output.PrintlnLog("No project directory found")
			return "", nil
		}

		directoryUnderTest = parentPath
	}
}
