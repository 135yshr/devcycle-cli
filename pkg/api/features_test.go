package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_Features(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		features := []Feature{
			{
				ID:        "feat-1",
				Key:       "feature-one",
				Name:      "Feature One",
				Type:      "release",
				Status:    "active",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:        "feat-2",
				Key:       "feature-two",
				Name:      "Feature Two",
				Type:      "experiment",
				Status:    "inactive",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/features" {
				t.Errorf("expected /projects/my-project/features, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(features)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.Features(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 features, got %d", len(result))
		}
		if result[0].Key != "feature-one" {
			t.Errorf("expected feature-one, got %s", result[0].Key)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]Feature{})
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		result, err := client.Features(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected 0 features, got %d", len(result))
		}
	})
}

func TestClient_Feature(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		feature := Feature{
			ID:          "feat-1",
			Key:         "feature-one",
			Name:        "Feature One",
			Description: "A test feature",
			Type:        "release",
			Status:      "active",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/features/feature-one" {
				t.Errorf("expected /projects/my-project/features/feature-one, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(feature)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.Feature(context.Background(), "my-project", "feature-one")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "feature-one" {
			t.Errorf("expected feature-one, got %s", result.Key)
		}
		if result.Type != "release" {
			t.Errorf("expected release, got %s", result.Type)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("feature not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.Feature(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_CreateFeature(t *testing.T) {
	t.Run("successful create", func(t *testing.T) {
		feature := Feature{
			ID:        "feat-new",
			Key:       "new-feature",
			Name:      "New Feature",
			Type:      "release",
			Status:    "active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/features" {
				t.Errorf("expected /projects/my-project/features, got %s", r.URL.Path)
			}

			var req CreateFeatureRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Name != "New Feature" {
				t.Errorf("expected name 'New Feature', got %s", req.Name)
			}
			if req.Key != "new-feature" {
				t.Errorf("expected key 'new-feature', got %s", req.Key)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(feature)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.CreateFeature(context.Background(), "my-project", &CreateFeatureRequest{
			Name: "New Feature",
			Key:  "new-feature",
			Type: "release",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "new-feature" {
			t.Errorf("expected new-feature, got %s", result.Key)
		}
	})

	t.Run("conflict error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("feature already exists"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.CreateFeature(context.Background(), "my-project", &CreateFeatureRequest{
			Name: "Duplicate Feature",
			Key:  "duplicate-feature",
			Type: "release",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestClient_UpdateFeature(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {
		feature := Feature{
			ID:        "feat-1",
			Key:       "feature-one",
			Name:      "Updated Feature",
			Type:      "release",
			Status:    "active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/features/feature-one" {
				t.Errorf("expected /projects/my-project/features/feature-one, got %s", r.URL.Path)
			}

			var req UpdateFeatureRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Name != "Updated Feature" {
				t.Errorf("expected name 'Updated Feature', got %s", req.Name)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(feature)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.UpdateFeature(context.Background(), "my-project", "feature-one", &UpdateFeatureRequest{
			Name: "Updated Feature",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Name != "Updated Feature" {
			t.Errorf("expected Updated Feature, got %s", result.Name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("feature not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.UpdateFeature(context.Background(), "my-project", "non-existent", &UpdateFeatureRequest{
			Name: "Updated Feature",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_DeleteFeature(t *testing.T) {
	t.Run("successful delete", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/features/feature-one" {
				t.Errorf("expected /projects/my-project/features/feature-one, got %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		err := client.DeleteFeature(context.Background(), "my-project", "feature-one")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("feature not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		err := client.DeleteFeature(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}
