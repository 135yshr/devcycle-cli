package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Audiences(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		audiences := []AudienceDefinition{
			{
				ID:          "aud-1",
				Key:         "beta-users",
				Name:        "Beta Users",
				Description: "Users in beta program",
				Filters: Filters{
					Operator: "and",
					Filters: []Filter{
						{Type: "user", SubType: "email", Comparator: "contain", Values: []any{"@beta.com"}},
					},
				},
			},
			{
				ID:          "aud-2",
				Key:         "internal-users",
				Name:        "Internal Users",
				Description: "Internal team members",
				Filters: Filters{
					Operator: "and",
					Filters: []Filter{
						{Type: "user", SubType: "email", Comparator: "contain", Values: []any{"@company.com"}},
					},
				},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/audiences" {
				t.Errorf("expected /projects/my-project/audiences, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(audiences)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.Audiences(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 audiences, got %d", len(result))
		}
		if result[0].Key != "beta-users" {
			t.Errorf("expected beta-users, got %s", result[0].Key)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]AudienceDefinition{})
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		result, err := client.Audiences(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected 0 audiences, got %d", len(result))
		}
	})
}

func TestClient_Audience(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		audience := AudienceDefinition{
			ID:          "aud-1",
			Key:         "beta-users",
			Name:        "Beta Users",
			Description: "Users in beta program",
			Filters: Filters{
				Operator: "and",
				Filters: []Filter{
					{Type: "user", SubType: "email", Comparator: "contain", Values: []any{"@beta.com"}},
				},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project/audiences/beta-users" {
				t.Errorf("expected /projects/my-project/audiences/beta-users, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(audience)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.Audience(context.Background(), "my-project", "beta-users")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "beta-users" {
			t.Errorf("expected beta-users, got %s", result.Key)
		}
		if result.Name != "Beta Users" {
			t.Errorf("expected Beta Users, got %s", result.Name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("audience not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.Audience(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_CreateAudience(t *testing.T) {
	t.Run("successful create", func(t *testing.T) {
		audience := AudienceDefinition{
			ID:          "aud-new",
			Key:         "new-audience",
			Name:        "New Audience",
			Description: "A new audience",
			Filters: Filters{
				Operator: "and",
				Filters: []Filter{
					{Type: "all"},
				},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/audiences" {
				t.Errorf("expected /projects/my-project/audiences, got %s", r.URL.Path)
			}

			var req CreateAudienceRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Name != "New Audience" {
				t.Errorf("expected name 'New Audience', got %s", req.Name)
			}
			if req.Key != "new-audience" {
				t.Errorf("expected key 'new-audience', got %s", req.Key)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(audience)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.CreateAudience(context.Background(), "my-project", &CreateAudienceRequest{
			Name:        "New Audience",
			Key:         "new-audience",
			Description: "A new audience",
			Filters: Filters{
				Operator: "and",
				Filters: []Filter{
					{Type: "all"},
				},
			},
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "new-audience" {
			t.Errorf("expected new-audience, got %s", result.Key)
		}
	})

	t.Run("conflict error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("audience already exists"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.CreateAudience(context.Background(), "my-project", &CreateAudienceRequest{
			Name: "Duplicate Audience",
			Key:  "duplicate-audience",
			Filters: Filters{
				Operator: "and",
				Filters:  []Filter{{Type: "all"}},
			},
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestClient_UpdateAudience(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {
		audience := AudienceDefinition{
			ID:          "aud-1",
			Key:         "beta-users",
			Name:        "Updated Audience",
			Description: "Updated description",
			Filters: Filters{
				Operator: "and",
				Filters: []Filter{
					{Type: "all"},
				},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/audiences/beta-users" {
				t.Errorf("expected /projects/my-project/audiences/beta-users, got %s", r.URL.Path)
			}

			var req UpdateAudienceRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Name != "Updated Audience" {
				t.Errorf("expected name 'Updated Audience', got %s", req.Name)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(audience)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.UpdateAudience(context.Background(), "my-project", "beta-users", &UpdateAudienceRequest{
			Name: "Updated Audience",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Name != "Updated Audience" {
			t.Errorf("expected Updated Audience, got %s", result.Name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("audience not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.UpdateAudience(context.Background(), "my-project", "non-existent", &UpdateAudienceRequest{
			Name: "Updated Audience",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_DeleteAudience(t *testing.T) {
	t.Run("successful delete", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/audiences/beta-users" {
				t.Errorf("expected /projects/my-project/audiences/beta-users, got %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		err := client.DeleteAudience(context.Background(), "my-project", "beta-users")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("audience not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		err := client.DeleteAudience(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}
