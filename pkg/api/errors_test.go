package api

import (
	"errors"
	"testing"
)

func TestAPIError_Error(t *testing.T) {
	err := &APIError{
		StatusCode: 404,
		Message:    "not found",
	}

	expected := "API error (status 404): not found"
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}
}

func TestIsNotFound(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "404 error",
			err:      &APIError{StatusCode: 404, Message: "not found"},
			expected: true,
		},
		{
			name:     "500 error",
			err:      &APIError{StatusCode: 500, Message: "server error"},
			expected: false,
		},
		{
			name:     "non-API error",
			err:      errors.New("some error"),
			expected: false,
		},
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotFound(tt.err)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsUnauthorized(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "401 error",
			err:      &APIError{StatusCode: 401, Message: "unauthorized"},
			expected: true,
		},
		{
			name:     "403 error",
			err:      &APIError{StatusCode: 403, Message: "forbidden"},
			expected: false,
		},
		{
			name:     "non-API error",
			err:      errors.New("some error"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsUnauthorized(tt.err)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsForbidden(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "403 error",
			err:      &APIError{StatusCode: 403, Message: "forbidden"},
			expected: true,
		},
		{
			name:     "401 error",
			err:      &APIError{StatusCode: 401, Message: "unauthorized"},
			expected: false,
		},
		{
			name:     "non-API error",
			err:      errors.New("some error"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsForbidden(tt.err)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
