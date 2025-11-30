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

// Variation represents a feature variation
type Variation struct {
	ID        string         `json:"_id"`
	Key       string         `json:"key"`
	Name      string         `json:"name"`
	Variables map[string]any `json:"variables,omitempty"`
}

// FeatureConfiguration represents a feature's targeting configuration
type FeatureConfiguration struct {
	ID      string                        `json:"_id,omitempty"`
	Feature string                        `json:"_feature,omitempty"`
	Status  string                        `json:"status,omitempty"`
	Targets []Target                      `json:"targets,omitempty"`
	Configs map[string]*EnvironmentConfig `json:"configurations,omitempty"`
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

// =============================================================================
// Phase 5: Operations & Monitoring
// =============================================================================

// AuditLog represents an audit log entry
type AuditLog struct {
	ID        string    `json:"_id"`
	Type      string    `json:"type"`
	User      AuditUser `json:"user"`
	Changes   []Change  `json:"changes"`
	CreatedAt time.Time `json:"createdAt"`
}

// AuditUser represents a user in audit logs
type AuditUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Change represents a change in audit logs
type Change struct {
	Type             string `json:"type"`
	NewContents      any    `json:"newContents"`
	PreviousContents any    `json:"previousContents"`
}

// Metric represents a metric definition
type Metric struct {
	ID          string    `json:"_id"`
	Key         string    `json:"key"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	EventType   string    `json:"eventType"`
	OptimizeFor string    `json:"optimizeFor"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// CreateMetricRequest represents a request to create a metric
type CreateMetricRequest struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	Type        string `json:"type"`
	EventType   string `json:"eventType"`
	OptimizeFor string `json:"optimizeFor"`
	Description string `json:"description,omitempty"`
}

// UpdateMetricRequest represents a request to update a metric
type UpdateMetricRequest struct {
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	EventType   string `json:"eventType,omitempty"`
	OptimizeFor string `json:"optimizeFor,omitempty"`
	Description string `json:"description,omitempty"`
}

// MetricResults represents metric results data
type MetricResults struct {
	Data []MetricResultData `json:"data"`
}

// MetricResultData represents a single metric result entry
type MetricResultData struct {
	VariationKey string  `json:"variationKey"`
	Count        int     `json:"count"`
	Value        float64 `json:"value"`
}

// Webhook represents a webhook configuration
type Webhook struct {
	ID          string    `json:"_id"`
	URL         string    `json:"url"`
	Description string    `json:"description,omitempty"`
	IsEnabled   bool      `json:"isEnabled"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// CreateWebhookRequest represents a request to create a webhook
type CreateWebhookRequest struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
	IsEnabled   bool   `json:"isEnabled"`
}

// UpdateWebhookRequest represents a request to update a webhook
type UpdateWebhookRequest struct {
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	IsEnabled   *bool  `json:"isEnabled,omitempty"`
}

// CustomProperty represents a custom property definition
type CustomProperty struct {
	ID          string    `json:"_id"`
	Key         string    `json:"key"`
	PropertyKey string    `json:"propertyKey"`
	DisplayName string    `json:"displayName"`
	Type        string    `json:"type"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// CreateCustomPropertyRequest represents a request to create a custom property
type CreateCustomPropertyRequest struct {
	Key         string `json:"key"`
	DisplayName string `json:"displayName"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}

// UpdateCustomPropertyRequest represents a request to update a custom property
type UpdateCustomPropertyRequest struct {
	DisplayName string `json:"displayName,omitempty"`
	Description string `json:"description,omitempty"`
}
