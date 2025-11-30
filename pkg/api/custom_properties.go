package api

import (
	"context"
	"fmt"
	"net/url"
)

// CustomProperties returns all custom properties for a project
func (c *Client) CustomProperties(ctx context.Context, projectKey string) ([]CustomProperty, error) {
	var properties []CustomProperty
	path := fmt.Sprintf("/projects/%s/customProperties", url.PathEscape(projectKey))
	if err := c.Get(ctx, path, &properties); err != nil {
		return nil, fmt.Errorf("failed to list custom properties: %w", err)
	}
	return properties, nil
}

// CustomProperty returns a specific custom property
func (c *Client) CustomProperty(ctx context.Context, projectKey, propertyKey string) (*CustomProperty, error) {
	var property CustomProperty
	path := fmt.Sprintf("/projects/%s/customProperties/%s", url.PathEscape(projectKey), url.PathEscape(propertyKey))
	if err := c.Get(ctx, path, &property); err != nil {
		return nil, fmt.Errorf("failed to get custom property: %w", err)
	}
	return &property, nil
}

// CreateCustomProperty creates a new custom property
func (c *Client) CreateCustomProperty(ctx context.Context, projectKey string, req *CreateCustomPropertyRequest) (*CustomProperty, error) {
	var property CustomProperty
	path := fmt.Sprintf("/projects/%s/customProperties", url.PathEscape(projectKey))
	if err := c.Post(ctx, path, req, &property); err != nil {
		return nil, fmt.Errorf("failed to create custom property: %w", err)
	}
	return &property, nil
}

// UpdateCustomProperty updates an existing custom property
func (c *Client) UpdateCustomProperty(ctx context.Context, projectKey, propertyKey string, req *UpdateCustomPropertyRequest) (*CustomProperty, error) {
	var property CustomProperty
	path := fmt.Sprintf("/projects/%s/customProperties/%s", url.PathEscape(projectKey), url.PathEscape(propertyKey))
	if err := c.Patch(ctx, path, req, &property); err != nil {
		return nil, fmt.Errorf("failed to update custom property: %w", err)
	}
	return &property, nil
}

// DeleteCustomProperty deletes a custom property
func (c *Client) DeleteCustomProperty(ctx context.Context, projectKey, propertyKey string) error {
	path := fmt.Sprintf("/projects/%s/customProperties/%s", url.PathEscape(projectKey), url.PathEscape(propertyKey))
	if err := c.Delete(ctx, path); err != nil {
		return fmt.Errorf("failed to delete custom property: %w", err)
	}
	return nil
}
