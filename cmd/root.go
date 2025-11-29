package cmd

import (
	"fmt"
	"os"

	"github.com/135yshr/devcycle-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	outputFormat string
)

var rootCmd = &cobra.Command{
	Use:     "dvcx",
	Short:   "DevCycle CLI - Unofficial command-line tool for DevCycle Management API",
	Version: GetVersion(),
	Long: `dvcx is an unofficial command-line tool that enables the use of
the DevCycle Management API from the command line.

It provides comprehensive access to DevCycle features including:
  - Projects management
  - Features and Variables
  - Environments
  - Targeting rules
  - Audiences and Overrides
  - Audit logs and Metrics`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .devcycle/config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "output format (table, json, yaml)")

	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		viper.AddConfigPath(pwd + "/.devcycle")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.SetEnvPrefix("DVCX")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("debug") {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}

	config.Load()
}

func GetOutput() string {
	return viper.GetString("output")
}
