package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_ListEnvironments(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		environments := []Environment{
			{
				ID:        "env-1",
				Key:       "development",
				Name:      "Development",
				Type:      "development",
				Color:     "#00ff00",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:        "env-2",
				Key:       "production",
				Name:      "Production",
				Type:      "production",
				Color:     "#ff0000",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/environments" {
				t.Errorf("expected /projects/my-project/environments, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(environments)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.ListEnvironments(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 environments, got %d", len(result))
		}
		if result[0].Type != "development" {
			t.Errorf("expected development, got %s", result[0].Type)
		}
		if result[1].Type != "production" {
			t.Errorf("expected production, got %s", result[1].Type)
		}
	})
}

func TestClient_GetEnvironment(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		environment := Environment{
			ID:          "env-1",
			Key:         "development",
			Name:        "Development",
			Description: "Development environment",
			Type:        "development",
			Color:       "#00ff00",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/environments/development" {
				t.Errorf("expected /projects/my-project/environments/development, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(environment)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.GetEnvironment(context.Background(), "my-project", "development")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "development" {
			t.Errorf("expected development, got %s", result.Key)
		}
		if result.Color != "#00ff00" {
			t.Errorf("expected #00ff00, got %s", result.Color)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("environment not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.GetEnvironment(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}
