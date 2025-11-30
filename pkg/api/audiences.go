package api

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// AudienceDefinition represents a reusable audience definition
type AudienceDefinition struct {
	ID          string    `json:"_id"`
	Key         string    `json:"key"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Filters     Filters   `json:"filters"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// CreateAudienceRequest represents the request body for creating an audience
type CreateAudienceRequest struct {
	Name        string  `json:"name"`
	Key         string  `json:"key"`
	Description string  `json:"description,omitempty"`
	Filters     Filters `json:"filters"`
}

// UpdateAudienceRequest represents the request body for updating an audience
type UpdateAudienceRequest struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Filters     *Filters `json:"filters,omitempty"`
}

// Audiences returns all audiences for a project
func (c *Client) Audiences(ctx context.Context, projectKey string) ([]AudienceDefinition, error) {
	var audiences []AudienceDefinition
	path := fmt.Sprintf("/projects/%s/audiences", url.PathEscape(projectKey))
	if err := c.Get(ctx, path, &audiences); err != nil {
		return nil, fmt.Errorf("failed to list audiences: %w", err)
	}
	return audiences, nil
}

// Audience returns a specific audience
func (c *Client) Audience(ctx context.Context, projectKey, audienceKey string) (*AudienceDefinition, error) {
	var audience AudienceDefinition
	path := fmt.Sprintf("/projects/%s/audiences/%s", url.PathEscape(projectKey), url.PathEscape(audienceKey))
	if err := c.Get(ctx, path, &audience); err != nil {
		return nil, fmt.Errorf("failed to get audience: %w", err)
	}
	return &audience, nil
}

// CreateAudience creates a new audience
func (c *Client) CreateAudience(ctx context.Context, projectKey string, req *CreateAudienceRequest) (*AudienceDefinition, error) {
	var audience AudienceDefinition
	path := fmt.Sprintf("/projects/%s/audiences", url.PathEscape(projectKey))
	if err := c.Post(ctx, path, req, &audience); err != nil {
		return nil, fmt.Errorf("failed to create audience: %w", err)
	}
	return &audience, nil
}

// UpdateAudience updates an existing audience
func (c *Client) UpdateAudience(ctx context.Context, projectKey, audienceKey string, req *UpdateAudienceRequest) (*AudienceDefinition, error) {
	var audience AudienceDefinition
	path := fmt.Sprintf("/projects/%s/audiences/%s", url.PathEscape(projectKey), url.PathEscape(audienceKey))
	if err := c.Patch(ctx, path, req, &audience); err != nil {
		return nil, fmt.Errorf("failed to update audience: %w", err)
	}
	return &audience, nil
}

// DeleteAudience deletes an audience
func (c *Client) DeleteAudience(ctx context.Context, projectKey, audienceKey string) error {
	path := fmt.Sprintf("/projects/%s/audiences/%s", url.PathEscape(projectKey), url.PathEscape(audienceKey))
	if err := c.Delete(ctx, path); err != nil {
		return fmt.Errorf("failed to delete audience: %w", err)
	}
	return nil
}
