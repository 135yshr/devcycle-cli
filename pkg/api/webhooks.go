package api

import (
	"context"
	"fmt"
)

// Webhooks returns all webhooks for a project
func (c *Client) Webhooks(ctx context.Context, projectKey string) ([]Webhook, error) {
	var webhooks []Webhook
	path := fmt.Sprintf("/projects/%s/webhooks", projectKey)
	if err := c.Get(ctx, path, &webhooks); err != nil {
		return nil, fmt.Errorf("failed to list webhooks: %w", err)
	}
	return webhooks, nil
}

// Webhook returns a specific webhook
func (c *Client) Webhook(ctx context.Context, projectKey, webhookID string) (*Webhook, error) {
	var webhook Webhook
	path := fmt.Sprintf("/projects/%s/webhooks/%s", projectKey, webhookID)
	if err := c.Get(ctx, path, &webhook); err != nil {
		return nil, fmt.Errorf("failed to get webhook: %w", err)
	}
	return &webhook, nil
}

// CreateWebhook creates a new webhook
func (c *Client) CreateWebhook(ctx context.Context, projectKey string, req *CreateWebhookRequest) (*Webhook, error) {
	var webhook Webhook
	path := fmt.Sprintf("/projects/%s/webhooks", projectKey)
	if err := c.Post(ctx, path, req, &webhook); err != nil {
		return nil, fmt.Errorf("failed to create webhook: %w", err)
	}
	return &webhook, nil
}

// UpdateWebhook updates an existing webhook
func (c *Client) UpdateWebhook(ctx context.Context, projectKey, webhookID string, req *UpdateWebhookRequest) (*Webhook, error) {
	var webhook Webhook
	path := fmt.Sprintf("/projects/%s/webhooks/%s", projectKey, webhookID)
	if err := c.Patch(ctx, path, req, &webhook); err != nil {
		return nil, fmt.Errorf("failed to update webhook: %w", err)
	}
	return &webhook, nil
}

// DeleteWebhook deletes a webhook
func (c *Client) DeleteWebhook(ctx context.Context, projectKey, webhookID string) error {
	path := fmt.Sprintf("/projects/%s/webhooks/%s", projectKey, webhookID)
	if err := c.Delete(ctx, path); err != nil {
		return fmt.Errorf("failed to delete webhook: %w", err)
	}
	return nil
}
