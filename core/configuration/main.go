package configuration

import (
	"log"
	"path/filepath"

	"github.com/SoloDeploy/solo/core/filesystem"
	"github.com/SoloDeploy/solo/core/project"
	"github.com/spf13/viper"
)

type ProviderConfiguration struct {
	Provider string
	Options  map[string]string
}

type ProviderCollection struct {
	Git                ProviderConfiguration `mapstructure:"git" yaml:"git"`
	ContainerArtifacts ProviderConfiguration `mapstructure:"container_artifacts" yaml:"container_artifacts"`
	ContainerRuntime   ProviderConfiguration `mapstructure:"container_runtime" yaml:"container_runtime"`
	Configuration      ProviderConfiguration `mapstructure:"configuration" yaml:"configuration"`
	Secrets            ProviderConfiguration `mapstructure:"secrets" yaml:"secrets"`
}

type SoloProjectConfiguration struct {
	Name       string `mapstructure:"name" yaml:"name"`
	RootFolder string `mapstructure:"root_folder" yaml:"root_folder"`
}

type Configuration struct {
	Providers ProviderCollection       `mapstructure:"providers" yaml:"providers"`
	Project   SoloProjectConfiguration `mapstructure:"project" yaml:"project"`
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
		v.SetDefault("project.name", filepath.Base(projectFolder))
		v.Set("project.root_folder", projectFolder)
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
