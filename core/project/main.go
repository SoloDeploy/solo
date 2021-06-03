package project

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/SoloDeploy/solo/core/output"
)

func FindRootFolder() (string, error) {
	output.PrintlnLog("Looking for project directory")
	directoryUnderTest, err := os.Getwd()

	for {
		configPath := filepath.Join(directoryUnderTest, ".solo/config.yml")

		_, err = ioutil.ReadFile(configPath)
		if err == nil {
			output.FPrintlnLog("Config file found. Project root folder at %v", directoryUnderTest)
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
