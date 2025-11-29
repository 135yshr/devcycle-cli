package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_ListFeatures(t *testing.T) {
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
		result, err := client.ListFeatures(context.Background(), "my-project")

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
		result, err := client.ListFeatures(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected 0 features, got %d", len(result))
		}
	})
}

func TestClient_GetFeature(t *testing.T) {
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
		result, err := client.GetFeature(context.Background(), "my-project", "feature-one")

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
		_, err := client.GetFeature(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}
