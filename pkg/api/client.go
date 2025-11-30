package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the base URL for DevCycle Management API v1.
	DefaultBaseURL = "https://api.devcycle.com/v1"
	// DefaultBaseURLV2 is the base URL for DevCycle Management API v2.
	DefaultBaseURLV2 = "https://api.devcycle.com/v2"
	// AuthURL is the OAuth2 token endpoint for DevCycle authentication.
	AuthURL = "https://auth.devcycle.com/oauth/token"
	// DefaultTimeout is the default HTTP client timeout for API requests.
	DefaultTimeout = 30 * time.Second
)

// Client provides methods for interacting with the DevCycle Management API.
// It handles authentication, request/response serialization, and error handling.
type Client struct {
	baseURL    string
	httpClient *http.Client
	token      string
}

// ClientOption is a function that configures a Client.
// Use the With* functions to create ClientOptions.
type ClientOption func(*Client)

// WithBaseURL returns a ClientOption that sets a custom base URL for API requests.
// Use this for testing or connecting to non-production environments.
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithTimeout returns a ClientOption that sets a custom timeout for HTTP requests.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithToken returns a ClientOption that sets the authentication token.
// The token should be obtained using the Authenticate function.
func WithToken(token string) ClientOption {
	return func(c *Client) {
		c.token = token
	}
}

// NewClient creates a new DevCycle API client with the given options.
// By default, it uses DefaultBaseURL and DefaultTimeout.
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		baseURL: DefaultBaseURL,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// SetToken updates the authentication token for the client.
// This is useful for token refresh scenarios.
func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) do(ctx context.Context, method, path string, body any, result any) error {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(respBody),
		}
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// Get sends a GET request to the specified path and unmarshals the response into result.
func (c *Client) Get(ctx context.Context, path string, result any) error {
	return c.do(ctx, http.MethodGet, path, nil, result)
}

// Post sends a POST request with the given body and unmarshals the response into result.
func (c *Client) Post(ctx context.Context, path string, body any, result any) error {
	return c.do(ctx, http.MethodPost, path, body, result)
}

// Patch sends a PATCH request with the given body and unmarshals the response into result.
func (c *Client) Patch(ctx context.Context, path string, body any, result any) error {
	return c.do(ctx, http.MethodPatch, path, body, result)
}

// Put sends a PUT request with the given body and unmarshals the response into result.
func (c *Client) Put(ctx context.Context, path string, body any, result any) error {
	return c.do(ctx, http.MethodPut, path, body, result)
}

// Delete sends a DELETE request to the specified path.
func (c *Client) Delete(ctx context.Context, path string) error {
	return c.do(ctx, http.MethodDelete, path, nil, nil)
}

// doV2 executes HTTP request against v2 API endpoint
func (c *Client) doV2(ctx context.Context, method, path string, body any, result any) error {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	url := DefaultBaseURLV2 + path
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(respBody),
		}
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// PostV2 sends a POST request to the v2 API endpoint with the given body
// and unmarshals the response into result.
func (c *Client) PostV2(ctx context.Context, path string, body any, result any) error {
	return c.doV2(ctx, http.MethodPost, path, body, result)
}

// PatchV2 sends a PATCH request to the v2 API endpoint with the given body
// and unmarshals the response into result.
func (c *Client) PatchV2(ctx context.Context, path string, body any, result any) error {
	return c.doV2(ctx, http.MethodPatch, path, body, result)
}
