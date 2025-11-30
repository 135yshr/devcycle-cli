package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/135yshr/devcycle-cli/internal/output"
	"github.com/135yshr/devcycle-cli/pkg/api"
	"github.com/spf13/cobra"
)

var overridesCmd = &cobra.Command{
	Use:   "overrides",
	Short: "Manage self-targeting overrides",
	Long:  `List, create, update, and delete self-targeting overrides for features.`,
}

var overridesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all overrides for a feature",
	Long:  `List all self-targeting overrides for a specific feature.`,
	RunE:  runOverridesList,
}

var overridesGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get current user's override for a feature",
	Long:  `Get the current user's self-targeting override for a specific feature.`,
	RunE:  runOverridesGet,
}

var overridesSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set override for current user",
	Long:  `Create or update a self-targeting override for the current user.`,
	RunE:  runOverridesSet,
}

var overridesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete override for current user",
	Long:  `Delete a self-targeting override for the current user in a specific environment.`,
	RunE:  runOverridesDelete,
}

var overridesListMineCmd = &cobra.Command{
	Use:   "list-mine",
	Short: "List all my overrides in project",
	Long:  `List all self-targeting overrides for the current user across all features in a project.`,
	RunE:  runOverridesListMine,
}

var overridesDeleteMineCmd = &cobra.Command{
	Use:   "delete-mine",
	Short: "Delete all my overrides in project",
	Long:  `Delete all self-targeting overrides for the current user in a project.`,
	RunE:  runOverridesDeleteMine,
}

var overrideProject string
var overrideFeature string
var overrideEnvironment string
var overrideVariation string
var overrideForce bool

func init() {
	rootCmd.AddCommand(overridesCmd)
	overridesCmd.AddCommand(overridesListCmd)
	overridesCmd.AddCommand(overridesGetCmd)
	overridesCmd.AddCommand(overridesSetCmd)
	overridesCmd.AddCommand(overridesDeleteCmd)
	overridesCmd.AddCommand(overridesListMineCmd)
	overridesCmd.AddCommand(overridesDeleteMineCmd)

	// Persistent flags for all overrides commands
	overridesCmd.PersistentFlags().StringVarP(&overrideProject, "project", "p", "", "project key (uses config default if not specified)")

	// Feature-scoped command flags
	overridesListCmd.Flags().StringVarP(&overrideFeature, "feature", "f", "", "feature key (required)")
	overridesGetCmd.Flags().StringVarP(&overrideFeature, "feature", "f", "", "feature key (required)")
	overridesSetCmd.Flags().StringVarP(&overrideFeature, "feature", "f", "", "feature key (required)")
	overridesDeleteCmd.Flags().StringVarP(&overrideFeature, "feature", "f", "", "feature key (required)")

	// Set command flags
	overridesSetCmd.Flags().StringVarP(&overrideEnvironment, "environment", "e", "", "environment key (required)")
	overridesSetCmd.Flags().StringVarP(&overrideVariation, "variation", "v", "", "variation key (required)")

	// Delete command flags
	overridesDeleteCmd.Flags().StringVarP(&overrideEnvironment, "environment", "e", "", "environment key (required)")

	// Delete-mine command flags
	overridesDeleteMineCmd.Flags().BoolVar(&overrideForce, "force", false, "skip confirmation prompt")
}

type overridesTableData struct {
	overrides []api.Override
}

func (d overridesTableData) Headers() []string {
	return []string{"FEATURE", "ENVIRONMENT", "VARIATION"}
}

func (d overridesTableData) Rows() [][]string {
	rows := make([][]string, len(d.overrides))
	for i, o := range d.overrides {
		feature := o.Feature
		if feature == "" {
			feature = "-"
		}
		env := o.Environment
		if env == "" {
			env = "-"
		}
		variation := o.Variation
		if variation == "" {
			variation = "-"
		}
		rows[i] = []string{
			feature,
			env,
			variation,
		}
	}
	return rows
}

func getOverrideProjectKey() string {
	if overrideProject != "" {
		return overrideProject
	}
	return getProjectKey()
}

func runOverridesList(cmd *cobra.Command, args []string) error {
	projectKey := getOverrideProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}
	if overrideFeature == "" {
		return fmt.Errorf("required flag \"feature\" not set")
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	overrides, err := client.FeatureOverrides(ctx, projectKey, overrideFeature)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(overridesTableData{overrides: overrides})
	}
	return printer.Print(overrides)
}

func runOverridesGet(cmd *cobra.Command, args []string) error {
	projectKey := getOverrideProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}
	if overrideFeature == "" {
		return fmt.Errorf("required flag \"feature\" not set")
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	override, err := client.CurrentOverride(ctx, projectKey, overrideFeature)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(override)
}

func runOverridesSet(cmd *cobra.Command, args []string) error {
	projectKey := getOverrideProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}
	if overrideFeature == "" {
		return fmt.Errorf("required flag \"feature\" not set")
	}
	if overrideEnvironment == "" {
		return fmt.Errorf("required flag \"environment\" not set")
	}
	if overrideVariation == "" {
		return fmt.Errorf("required flag \"variation\" not set")
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &api.SetOverrideRequest{
		Environment: overrideEnvironment,
		Variation:   overrideVariation,
	}

	override, err := client.SetOverride(ctx, projectKey, overrideFeature, req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(override)
}

func runOverridesDelete(cmd *cobra.Command, args []string) error {
	projectKey := getOverrideProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}
	if overrideFeature == "" {
		return fmt.Errorf("required flag \"feature\" not set")
	}
	if overrideEnvironment == "" {
		return fmt.Errorf("required flag \"environment\" not set")
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.DeleteOverride(ctx, projectKey, overrideFeature, overrideEnvironment); err != nil {
		return err
	}

	cmd.Printf("Override for feature '%s' in environment '%s' deleted successfully\n", overrideFeature, overrideEnvironment)
	return nil
}

func runOverridesListMine(cmd *cobra.Command, args []string) error {
	projectKey := getOverrideProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	overrides, err := client.MyOverrides(ctx, projectKey)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(overridesTableData{overrides: overrides})
	}
	return printer.Print(overrides)
}

func runOverridesDeleteMine(cmd *cobra.Command, args []string) error {
	projectKey := getOverrideProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	if !overrideForce {
		fmt.Print("Are you sure you want to delete all your overrides in this project? [y/N]: ")
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" && response != "yes" && response != "Yes" {
			cmd.Println("Delete cancelled")
			return nil
		}
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.DeleteAllMyOverrides(ctx, projectKey); err != nil {
		return err
	}

	cmd.Println("All overrides deleted successfully")
	return nil
}
