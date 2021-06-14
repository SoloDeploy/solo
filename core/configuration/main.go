package configuration

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/SoloDeploy/solo/core/filesystem"
	"github.com/SoloDeploy/solo/core/output"
	"github.com/SoloDeploy/solo/core/project"
	"github.com/spf13/viper"
)

// ProviderConfiguration holds the generic configuration read from various SoloDeploy
// configuration sources.
type ProviderConfiguration struct {

	// Provider is the source of the Provider executable. It can be defined in any of
	// the following sources:
	//   * Local file path (relative or absolute).
	//   * URL (including protocol).
	//   * Standardised keys for Solo published providers. For example, "solo-providers-git-github" will be mapped to
	//     https://github.com/SoloDeploy/solo-providers-git-github/releases/download/v0.0.1/git/github.$OS-$ARCH
	Provider string

	// Options are key value pairs that are specific to the provider implementation. See
	// the Provider's documentation for more info.
	Options map[string]string
}

// ProviderCollection contains the configuration of all 5 possible Solo Providers. This
// configuration is populated in various input sources. For example, in a project folder
// the `./.solo/config.yml` file can be used, but in a Solo manifest, controller or deployable
// repository the `./.solo.yml` file can be used. Alternatively, environment variables can
// be used.
type ProviderCollection struct {
	Git               ProviderConfiguration `mapstructure:"git" yaml:"git"`
	DeployerArtifacts ProviderConfiguration `mapstructure:"deployer_artifacts" yaml:"deployer_artifacts"`
	DeployerRuntime   ProviderConfiguration `mapstructure:"deployer_runtime" yaml:"deployer_runtime"`
	Configuration     ProviderConfiguration `mapstructure:"configuration" yaml:"configuration"`
	Secrets           ProviderConfiguration `mapstructure:"secrets" yaml:"secrets"`
}

// SoloProjectConfiguration contains the configuration for project info when running in
// the context of a Solo project folder.
type SoloProjectConfiguration struct {

	// Name sets the name of the project. Optional. When empty, this defaults to the project
	// folder name. The project name is used in the repository naming conventions. For example,
	// a manifest naming convention is `${projectName}.manifest.${layerName}.${environmentName}`
	Name string `mapstructure:"name" yaml:"name"`

	// RootFolder is set dynamically if SoloDeploy CLI is executed in context of a Solo
	// project, i.e. if a parent directory with the project configuration file is found
	// (./.solo/config.yml)
	RootFolder string `mapstructure:"root_folder" yaml:"root_folder"`
}

// Configuration contains all the configuration found in various sources. Configuration values
// are applied in the following order, duplicate values are overwritten:
//   * User config (~/.solo.yml)
//   * Project folder config (${projectRootFolder}/.solo/config.yml)
//   * Repository config (${repositoryRoot}/.solo.yml)
//   * Environment variables (prefixed with SOLO_* and ALL_CAPS - example `SOLO_PROJECT_NAME`)
type Configuration struct {

	// Providers contains the configuration for the 5 possible Solo providers.
	Providers ProviderCollection `mapstructure:"providers" yaml:"providers"`

	// Project contains the configuration for Solo project settings.
	Project SoloProjectConfiguration `mapstructure:"project" yaml:"project"`
}

func loadConfigurationFile(configPath string) error {
	if exists, err := filesystem.FileExists(configPath); !exists || err != nil {
		return err
	}

	output.PrintlnfLog("Loading configuration from %v", configPath)
	v := viper.New()
	v.SetConfigFile(configPath)
	v.ReadInConfig()
	viper.MergeConfigMap(v.AllSettings())
	return nil
}

func loadProjectConfiguration() error {
	projectFolder, err := project.FindRootFolder()
	if err != nil {
		return err
	}
	if projectFolder != "" {
		viper.SetDefault("project.name", filepath.Base(projectFolder))
		viper.Set("project.root_folder", projectFolder)
		projectConfigPath := filepath.Join(projectFolder, ".solo", "config.yml")
		return loadConfigurationFile(projectConfigPath)
	}
	return nil
}

// LoadConfiguration looks at all the configuration sources and loads the Solo configuration
// in the following order, duplicate values are overwritten:
//   * User config (~/.solo.yml)
//   * Project folder config (${projectRootFolder}/.solo/config.yml)
//   * Repository config (${repositoryRoot}/.solo.yml)
//   * Environment variables (prefixed with SOLO_* and ALL_CAPS - example `SOLO_PROJECT_NAME`)
func LoadConfiguration() (configuration *Configuration, err error) {
	output.PrintlnLog("Loading configuration")

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	err = loadConfigurationFile(path.Join(userHomeDir, ".solo.yml"))
	if err != nil {
		return
	}

	err = loadConfigurationFile("./.solo.yml")
	if err != nil {
		return
	}

	err = loadProjectConfiguration()
	if err != nil {
		return
	}

	viper.SetEnvPrefix("SOLO")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err = viper.Unmarshal(&configuration)

	return
}
