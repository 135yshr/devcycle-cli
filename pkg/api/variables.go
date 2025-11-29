package api

import (
	"context"
	"fmt"
)

func (c *Client) Variables(ctx context.Context, projectKey string) ([]Variable, error) {
	var variables []Variable
	path := fmt.Sprintf("/projects/%s/variables", projectKey)
	if err := c.Get(ctx, path, &variables); err != nil {
		return nil, fmt.Errorf("failed to list variables: %w", err)
	}
	return variables, nil
}

func (c *Client) Variable(ctx context.Context, projectKey, variableKey string) (*Variable, error) {
	var variable Variable
	path := fmt.Sprintf("/projects/%s/variables/%s", projectKey, variableKey)
	if err := c.Get(ctx, path, &variable); err != nil {
		return nil, fmt.Errorf("failed to get variable: %w", err)
	}
	return &variable, nil
}
