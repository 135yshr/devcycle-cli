package api

import (
	"context"
	"fmt"
)

// Override represents a self-targeting override
type Override struct {
	Feature     string         `json:"feature,omitempty"`
	Environment string         `json:"environment,omitempty"`
	Variation   string         `json:"variation,omitempty"`
	Variables   map[string]any `json:"variables,omitempty"`
}

// SetOverrideRequest represents the request body for setting an override
type SetOverrideRequest struct {
	Environment string `json:"environment"`
	Variation   string `json:"variation"`
}

// FeatureOverrides returns all overrides for a specific feature
func (c *Client) FeatureOverrides(ctx context.Context, projectKey, featureKey string) ([]Override, error) {
	var overrides []Override
	path := fmt.Sprintf("/projects/%s/features/%s/overrides", projectKey, featureKey)
	if err := c.Get(ctx, path, &overrides); err != nil {
		return nil, fmt.Errorf("failed to list feature overrides: %w", err)
	}
	return overrides, nil
}

// CurrentOverride returns the current user's override for a specific feature
func (c *Client) CurrentOverride(ctx context.Context, projectKey, featureKey string) (*Override, error) {
	var override Override
	path := fmt.Sprintf("/projects/%s/features/%s/overrides/current", projectKey, featureKey)
	if err := c.Get(ctx, path, &override); err != nil {
		return nil, fmt.Errorf("failed to get current override: %w", err)
	}
	return &override, nil
}

// SetOverride creates or updates the current user's override for a feature
func (c *Client) SetOverride(ctx context.Context, projectKey, featureKey string, req *SetOverrideRequest) (*Override, error) {
	var override Override
	path := fmt.Sprintf("/projects/%s/features/%s/overrides/current", projectKey, featureKey)
	if err := c.Put(ctx, path, req, &override); err != nil {
		return nil, fmt.Errorf("failed to set override: %w", err)
	}
	return &override, nil
}

// DeleteOverride deletes the current user's override for a feature in a specific environment
func (c *Client) DeleteOverride(ctx context.Context, projectKey, featureKey, environment string) error {
	path := fmt.Sprintf("/projects/%s/features/%s/overrides/current?environment=%s", projectKey, featureKey, environment)
	if err := c.Delete(ctx, path); err != nil {
		return fmt.Errorf("failed to delete override: %w", err)
	}
	return nil
}

// MyOverrides returns all overrides for the current user in a project
func (c *Client) MyOverrides(ctx context.Context, projectKey string) ([]Override, error) {
	var overrides []Override
	path := fmt.Sprintf("/projects/%s/overrides/current", projectKey)
	if err := c.Get(ctx, path, &overrides); err != nil {
		return nil, fmt.Errorf("failed to list my overrides: %w", err)
	}
	return overrides, nil
}

// DeleteAllMyOverrides deletes all overrides for the current user in a project
func (c *Client) DeleteAllMyOverrides(ctx context.Context, projectKey string) error {
	path := fmt.Sprintf("/projects/%s/overrides/current", projectKey)
	if err := c.Delete(ctx, path); err != nil {
		return fmt.Errorf("failed to delete all my overrides: %w", err)
	}
	return nil
}
