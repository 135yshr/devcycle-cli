package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_AuditLogs(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		logs := []AuditLog{
			{
				ID:   "log-1",
				Type: "modifiedFeature",
				User: AuditUser{
					Name:  "John Doe",
					Email: "john@example.com",
				},
				Changes: []Change{
					{
						Type:             "modifiedFeature",
						NewContents:      map[string]any{"key": "new-key"},
						PreviousContents: map[string]any{"key": "old-key"},
					},
				},
				CreatedAt: time.Now(),
			},
			{
				ID:   "log-2",
				Type: "addedVariable",
				User: AuditUser{
					Name:  "Jane Doe",
					Email: "jane@example.com",
				},
				Changes: []Change{
					{
						Type:             "addedVariable",
						NewContents:      map[string]any{"key": "new-var"},
						PreviousContents: nil,
					},
				},
				CreatedAt: time.Now(),
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/audit" {
				t.Errorf("expected /projects/my-project/audit, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(logs)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.AuditLogs(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 logs, got %d", len(result))
		}
		if result[0].Type != "modifiedFeature" {
			t.Errorf("expected modifiedFeature, got %s", result[0].Type)
		}
		if result[0].User.Name != "John Doe" {
			t.Errorf("expected John Doe, got %s", result[0].User.Name)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]AuditLog{})
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		result, err := client.AuditLogs(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected 0 logs, got %d", len(result))
		}
	})

	t.Run("server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.AuditLogs(context.Background(), "my-project")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestClient_FeatureAuditLogs(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		logs := []AuditLog{
			{
				ID:   "log-1",
				Type: "modifiedFeature",
				User: AuditUser{
					Name:  "John Doe",
					Email: "john@example.com",
				},
				Changes: []Change{
					{
						Type:             "modifiedFeature",
						NewContents:      map[string]any{"name": "Updated Feature"},
						PreviousContents: map[string]any{"name": "Old Feature"},
					},
				},
				CreatedAt: time.Now(),
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/features/my-feature/audit" {
				t.Errorf("expected /projects/my-project/features/my-feature/audit, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(logs)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.FeatureAuditLogs(context.Background(), "my-project", "my-feature")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 1 {
			t.Errorf("expected 1 log, got %d", len(result))
		}
		if result[0].Type != "modifiedFeature" {
			t.Errorf("expected modifiedFeature, got %s", result[0].Type)
		}
	})

	t.Run("feature not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("feature not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.FeatureAuditLogs(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}
