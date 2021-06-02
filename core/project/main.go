package project

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func FindRootFolder() (string, error) {
	log.Println("Looking for project directory")
	directoryUnderTest, err := os.Getwd()

	for {
		configPath := filepath.Join(directoryUnderTest, ".solo/config.yml")

		_, err = ioutil.ReadFile(configPath)
		if err == nil {
			log.Printf("Config file found. Project root folder at %v", directoryUnderTest)
			return directoryUnderTest, nil
		}

		parentPath := filepath.Dir(directoryUnderTest)
		if parentPath == directoryUnderTest {
			log.Println("No project directory found")
			return "", nil
		}

		directoryUnderTest = parentPath
	}
}
