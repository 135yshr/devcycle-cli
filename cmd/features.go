package cmd

import (
	"context"
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

func init() {
	rootCmd.AddCommand(featuresCmd)
	featuresCmd.AddCommand(featuresListCmd)
	featuresCmd.AddCommand(featuresGetCmd)
	featuresCmd.AddCommand(featuresCreateCmd)
	featuresCmd.AddCommand(featuresUpdateCmd)
	featuresCmd.AddCommand(featuresDeleteCmd)

	featuresCmd.PersistentFlags().StringVarP(&featureProject, "project", "p", "", "project key (uses config default if not specified)")

	// Create command flags
	featuresCreateCmd.Flags().StringVarP(&featureName, "name", "n", "", "feature name (required)")
	featuresCreateCmd.Flags().StringVarP(&featureKey, "key", "k", "", "feature key (required)")
	featuresCreateCmd.Flags().StringVarP(&featureDescription, "description", "d", "", "feature description")
	featuresCreateCmd.Flags().StringVarP(&featureType, "type", "t", "release", "feature type (release, experiment, permission, ops)")
	featuresCreateCmd.MarkFlagRequired("name")
	featuresCreateCmd.MarkFlagRequired("key")

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
