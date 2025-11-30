package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Variations(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		variations := []Variation{
			{
				ID:        "var-1",
				Key:       "off",
				Name:      "Off",
				Variables: map[string]any{"enabled": false},
			},
			{
				ID:        "var-2",
				Key:       "on",
				Name:      "On",
				Variables: map[string]any{"enabled": true},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/features/my-feature/variations" {
				t.Errorf("expected /projects/my-project/features/my-feature/variations, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(variations)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.Variations(context.Background(), "my-project", "my-feature")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 variations, got %d", len(result))
		}
		if result[0].Key != "off" {
			t.Errorf("expected off, got %s", result[0].Key)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]Variation{})
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		result, err := client.Variations(context.Background(), "my-project", "my-feature")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected 0 variations, got %d", len(result))
		}
	})
}

func TestClient_Variation(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		variation := Variation{
			ID:        "var-1",
			Key:       "on",
			Name:      "On",
			Variables: map[string]any{"enabled": true},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/features/my-feature/variations/on" {
				t.Errorf("expected /projects/my-project/features/my-feature/variations/on, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(variation)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.Variation(context.Background(), "my-project", "my-feature", "on")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "on" {
			t.Errorf("expected on, got %s", result.Key)
		}
		if result.Name != "On" {
			t.Errorf("expected On, got %s", result.Name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("variation not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.Variation(context.Background(), "my-project", "my-feature", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_CreateVariation(t *testing.T) {
	t.Run("successful create", func(t *testing.T) {
		variation := Variation{
			ID:        "var-new",
			Key:       "new-variation",
			Name:      "New Variation",
			Variables: map[string]any{"enabled": true},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/features/my-feature/variations" {
				t.Errorf("expected /projects/my-project/features/my-feature/variations, got %s", r.URL.Path)
			}

			var req CreateVariationRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Name != "New Variation" {
				t.Errorf("expected name 'New Variation', got %s", req.Name)
			}
			if req.Key != "new-variation" {
				t.Errorf("expected key 'new-variation', got %s", req.Key)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(variation)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.CreateVariation(context.Background(), "my-project", "my-feature", &CreateVariationRequest{
			Name:      "New Variation",
			Key:       "new-variation",
			Variables: map[string]any{"enabled": true},
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "new-variation" {
			t.Errorf("expected new-variation, got %s", result.Key)
		}
	})

	t.Run("conflict error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("variation already exists"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.CreateVariation(context.Background(), "my-project", "my-feature", &CreateVariationRequest{
			Name: "Duplicate Variation",
			Key:  "duplicate-variation",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestClient_UpdateVariation(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {
		variation := Variation{
			ID:        "var-1",
			Key:       "on",
			Name:      "Updated Variation",
			Variables: map[string]any{"enabled": true, "percentage": 50},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/features/my-feature/variations/on" {
				t.Errorf("expected /projects/my-project/features/my-feature/variations/on, got %s", r.URL.Path)
			}

			var req UpdateVariationRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Name != "Updated Variation" {
				t.Errorf("expected name 'Updated Variation', got %s", req.Name)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(variation)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.UpdateVariation(context.Background(), "my-project", "my-feature", "on", &UpdateVariationRequest{
			Name: "Updated Variation",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Name != "Updated Variation" {
			t.Errorf("expected Updated Variation, got %s", result.Name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("variation not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.UpdateVariation(context.Background(), "my-project", "my-feature", "non-existent", &UpdateVariationRequest{
			Name: "Updated Variation",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_DeleteVariation(t *testing.T) {
	t.Run("successful delete", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/features/my-feature/variations/on" {
				t.Errorf("expected /projects/my-project/features/my-feature/variations/on, got %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		err := client.DeleteVariation(context.Background(), "my-project", "my-feature", "on")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("variation not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		err := client.DeleteVariation(context.Background(), "my-project", "my-feature", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}
