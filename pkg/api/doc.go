// Package api provides a Go client for the DevCycle Management API.
//
// The api package enables programmatic access to DevCycle feature flag
// management functionality, including projects, environments, features,
// variables, variations, and targeting configurations.
//
// # Getting Started
//
// Create a client and authenticate:
//
//	token, err := api.Authenticate(ctx, clientID, clientSecret)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	client := api.NewClient(api.WithToken(token.AccessToken))
//
// # Working with Projects
//
// List all projects:
//
//	projects, err := client.Projects(ctx)
//
// Get a specific project:
//
//	project, err := client.Project(ctx, "my-project-key")
//
// # Working with Features
//
// List features for a project:
//
//	features, err := client.Features(ctx, "my-project-key")
//
// Get a specific feature:
//
//	feature, err := client.Feature(ctx, "my-project-key", "my-feature-key")
//
// # Working with Variables
//
// List variables for a project:
//
//	variables, err := client.Variables(ctx, "my-project-key")
//
// # Working with Environments
//
// List environments for a project:
//
//	environments, err := client.Environments(ctx, "my-project-key")
//
// # Error Handling
//
// API errors are returned as [APIError] which includes the HTTP status code
// and error message. Helper functions are provided for common error checks:
//
//	_, err := client.Project(ctx, "nonexistent")
//	if api.IsNotFound(err) {
//	    // Handle 404 Not Found
//	}
//
// # API Reference
//
// For complete DevCycle API documentation, see https://docs.devcycle.com/management-api/
package api
