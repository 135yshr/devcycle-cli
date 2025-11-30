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

var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Manage SDK keys",
	Long:  `List and rotate SDK keys for DevCycle environments.`,
}

var keysListCmd = &cobra.Command{
	Use:   "list",
	Short: "List SDK keys",
	Long:  `List all SDK keys for an environment.`,
	RunE:  runKeysList,
}

var keysRotateCmd = &cobra.Command{
	Use:   "rotate",
	Short: "Rotate an SDK key",
	Long:  `Rotate an SDK key for an environment. This will invalidate the old key.`,
	RunE:  runKeysRotate,
}

var keyProject string
var keyEnvironment string
var keyType string
var keyForce bool

func init() {
	rootCmd.AddCommand(keysCmd)
	keysCmd.AddCommand(keysListCmd)
	keysCmd.AddCommand(keysRotateCmd)

	keysCmd.PersistentFlags().StringVarP(&keyProject, "project", "p", "", "project key (uses config default if not specified)")
	keysCmd.PersistentFlags().StringVarP(&keyEnvironment, "environment", "e", "", "environment key (required)")

	// Rotate command flags
	keysRotateCmd.Flags().StringVarP(&keyType, "type", "t", "", "key type to rotate (client, server, mobile) (required)")
	keysRotateCmd.Flags().BoolVarP(&keyForce, "force", "f", false, "skip confirmation prompt")
}

type keysTableData struct {
	environment string
	keys        map[string]string
}

func (d keysTableData) Headers() []string {
	return []string{"ENVIRONMENT", "TYPE", "KEY"}
}

func (d keysTableData) Rows() [][]string {
	rows := [][]string{}
	for keyType, key := range d.keys {
		rows = append(rows, []string{
			d.environment,
			keyType,
			key,
		})
	}
	return rows
}

func runKeysList(cmd *cobra.Command, args []string) error {
	projectKey := getKeyProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	if keyEnvironment == "" {
		return fmt.Errorf("required flag \"environment\" not set")
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Get environment details which includes SDK keys
	environment, err := client.Environment(ctx, projectKey, keyEnvironment)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))

	if output.ParseFormat(GetOutput()) == output.FormatTable {
		// Extract SDK keys from environment
		keys := extractSDKKeys(environment)
		return printer.Print(keysTableData{environment: keyEnvironment, keys: keys})
	}

	return printer.Print(environment)
}

func extractSDKKeys(env *api.Environment) map[string]string {
	// Note: The actual SDK keys structure depends on the API response
	// This is a placeholder that may need adjustment based on the actual API
	return map[string]string{
		"client": "(SDK keys are included in environment details)",
		"server": "(SDK keys are included in environment details)",
		"mobile": "(SDK keys are included in environment details)",
	}
}

func runKeysRotate(cmd *cobra.Command, args []string) error {
	projectKey := getKeyProjectKey()
	if projectKey == "" {
		return errProjectRequired
	}

	if keyEnvironment == "" {
		return fmt.Errorf("required flag \"environment\" not set")
	}

	if keyType == "" {
		return fmt.Errorf("required flag \"type\" not set")
	}

	// Validate key type
	validTypes := map[string]bool{"client": true, "server": true, "mobile": true}
	if !validTypes[keyType] {
		return fmt.Errorf("invalid key type: %s (must be client, server, or mobile)", keyType)
	}

	if !keyForce {
		fmt.Printf("Are you sure you want to rotate the %s SDK key for environment '%s'?\n", keyType, keyEnvironment)
		fmt.Print("This will invalidate the existing key. [y/N]: ")
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "yes" {
			cmd.Println("Rotate cancelled")
			return nil
		}
	}

	client, err := getClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &api.RotateKeyRequest{
		Type: keyType,
	}

	result, err := client.RotateSDKKey(ctx, projectKey, keyEnvironment, req)
	if err != nil {
		return err
	}

	printer := output.NewPrinter(output.ParseFormat(GetOutput()))
	return printer.Print(result)
}

func getKeyProjectKey() string {
	if keyProject != "" {
		return keyProject
	}
	return config.Project()
}
