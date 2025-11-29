package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/135yshr/devcycle-cli/internal/api"
	"github.com/135yshr/devcycle-cli/internal/output"
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

func init() {
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(projectsListCmd)
	projectsCmd.AddCommand(projectsGetCmd)
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

	projects, err := client.ListProjects(ctx)
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

	project, err := client.GetProject(ctx, args[0])
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
