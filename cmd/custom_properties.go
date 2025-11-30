package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/135yshr/devcycle-cli/internal/output"
	"github.com/135yshr/devcycle-cli/pkg/api"
	"github.com/spf13/cobra"
)

var customPropertiesCmd = &cobra.Command{
	Use:     "custom-properties",
	Aliases: []string{"cp"},
	Short:   "Manage custom properties",
	Long:    `List, create, update, and delete custom properties for targeting.`,
}

var customPropertiesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all custom properties",
	Long:  `List all custom properties in a project.`,
	RunE:  runCustomPropertiesList,
}

var customPropertiesGetCmd = &cobra.Command{
	Use:   "get [property-key]",
	Short: "Get custom property details",
	Long:  `Get detailed information about a specific custom property.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runCustomPropertiesGet,
}

var customPropertiesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new custom property",
	Long:  `Create a new custom property for targeting rules.`,
	RunE:  runCustomPropertiesCreate,
}

var customPropertiesUpdateCmd = &cobra.Command{
	Use:   "update [property-key]",
	Short: "Update a custom property",
	Long:  `Update an existing custom property.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runCustomPropertiesUpdate,
}

var customPropertiesDeleteCmd = &cobra.Command{
	Use:   "delete [property-key]",
	Short: "Delete a custom property",
	Long:  `Delete a custom property from a project.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runCustomPropertiesDelete,
}

var customPropertyProject string
var customPropertyKey string
var customPropertyDisplayName string
var customPropertyType string
var customPropertyDescription string
var customPropertyForce bool

func init() {
	rootCmd.AddCommand(customPropertiesCmd)
	customPropertiesCmd.AddCommand(customPropertiesListCmd)
	customPropertiesCmd.AddCommand(customPropertiesGetCmd)
	customPropertiesCmd.AddCommand(customPropertiesCreateCmd)
	customPropertiesCmd.AddCommand(customPropertiesUpdateCmd)
	customPropertiesCmd.AddCommand(customPropertiesDeleteCmd)

	// Persistent flags for all custom-properties commands
	customPropertiesCmd.PersistentFlags().StringVarP(&customPropertyProject, "project", "p", "", "project key (uses config default if not specified)")

	// Create command flags
	customPropertiesCreateCmd.Flags().StringVarP(&customPropertyKey, "key", "k", "", "property key (required, must match SDK property name)")
	customPropertiesCreateCmd.Flags().StringVar(&customPropertyDisplayName, "display-name", "", "display name (required)")
	customPropertiesCreateCmd.Flags().StringVarP(&customPropertyType, "type", "t", "", "property type: Boolean, Number, String (required)")
	customPropertiesCreateCmd.Flags().StringVar(&customPropertyDescription, "description", "", "property description")

	// Update command flags
	customPropertiesUpdateCmd.Flags().StringVar(&customPropertyDisplayName, "display-name", "", "display name")
	customPropertiesUpdateCmd.Flags().StringVar(&customPropertyDescription, "description", "", "property description")

	// Delete command flags
	customPropertiesDeleteCmd.Flags().BoolVar(&customPropertyForce, "force", false, "skip confirmation prompt")
}

type customPropertiesTableData struct {
	properties []api.CustomProperty
}

func (d customPropertiesTableData) Headers() []string {
	return []string{"KEY", "DISPLAY NAME", "TYPE", "DESCRIPTION", "CREATED AT"}
}

func (d customPropertiesTableData) Rows() [][]string {
	rows := make([][]string, len(d.properties))
	for i, p := range d.properties {
		desc := p.Description
		if len(desc) > 30 {
			desc = desc[:27] + "..."
		}
		if desc == "" {
			desc = "-"
		}

		rows[i] = []string{
			p.Key,
			p.DisplayName,
			p.Type,
			desc,
			p.CreatedAt.Format(time.RFC3339),
		}
	}
	return rows
}

func getCustomPropertyProjectKey() string {
	if customPropertyProject != "" {
		return customPropertyProject
	}
	return getProjectKey()
}

func runCustomPropertiesList(cmd *cobra.Command, args []string) error {
	projectKey := getCustomPropertyProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	properties, err := client.CustomProperties(ctx, projectKey)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(customPropertiesTableData{properties: properties})
	}
	return printer.Print(properties)
}

func runCustomPropertiesGet(cmd *cobra.Command, args []string) error {
	projectKey := getCustomPropertyProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	property, err := client.CustomProperty(ctx, projectKey, args[0])
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(property)
}

func runCustomPropertiesCreate(cmd *cobra.Command, args []string) error {
	projectKey := getCustomPropertyProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	// Validate required fields
	if customPropertyKey == "" {
		return fmt.Errorf("required flag \"key\" not set")
	}
	if customPropertyDisplayName == "" {
		return fmt.Errorf("required flag \"display-name\" not set")
	}
	if customPropertyType == "" {
		return fmt.Errorf("required flag \"type\" not set")
	}

	// Validate type
	validTypes := map[string]bool{"Boolean": true, "Number": true, "String": true}
	if !validTypes[customPropertyType] {
		return fmt.Errorf("invalid type: %s (must be Boolean, Number, or String)", customPropertyType)
	}

	req := &api.CreateCustomPropertyRequest{
		Key:         customPropertyKey,
		DisplayName: customPropertyDisplayName,
		Type:        customPropertyType,
		Description: customPropertyDescription,
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	property, err := client.CreateCustomProperty(ctx, projectKey, req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(property)
}

func runCustomPropertiesUpdate(cmd *cobra.Command, args []string) error {
	projectKey := getCustomPropertyProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	req := &api.UpdateCustomPropertyRequest{}

	// Set fields only if provided
	if customPropertyDisplayName != "" {
		req.DisplayName = customPropertyDisplayName
	}
	if customPropertyDescription != "" {
		req.Description = customPropertyDescription
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	property, err := client.UpdateCustomProperty(ctx, projectKey, args[0], req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(property)
}

func runCustomPropertiesDelete(cmd *cobra.Command, args []string) error {
	projectKey := getCustomPropertyProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	propertyKey := args[0]
	if !confirmDelete("custom property", propertyKey, customPropertyForce) {
		cmd.Println("Delete cancelled")
		return nil
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.DeleteCustomProperty(ctx, projectKey, propertyKey); err != nil {
		return err
	}

	cmd.Printf("Custom property '%s' deleted successfully\n", propertyKey)
	return nil
}
