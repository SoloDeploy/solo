package configuration

import (
	"log"
	"path/filepath"

	"github.com/SoloDeploy/solo/core/filesystem"
	"github.com/SoloDeploy/solo/core/project"
	"github.com/spf13/viper"
)

type ProviderConfiguration struct {
	Provider      string            `yaml:"provider"`
	Options       map[string]string `yaml:"options"`
}

type ProviderCollection struct {
	Git                ProviderConfiguration `yaml:"git"`
	ContainerArtifacts ProviderConfiguration `yaml:"container_artifacts"`
	ContainerRuntime   ProviderConfiguration `yaml:"container_runtime"`
	Configuration      ProviderConfiguration `yaml:"configuration"`
	Secrets            ProviderConfiguration `yaml:"secrets"`
}

type Configuration struct {
	Providers ProviderCollection `yaml:"providers"`
}

func loadLocalConfiguration() (*viper.Viper, error) {
	localConfigPath := "./.solo.yml"
	if exists, err := filesystem.FileExists(localConfigPath); !exists || err != nil {
		return nil, err
	}

	// the config file exists, load it
	log.Println("Loading local configuration")
	v := viper.New()
	v.SetConfigFile(localConfigPath)
	v.ReadInConfig()
	return v, nil
}

func loadProjectConfiguration() (*viper.Viper, error) {
	projectFolder, err := project.FindRootFolder()
	if err != nil {
		return nil, err
	}
	if projectFolder != "" {
		log.Println("Loading project configuration")
		v := viper.New()
		projectConfigPath := filepath.Join(projectFolder, ".solo", "config.yml")
		v.SetConfigFile(projectConfigPath)
		v.ReadInConfig()
		return v, nil
	}
	return nil, nil
}

func LoadConfiguration() (configuration *Configuration, err error) {
	log.Printf("Loading configuration")

	v, err := loadLocalConfiguration()
	if err != nil {
		return
	}
	if v != nil {
		viper.MergeConfigMap(v.AllSettings())
	}

	v, err = loadProjectConfiguration()
	if err != nil {
		return
	}
	if v != nil {
		viper.MergeConfigMap(v.AllSettings())
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("SOLO")
	err = viper.Unmarshal(&configuration)

	return
}
