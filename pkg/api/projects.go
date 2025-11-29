package api

import (
	"context"
	"fmt"
)

func (c *Client) ListProjects(ctx context.Context) ([]Project, error) {
	var projects []Project
	if err := c.Get(ctx, "/projects", &projects); err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}
	return projects, nil
}

func (c *Client) GetProject(ctx context.Context, projectKey string) (*Project, error) {
	var project Project
	path := fmt.Sprintf("/projects/%s", projectKey)
	if err := c.Get(ctx, path, &project); err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}
	return &project, nil
}
