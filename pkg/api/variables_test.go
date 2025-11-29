package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_Variables(t *testing.T) {
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
		result, err := client.Variables(context.Background(), "my-project")

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

func TestClient_Variable(t *testing.T) {
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
		result, err := client.Variable(context.Background(), "my-project", "enable-feature")

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
		_, err := client.Variable(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_CreateVariable(t *testing.T) {
	t.Run("successful create", func(t *testing.T) {
		variable := Variable{
			ID:        "var-new",
			Key:       "new-variable",
			Name:      "New Variable",
			Type:      "Boolean",
			Status:    "active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/variables" {
				t.Errorf("expected /projects/my-project/variables, got %s", r.URL.Path)
			}

			var req CreateVariableRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Name != "New Variable" {
				t.Errorf("expected name 'New Variable', got %s", req.Name)
			}
			if req.Key != "new-variable" {
				t.Errorf("expected key 'new-variable', got %s", req.Key)
			}
			if req.Type != "Boolean" {
				t.Errorf("expected type 'Boolean', got %s", req.Type)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(variable)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.CreateVariable(context.Background(), "my-project", &CreateVariableRequest{
			Name: "New Variable",
			Key:  "new-variable",
			Type: "Boolean",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "new-variable" {
			t.Errorf("expected new-variable, got %s", result.Key)
		}
	})

	t.Run("conflict error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("variable already exists"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.CreateVariable(context.Background(), "my-project", &CreateVariableRequest{
			Name: "Duplicate Variable",
			Key:  "duplicate-variable",
			Type: "Boolean",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestClient_UpdateVariable(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {
		variable := Variable{
			ID:        "var-1",
			Key:       "enable-feature",
			Name:      "Updated Variable",
			Type:      "Boolean",
			Status:    "active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/variables/enable-feature" {
				t.Errorf("expected /projects/my-project/variables/enable-feature, got %s", r.URL.Path)
			}

			var req UpdateVariableRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Name != "Updated Variable" {
				t.Errorf("expected name 'Updated Variable', got %s", req.Name)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(variable)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.UpdateVariable(context.Background(), "my-project", "enable-feature", &UpdateVariableRequest{
			Name: "Updated Variable",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Name != "Updated Variable" {
			t.Errorf("expected Updated Variable, got %s", result.Name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("variable not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.UpdateVariable(context.Background(), "my-project", "non-existent", &UpdateVariableRequest{
			Name: "Updated Variable",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_DeleteVariable(t *testing.T) {
	t.Run("successful delete", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/variables/enable-feature" {
				t.Errorf("expected /projects/my-project/variables/enable-feature, got %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		err := client.DeleteVariable(context.Background(), "my-project", "enable-feature")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("variable not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		err := client.DeleteVariable(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}
