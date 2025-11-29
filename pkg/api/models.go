package api

import "time"

type Project struct {
	ID          string    `json:"_id"`
	Key         string    `json:"key"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Environment struct {
	ID          string    `json:"_id"`
	Key         string    `json:"key"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Color       string    `json:"color,omitempty"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Feature struct {
	ID          string    `json:"_id"`
	Key         string    `json:"key"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type"`
	Status      string    `json:"status,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Variable struct {
	ID          string    `json:"_id"`
	Key         string    `json:"key"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type"`
	Status      string    `json:"status,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// v2 API types

// CreateFeatureV2Request represents a request to create a feature using v2 API
type CreateFeatureV2Request struct {
	Name             string                        `json:"name"`
	Key              string                        `json:"key"`
	Description      string                        `json:"description,omitempty"`
	Type             string                        `json:"type"`
	Tags             []string                      `json:"tags,omitempty"`
	ControlVariation string                        `json:"controlVariation,omitempty"`
	SDKVisibility    *SDKVisibility                `json:"sdkVisibility,omitempty"`
	Settings         *FeatureSettings              `json:"settings,omitempty"`
	Variables        []VariableDefinition          `json:"variables,omitempty"`
	Variations       []VariationDefinition         `json:"variations,omitempty"`
	Configurations   map[string]*EnvironmentConfig `json:"configurations,omitempty"`
}

// SDKVisibility represents SDK visibility settings
type SDKVisibility struct {
	Mobile bool `json:"mobile"`
	Client bool `json:"client"`
	Server bool `json:"server"`
}

// FeatureSettings represents feature-level settings
type FeatureSettings struct {
	PublicName        string `json:"publicName"`
	PublicDescription string `json:"publicDescription"`
	OptInEnabled      bool   `json:"optInEnabled"`
}

// VariableDefinition represents a variable definition for v2 API
type VariableDefinition struct {
	Key         string `json:"key"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type"` // String, Boolean, Number, JSON
	Description string `json:"description,omitempty"`
}

// VariationDefinition represents a variation definition for v2 API
type VariationDefinition struct {
	Key       string         `json:"key"`
	Name      string         `json:"name"`
	Variables map[string]any `json:"variables,omitempty"`
}

// EnvironmentConfig represents environment-specific configuration
type EnvironmentConfig struct {
	Status  string   `json:"status"` // active, inactive
	Targets []Target `json:"targets,omitempty"`
}

// Target represents a targeting rule
type Target struct {
	Name         string         `json:"name,omitempty"`
	Audience     Audience       `json:"audience"`
	Distribution []Distribution `json:"distribution"`
}

// Audience represents audience definition
type Audience struct {
	Name    string  `json:"name,omitempty"`
	Filters Filters `json:"filters"`
}

// Filters represents filter settings
type Filters struct {
	Operator string   `json:"operator"` // and, or
	Filters  []Filter `json:"filters"`
}

// Filter represents an individual filter
type Filter struct {
	Type       string `json:"type"` // all, user, audienceMatch, optIn
	SubType    string `json:"subType,omitempty"`
	Comparator string `json:"comparator,omitempty"`
	Values     any    `json:"values,omitempty"`
}

// Distribution represents variation distribution
type Distribution struct {
	Variation  string  `json:"_variation"`
	Percentage float64 `json:"percentage"` // 0.0 - 1.0
}

// FeatureV2 represents a feature response from v2 API
type FeatureV2 struct {
	ID               string                        `json:"_id"`
	Key              string                        `json:"key"`
	Name             string                        `json:"name"`
	Description      string                        `json:"description,omitempty"`
	Type             string                        `json:"type"`
	Status           string                        `json:"status,omitempty"`
	Tags             []string                      `json:"tags,omitempty"`
	ControlVariation string                        `json:"controlVariation,omitempty"`
	SDKVisibility    *SDKVisibility                `json:"sdkVisibility,omitempty"`
	Settings         *FeatureSettings              `json:"settings,omitempty"`
	Variables        []VariableDefinition          `json:"variables,omitempty"`
	Variations       []VariationDefinition         `json:"variations,omitempty"`
	Configurations   map[string]*EnvironmentConfig `json:"configurations,omitempty"`
	CreatedAt        time.Time                     `json:"createdAt"`
	UpdatedAt        time.Time                     `json:"updatedAt"`
}
