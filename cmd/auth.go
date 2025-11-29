package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/135yshr/devcycle-cli/internal/api"
	"github.com/135yshr/devcycle-cli/internal/config"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication commands",
	Long:  `Manage authentication with the DevCycle Management API.`,
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with DevCycle",
	Long: `Authenticate with DevCycle using OAuth2 client credentials.

You can provide credentials via:
  1. Command flags: --client-id and --client-secret
  2. Environment variables: DVCX_CLIENT_ID and DVCX_CLIENT_SECRET
  3. Config file: .devcycle/config.yaml`,
	RunE: runLogin,
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored authentication",
	Long:  `Remove the stored access token from the local configuration.`,
	RunE:  runLogout,
}

var (
	clientID     string
	clientSecret string
)

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(loginCmd)
	authCmd.AddCommand(logoutCmd)

	loginCmd.Flags().StringVar(&clientID, "client-id", "", "DevCycle API client ID")
	loginCmd.Flags().StringVar(&clientSecret, "client-secret", "", "DevCycle API client secret")
}

func runLogin(cmd *cobra.Command, args []string) error {
	id := clientID
	if id == "" {
		id = config.GetClientID()
	}
	if id == "" {
		return fmt.Errorf("client ID is required (use --client-id flag, DVCX_CLIENT_ID env var, or config file)")
	}

	secret := clientSecret
	if secret == "" {
		secret = config.GetClientSecret()
	}
	if secret == "" {
		return fmt.Errorf("client secret is required (use --client-secret flag, DVCX_CLIENT_SECRET env var, or config file)")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("Authenticating with DevCycle...")

	token, err := api.Authenticate(ctx, id, secret)
	if err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}

	if err := config.EnsureConfigDir(); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	tokenPath, err := config.GetTokenPath()
	if err != nil {
		return fmt.Errorf("failed to get token path: %w", err)
	}

	if err := api.SaveToken(token, tokenPath); err != nil {
		return fmt.Errorf("failed to save token: %w", err)
	}

	fmt.Println("Successfully authenticated!")
	fmt.Printf("Token expires at: %s\n", token.ExpiresAt.Format(time.RFC3339))

	return nil
}

func runLogout(cmd *cobra.Command, args []string) error {
	tokenPath, err := config.GetTokenPath()
	if err != nil {
		return fmt.Errorf("failed to get token path: %w", err)
	}

	if _, err := os.Stat(tokenPath); os.IsNotExist(err) {
		fmt.Println("No authentication token found.")
		return nil
	}

	if err := api.RemoveFile(tokenPath); err != nil {
		return fmt.Errorf("failed to remove token: %w", err)
	}

	fmt.Println("Successfully logged out.")
	return nil
}
