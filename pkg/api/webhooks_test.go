package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_Webhooks(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		webhooks := []Webhook{
			{
				ID:          "wh-1",
				URL:         "https://example.com/webhook1",
				Description: "First webhook",
				IsEnabled:   true,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				ID:          "wh-2",
				URL:         "https://example.com/webhook2",
				Description: "Second webhook",
				IsEnabled:   false,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/webhooks" {
				t.Errorf("expected /projects/my-project/webhooks, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(webhooks)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.Webhooks(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 webhooks, got %d", len(result))
		}
		if result[0].URL != "https://example.com/webhook1" {
			t.Errorf("expected https://example.com/webhook1, got %s", result[0].URL)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]Webhook{})
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		result, err := client.Webhooks(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected 0 webhooks, got %d", len(result))
		}
	})
}

func TestClient_Webhook(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		webhook := Webhook{
			ID:          "wh-1",
			URL:         "https://example.com/webhook1",
			Description: "First webhook",
			IsEnabled:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/webhooks/wh-1" {
				t.Errorf("expected /projects/my-project/webhooks/wh-1, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(webhook)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.Webhook(context.Background(), "my-project", "wh-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.ID != "wh-1" {
			t.Errorf("expected wh-1, got %s", result.ID)
		}
		if result.URL != "https://example.com/webhook1" {
			t.Errorf("expected https://example.com/webhook1, got %s", result.URL)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("webhook not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.Webhook(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_CreateWebhook(t *testing.T) {
	t.Run("successful create", func(t *testing.T) {
		webhook := Webhook{
			ID:          "wh-new",
			URL:         "https://example.com/new-webhook",
			Description: "A new webhook",
			IsEnabled:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/webhooks" {
				t.Errorf("expected /projects/my-project/webhooks, got %s", r.URL.Path)
			}

			var req CreateWebhookRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.URL != "https://example.com/new-webhook" {
				t.Errorf("expected URL 'https://example.com/new-webhook', got %s", req.URL)
			}
			if !req.IsEnabled {
				t.Errorf("expected IsEnabled true, got false")
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(webhook)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.CreateWebhook(context.Background(), "my-project", &CreateWebhookRequest{
			URL:         "https://example.com/new-webhook",
			Description: "A new webhook",
			IsEnabled:   true,
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.URL != "https://example.com/new-webhook" {
			t.Errorf("expected https://example.com/new-webhook, got %s", result.URL)
		}
	})

	t.Run("bad request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid URL"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.CreateWebhook(context.Background(), "my-project", &CreateWebhookRequest{
			URL:       "invalid-url",
			IsEnabled: true,
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestClient_UpdateWebhook(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {
		webhook := Webhook{
			ID:          "wh-1",
			URL:         "https://example.com/updated-webhook",
			Description: "Updated description",
			IsEnabled:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/webhooks/wh-1" {
				t.Errorf("expected /projects/my-project/webhooks/wh-1, got %s", r.URL.Path)
			}

			var req UpdateWebhookRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.URL != "https://example.com/updated-webhook" {
				t.Errorf("expected URL 'https://example.com/updated-webhook', got %s", req.URL)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(webhook)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.UpdateWebhook(context.Background(), "my-project", "wh-1", &UpdateWebhookRequest{
			URL: "https://example.com/updated-webhook",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.URL != "https://example.com/updated-webhook" {
			t.Errorf("expected https://example.com/updated-webhook, got %s", result.URL)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("webhook not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.UpdateWebhook(context.Background(), "my-project", "non-existent", &UpdateWebhookRequest{
			URL: "https://example.com/updated-webhook",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_DeleteWebhook(t *testing.T) {
	t.Run("successful delete", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/webhooks/wh-1" {
				t.Errorf("expected /projects/my-project/webhooks/wh-1, got %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		err := client.DeleteWebhook(context.Background(), "my-project", "wh-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("webhook not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		err := client.DeleteWebhook(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}
