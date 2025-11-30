package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_FeatureOverrides(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		overrides := []Override{
			{
				Feature:     "my-feature",
				Environment: "development",
				Variation:   "on",
			},
			{
				Feature:     "my-feature",
				Environment: "staging",
				Variation:   "off",
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/features/my-feature/overrides" {
				t.Errorf("expected /projects/my-project/features/my-feature/overrides, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(overrides)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.FeatureOverrides(context.Background(), "my-project", "my-feature")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 overrides, got %d", len(result))
		}
		if result[0].Environment != "development" {
			t.Errorf("expected development, got %s", result[0].Environment)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]Override{})
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		result, err := client.FeatureOverrides(context.Background(), "my-project", "my-feature")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected 0 overrides, got %d", len(result))
		}
	})
}

func TestClient_CurrentOverride(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		override := Override{
			Feature:     "my-feature",
			Environment: "development",
			Variation:   "on",
			Variables:   map[string]any{"enabled": true},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/features/my-feature/overrides/current" {
				t.Errorf("expected /projects/my-project/features/my-feature/overrides/current, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(override)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.CurrentOverride(context.Background(), "my-project", "my-feature")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Variation != "on" {
			t.Errorf("expected on, got %s", result.Variation)
		}
		if result.Environment != "development" {
			t.Errorf("expected development, got %s", result.Environment)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("override not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.CurrentOverride(context.Background(), "my-project", "my-feature")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_SetOverride(t *testing.T) {
	t.Run("successful set", func(t *testing.T) {
		override := Override{
			Feature:     "my-feature",
			Environment: "development",
			Variation:   "on",
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/features/my-feature/overrides/current" {
				t.Errorf("expected /projects/my-project/features/my-feature/overrides/current, got %s", r.URL.Path)
			}

			var req SetOverrideRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Environment != "development" {
				t.Errorf("expected environment 'development', got %s", req.Environment)
			}
			if req.Variation != "on" {
				t.Errorf("expected variation 'on', got %s", req.Variation)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(override)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.SetOverride(context.Background(), "my-project", "my-feature", &SetOverrideRequest{
			Environment: "development",
			Variation:   "on",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Variation != "on" {
			t.Errorf("expected on, got %s", result.Variation)
		}
	})

	t.Run("feature not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("feature not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.SetOverride(context.Background(), "my-project", "non-existent", &SetOverrideRequest{
			Environment: "development",
			Variation:   "on",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_DeleteOverride(t *testing.T) {
	t.Run("successful delete", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			expectedPath := "/projects/my-project/features/my-feature/overrides/current"
			if r.URL.Path != expectedPath {
				t.Errorf("expected %s, got %s", expectedPath, r.URL.Path)
			}
			if r.URL.Query().Get("environment") != "development" {
				t.Errorf("expected environment=development, got %s", r.URL.Query().Get("environment"))
			}
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		err := client.DeleteOverride(context.Background(), "my-project", "my-feature", "development")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("override not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		err := client.DeleteOverride(context.Background(), "my-project", "my-feature", "development")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_MyOverrides(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		overrides := []Override{
			{
				Feature:     "feature-1",
				Environment: "development",
				Variation:   "on",
			},
			{
				Feature:     "feature-2",
				Environment: "development",
				Variation:   "off",
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/overrides/current" {
				t.Errorf("expected /projects/my-project/overrides/current, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(overrides)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.MyOverrides(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 overrides, got %d", len(result))
		}
		if result[0].Feature != "feature-1" {
			t.Errorf("expected feature-1, got %s", result[0].Feature)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]Override{})
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		result, err := client.MyOverrides(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected 0 overrides, got %d", len(result))
		}
	})
}

func TestClient_DeleteAllMyOverrides(t *testing.T) {
	t.Run("successful delete", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/overrides/current" {
				t.Errorf("expected /projects/my-project/overrides/current, got %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		err := client.DeleteAllMyOverrides(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("unauthorized", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		err := client.DeleteAllMyOverrides(context.Background(), "my-project")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}
