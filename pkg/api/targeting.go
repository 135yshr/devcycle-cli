package api

import (
	"context"
	"fmt"
	"net/url"
)

// FeatureConfigurations returns the targeting configurations for a feature
func (c *Client) FeatureConfigurations(ctx context.Context, projectKey, featureKey string) (map[string]*EnvironmentConfig, error) {
	var configs map[string]*EnvironmentConfig
	path := fmt.Sprintf("/projects/%s/features/%s/configurations", url.PathEscape(projectKey), url.PathEscape(featureKey))
	if err := c.Get(ctx, path, &configs); err != nil {
		return nil, fmt.Errorf("failed to get feature configurations: %w", err)
	}
	return configs, nil
}

// UpdateFeatureConfigurationsRequest represents the request body for updating targeting configurations.
type UpdateFeatureConfigurationsRequest struct {
	Configurations map[string]*EnvironmentConfig `json:"configurations"`
}

// UpdateFeatureConfigurations updates the targeting configurations for a feature
func (c *Client) UpdateFeatureConfigurations(ctx context.Context, projectKey, featureKey string, req *UpdateFeatureConfigurationsRequest) (map[string]*EnvironmentConfig, error) {
	var configs map[string]*EnvironmentConfig
	path := fmt.Sprintf("/projects/%s/features/%s/configurations", url.PathEscape(projectKey), url.PathEscape(featureKey))
	if err := c.Patch(ctx, path, req, &configs); err != nil {
		return nil, fmt.Errorf("failed to update feature configurations: %w", err)
	}
	return configs, nil
}

// EnableFeature enables a feature for a specific environment
func (c *Client) EnableFeature(ctx context.Context, projectKey, featureKey, environmentKey string) error {
	req := &UpdateFeatureConfigurationsRequest{
		Configurations: map[string]*EnvironmentConfig{
			environmentKey: {
				Status: "active",
			},
		},
	}
	_, err := c.UpdateFeatureConfigurations(ctx, projectKey, featureKey, req)
	return err
}

// DisableFeature disables a feature for a specific environment
func (c *Client) DisableFeature(ctx context.Context, projectKey, featureKey, environmentKey string) error {
	req := &UpdateFeatureConfigurationsRequest{
		Configurations: map[string]*EnvironmentConfig{
			environmentKey: {
				Status: "inactive",
			},
		},
	}
	_, err := c.UpdateFeatureConfigurations(ctx, projectKey, featureKey, req)
	return err
}
