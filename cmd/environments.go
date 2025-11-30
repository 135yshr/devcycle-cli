package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/135yshr/devcycle-cli/internal/config"
	"github.com/135yshr/devcycle-cli/internal/output"
	"github.com/135yshr/devcycle-cli/pkg/api"
	"github.com/spf13/cobra"
)

var environmentsCmd = &cobra.Command{
	Use:     "environments",
	Aliases: []string{"envs", "env"},
	Short:   "Manage DevCycle environments",
	Long:    `List and view DevCycle environments.`,
}

var environmentsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all environments",
	Long:  `List all environments in a project.`,
	RunE:  runEnvironmentsList,
}

var environmentsGetCmd = &cobra.Command{
	Use:   "get [environment-key]",
	Short: "Get environment details",
	Long:  `Get detailed information about a specific environment.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runEnvironmentsGet,
}

var environmentsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new environment",
	Long:  `Create a new environment in a project.`,
	RunE:  runEnvironmentsCreate,
}

var environmentsUpdateCmd = &cobra.Command{
	Use:   "update [environment-key]",
	Short: "Update an environment",
	Long:  `Update an existing environment.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runEnvironmentsUpdate,
}

var environmentsDeleteCmd = &cobra.Command{
	Use:   "delete [environment-key]",
	Short: "Delete an environment",
	Long:  `Delete an environment from a project.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runEnvironmentsDelete,
}

var envProject string
var envName string
var envKey string
var envDescription string
var envColor string
var envType string
var envForce bool

func init() {
	rootCmd.AddCommand(environmentsCmd)
	environmentsCmd.AddCommand(environmentsListCmd)
	environmentsCmd.AddCommand(environmentsGetCmd)
	environmentsCmd.AddCommand(environmentsCreateCmd)
	environmentsCmd.AddCommand(environmentsUpdateCmd)
	environmentsCmd.AddCommand(environmentsDeleteCmd)

	environmentsCmd.PersistentFlags().StringVarP(&envProject, "project", "p", "", "project key (uses config default if not specified)")

	// Create command flags
	environmentsCreateCmd.Flags().StringVarP(&envName, "name", "n", "", "environment name (required)")
	environmentsCreateCmd.Flags().StringVarP(&envKey, "key", "k", "", "environment key (required)")
	environmentsCreateCmd.Flags().StringVarP(&envDescription, "description", "d", "", "environment description")
	environmentsCreateCmd.Flags().StringVar(&envColor, "color", "", "environment color (e.g., #00ff00)")
	environmentsCreateCmd.Flags().StringVarP(&envType, "type", "t", "development", "environment type (development, staging, production)")

	// Update command flags
	environmentsUpdateCmd.Flags().StringVarP(&envName, "name", "n", "", "environment name")
	environmentsUpdateCmd.Flags().StringVarP(&envDescription, "description", "d", "", "environment description")
	environmentsUpdateCmd.Flags().StringVar(&envColor, "color", "", "environment color (e.g., #00ff00)")

	// Delete command flags
	environmentsDeleteCmd.Flags().BoolVarP(&envForce, "force", "f", false, "skip confirmation prompt")
}

type environmentsTableData struct {
	environments []api.Environment
}

func (d environmentsTableData) Headers() []string {
	return []string{"KEY", "NAME", "TYPE", "COLOR", "CREATED"}
}

func (d environmentsTableData) Rows() [][]string {
	rows := make([][]string, len(d.environments))
	for i, e := range d.environments {
		rows[i] = []string{
			e.Key,
			e.Name,
			e.Type,
			e.Color,
			e.CreatedAt.Format("2006-01-02"),
		}
	}
	return rows
}

func runEnvironmentsList(cmd *cobra.Command, args []string) error {
	projectKey := getEnvProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	environments, err := client.Environments(ctx, projectKey)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(environmentsTableData{environments: environments})
	}
	return printer.Print(environments)
}

func runEnvironmentsGet(cmd *cobra.Command, args []string) error {
	projectKey := getEnvProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	environment, err := client.Environment(ctx, projectKey, args[0])
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(environment)
}

func getEnvProjectKey() string {
	if envProject != "" {
		return envProject
	}
	return config.Project()
}

func runEnvironmentsCreate(cmd *cobra.Command, args []string) error {
	if envName == "" {
		return fmt.Errorf("required flag \"name\" not set")
	}
	if envKey == "" {
		return fmt.Errorf("required flag \"key\" not set")
	}

	projectKey := getEnvProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &api.CreateEnvironmentRequest{
		Name:        envName,
		Key:         envKey,
		Description: envDescription,
		Color:       envColor,
		Type:        envType,
	}

	environment, err := client.CreateEnvironment(ctx, projectKey, req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(environment)
}

func runEnvironmentsUpdate(cmd *cobra.Command, args []string) error {
	projectKey := getEnvProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &api.UpdateEnvironmentRequest{
		Name:        envName,
		Description: envDescription,
		Color:       envColor,
	}

	environment, err := client.UpdateEnvironment(ctx, projectKey, args[0], req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(environment)
}

func runEnvironmentsDelete(cmd *cobra.Command, args []string) error {
	projectKey := getEnvProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	environmentKey := args[0]
	if !confirmDelete("environment", environmentKey, envForce) {
		cmd.Println("Delete cancelled")
		return nil
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.DeleteEnvironment(ctx, projectKey, environmentKey); err != nil {
		return err
	}

	cmd.Printf("Environment '%s' deleted successfully\n", environmentKey)
	return nil
}
