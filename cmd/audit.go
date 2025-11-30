package cmd

import (
	"context"
	"time"

	"github.com/135yshr/devcycle-cli/internal/output"
	"github.com/135yshr/devcycle-cli/pkg/api"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "View audit logs",
	Long:  `View audit logs for a project or specific feature.`,
}

var auditListCmd = &cobra.Command{
	Use:   "list",
	Short: "List project audit logs",
	Long:  `List all audit logs for a project.`,
	RunE:  runAuditList,
}

var auditFeatureCmd = &cobra.Command{
	Use:   "feature [feature-key]",
	Short: "List feature audit logs",
	Long:  `List audit logs for a specific feature.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runAuditFeature,
}

var auditProject string

func init() {
	rootCmd.AddCommand(auditCmd)
	auditCmd.AddCommand(auditListCmd)
	auditCmd.AddCommand(auditFeatureCmd)

	// Persistent flags for all audit commands
	auditCmd.PersistentFlags().StringVarP(&auditProject, "project", "p", "", "project key (uses config default if not specified)")
}

type auditTableData struct {
	logs []api.AuditLog
}

func (d auditTableData) Headers() []string {
	return []string{"ID", "TYPE", "USER", "EMAIL", "CHANGES", "CREATED AT"}
}

func (d auditTableData) Rows() [][]string {
	rows := make([][]string, len(d.logs))
	for i, log := range d.logs {
		changesStr := "-"
		if len(log.Changes) > 0 {
			changesStr = log.Changes[0].Type
			if len(log.Changes) > 1 {
				changesStr += " ..."
			}
		}

		rows[i] = []string{
			log.ID,
			log.Type,
			log.User.Name,
			log.User.Email,
			changesStr,
			log.CreatedAt.Format(time.RFC3339),
		}
	}
	return rows
}

func getAuditProjectKey() string {
	if auditProject != "" {
		return auditProject
	}
	return getProjectKey()
}

func runAuditList(cmd *cobra.Command, args []string) error {
	projectKey := getAuditProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logs, err := client.AuditLogs(ctx, projectKey)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(auditTableData{logs: logs})
	}
	return printer.Print(logs)
}

func runAuditFeature(cmd *cobra.Command, args []string) error {
	projectKey := getAuditProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logs, err := client.FeatureAuditLogs(ctx, projectKey, args[0])
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(auditTableData{logs: logs})
	}
	return printer.Print(logs)
}
