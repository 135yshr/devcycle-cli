package api

import (
	"context"
	"fmt"
)

// Variations returns all variations for a feature
func (c *Client) Variations(ctx context.Context, projectKey, featureKey string) ([]Variation, error) {
	var variations []Variation
	path := fmt.Sprintf("/projects/%s/features/%s/variations", projectKey, featureKey)
	if err := c.Get(ctx, path, &variations); err != nil {
		return nil, fmt.Errorf("failed to list variations: %w", err)
	}
	return variations, nil
}

// Variation returns a specific variation for a feature
func (c *Client) Variation(ctx context.Context, projectKey, featureKey, variationKey string) (*Variation, error) {
	var variation Variation
	path := fmt.Sprintf("/projects/%s/features/%s/variations/%s", projectKey, featureKey, variationKey)
	if err := c.Get(ctx, path, &variation); err != nil {
		return nil, fmt.Errorf("failed to get variation: %w", err)
	}
	return &variation, nil
}

// CreateVariationRequest represents the request body for creating a variation.
type CreateVariationRequest struct {
	Name      string         `json:"name"`
	Key       string         `json:"key"`
	Variables map[string]any `json:"variables,omitempty"`
}

// UpdateVariationRequest represents the request body for updating a variation.
type UpdateVariationRequest struct {
	Name      string         `json:"name,omitempty"`
	Key       string         `json:"key,omitempty"`
	Variables map[string]any `json:"variables,omitempty"`
}

// CreateVariation creates a new variation for a feature
func (c *Client) CreateVariation(ctx context.Context, projectKey, featureKey string, req *CreateVariationRequest) (*Variation, error) {
	var variation Variation
	path := fmt.Sprintf("/projects/%s/features/%s/variations", projectKey, featureKey)
	if err := c.Post(ctx, path, req, &variation); err != nil {
		return nil, fmt.Errorf("failed to create variation: %w", err)
	}
	return &variation, nil
}

// UpdateVariation updates an existing variation
func (c *Client) UpdateVariation(ctx context.Context, projectKey, featureKey, variationKey string, req *UpdateVariationRequest) (*Variation, error) {
	var variation Variation
	path := fmt.Sprintf("/projects/%s/features/%s/variations/%s", projectKey, featureKey, variationKey)
	if err := c.Patch(ctx, path, req, &variation); err != nil {
		return nil, fmt.Errorf("failed to update variation: %w", err)
	}
	return &variation, nil
}

// DeleteVariation deletes a variation from a feature
func (c *Client) DeleteVariation(ctx context.Context, projectKey, featureKey, variationKey string) error {
	path := fmt.Sprintf("/projects/%s/features/%s/variations/%s", projectKey, featureKey, variationKey)
	if err := c.Delete(ctx, path); err != nil {
		return fmt.Errorf("failed to delete variation: %w", err)
	}
	return nil
}
