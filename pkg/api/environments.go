package api

import (
	"context"
	"fmt"
)

func (c *Client) Environments(ctx context.Context, projectKey string) ([]Environment, error) {
	var environments []Environment
	path := fmt.Sprintf("/projects/%s/environments", projectKey)
	if err := c.Get(ctx, path, &environments); err != nil {
		return nil, fmt.Errorf("failed to list environments: %w", err)
	}
	return environments, nil
}

func (c *Client) Environment(ctx context.Context, projectKey, environmentKey string) (*Environment, error) {
	var environment Environment
	path := fmt.Sprintf("/projects/%s/environments/%s", projectKey, environmentKey)
	if err := c.Get(ctx, path, &environment); err != nil {
		return nil, fmt.Errorf("failed to get environment: %w", err)
	}
	return &environment, nil
}
