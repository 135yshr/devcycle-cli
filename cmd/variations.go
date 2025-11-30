package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/135yshr/devcycle-cli/internal/output"
	"github.com/135yshr/devcycle-cli/pkg/api"
	"github.com/spf13/cobra"
)

var variationsCmd = &cobra.Command{
	Use:   "variations",
	Short: "Manage feature variations",
	Long:  `List, create, update, and delete feature variations.`,
}

var variationsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all variations for a feature",
	Long:  `List all variations for a specific feature.`,
	RunE:  runVariationsList,
}

var variationsGetCmd = &cobra.Command{
	Use:   "get [variation-key]",
	Short: "Get variation details",
	Long:  `Get detailed information about a specific variation.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runVariationsGet,
}

var variationsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new variation",
	Long:  `Create a new variation for a feature.`,
	RunE:  runVariationsCreate,
}

var variationsUpdateCmd = &cobra.Command{
	Use:   "update [variation-key]",
	Short: "Update a variation",
	Long:  `Update an existing variation.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runVariationsUpdate,
}

var variationsDeleteCmd = &cobra.Command{
	Use:   "delete [variation-key]",
	Short: "Delete a variation",
	Long:  `Delete a variation from a feature.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runVariationsDelete,
}

var variationProject string
var variationFeature string
var variationName string
var variationKey string
var variationVariables string
var variationForce bool

func init() {
	rootCmd.AddCommand(variationsCmd)
	variationsCmd.AddCommand(variationsListCmd)
	variationsCmd.AddCommand(variationsGetCmd)
	variationsCmd.AddCommand(variationsCreateCmd)
	variationsCmd.AddCommand(variationsUpdateCmd)
	variationsCmd.AddCommand(variationsDeleteCmd)

	// Persistent flags for all variations commands
	variationsCmd.PersistentFlags().StringVarP(&variationProject, "project", "p", "", "project key (uses config default if not specified)")
	variationsCmd.PersistentFlags().StringVarP(&variationFeature, "feature", "f", "", "feature key (required)")

	// Create command flags
	variationsCreateCmd.Flags().StringVarP(&variationName, "name", "n", "", "variation name (required)")
	variationsCreateCmd.Flags().StringVarP(&variationKey, "key", "k", "", "variation key (required)")
	variationsCreateCmd.Flags().StringVarP(&variationVariables, "variables", "v", "", "variation variables as JSON (e.g., '{\"enabled\": true}')")

	// Update command flags
	variationsUpdateCmd.Flags().StringVarP(&variationName, "name", "n", "", "variation name")
	variationsUpdateCmd.Flags().StringVarP(&variationVariables, "variables", "v", "", "variation variables as JSON (e.g., '{\"enabled\": true}')")

	// Delete command flags
	variationsDeleteCmd.Flags().BoolVar(&variationForce, "force", false, "skip confirmation prompt")
}

type variationsTableData struct {
	variations []api.Variation
}

func (d variationsTableData) Headers() []string {
	return []string{"KEY", "NAME", "VARIABLES"}
}

func (d variationsTableData) Rows() [][]string {
	rows := make([][]string, len(d.variations))
	for i, v := range d.variations {
		varsStr := "-"
		if len(v.Variables) > 0 {
			varsJSON, _ := json.Marshal(v.Variables)
			varsStr = string(varsJSON)
			if len(varsStr) > 50 {
				varsStr = varsStr[:47] + "..."
			}
		}
		rows[i] = []string{
			v.Key,
			v.Name,
			varsStr,
		}
	}
	return rows
}

func getVariationProjectKey() string {
	if variationProject != "" {
		return variationProject
	}
	return getProjectKey()
}

func runVariationsList(cmd *cobra.Command, args []string) error {
	projectKey := getVariationProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}
	if variationFeature == "" {
		return fmt.Errorf("required flag \"feature\" not set")
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	variations, err := client.Variations(ctx, projectKey, variationFeature)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(variationsTableData{variations: variations})
	}
	return printer.Print(variations)
}

func runVariationsGet(cmd *cobra.Command, args []string) error {
	projectKey := getVariationProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}
	if variationFeature == "" {
		return fmt.Errorf("required flag \"feature\" not set")
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	variation, err := client.Variation(ctx, projectKey, variationFeature, args[0])
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(variation)
}

func runVariationsCreate(cmd *cobra.Command, args []string) error {
	if variationName == "" {
		return fmt.Errorf("required flag \"name\" not set")
	}
	if variationKey == "" {
		return fmt.Errorf("required flag \"key\" not set")
	}

	projectKey := getVariationProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}
	if variationFeature == "" {
		return fmt.Errorf("required flag \"feature\" not set")
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &api.CreateVariationRequest{
		Name: variationName,
		Key:  variationKey,
	}

	if variationVariables != "" {
		var vars map[string]any
		if err := json.Unmarshal([]byte(variationVariables), &vars); err != nil {
			return fmt.Errorf("invalid JSON for variables: %w", err)
		}
		req.Variables = vars
	}

	variation, err := client.CreateVariation(ctx, projectKey, variationFeature, req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(variation)
}

func runVariationsUpdate(cmd *cobra.Command, args []string) error {
	projectKey := getVariationProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}
	if variationFeature == "" {
		return fmt.Errorf("required flag \"feature\" not set")
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &api.UpdateVariationRequest{
		Name: variationName,
	}

	if variationVariables != "" {
		var vars map[string]any
		if err := json.Unmarshal([]byte(variationVariables), &vars); err != nil {
			return fmt.Errorf("invalid JSON for variables: %w", err)
		}
		req.Variables = vars
	}

	variation, err := client.UpdateVariation(ctx, projectKey, variationFeature, args[0], req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(variation)
}

func runVariationsDelete(cmd *cobra.Command, args []string) error {
	projectKey := getVariationProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}
	if variationFeature == "" {
		return fmt.Errorf("required flag \"feature\" not set")
	}

	variationKeyArg := args[0]
	if !confirmDelete("variation", variationKeyArg, variationForce) {
		cmd.Println("Delete cancelled")
		return nil
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.DeleteVariation(ctx, projectKey, variationFeature, variationKeyArg); err != nil {
		return err
	}

	cmd.Printf("Variation '%s' deleted successfully\n", variationKeyArg)
	return nil
}
