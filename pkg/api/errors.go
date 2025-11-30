package api

import (
	"errors"
	"fmt"
)

// APIError represents an error response from the DevCycle API.
// It contains the HTTP status code and the error message returned by the API.
type APIError struct {
	StatusCode int
	Message    string
}

// Error implements the error interface for APIError.
func (e *APIError) Error() string {
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
}

// IsNotFound reports whether the error is a 404 Not Found response from the API.
func IsNotFound(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == 404
	}
	return false
}

// IsUnauthorized reports whether the error is a 401 Unauthorized response from the API.
// This typically indicates an invalid or expired authentication token.
func IsUnauthorized(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == 401
	}
	return false
}

// IsForbidden reports whether the error is a 403 Forbidden response from the API.
// This typically indicates insufficient permissions for the requested operation.
func IsForbidden(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == 403
	}
	return false
}
