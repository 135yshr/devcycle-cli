package cmd

import (
	"context"
	"time"

	"github.com/135yshr/devcycle-cli/internal/api"
	"github.com/135yshr/devcycle-cli/internal/config"
	"github.com/135yshr/devcycle-cli/internal/output"
	"github.com/spf13/cobra"
)

var variablesCmd = &cobra.Command{
	Use:   "variables",
	Short: "Manage DevCycle variables",
	Long:  `List and view DevCycle variables.`,
}

var variablesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all variables",
	Long:  `List all variables in a project.`,
	RunE:  runVariablesList,
}

var variablesGetCmd = &cobra.Command{
	Use:   "get [variable-key]",
	Short: "Get variable details",
	Long:  `Get detailed information about a specific variable.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runVariablesGet,
}

var variableProject string

func init() {
	rootCmd.AddCommand(variablesCmd)
	variablesCmd.AddCommand(variablesListCmd)
	variablesCmd.AddCommand(variablesGetCmd)

	variablesCmd.PersistentFlags().StringVarP(&variableProject, "project", "p", "", "project key (uses config default if not specified)")
}

type variablesTableData struct {
	variables []api.Variable
}

func (d variablesTableData) Headers() []string {
	return []string{"KEY", "NAME", "TYPE", "STATUS", "CREATED"}
}

func (d variablesTableData) Rows() [][]string {
	rows := make([][]string, len(d.variables))
	for i, v := range d.variables {
		rows[i] = []string{
			v.Key,
			v.Name,
			v.Type,
			v.Status,
			v.CreatedAt.Format("2006-01-02"),
		}
	}
	return rows
}

func runVariablesList(cmd *cobra.Command, args []string) error {
	projectKey := getVariableProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	variables, err := client.ListVariables(ctx, projectKey)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(variablesTableData{variables: variables})
	}
	return printer.Print(variables)
}

func runVariablesGet(cmd *cobra.Command, args []string) error {
	projectKey := getVariableProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	variable, err := client.GetVariable(ctx, projectKey, args[0])
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(variable)
}

func getVariableProjectKey() string {
	if variableProject != "" {
		return variableProject
	}
	return config.GetProject()
}
