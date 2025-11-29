package api

import (
	"context"
	"fmt"
)

func (c *Client) Features(ctx context.Context, projectKey string) ([]Feature, error) {
	var features []Feature
	path := fmt.Sprintf("/projects/%s/features", projectKey)
	if err := c.Get(ctx, path, &features); err != nil {
		return nil, fmt.Errorf("failed to list features: %w", err)
	}
	return features, nil
}

func (c *Client) Feature(ctx context.Context, projectKey, featureKey string) (*Feature, error) {
	var feature Feature
	path := fmt.Sprintf("/projects/%s/features/%s", projectKey, featureKey)
	if err := c.Get(ctx, path, &feature); err != nil {
		return nil, fmt.Errorf("failed to get feature: %w", err)
	}
	return &feature, nil
}

// CreateFeatureRequest represents the request body for creating a feature.
type CreateFeatureRequest struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
}

// UpdateFeatureRequest represents the request body for updating a feature.
type UpdateFeatureRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c *Client) CreateFeature(ctx context.Context, projectKey string, req *CreateFeatureRequest) (*Feature, error) {
	var feature Feature
	path := fmt.Sprintf("/projects/%s/features", projectKey)
	if err := c.Post(ctx, path, req, &feature); err != nil {
		return nil, fmt.Errorf("failed to create feature: %w", err)
	}
	return &feature, nil
}

func (c *Client) UpdateFeature(ctx context.Context, projectKey, featureKey string, req *UpdateFeatureRequest) (*Feature, error) {
	var feature Feature
	path := fmt.Sprintf("/projects/%s/features/%s", projectKey, featureKey)
	if err := c.Patch(ctx, path, req, &feature); err != nil {
		return nil, fmt.Errorf("failed to update feature: %w", err)
	}
	return &feature, nil
}

func (c *Client) DeleteFeature(ctx context.Context, projectKey, featureKey string) error {
	path := fmt.Sprintf("/projects/%s/features/%s", projectKey, featureKey)
	if err := c.Delete(ctx, path); err != nil {
		return fmt.Errorf("failed to delete feature: %w", err)
	}
	return nil
}
