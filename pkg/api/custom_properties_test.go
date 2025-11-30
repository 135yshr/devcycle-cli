package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_CustomProperties(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		properties := []CustomProperty{
			{
				ID:          "cp-1",
				Key:         "user-type",
				PropertyKey: "user-type",
				DisplayName: "User Type",
				Type:        "String",
				Description: "Type of user",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				ID:          "cp-2",
				Key:         "is-premium",
				PropertyKey: "is-premium",
				DisplayName: "Is Premium",
				Type:        "Boolean",
				Description: "Whether user is premium",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/customProperties" {
				t.Errorf("expected /projects/my-project/customProperties, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(properties)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.CustomProperties(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 properties, got %d", len(result))
		}
		if result[0].Key != "user-type" {
			t.Errorf("expected user-type, got %s", result[0].Key)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]CustomProperty{})
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		result, err := client.CustomProperties(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected 0 properties, got %d", len(result))
		}
	})
}

func TestClient_CustomProperty(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		property := CustomProperty{
			ID:          "cp-1",
			Key:         "user-type",
			PropertyKey: "user-type",
			DisplayName: "User Type",
			Type:        "String",
			Description: "Type of user",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/customProperties/user-type" {
				t.Errorf("expected /projects/my-project/customProperties/user-type, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(property)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.CustomProperty(context.Background(), "my-project", "user-type")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "user-type" {
			t.Errorf("expected user-type, got %s", result.Key)
		}
		if result.DisplayName != "User Type" {
			t.Errorf("expected User Type, got %s", result.DisplayName)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("custom property not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.CustomProperty(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_CreateCustomProperty(t *testing.T) {
	t.Run("successful create", func(t *testing.T) {
		property := CustomProperty{
			ID:          "cp-new",
			Key:         "new-property",
			PropertyKey: "new-property",
			DisplayName: "New Property",
			Type:        "String",
			Description: "A new property",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/customProperties" {
				t.Errorf("expected /projects/my-project/customProperties, got %s", r.URL.Path)
			}

			var req CreateCustomPropertyRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Key != "new-property" {
				t.Errorf("expected key 'new-property', got %s", req.Key)
			}
			if req.DisplayName != "New Property" {
				t.Errorf("expected displayName 'New Property', got %s", req.DisplayName)
			}
			if req.Type != "String" {
				t.Errorf("expected type 'String', got %s", req.Type)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(property)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.CreateCustomProperty(context.Background(), "my-project", &CreateCustomPropertyRequest{
			Key:         "new-property",
			DisplayName: "New Property",
			Type:        "String",
			Description: "A new property",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "new-property" {
			t.Errorf("expected new-property, got %s", result.Key)
		}
	})

	t.Run("conflict error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("custom property already exists"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.CreateCustomProperty(context.Background(), "my-project", &CreateCustomPropertyRequest{
			Key:         "duplicate-property",
			DisplayName: "Duplicate Property",
			Type:        "String",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestClient_UpdateCustomProperty(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {
		property := CustomProperty{
			ID:          "cp-1",
			Key:         "user-type",
			PropertyKey: "user-type",
			DisplayName: "Updated Display Name",
			Type:        "String",
			Description: "Updated description",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/customProperties/user-type" {
				t.Errorf("expected /projects/my-project/customProperties/user-type, got %s", r.URL.Path)
			}

			var req UpdateCustomPropertyRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.DisplayName != "Updated Display Name" {
				t.Errorf("expected displayName 'Updated Display Name', got %s", req.DisplayName)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(property)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.UpdateCustomProperty(context.Background(), "my-project", "user-type", &UpdateCustomPropertyRequest{
			DisplayName: "Updated Display Name",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.DisplayName != "Updated Display Name" {
			t.Errorf("expected Updated Display Name, got %s", result.DisplayName)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("custom property not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.UpdateCustomProperty(context.Background(), "my-project", "non-existent", &UpdateCustomPropertyRequest{
			DisplayName: "Updated Display Name",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_DeleteCustomProperty(t *testing.T) {
	t.Run("successful delete", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/customProperties/user-type" {
				t.Errorf("expected /projects/my-project/customProperties/user-type, got %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		err := client.DeleteCustomProperty(context.Background(), "my-project", "user-type")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("custom property not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		err := client.DeleteCustomProperty(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}
