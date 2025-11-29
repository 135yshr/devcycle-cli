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

var featuresCmd = &cobra.Command{
	Use:   "features",
	Short: "Manage DevCycle features",
	Long:  `List and view DevCycle features.`,
}

var featuresListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all features",
	Long:  `List all features in a project.`,
	RunE:  runFeaturesList,
}

var featuresGetCmd = &cobra.Command{
	Use:   "get [feature-key]",
	Short: "Get feature details",
	Long:  `Get detailed information about a specific feature.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runFeaturesGet,
}

var featuresCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new feature",
	Long:  `Create a new feature in a project.`,
	RunE:  runFeaturesCreate,
}

var featuresUpdateCmd = &cobra.Command{
	Use:   "update [feature-key]",
	Short: "Update a feature",
	Long:  `Update an existing feature.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runFeaturesUpdate,
}

var featuresDeleteCmd = &cobra.Command{
	Use:   "delete [feature-key]",
	Short: "Delete a feature",
	Long:  `Delete a feature from a project.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runFeaturesDelete,
}

var featureProject string
var featureName string
var featureKey string
var featureDescription string
var featureType string
var featureForce bool
var featureFromFile string
var featureDryRun bool

func init() {
	rootCmd.AddCommand(featuresCmd)
	featuresCmd.AddCommand(featuresListCmd)
	featuresCmd.AddCommand(featuresGetCmd)
	featuresCmd.AddCommand(featuresCreateCmd)
	featuresCmd.AddCommand(featuresUpdateCmd)
	featuresCmd.AddCommand(featuresDeleteCmd)

	featuresCmd.PersistentFlags().StringVarP(&featureProject, "project", "p", "", "project key (uses config default if not specified)")

	// Create command flags
	featuresCreateCmd.Flags().StringVarP(&featureName, "name", "n", "", "feature name (required for simple create)")
	featuresCreateCmd.Flags().StringVarP(&featureKey, "key", "k", "", "feature key (required for simple create)")
	featuresCreateCmd.Flags().StringVarP(&featureDescription, "description", "d", "", "feature description")
	featuresCreateCmd.Flags().StringVarP(&featureType, "type", "t", "release", "feature type (release, experiment, permission, ops)")
	featuresCreateCmd.Flags().StringVarP(&featureFromFile, "from-file", "F", "", "JSON input file for feature creation (uses v2 API), use '-' for stdin")
	featuresCreateCmd.Flags().BoolVar(&featureDryRun, "dry-run", false, "validate configuration without creating")

	// Update command flags
	featuresUpdateCmd.Flags().StringVarP(&featureName, "name", "n", "", "feature name")
	featuresUpdateCmd.Flags().StringVarP(&featureDescription, "description", "d", "", "feature description")

	// Delete command flags
	featuresDeleteCmd.Flags().BoolVarP(&featureForce, "force", "f", false, "skip confirmation prompt")
}

type featuresTableData struct {
	features []api.Feature
}

func (d featuresTableData) Headers() []string {
	return []string{"KEY", "NAME", "TYPE", "STATUS", "CREATED"}
}

func (d featuresTableData) Rows() [][]string {
	rows := make([][]string, len(d.features))
	for i, f := range d.features {
		rows[i] = []string{
			f.Key,
			f.Name,
			f.Type,
			f.Status,
			f.CreatedAt.Format("2006-01-02"),
		}
	}
	return rows
}

func runFeaturesList(cmd *cobra.Command, args []string) error {
	projectKey := getProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	features, err := client.Features(ctx, projectKey)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(featuresTableData{features: features})
	}
	return printer.Print(features)
}

func runFeaturesGet(cmd *cobra.Command, args []string) error {
	projectKey := getProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	feature, err := client.Feature(ctx, projectKey, args[0])
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(feature)
}

func getProjectKey() string {
	if featureProject != "" {
		return featureProject
	}
	return config.Project()
}

func runFeaturesCreate(cmd *cobra.Command, args []string) error {
	// Use v2 API when --from-file is specified
	if featureFromFile != "" {
		return runFeaturesCreateV2(cmd, args)
	}

	// Validate required flags for simple create
	if featureName == "" {
		return fmt.Errorf("required flag \"name\" not set")
	}
	if featureKey == "" {
		return fmt.Errorf("required flag \"key\" not set")
	}

	projectKey := getProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &api.CreateFeatureRequest{
		Name:        featureName,
		Key:         featureKey,
		Description: featureDescription,
		Type:        featureType,
	}

	feature, err := client.CreateFeature(ctx, projectKey, req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(feature)
}

func runFeaturesCreateV2(cmd *cobra.Command, args []string) error {
	projectKey := getProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	// Load feature request from file
	req, err := api.LoadFeatureRequestFromFile(featureFromFile)
	if err != nil {
		return err
	}

	// Validate the request
	if err := api.ValidateFeatureRequest(req); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	// Dry-run mode: validate and show preview without creating
	if featureDryRun {
		return printDryRunPreview(cmd, req)
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	feature, err := client.CreateFeatureV2(ctx, projectKey, req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(feature)
}

func printDryRunPreview(cmd *cobra.Command, req *api.CreateFeatureV2Request) error {
	cmd.Println("Validating feature configuration...")
	cmd.Println()
	cmd.Println("[OK] JSON syntax valid")
	cmd.Println("[OK] Schema validation passed")
	cmd.Println()
	cmd.Println("Feature Preview:")
	cmd.Println("----------------")
	cmd.Printf("Name:        %s\n", req.Name)
	cmd.Printf("Key:         %s\n", req.Key)
	cmd.Printf("Type:        %s\n", req.Type)
	if req.Description != "" {
		cmd.Printf("Description: %s\n", req.Description)
	}

	if len(req.Variables) > 0 {
		cmd.Println()
		cmd.Printf("Variables (%d):\n", len(req.Variables))
		for _, v := range req.Variables {
			cmd.Printf("  - %s (%s)\n", v.Key, v.Type)
		}
	}

	if len(req.Variations) > 0 {
		cmd.Println()
		cmd.Printf("Variations (%d):\n", len(req.Variations))
		for _, v := range req.Variations {
			if len(v.Variables) > 0 {
				cmd.Printf("  - %s: ", v.Key)
				first := true
				for key, val := range v.Variables {
					if !first {
						cmd.Print(", ")
					}
					cmd.Printf("%s = %v", key, val)
					first = false
				}
				cmd.Println()
			} else {
				cmd.Printf("  - %s\n", v.Key)
			}
		}
	}

	if len(req.Configurations) > 0 {
		cmd.Println()
		cmd.Println("Environment Configurations:")
		for envKey, config := range req.Configurations {
			if config == nil {
				continue
			}
			status := "Inactive"
			if config.Status == "active" {
				status = "Active"
			}
			cmd.Printf("  %s: %s\n", envKey, status)
		}
	}

	cmd.Println()
	cmd.Println("Dry-run complete. No changes were made.")
	return nil
}

func runFeaturesUpdate(cmd *cobra.Command, args []string) error {
	projectKey := getProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &api.UpdateFeatureRequest{
		Name:        featureName,
		Description: featureDescription,
	}

	feature, err := client.UpdateFeature(ctx, projectKey, args[0], req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(feature)
}

func runFeaturesDelete(cmd *cobra.Command, args []string) error {
	projectKey := getProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	featureKey := args[0]
	if !confirmDelete("feature", featureKey, featureForce) {
		cmd.Println("Delete cancelled")
		return nil
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.DeleteFeature(ctx, projectKey, featureKey); err != nil {
		return err
	}

	cmd.Printf("Feature '%s' deleted successfully\n", featureKey)
	return nil
}
