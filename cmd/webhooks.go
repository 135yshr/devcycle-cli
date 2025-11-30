package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/135yshr/devcycle-cli/internal/output"
	"github.com/135yshr/devcycle-cli/pkg/api"
	"github.com/spf13/cobra"
)

var webhooksCmd = &cobra.Command{
	Use:   "webhooks",
	Short: "Manage webhooks",
	Long:  `List, create, update, and delete webhooks.`,
}

var webhooksListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all webhooks",
	Long:  `List all webhooks in a project.`,
	RunE:  runWebhooksList,
}

var webhooksGetCmd = &cobra.Command{
	Use:   "get [webhook-id]",
	Short: "Get webhook details",
	Long:  `Get detailed information about a specific webhook.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runWebhooksGet,
}

var webhooksCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new webhook",
	Long:  `Create a new webhook to receive event notifications.`,
	RunE:  runWebhooksCreate,
}

var webhooksUpdateCmd = &cobra.Command{
	Use:   "update [webhook-id]",
	Short: "Update a webhook",
	Long:  `Update an existing webhook.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runWebhooksUpdate,
}

var webhooksDeleteCmd = &cobra.Command{
	Use:   "delete [webhook-id]",
	Short: "Delete a webhook",
	Long:  `Delete a webhook from a project.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runWebhooksDelete,
}

var webhookProject string
var webhookURL string
var webhookDescription string
var webhookEnabled bool
var webhookDisabled bool
var webhookForce bool

func init() {
	rootCmd.AddCommand(webhooksCmd)
	webhooksCmd.AddCommand(webhooksListCmd)
	webhooksCmd.AddCommand(webhooksGetCmd)
	webhooksCmd.AddCommand(webhooksCreateCmd)
	webhooksCmd.AddCommand(webhooksUpdateCmd)
	webhooksCmd.AddCommand(webhooksDeleteCmd)

	// Persistent flags for all webhooks commands
	webhooksCmd.PersistentFlags().StringVarP(&webhookProject, "project", "p", "", "project key (uses config default if not specified)")

	// Create command flags
	webhooksCreateCmd.Flags().StringVar(&webhookURL, "url", "", "webhook URL (required)")
	webhooksCreateCmd.Flags().StringVar(&webhookDescription, "description", "", "webhook description")
	webhooksCreateCmd.Flags().BoolVar(&webhookEnabled, "enabled", true, "enable the webhook (default: true)")

	// Update command flags
	webhooksUpdateCmd.Flags().StringVar(&webhookURL, "url", "", "webhook URL")
	webhooksUpdateCmd.Flags().StringVar(&webhookDescription, "description", "", "webhook description")
	webhooksUpdateCmd.Flags().BoolVar(&webhookEnabled, "enabled", false, "enable the webhook")
	webhooksUpdateCmd.Flags().BoolVar(&webhookDisabled, "disabled", false, "disable the webhook")

	// Delete command flags
	webhooksDeleteCmd.Flags().BoolVar(&webhookForce, "force", false, "skip confirmation prompt")
}

type webhooksTableData struct {
	webhooks []api.Webhook
}

func (d webhooksTableData) Headers() []string {
	return []string{"ID", "URL", "DESCRIPTION", "ENABLED", "CREATED AT"}
}

func (d webhooksTableData) Rows() [][]string {
	rows := make([][]string, len(d.webhooks))
	for i, w := range d.webhooks {
		desc := w.Description
		if len(desc) > 30 {
			desc = desc[:27] + "..."
		}
		if desc == "" {
			desc = "-"
		}

		url := w.URL
		if len(url) > 40 {
			url = url[:37] + "..."
		}

		enabled := "No"
		if w.IsEnabled {
			enabled = "Yes"
		}

		rows[i] = []string{
			w.ID,
			url,
			desc,
			enabled,
			w.CreatedAt.Format(time.RFC3339),
		}
	}
	return rows
}

func getWebhookProjectKey() string {
	if webhookProject != "" {
		return webhookProject
	}
	return getProjectKey()
}

func runWebhooksList(cmd *cobra.Command, args []string) error {
	projectKey := getWebhookProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	webhooks, err := client.Webhooks(ctx, projectKey)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		return printer.Print(webhooksTableData{webhooks: webhooks})
	}
	return printer.Print(webhooks)
}

func runWebhooksGet(cmd *cobra.Command, args []string) error {
	projectKey := getWebhookProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	webhook, err := client.Webhook(ctx, projectKey, args[0])
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(webhook)
}

func runWebhooksCreate(cmd *cobra.Command, args []string) error {
	projectKey := getWebhookProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	// Validate required fields
	if webhookURL == "" {
		return fmt.Errorf("required flag \"url\" not set")
	}

	req := &api.CreateWebhookRequest{
		URL:         webhookURL,
		Description: webhookDescription,
		IsEnabled:   webhookEnabled,
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	webhook, err := client.CreateWebhook(ctx, projectKey, req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(webhook)
}

func runWebhooksUpdate(cmd *cobra.Command, args []string) error {
	projectKey := getWebhookProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	req := &api.UpdateWebhookRequest{}

	// Set fields only if provided
	if webhookURL != "" {
		req.URL = webhookURL
	}
	if webhookDescription != "" {
		req.Description = webhookDescription
	}

	// Handle enabled/disabled flags
	if webhookEnabled && webhookDisabled {
		return fmt.Errorf("cannot set both --enabled and --disabled flags")
	}
	if cmd.Flags().Changed("enabled") {
		req.IsEnabled = &webhookEnabled
	}
	if webhookDisabled {
		disabled := false
		req.IsEnabled = &disabled
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	webhook, err := client.UpdateWebhook(ctx, projectKey, args[0], req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(webhook)
}

func runWebhooksDelete(cmd *cobra.Command, args []string) error {
	projectKey := getWebhookProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	webhookID := args[0]
	if !confirmDelete("webhook", webhookID, webhookForce) {
		cmd.Println("Delete cancelled")
		return nil
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.DeleteWebhook(ctx, projectKey, webhookID); err != nil {
		return err
	}

	cmd.Printf("Webhook '%s' deleted successfully\n", webhookID)
	return nil
}
