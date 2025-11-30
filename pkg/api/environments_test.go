package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_Environments(t *testing.T) {
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
		result, err := client.Environments(context.Background(), "my-project")

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

func TestClient_Environment(t *testing.T) {
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
		result, err := client.Environment(context.Background(), "my-project", "development")

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
		_, err := client.Environment(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_CreateEnvironment(t *testing.T) {
	t.Run("successful create", func(t *testing.T) {
		environment := Environment{
			ID:          "env-new",
			Key:         "staging",
			Name:        "Staging",
			Description: "Staging environment",
			Type:        "staging",
			Color:       "#ffff00",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST method, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/environments" {
				t.Errorf("expected /projects/my-project/environments, got %s", r.URL.Path)
			}

			var req CreateEnvironmentRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				t.Errorf("failed to decode request body: %v", err)
			}
			if req.Key != "staging" {
				t.Errorf("expected key 'staging', got %s", req.Key)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(environment)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		req := &CreateEnvironmentRequest{
			Name:        "Staging",
			Key:         "staging",
			Description: "Staging environment",
			Type:        "staging",
			Color:       "#ffff00",
		}

		result, err := client.CreateEnvironment(context.Background(), "my-project", req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "staging" {
			t.Errorf("expected staging, got %s", result.Key)
		}
	})
}

func TestClient_UpdateEnvironment(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {
		environment := Environment{
			ID:          "env-1",
			Key:         "development",
			Name:        "Dev Environment Updated",
			Description: "Updated description",
			Type:        "development",
			Color:       "#0000ff",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH method, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/environments/development" {
				t.Errorf("expected /projects/my-project/environments/development, got %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(environment)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		req := &UpdateEnvironmentRequest{
			Name:        "Dev Environment Updated",
			Description: "Updated description",
			Color:       "#0000ff",
		}

		result, err := client.UpdateEnvironment(context.Background(), "my-project", "development", req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Name != "Dev Environment Updated" {
			t.Errorf("expected 'Dev Environment Updated', got %s", result.Name)
		}
	})
}

func TestClient_DeleteEnvironment(t *testing.T) {
	t.Run("successful delete", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE method, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/environments/staging" {
				t.Errorf("expected /projects/my-project/environments/staging, got %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		err := client.DeleteEnvironment(context.Background(), "my-project", "staging")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("environment not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		err := client.DeleteEnvironment(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_RotateSDKKey(t *testing.T) {
	t.Run("successful rotate", func(t *testing.T) {
		response := RotateKeyResponse{
			PreviousKey: "old-key-123",
			NewKey: SDKKeyInfo{
				Key:       "new-key-456",
				CreatedAt: time.Now(),
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST method, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/environments/development/keys" {
				t.Errorf("expected /projects/my-project/environments/development/keys, got %s", r.URL.Path)
			}

			var req RotateKeyRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				t.Errorf("failed to decode request body: %v", err)
			}
			if req.Type != "client" {
				t.Errorf("expected type 'client', got %s", req.Type)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		req := &RotateKeyRequest{
			Type: "client",
		}

		result, err := client.RotateSDKKey(context.Background(), "my-project", "development", req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.PreviousKey != "old-key-123" {
			t.Errorf("expected previous key 'old-key-123', got %s", result.PreviousKey)
		}
		if result.NewKey.Key != "new-key-456" {
			t.Errorf("expected new key 'new-key-456', got %s", result.NewKey.Key)
		}
	})
}
