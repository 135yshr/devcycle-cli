package api

import (
	"context"
	"fmt"
)

// Variables returns all variables for a project.
func (c *Client) Variables(ctx context.Context, projectKey string) ([]Variable, error) {
	var variables []Variable
	path := fmt.Sprintf("/projects/%s/variables", projectKey)
	if err := c.Get(ctx, path, &variables); err != nil {
		return nil, fmt.Errorf("failed to list variables: %w", err)
	}
	return variables, nil
}

// Variable returns a specific variable by its key.
func (c *Client) Variable(ctx context.Context, projectKey, variableKey string) (*Variable, error) {
	var variable Variable
	path := fmt.Sprintf("/projects/%s/variables/%s", projectKey, variableKey)
	if err := c.Get(ctx, path, &variable); err != nil {
		return nil, fmt.Errorf("failed to get variable: %w", err)
	}
	return &variable, nil
}

// CreateVariableRequest represents the request body for creating a variable.
type CreateVariableRequest struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
	Feature     string `json:"_feature,omitempty"`
}

// UpdateVariableRequest represents the request body for updating a variable.
type UpdateVariableRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// CreateVariable creates a new variable in a project.
func (c *Client) CreateVariable(ctx context.Context, projectKey string, req *CreateVariableRequest) (*Variable, error) {
	var variable Variable
	path := fmt.Sprintf("/projects/%s/variables", projectKey)
	if err := c.Post(ctx, path, req, &variable); err != nil {
		return nil, fmt.Errorf("failed to create variable: %w", err)
	}
	return &variable, nil
}

// UpdateVariable updates an existing variable's properties.
func (c *Client) UpdateVariable(ctx context.Context, projectKey, variableKey string, req *UpdateVariableRequest) (*Variable, error) {
	var variable Variable
	path := fmt.Sprintf("/projects/%s/variables/%s", projectKey, variableKey)
	if err := c.Patch(ctx, path, req, &variable); err != nil {
		return nil, fmt.Errorf("failed to update variable: %w", err)
	}
	return &variable, nil
}

// DeleteVariable removes a variable from a project.
// Warning: This action cannot be undone.
func (c *Client) DeleteVariable(ctx context.Context, projectKey, variableKey string) error {
	path := fmt.Sprintf("/projects/%s/variables/%s", projectKey, variableKey)
	if err := c.Delete(ctx, path); err != nil {
		return fmt.Errorf("failed to delete variable: %w", err)
	}
	return nil
}
