package api

import (
	"context"
	"fmt"
	"net/url"
)

// AuditLogs returns all audit logs for a project
func (c *Client) AuditLogs(ctx context.Context, projectKey string) ([]AuditLog, error) {
	var logs []AuditLog
	path := fmt.Sprintf("/projects/%s/audit", url.PathEscape(projectKey))
	if err := c.Get(ctx, path, &logs); err != nil {
		return nil, fmt.Errorf("failed to list audit logs: %w", err)
	}
	return logs, nil
}

// FeatureAuditLogs returns audit logs for a specific feature
func (c *Client) FeatureAuditLogs(ctx context.Context, projectKey, featureKey string) ([]AuditLog, error) {
	var logs []AuditLog
	path := fmt.Sprintf("/projects/%s/features/%s/audit", url.PathEscape(projectKey), url.PathEscape(featureKey))
	if err := c.Get(ctx, path, &logs); err != nil {
		return nil, fmt.Errorf("failed to list feature audit logs: %w", err)
	}
	return logs, nil
}
