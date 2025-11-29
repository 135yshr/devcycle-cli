package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_ListVariables(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		variables := []Variable{
			{
				ID:        "var-1",
				Key:       "enable-feature",
				Name:      "Enable Feature",
				Type:      "Boolean",
				Status:    "active",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:        "var-2",
				Key:       "feature-config",
				Name:      "Feature Config",
				Type:      "JSON",
				Status:    "active",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/variables" {
				t.Errorf("expected /projects/my-project/variables, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(variables)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.ListVariables(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 variables, got %d", len(result))
		}
		if result[0].Type != "Boolean" {
			t.Errorf("expected Boolean, got %s", result[0].Type)
		}
	})
}

func TestClient_GetVariable(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		variable := Variable{
			ID:          "var-1",
			Key:         "enable-feature",
			Name:        "Enable Feature",
			Description: "Enables the new feature",
			Type:        "Boolean",
			Status:      "active",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/variables/enable-feature" {
				t.Errorf("expected /projects/my-project/variables/enable-feature, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(variable)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.GetVariable(context.Background(), "my-project", "enable-feature")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "enable-feature" {
			t.Errorf("expected enable-feature, got %s", result.Key)
		}
		if result.Type != "Boolean" {
			t.Errorf("expected Boolean, got %s", result.Type)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("variable not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.GetVariable(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}
