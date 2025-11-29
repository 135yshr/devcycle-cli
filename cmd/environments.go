package cmd

import (
	"context"
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

var envProject string

func init() {
	rootCmd.AddCommand(environmentsCmd)
	environmentsCmd.AddCommand(environmentsListCmd)
	environmentsCmd.AddCommand(environmentsGetCmd)

	environmentsCmd.PersistentFlags().StringVarP(&envProject, "project", "p", "", "project key (uses config default if not specified)")
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

	environments, err := client.ListEnvironments(ctx, projectKey)
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

	environment, err := client.GetEnvironment(ctx, projectKey, args[0])
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
	return config.GetProject()
}
