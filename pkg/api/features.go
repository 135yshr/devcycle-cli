package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func (c *Client) Features(ctx context.Context, projectKey string) ([]Feature, error) {
	var features []Feature
	path := fmt.Sprintf("/projects/%s/features", projectKey)
	if err := c.Get(ctx, path, &features); err != nil {
		return nil, fmt.Errorf("failed to list features: %w", err)
	}
	return features, nil
}

func (c *Client) Feature(ctx context.Context, projectKey, featureKey string) (*Feature, error) {
	var feature Feature
	path := fmt.Sprintf("/projects/%s/features/%s", projectKey, featureKey)
	if err := c.Get(ctx, path, &feature); err != nil {
		return nil, fmt.Errorf("failed to get feature: %w", err)
	}
	return &feature, nil
}

// CreateFeatureRequest represents the request body for creating a feature.
type CreateFeatureRequest struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
}

// UpdateFeatureRequest represents the request body for updating a feature.
type UpdateFeatureRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c *Client) CreateFeature(ctx context.Context, projectKey string, req *CreateFeatureRequest) (*Feature, error) {
	var feature Feature
	path := fmt.Sprintf("/projects/%s/features", projectKey)
	if err := c.Post(ctx, path, req, &feature); err != nil {
		return nil, fmt.Errorf("failed to create feature: %w", err)
	}
	return &feature, nil
}

func (c *Client) UpdateFeature(ctx context.Context, projectKey, featureKey string, req *UpdateFeatureRequest) (*Feature, error) {
	var feature Feature
	path := fmt.Sprintf("/projects/%s/features/%s", projectKey, featureKey)
	if err := c.Patch(ctx, path, req, &feature); err != nil {
		return nil, fmt.Errorf("failed to update feature: %w", err)
	}
	return &feature, nil
}

func (c *Client) DeleteFeature(ctx context.Context, projectKey, featureKey string) error {
	path := fmt.Sprintf("/projects/%s/features/%s", projectKey, featureKey)
	if err := c.Delete(ctx, path); err != nil {
		return fmt.Errorf("failed to delete feature: %w", err)
	}
	return nil
}

// v2 API methods

// CreateFeatureV2 creates a feature using v2 API with full configuration support
func (c *Client) CreateFeatureV2(ctx context.Context, projectKey string, req *CreateFeatureV2Request) (*FeatureV2, error) {
	var feature FeatureV2
	path := fmt.Sprintf("/projects/%s/features", projectKey)
	if err := c.PostV2(ctx, path, req, &feature); err != nil {
		return nil, fmt.Errorf("failed to create feature (v2): %w", err)
	}
	return &feature, nil
}

// CreateFeatureFromFile creates a feature from a JSON file using v2 API
func (c *Client) CreateFeatureFromFile(ctx context.Context, projectKey, filePath string) (*FeatureV2, error) {
	req, err := LoadFeatureRequestFromFile(filePath)
	if err != nil {
		return nil, err
	}
	return c.CreateFeatureV2(ctx, projectKey, req)
}

// UpdateFeatureV2 updates a feature using v2 API with full configuration support
func (c *Client) UpdateFeatureV2(ctx context.Context, projectKey, featureKey string, req *CreateFeatureV2Request) (*FeatureV2, error) {
	var feature FeatureV2
	path := fmt.Sprintf("/projects/%s/features/%s", projectKey, featureKey)
	if err := c.PatchV2(ctx, path, req, &feature); err != nil {
		return nil, fmt.Errorf("failed to update feature (v2): %w", err)
	}
	return &feature, nil
}

// LoadFeatureRequestFromFile loads a CreateFeatureV2Request from a JSON file
// If filePath is "-", it reads from stdin
func LoadFeatureRequestFromFile(filePath string) (*CreateFeatureV2Request, error) {
	var data []byte
	var err error

	if filePath == "-" {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("failed to read from stdin: %w", err)
		}
	} else {
		data, err = os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
		}
	}

	var req CreateFeatureV2Request
	if err := json.Unmarshal(data, &req); err != nil {
		if filePath == "-" {
			return nil, fmt.Errorf("failed to parse JSON from stdin: %w", err)
		}
		return nil, fmt.Errorf("failed to parse JSON from %s: %w", filePath, err)
	}

	return &req, nil
}

// ValidateFeatureRequest validates a CreateFeatureV2Request
func ValidateFeatureRequest(req *CreateFeatureV2Request) error {
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	if req.Key == "" {
		return fmt.Errorf("key is required")
	}
	if req.Type == "" {
		req.Type = "release" // default type
	}

	validTypes := map[string]bool{
		"release":    true,
		"experiment": true,
		"permission": true,
		"ops":        true,
	}
	if !validTypes[req.Type] {
		return fmt.Errorf("invalid type: %s (must be one of: release, experiment, permission, ops)", req.Type)
	}

	// Validate variable types
	validVarTypes := map[string]bool{
		"String":  true,
		"Boolean": true,
		"Number":  true,
		"JSON":    true,
	}
	for _, v := range req.Variables {
		if v.Key == "" {
			return fmt.Errorf("variable key is required")
		}
		if v.Type == "" {
			return fmt.Errorf("variable type is required for variable %s", v.Key)
		}
		if !validVarTypes[v.Type] {
			return fmt.Errorf("invalid variable type for %s: %s (must be one of: String, Boolean, Number, JSON)", v.Key, v.Type)
		}
	}

	// Validate variations
	for _, v := range req.Variations {
		if v.Key == "" {
			return fmt.Errorf("variation key is required")
		}
		if v.Name == "" {
			return fmt.Errorf("variation name is required for variation %s", v.Key)
		}
	}

	// Validate configurations
	validStatuses := map[string]bool{
		"active":   true,
		"inactive": true,
	}
	for envKey, config := range req.Configurations {
		if config == nil {
			continue
		}
		if config.Status != "" && !validStatuses[config.Status] {
			return fmt.Errorf("invalid status for environment %s: %s (must be active or inactive)", envKey, config.Status)
		}
		for _, target := range config.Targets {
			if len(target.Distribution) == 0 {
				return fmt.Errorf("distribution is required for target in environment %s", envKey)
			}
			var totalPercentage float64
			for _, dist := range target.Distribution {
				if dist.Variation == "" {
					return fmt.Errorf("_variation is required in distribution for environment %s", envKey)
				}
				if dist.Percentage < 0 || dist.Percentage > 1 {
					return fmt.Errorf("percentage must be between 0 and 1 for environment %s", envKey)
				}
				totalPercentage += dist.Percentage
			}
			if totalPercentage < 0.99 || totalPercentage > 1.01 {
				return fmt.Errorf("total distribution percentage must equal 1.0 for environment %s (got %.2f)", envKey, totalPercentage)
			}
		}
	}

	return nil
}
