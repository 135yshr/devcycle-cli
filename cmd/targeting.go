package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/135yshr/devcycle-cli/internal/output"
	"github.com/135yshr/devcycle-cli/pkg/api"
	"github.com/spf13/cobra"
)

var targetingCmd = &cobra.Command{
	Use:   "targeting",
	Short: "Manage feature targeting configurations",
	Long:  `Get and update feature targeting configurations.`,
}

var targetingGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get targeting configuration for a feature",
	Long:  `Get the targeting configuration for a specific feature.`,
	RunE:  runTargetingGet,
}

var targetingUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update targeting configuration",
	Long:  `Update the targeting configuration for a feature.`,
	RunE:  runTargetingUpdate,
}

var targetingEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a feature for an environment",
	Long:  `Enable a feature for a specific environment.`,
	RunE:  runTargetingEnable,
}

var targetingDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable a feature for an environment",
	Long:  `Disable a feature for a specific environment.`,
	RunE:  runTargetingDisable,
}

var targetingProject string
var targetingFeature string
var targetingEnvironment string
var targetingFromFile string

func init() {
	rootCmd.AddCommand(targetingCmd)
	targetingCmd.AddCommand(targetingGetCmd)
	targetingCmd.AddCommand(targetingUpdateCmd)
	targetingCmd.AddCommand(targetingEnableCmd)
	targetingCmd.AddCommand(targetingDisableCmd)

	// Persistent flags for all targeting commands
	targetingCmd.PersistentFlags().StringVarP(&targetingProject, "project", "p", "", "project key (uses config default if not specified)")
	targetingCmd.PersistentFlags().StringVarP(&targetingFeature, "feature", "f", "", "feature key (required)")

	// Update command flags
	targetingUpdateCmd.Flags().StringVarP(&targetingFromFile, "from-file", "F", "", "JSON input file for configuration update, use '-' for stdin")

	// Enable/Disable command flags
	targetingEnableCmd.Flags().StringVarP(&targetingEnvironment, "environment", "e", "", "environment key (required)")
	targetingDisableCmd.Flags().StringVarP(&targetingEnvironment, "environment", "e", "", "environment key (required)")
}

type targetingTableData struct {
	configs map[string]*api.EnvironmentConfig
}

func (d targetingTableData) Headers() []string {
	return []string{"ENVIRONMENT", "STATUS", "TARGETS"}
}

func (d targetingTableData) Rows() [][]string {
	rows := make([][]string, 0, len(d.configs))
	for envKey, config := range d.configs {
		if config == nil {
			continue
		}
		targetCount := len(config.Targets)
		rows = append(rows, []string{
			envKey,
			config.Status,
			fmt.Sprintf("%d rule(s)", targetCount),
		})
	}
	return rows
}

func getTargetingProjectKey() string {
	if targetingProject != "" {
		return targetingProject
	}
	return getProjectKey()
}

func runTargetingGet(cmd *cobra.Command, args []string) error {
	projectKey := getTargetingProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}
	if targetingFeature == "" {
		return fmt.Errorf("required flag \"feature\" not set")
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	configs, err := client.FeatureConfigurations(ctx, projectKey, targetingFeature)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(targetingTableData{configs: configs})
	}
	return printer.Print(configs)
}

func runTargetingUpdate(cmd *cobra.Command, args []string) error {
	projectKey := getTargetingProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}
	if targetingFeature == "" {
		return fmt.Errorf("required flag \"feature\" not set")
	}
	if targetingFromFile == "" {
		return fmt.Errorf("required flag \"from-file\" not set")
	}

	// Load configuration from file
	var data []byte
	var err error

	if targetingFromFile == "-" {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("failed to read from stdin: %w", err)
		}
	} else {
		data, err = os.ReadFile(targetingFromFile)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", targetingFromFile, err)
		}
	}

	var configs map[string]*api.EnvironmentConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &api.UpdateFeatureConfigurationsRequest{
		Configurations: configs,
	}

	result, err := client.UpdateFeatureConfigurations(ctx, projectKey, targetingFeature, req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(targetingTableData{configs: result})
	}
	return printer.Print(result)
}

func runTargetingEnable(cmd *cobra.Command, args []string) error {
	projectKey := getTargetingProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}
	if targetingFeature == "" {
		return fmt.Errorf("required flag \"feature\" not set")
	}
	if targetingEnvironment == "" {
		return fmt.Errorf("required flag \"environment\" not set")
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.EnableFeature(ctx, projectKey, targetingFeature, targetingEnvironment); err != nil {
		return err
	}

	cmd.Printf("Feature '%s' enabled for environment '%s'\n", targetingFeature, targetingEnvironment)
	return nil
}

func runTargetingDisable(cmd *cobra.Command, args []string) error {
	projectKey := getTargetingProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}
	if targetingFeature == "" {
		return fmt.Errorf("required flag \"feature\" not set")
	}
	if targetingEnvironment == "" {
		return fmt.Errorf("required flag \"environment\" not set")
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.DisableFeature(ctx, projectKey, targetingFeature, targetingEnvironment); err != nil {
		return err
	}

	cmd.Printf("Feature '%s' disabled for environment '%s'\n", targetingFeature, targetingEnvironment)
	return nil
}
