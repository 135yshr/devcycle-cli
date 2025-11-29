package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/135yshr/devcycle-cli/internal/output"
	"github.com/135yshr/devcycle-cli/pkg/api"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Manage DevCycle projects",
	Long:  `List and view DevCycle projects.`,
}

var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	Long:  `List all projects accessible with your credentials.`,
	RunE:  runProjectsList,
}

var projectsGetCmd = &cobra.Command{
	Use:   "get [project-key]",
	Short: "Get project details",
	Long:  `Get detailed information about a specific project.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runProjectsGet,
}

var projectsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Long:  `Create a new DevCycle project.`,
	RunE:  runProjectsCreate,
}

var projectsUpdateCmd = &cobra.Command{
	Use:   "update [project-key]",
	Short: "Update a project",
	Long:  `Update an existing project.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runProjectsUpdate,
}

var projectName string
var projectKey string
var projectDescription string

func init() {
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(projectsListCmd)
	projectsCmd.AddCommand(projectsGetCmd)
	projectsCmd.AddCommand(projectsCreateCmd)
	projectsCmd.AddCommand(projectsUpdateCmd)

	// Create command flags
	projectsCreateCmd.Flags().StringVarP(&projectName, "name", "n", "", "project name (required)")
	projectsCreateCmd.Flags().StringVarP(&projectKey, "key", "k", "", "project key (required)")
	projectsCreateCmd.Flags().StringVarP(&projectDescription, "description", "d", "", "project description")
	projectsCreateCmd.MarkFlagRequired("name")
	projectsCreateCmd.MarkFlagRequired("key")

	// Update command flags
	projectsUpdateCmd.Flags().StringVarP(&projectName, "name", "n", "", "project name")
	projectsUpdateCmd.Flags().StringVarP(&projectDescription, "description", "d", "", "project description")
}

type projectsTableData struct {
	projects []api.Project
}

func (d projectsTableData) Headers() []string {
	return []string{"KEY", "NAME", "DESCRIPTION", "CREATED"}
}

func (d projectsTableData) Rows() [][]string {
	rows := make([][]string, len(d.projects))
	for i, p := range d.projects {
		rows[i] = []string{
			p.Key,
			p.Name,
			truncate(p.Description, 40),
			p.CreatedAt.Format("2006-01-02"),
		}
	}
	return rows
}

func runProjectsList(cmd *cobra.Command, args []string) error {
	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	projects, err := client.Projects(ctx)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(projectsTableData{projects: projects})
	}
	return printer.Print(projects)
}

func runProjectsGet(cmd *cobra.Command, args []string) error {
	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	project, err := client.Project(ctx, args[0])
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(project)
}

func getClient() (*api.Client, error) {
	token, err := loadToken()
	if err != nil {
		return nil, fmt.Errorf("not authenticated. Run 'dvcx auth login' first")
	}

	if token.IsExpired() {
		return nil, fmt.Errorf("token expired. Run 'dvcx auth login' to refresh")
	}

	return api.NewClient(api.WithToken(token.AccessToken)), nil
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func runProjectsCreate(cmd *cobra.Command, args []string) error {
	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &api.CreateProjectRequest{
		Name:        projectName,
		Key:         projectKey,
		Description: projectDescription,
	}

	project, err := client.CreateProject(ctx, req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(project)
}

func runProjectsUpdate(cmd *cobra.Command, args []string) error {
	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &api.UpdateProjectRequest{
		Name:        projectName,
		Description: projectDescription,
	}

	project, err := client.UpdateProject(ctx, args[0], req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(project)
}
