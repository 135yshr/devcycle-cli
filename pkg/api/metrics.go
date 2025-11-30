package api

import (
	"context"
	"fmt"
	"net/url"
)

// Metrics returns all metrics for a project
func (c *Client) Metrics(ctx context.Context, projectKey string) ([]Metric, error) {
	var metrics []Metric
	path := fmt.Sprintf("/projects/%s/metrics", url.PathEscape(projectKey))
	if err := c.Get(ctx, path, &metrics); err != nil {
		return nil, fmt.Errorf("failed to list metrics: %w", err)
	}
	return metrics, nil
}

// Metric returns a specific metric
func (c *Client) Metric(ctx context.Context, projectKey, metricKey string) (*Metric, error) {
	var metric Metric
	path := fmt.Sprintf("/projects/%s/metrics/%s", url.PathEscape(projectKey), url.PathEscape(metricKey))
	if err := c.Get(ctx, path, &metric); err != nil {
		return nil, fmt.Errorf("failed to get metric: %w", err)
	}
	return &metric, nil
}

// CreateMetric creates a new metric
func (c *Client) CreateMetric(ctx context.Context, projectKey string, req *CreateMetricRequest) (*Metric, error) {
	var metric Metric
	path := fmt.Sprintf("/projects/%s/metrics", url.PathEscape(projectKey))
	if err := c.Post(ctx, path, req, &metric); err != nil {
		return nil, fmt.Errorf("failed to create metric: %w", err)
	}
	return &metric, nil
}

// UpdateMetric updates an existing metric
func (c *Client) UpdateMetric(ctx context.Context, projectKey, metricKey string, req *UpdateMetricRequest) (*Metric, error) {
	var metric Metric
	path := fmt.Sprintf("/projects/%s/metrics/%s", url.PathEscape(projectKey), url.PathEscape(metricKey))
	if err := c.Patch(ctx, path, req, &metric); err != nil {
		return nil, fmt.Errorf("failed to update metric: %w", err)
	}
	return &metric, nil
}

// DeleteMetric deletes a metric
func (c *Client) DeleteMetric(ctx context.Context, projectKey, metricKey string) error {
	path := fmt.Sprintf("/projects/%s/metrics/%s", url.PathEscape(projectKey), url.PathEscape(metricKey))
	if err := c.Delete(ctx, path); err != nil {
		return fmt.Errorf("failed to delete metric: %w", err)
	}
	return nil
}

// MetricResultsOptions contains options for fetching metric results
type MetricResultsOptions struct {
	Environment string
	Feature     string
	StartDate   string
	EndDate     string
}

// MetricResults returns results for a specific metric
func (c *Client) MetricResults(ctx context.Context, projectKey, metricKey string, opts *MetricResultsOptions) (*MetricResults, error) {
	var results MetricResults
	path := fmt.Sprintf("/projects/%s/metrics/%s/results", url.PathEscape(projectKey), url.PathEscape(metricKey))

	// Add query parameters if options are provided
	if opts != nil {
		params := url.Values{}
		if opts.Environment != "" {
			params.Set("environment", opts.Environment)
		}
		if opts.Feature != "" {
			params.Set("feature", opts.Feature)
		}
		if opts.StartDate != "" {
			params.Set("startDate", opts.StartDate)
		}
		if opts.EndDate != "" {
			params.Set("endDate", opts.EndDate)
		}
		if len(params) > 0 {
			path += "?" + params.Encode()
		}
	}

	if err := c.Get(ctx, path, &results); err != nil {
		return nil, fmt.Errorf("failed to get metric results: %w", err)
	}
	return &results, nil
}
