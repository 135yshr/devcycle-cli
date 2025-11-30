package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/135yshr/devcycle-cli/internal/output"
	"github.com/135yshr/devcycle-cli/pkg/api"
	"github.com/spf13/cobra"
)

// maxAudienceFileSize is the maximum file size allowed for JSON audience files (10MB)
const maxAudienceFileSize = 10 * 1024 * 1024

var audiencesCmd = &cobra.Command{
	Use:   "audiences",
	Short: "Manage audiences",
	Long:  `List, create, update, and delete reusable audiences.`,
}

var audiencesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all audiences",
	Long:  `List all audiences in a project.`,
	RunE:  runAudiencesList,
}

var audiencesGetCmd = &cobra.Command{
	Use:   "get [audience-key]",
	Short: "Get audience details",
	Long:  `Get detailed information about a specific audience.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runAudiencesGet,
}

var audiencesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new audience",
	Long:  `Create a new reusable audience.`,
	RunE:  runAudiencesCreate,
}

var audiencesUpdateCmd = &cobra.Command{
	Use:   "update [audience-key]",
	Short: "Update an audience",
	Long:  `Update an existing audience.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runAudiencesUpdate,
}

var audiencesDeleteCmd = &cobra.Command{
	Use:   "delete [audience-key]",
	Short: "Delete an audience",
	Long:  `Delete an audience from a project.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runAudiencesDelete,
}

var audienceProject string
var audienceName string
var audienceKey string
var audienceDescription string
var audienceFilters string
var audienceFromFile string
var audienceForce bool

func init() {
	rootCmd.AddCommand(audiencesCmd)
	audiencesCmd.AddCommand(audiencesListCmd)
	audiencesCmd.AddCommand(audiencesGetCmd)
	audiencesCmd.AddCommand(audiencesCreateCmd)
	audiencesCmd.AddCommand(audiencesUpdateCmd)
	audiencesCmd.AddCommand(audiencesDeleteCmd)

	// Persistent flags for all audiences commands
	audiencesCmd.PersistentFlags().StringVarP(&audienceProject, "project", "p", "", "project key (uses config default if not specified)")

	// Create command flags
	audiencesCreateCmd.Flags().StringVarP(&audienceName, "name", "n", "", "audience name (required)")
	audiencesCreateCmd.Flags().StringVarP(&audienceKey, "key", "k", "", "audience key (required)")
	audiencesCreateCmd.Flags().StringVar(&audienceDescription, "description", "", "audience description")
	audiencesCreateCmd.Flags().StringVar(&audienceFilters, "filters", "", "audience filters as JSON")
	audiencesCreateCmd.Flags().StringVarP(&audienceFromFile, "from-file", "F", "", "JSON file containing audience definition")

	// Update command flags
	audiencesUpdateCmd.Flags().StringVarP(&audienceName, "name", "n", "", "audience name")
	audiencesUpdateCmd.Flags().StringVar(&audienceDescription, "description", "", "audience description")
	audiencesUpdateCmd.Flags().StringVar(&audienceFilters, "filters", "", "audience filters as JSON")
	audiencesUpdateCmd.Flags().StringVarP(&audienceFromFile, "from-file", "F", "", "JSON file containing audience definition")

	// Delete command flags
	audiencesDeleteCmd.Flags().BoolVar(&audienceForce, "force", false, "skip confirmation prompt")
}

type audiencesTableData struct {
	audiences []api.AudienceDefinition
}

func (d audiencesTableData) Headers() []string {
	return []string{"KEY", "NAME", "DESCRIPTION", "FILTERS"}
}

func (d audiencesTableData) Rows() [][]string {
	rows := make([][]string, len(d.audiences))
	for i, a := range d.audiences {
		desc := a.Description
		if len(desc) > 30 {
			desc = desc[:27] + "..."
		}
		if desc == "" {
			desc = "-"
		}

		filtersStr := "-"
		if len(a.Filters.Filters) > 0 {
			filtersStr = fmt.Sprintf("%d filter(s)", len(a.Filters.Filters))
		}

		rows[i] = []string{
			a.Key,
			a.Name,
			desc,
			filtersStr,
		}
	}
	return rows
}

func getAudienceProjectKey() string {
	if audienceProject != "" {
		return audienceProject
	}
	return getProjectKey()
}

func runAudiencesList(cmd *cobra.Command, args []string) error {
	projectKey := getAudienceProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	audiences, err := client.Audiences(ctx, projectKey)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(audiencesTableData{audiences: audiences})
	}
	return printer.Print(audiences)
}

func runAudiencesGet(cmd *cobra.Command, args []string) error {
	projectKey := getAudienceProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	audience, err := client.Audience(ctx, projectKey, args[0])
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(audience)
}

func runAudiencesCreate(cmd *cobra.Command, args []string) error {
	projectKey := getAudienceProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	var req api.CreateAudienceRequest

	// Load from file if specified
	if audienceFromFile != "" {
		file, err := os.Open(audienceFromFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat file: %w", err)
		}
		if info.Size() > maxAudienceFileSize {
			return fmt.Errorf("file %s exceeds maximum allowed size (%d bytes)", audienceFromFile, maxAudienceFileSize)
		}
		data, err := io.ReadAll(file)
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}
		if err := json.Unmarshal(data, &req); err != nil {
			return fmt.Errorf("invalid JSON in file: %w", err)
		}
	}

	// Override with flags if provided
	if audienceName != "" {
		req.Name = audienceName
	}
	if audienceKey != "" {
		req.Key = audienceKey
	}
	if audienceDescription != "" {
		req.Description = audienceDescription
	}
	if audienceFilters != "" {
		var filters api.Filters
		if err := json.Unmarshal([]byte(audienceFilters), &filters); err != nil {
			return fmt.Errorf("invalid JSON for filters: %w", err)
		}
		req.Filters = filters
	}

	// Validate required fields
	if req.Name == "" {
		return fmt.Errorf("required flag \"name\" not set")
	}
	if req.Key == "" {
		return fmt.Errorf("required flag \"key\" not set")
	}
	if len(req.Filters.Filters) == 0 {
		// Default to "all users" filter if not specified
		req.Filters = api.Filters{
			Operator: "and",
			Filters:  []api.Filter{{Type: "all"}},
		}
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	audience, err := client.CreateAudience(ctx, projectKey, &req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(audience)
}

func runAudiencesUpdate(cmd *cobra.Command, args []string) error {
	projectKey := getAudienceProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	var req api.UpdateAudienceRequest

	// Load from file if specified
	if audienceFromFile != "" {
		file, err := os.Open(audienceFromFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat file: %w", err)
		}
		if info.Size() > maxAudienceFileSize {
			return fmt.Errorf("file %s exceeds maximum allowed size (%d bytes)", audienceFromFile, maxAudienceFileSize)
		}
		data, err := io.ReadAll(file)
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}
		if err := json.Unmarshal(data, &req); err != nil {
			return fmt.Errorf("invalid JSON in file: %w", err)
		}
	}

	// Override with flags if provided
	if audienceName != "" {
		req.Name = audienceName
	}
	if audienceDescription != "" {
		req.Description = audienceDescription
	}
	if audienceFilters != "" {
		var filters api.Filters
		if err := json.Unmarshal([]byte(audienceFilters), &filters); err != nil {
			return fmt.Errorf("invalid JSON for filters: %w", err)
		}
		req.Filters = &filters
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	audience, err := client.UpdateAudience(ctx, projectKey, args[0], &req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(audience)
}

func runAudiencesDelete(cmd *cobra.Command, args []string) error {
	projectKey := getAudienceProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	audienceKeyArg := args[0]
	if !confirmDelete("audience", audienceKeyArg, audienceForce) {
		cmd.Println("Delete cancelled")
		return nil
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.DeleteAudience(ctx, projectKey, audienceKeyArg); err != nil {
		return err
	}

	cmd.Printf("Audience '%s' deleted successfully\n", audienceKeyArg)
	return nil
}
