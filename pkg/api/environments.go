package api

import (
	"context"
	"fmt"
	"net/url"
)

// Environments returns all environments for a project.
func (c *Client) Environments(ctx context.Context, projectKey string) ([]Environment, error) {
	var environments []Environment
	path := fmt.Sprintf("/projects/%s/environments", url.PathEscape(projectKey))
	if err := c.Get(ctx, path, &environments); err != nil {
		return nil, fmt.Errorf("failed to list environments: %w", err)
	}
	return environments, nil
}

// Environment returns a specific environment by its key.
func (c *Client) Environment(ctx context.Context, projectKey, environmentKey string) (*Environment, error) {
	var environment Environment
	path := fmt.Sprintf("/projects/%s/environments/%s", url.PathEscape(projectKey), url.PathEscape(environmentKey))
	if err := c.Get(ctx, path, &environment); err != nil {
		return nil, fmt.Errorf("failed to get environment: %w", err)
	}
	return &environment, nil
}

// CreateEnvironment creates a new environment in a project.
func (c *Client) CreateEnvironment(ctx context.Context, projectKey string, req *CreateEnvironmentRequest) (*Environment, error) {
	var environment Environment
	path := fmt.Sprintf("/projects/%s/environments", url.PathEscape(projectKey))
	if err := c.Post(ctx, path, req, &environment); err != nil {
		return nil, fmt.Errorf("failed to create environment: %w", err)
	}
	return &environment, nil
}

// UpdateEnvironment updates an existing environment's properties.
func (c *Client) UpdateEnvironment(ctx context.Context, projectKey, environmentKey string, req *UpdateEnvironmentRequest) (*Environment, error) {
	var environment Environment
	path := fmt.Sprintf("/projects/%s/environments/%s", url.PathEscape(projectKey), url.PathEscape(environmentKey))
	if err := c.Patch(ctx, path, req, &environment); err != nil {
		return nil, fmt.Errorf("failed to update environment: %w", err)
	}
	return &environment, nil
}

// DeleteEnvironment removes an environment from a project.
// Warning: This action cannot be undone.
func (c *Client) DeleteEnvironment(ctx context.Context, projectKey, environmentKey string) error {
	path := fmt.Sprintf("/projects/%s/environments/%s", url.PathEscape(projectKey), url.PathEscape(environmentKey))
	if err := c.Delete(ctx, path); err != nil {
		return fmt.Errorf("failed to delete environment: %w", err)
	}
	return nil
}

// RotateSDKKey rotates an SDK key for an environment.
// Returns both the previous key and the newly generated key.
func (c *Client) RotateSDKKey(ctx context.Context, projectKey, environmentKey string, req *RotateKeyRequest) (*RotateKeyResponse, error) {
	var response RotateKeyResponse
	path := fmt.Sprintf("/projects/%s/environments/%s/keys", url.PathEscape(projectKey), url.PathEscape(environmentKey))
	if err := c.Post(ctx, path, req, &response); err != nil {
		return nil, fmt.Errorf("failed to rotate SDK key: %w", err)
	}
	return &response, nil
}
