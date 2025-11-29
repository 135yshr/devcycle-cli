package cmd

import (
	"context"
	"time"

	"github.com/135yshr/devcycle-cli/internal/config"
	"github.com/135yshr/devcycle-cli/internal/output"
	"github.com/135yshr/devcycle-cli/pkg/api"
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

var variablesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new variable",
	Long:  `Create a new variable in a project.`,
	RunE:  runVariablesCreate,
}

var variablesUpdateCmd = &cobra.Command{
	Use:   "update [variable-key]",
	Short: "Update a variable",
	Long:  `Update an existing variable.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runVariablesUpdate,
}

var variablesDeleteCmd = &cobra.Command{
	Use:   "delete [variable-key]",
	Short: "Delete a variable",
	Long:  `Delete a variable from a project.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runVariablesDelete,
}

var variableProject string
var variableName string
var variableKey string
var variableDescription string
var variableType string
var variableFeature string
var variableForce bool

func init() {
	rootCmd.AddCommand(variablesCmd)
	variablesCmd.AddCommand(variablesListCmd)
	variablesCmd.AddCommand(variablesGetCmd)
	variablesCmd.AddCommand(variablesCreateCmd)
	variablesCmd.AddCommand(variablesUpdateCmd)
	variablesCmd.AddCommand(variablesDeleteCmd)

	variablesCmd.PersistentFlags().StringVarP(&variableProject, "project", "p", "", "project key (uses config default if not specified)")

	// Create command flags
	variablesCreateCmd.Flags().StringVarP(&variableName, "name", "n", "", "variable name (required)")
	variablesCreateCmd.Flags().StringVarP(&variableKey, "key", "k", "", "variable key (required)")
	variablesCreateCmd.Flags().StringVarP(&variableDescription, "description", "d", "", "variable description")
	variablesCreateCmd.Flags().StringVarP(&variableType, "type", "t", "Boolean", "variable type (String, Boolean, Number, JSON)")
	variablesCreateCmd.Flags().StringVar(&variableFeature, "feature", "", "associated feature key")
	variablesCreateCmd.MarkFlagRequired("name")
	variablesCreateCmd.MarkFlagRequired("key")

	// Update command flags
	variablesUpdateCmd.Flags().StringVarP(&variableName, "name", "n", "", "variable name")
	variablesUpdateCmd.Flags().StringVarP(&variableDescription, "description", "d", "", "variable description")

	// Delete command flags
	variablesDeleteCmd.Flags().BoolVarP(&variableForce, "force", "f", false, "skip confirmation prompt")
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

	variables, err := client.Variables(ctx, projectKey)
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

	variable, err := client.Variable(ctx, projectKey, args[0])
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
	return config.Project()
}

func runVariablesCreate(cmd *cobra.Command, args []string) error {
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

	req := &api.CreateVariableRequest{
		Name:        variableName,
		Key:         variableKey,
		Description: variableDescription,
		Type:        variableType,
		Feature:     variableFeature,
	}

	variable, err := client.CreateVariable(ctx, projectKey, req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(variable)
}

func runVariablesUpdate(cmd *cobra.Command, args []string) error {
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

	req := &api.UpdateVariableRequest{
		Name:        variableName,
		Description: variableDescription,
	}

	variable, err := client.UpdateVariable(ctx, projectKey, args[0], req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(variable)
}

func runVariablesDelete(cmd *cobra.Command, args []string) error {
	projectKey := getVariableProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	varKey := args[0]
	if !confirmDelete("variable", varKey, variableForce) {
		cmd.Println("Delete cancelled")
		return nil
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.DeleteVariable(ctx, projectKey, varKey); err != nil {
		return err
	}

	cmd.Printf("Variable '%s' deleted successfully\n", varKey)
	return nil
}
