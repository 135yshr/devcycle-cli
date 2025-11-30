package api

import (
	"context"
	"fmt"
)

// Projects returns all projects accessible to the authenticated user.
func (c *Client) Projects(ctx context.Context) ([]Project, error) {
	var projects []Project
	if err := c.Get(ctx, "/projects", &projects); err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}
	return projects, nil
}

// Project returns a specific project by its key.
func (c *Client) Project(ctx context.Context, projectKey string) (*Project, error) {
	var project Project
	path := fmt.Sprintf("/projects/%s", projectKey)
	if err := c.Get(ctx, path, &project); err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}
	return &project, nil
}

// CreateProjectRequest represents the request body for creating a project.
type CreateProjectRequest struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description,omitempty"`
}

// UpdateProjectRequest represents the request body for updating a project.
type UpdateProjectRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// CreateProject creates a new project with the specified configuration.
func (c *Client) CreateProject(ctx context.Context, req *CreateProjectRequest) (*Project, error) {
	var project Project
	if err := c.Post(ctx, "/projects", req, &project); err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}
	return &project, nil
}

// UpdateProject updates an existing project's name and/or description.
func (c *Client) UpdateProject(ctx context.Context, projectKey string, req *UpdateProjectRequest) (*Project, error) {
	var project Project
	path := fmt.Sprintf("/projects/%s", projectKey)
	if err := c.Patch(ctx, path, req, &project); err != nil {
		return nil, fmt.Errorf("failed to update project: %w", err)
	}
	return &project, nil
}
