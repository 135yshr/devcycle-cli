package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_FeatureConfigurations(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		configs := map[string]*EnvironmentConfig{
			"development": {
				Status: "active",
				Targets: []Target{
					{
						Name: "All Users",
						Audience: Audience{
							Filters: Filters{
								Operator: "and",
								Filters:  []Filter{{Type: "all"}},
							},
						},
						Distribution: []Distribution{
							{Variation: "on", Percentage: 1.0},
						},
					},
				},
			},
			"production": {
				Status:  "inactive",
				Targets: []Target{},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/features/my-feature/configurations" {
				t.Errorf("expected /projects/my-project/features/my-feature/configurations, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(configs)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.FeatureConfigurations(context.Background(), "my-project", "my-feature")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 environments, got %d", len(result))
		}
		if result["development"] == nil {
			t.Error("expected development config")
		}
		if result["development"].Status != "active" {
			t.Errorf("expected active, got %s", result["development"].Status)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("feature not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.FeatureConfigurations(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_UpdateFeatureConfigurations(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {
		configs := map[string]*EnvironmentConfig{
			"development": {
				Status: "active",
				Targets: []Target{
					{
						Name: "All Users",
						Audience: Audience{
							Filters: Filters{
								Operator: "and",
								Filters:  []Filter{{Type: "all"}},
							},
						},
						Distribution: []Distribution{
							{Variation: "on", Percentage: 1.0},
						},
					},
				},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/features/my-feature/configurations" {
				t.Errorf("expected /projects/my-project/features/my-feature/configurations, got %s", r.URL.Path)
			}

			var req UpdateFeatureConfigurationsRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Configurations["development"] == nil {
				t.Error("expected development config in request")
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(configs)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.UpdateFeatureConfigurations(context.Background(), "my-project", "my-feature", &UpdateFeatureConfigurationsRequest{
			Configurations: configs,
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result["development"] == nil {
			t.Error("expected development config")
		}
		if result["development"].Status != "active" {
			t.Errorf("expected active, got %s", result["development"].Status)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("feature not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.UpdateFeatureConfigurations(context.Background(), "my-project", "non-existent", &UpdateFeatureConfigurationsRequest{
			Configurations: map[string]*EnvironmentConfig{
				"development": {Status: "active"},
			},
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_EnableFeature(t *testing.T) {
	t.Run("successful enable", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH, got %s", r.Method)
			}

			var req UpdateFeatureConfigurationsRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Configurations["development"] == nil {
				t.Error("expected development config in request")
			}
			if req.Configurations["development"].Status != "active" {
				t.Errorf("expected status 'active', got %s", req.Configurations["development"].Status)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]*EnvironmentConfig{
				"development": {Status: "active"},
			})
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		err := client.EnableFeature(context.Background(), "my-project", "my-feature", "development")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestClient_DisableFeature(t *testing.T) {
	t.Run("successful disable", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH, got %s", r.Method)
			}

			var req UpdateFeatureConfigurationsRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Configurations["development"] == nil {
				t.Error("expected development config in request")
			}
			if req.Configurations["development"].Status != "inactive" {
				t.Errorf("expected status 'inactive', got %s", req.Configurations["development"].Status)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]*EnvironmentConfig{
				"development": {Status: "inactive"},
			})
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		err := client.DisableFeature(context.Background(), "my-project", "my-feature", "development")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
