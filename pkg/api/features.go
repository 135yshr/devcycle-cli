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
