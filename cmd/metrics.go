package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/135yshr/devcycle-cli/internal/output"
	"github.com/135yshr/devcycle-cli/pkg/api"
	"github.com/spf13/cobra"
)

var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Manage metrics",
	Long:  `List, create, update, delete metrics and view metric results.`,
}

var metricsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all metrics",
	Long:  `List all metrics in a project.`,
	RunE:  runMetricsList,
}

var metricsGetCmd = &cobra.Command{
	Use:   "get [metric-key]",
	Short: "Get metric details",
	Long:  `Get detailed information about a specific metric.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runMetricsGet,
}

var metricsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new metric",
	Long:  `Create a new metric for experimentation.`,
	RunE:  runMetricsCreate,
}

var metricsUpdateCmd = &cobra.Command{
	Use:   "update [metric-key]",
	Short: "Update a metric",
	Long:  `Update an existing metric.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runMetricsUpdate,
}

var metricsDeleteCmd = &cobra.Command{
	Use:   "delete [metric-key]",
	Short: "Delete a metric",
	Long:  `Delete a metric from a project.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runMetricsDelete,
}

var metricsResultsCmd = &cobra.Command{
	Use:   "results [metric-key]",
	Short: "Get metric results",
	Long:  `Get results data for a specific metric.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runMetricsResults,
}

var metricProject string
var metricName string
var metricKey string
var metricType string
var metricEventType string
var metricOptimizeFor string
var metricDescription string
var metricForce bool

// Results options
var metricResultsEnvironment string
var metricResultsFeature string
var metricResultsStartDate string
var metricResultsEndDate string

func init() {
	rootCmd.AddCommand(metricsCmd)
	metricsCmd.AddCommand(metricsListCmd)
	metricsCmd.AddCommand(metricsGetCmd)
	metricsCmd.AddCommand(metricsCreateCmd)
	metricsCmd.AddCommand(metricsUpdateCmd)
	metricsCmd.AddCommand(metricsDeleteCmd)
	metricsCmd.AddCommand(metricsResultsCmd)

	// Persistent flags for all metrics commands
	metricsCmd.PersistentFlags().StringVarP(&metricProject, "project", "p", "", "project key (uses config default if not specified)")

	// Create command flags
	metricsCreateCmd.Flags().StringVarP(&metricName, "name", "n", "", "metric name (required)")
	metricsCreateCmd.Flags().StringVarP(&metricKey, "key", "k", "", "metric key (required)")
	metricsCreateCmd.Flags().StringVarP(&metricType, "type", "t", "", "metric type: count, countPerEval, sum, average (required)")
	metricsCreateCmd.Flags().StringVar(&metricEventType, "event-type", "", "event type to track (required)")
	metricsCreateCmd.Flags().StringVar(&metricOptimizeFor, "optimize-for", "", "optimization direction: increase, decrease (required)")
	metricsCreateCmd.Flags().StringVar(&metricDescription, "description", "", "metric description")

	// Update command flags
	metricsUpdateCmd.Flags().StringVarP(&metricName, "name", "n", "", "metric name")
	metricsUpdateCmd.Flags().StringVarP(&metricType, "type", "t", "", "metric type: count, countPerEval, sum, average")
	metricsUpdateCmd.Flags().StringVar(&metricEventType, "event-type", "", "event type to track")
	metricsUpdateCmd.Flags().StringVar(&metricOptimizeFor, "optimize-for", "", "optimization direction: increase, decrease")
	metricsUpdateCmd.Flags().StringVar(&metricDescription, "description", "", "metric description")

	// Delete command flags
	metricsDeleteCmd.Flags().BoolVar(&metricForce, "force", false, "skip confirmation prompt")

	// Results command flags
	metricsResultsCmd.Flags().StringVarP(&metricResultsEnvironment, "environment", "e", "", "filter by environment")
	metricsResultsCmd.Flags().StringVarP(&metricResultsFeature, "feature", "f", "", "filter by feature")
	metricsResultsCmd.Flags().StringVar(&metricResultsStartDate, "start-date", "", "start date (ISO 8601 format)")
	metricsResultsCmd.Flags().StringVar(&metricResultsEndDate, "end-date", "", "end date (ISO 8601 format)")
}

type metricsTableData struct {
	metrics []api.Metric
}

func (d metricsTableData) Headers() []string {
	return []string{"KEY", "NAME", "TYPE", "EVENT TYPE", "OPTIMIZE FOR", "DESCRIPTION"}
}

func (d metricsTableData) Rows() [][]string {
	rows := make([][]string, len(d.metrics))
	for i, m := range d.metrics {
		desc := m.Description
		if len(desc) > 30 {
			desc = desc[:27] + "..."
		}
		if desc == "" {
			desc = "-"
		}

		rows[i] = []string{
			m.Key,
			m.Name,
			m.Type,
			m.EventType,
			m.OptimizeFor,
			desc,
		}
	}
	return rows
}

type metricResultsTableData struct {
	results *api.MetricResults
}

func (d metricResultsTableData) Headers() []string {
	return []string{"VARIATION", "COUNT", "VALUE"}
}

func (d metricResultsTableData) Rows() [][]string {
	rows := make([][]string, len(d.results.Data))
	for i, r := range d.results.Data {
		rows[i] = []string{
			r.VariationKey,
			fmt.Sprintf("%d", r.Count),
			fmt.Sprintf("%.4f", r.Value),
		}
	}
	return rows
}

func getMetricProjectKey() string {
	if metricProject != "" {
		return metricProject
	}
	return getProjectKey()
}

func runMetricsList(cmd *cobra.Command, args []string) error {
	projectKey := getMetricProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	metrics, err := client.Metrics(ctx, projectKey)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(metricsTableData{metrics: metrics})
	}
	return printer.Print(metrics)
}

func runMetricsGet(cmd *cobra.Command, args []string) error {
	projectKey := getMetricProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	metric, err := client.Metric(ctx, projectKey, args[0])
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(metric)
}

func runMetricsCreate(cmd *cobra.Command, args []string) error {
	projectKey := getMetricProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	// Validate required fields
	if metricName == "" {
		return fmt.Errorf("required flag \"name\" not set")
	}
	if metricKey == "" {
		return fmt.Errorf("required flag \"key\" not set")
	}
	if metricType == "" {
		return fmt.Errorf("required flag \"type\" not set")
	}
	if metricEventType == "" {
		return fmt.Errorf("required flag \"event-type\" not set")
	}
	if metricOptimizeFor == "" {
		return fmt.Errorf("required flag \"optimize-for\" not set")
	}

	req := &api.CreateMetricRequest{
		Name:        metricName,
		Key:         metricKey,
		Type:        metricType,
		EventType:   metricEventType,
		OptimizeFor: metricOptimizeFor,
		Description: metricDescription,
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	metric, err := client.CreateMetric(ctx, projectKey, req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(metric)
}

func runMetricsUpdate(cmd *cobra.Command, args []string) error {
	projectKey := getMetricProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	req := &api.UpdateMetricRequest{}

	// Set fields only if provided
	if metricName != "" {
		req.Name = metricName
	}
	if metricType != "" {
		req.Type = metricType
	}
	if metricEventType != "" {
		req.EventType = metricEventType
	}
	if metricOptimizeFor != "" {
		req.OptimizeFor = metricOptimizeFor
	}
	if metricDescription != "" {
		req.Description = metricDescription
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	metric, err := client.UpdateMetric(ctx, projectKey, args[0], req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(metric)
}

func runMetricsDelete(cmd *cobra.Command, args []string) error {
	projectKey := getMetricProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	metricKeyArg := args[0]
	if !confirmDelete("metric", metricKeyArg, metricForce) {
		cmd.Println("Delete cancelled")
		return nil
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.DeleteMetric(ctx, projectKey, metricKeyArg); err != nil {
		return err
	}

	cmd.Printf("Metric '%s' deleted successfully\n", metricKeyArg)
	return nil
}

func runMetricsResults(cmd *cobra.Command, args []string) error {
	projectKey := getMetricProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	opts := &api.MetricResultsOptions{
		Environment: metricResultsEnvironment,
		Feature:     metricResultsFeature,
		StartDate:   metricResultsStartDate,
		EndDate:     metricResultsEndDate,
	}

	results, err := client.MetricResults(ctx, projectKey, args[0], opts)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(metricResultsTableData{results: results})
	}
	return printer.Print(results)
}
