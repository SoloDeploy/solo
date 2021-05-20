package configuration

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type ProviderConfiguration struct {
	Provider string            `yaml:"provider"`
	Options  map[string]string `yaml:"options"`
}

type ProviderCollection struct {
	Git              ProviderConfiguration `yaml:"git"`
	ArtifactRegistry ProviderConfiguration `yaml:"artifacts_registry"`
	ContainerCluster ProviderConfiguration `yaml:"container_cluster"`
	Configuration    ProviderConfiguration `yaml:"configuration"`
	Secrets          ProviderConfiguration `yaml:"secrets"`
}

type Configuration struct {
	Providers ProviderCollection `yaml:"providers"`
}

func loadConfiguration() (config *Configuration, err error) {
	config = &Configuration{}

	filename, _ := filepath.Abs("./.solo/config.yml")
	yamlContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlContent, &config)
	if err != nil {
		return nil, err
	}

	fmt.Print("Providers: ", config.Providers)
	fmt.Print("Providers.Git: ", config.Providers.Git)
	fmt.Print("Providers.Git.Options.auth: ", config.Providers.Git.Options["auth"])

	return
}
