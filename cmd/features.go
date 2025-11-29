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

var featureProject string

func init() {
	rootCmd.AddCommand(featuresCmd)
	featuresCmd.AddCommand(featuresListCmd)
	featuresCmd.AddCommand(featuresGetCmd)

	featuresCmd.PersistentFlags().StringVarP(&featureProject, "project", "p", "", "project key (uses config default if not specified)")
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
