package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	ConfigDir  = ".devcycle"
	ConfigFile = "config.yaml"
	TokenFile  = "token.json"
)

type Config struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	Project      string `mapstructure:"project"`
	Environment  string `mapstructure:"environment"`
	Output       string `mapstructure:"output"`
	Debug        bool   `mapstructure:"debug"`
}

var current Config

func Load() {
	viper.Unmarshal(&current)
}

func Get() *Config {
	return &current
}

func GetConfigDir() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(pwd, ConfigDir), nil
}

func GetConfigPath() (string, error) {
	dir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, ConfigFile), nil
}

func GetTokenPath() (string, error) {
	dir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, TokenFile), nil
}

func EnsureConfigDir() error {
	dir, err := GetConfigDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(dir, 0700)
}

func SetProject(project string) {
	current.Project = project
	viper.Set("project", project)
}

func SetEnvironment(env string) {
	current.Environment = env
	viper.Set("environment", env)
}

func GetProject() string {
	return viper.GetString("project")
}

func GetEnvironment() string {
	return viper.GetString("environment")
}

func GetClientID() string {
	return viper.GetString("client_id")
}

func GetClientSecret() string {
	return viper.GetString("client_secret")
}

func IsDebug() bool {
	return viper.GetBool("debug")
}
