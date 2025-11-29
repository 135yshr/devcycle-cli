package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	t.Run("default options", func(t *testing.T) {
		client := NewClient()

		if client.baseURL != DefaultBaseURL {
			t.Errorf("expected baseURL %s, got %s", DefaultBaseURL, client.baseURL)
		}
		if client.token != "" {
			t.Errorf("expected empty token, got %s", client.token)
		}
	})

	t.Run("with custom options", func(t *testing.T) {
		customURL := "https://custom.api.com"
		customToken := "test-token"

		client := NewClient(
			WithBaseURL(customURL),
			WithToken(customToken),
			WithTimeout(60*time.Second),
		)

		if client.baseURL != customURL {
			t.Errorf("expected baseURL %s, got %s", customURL, client.baseURL)
		}
		if client.token != customToken {
			t.Errorf("expected token %s, got %s", customToken, client.token)
		}
	})
}

func TestClient_SetToken(t *testing.T) {
	client := NewClient()
	token := "new-token"

	client.SetToken(token)

	if client.token != token {
		t.Errorf("expected token %s, got %s", token, client.token)
	}
}

func TestClient_Get(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		expected := map[string]string{"key": "value"}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET method, got %s", r.Method)
			}
			if r.Header.Get("Authorization") != "Bearer test-token" {
				t.Errorf("expected Bearer token, got %s", r.Header.Get("Authorization"))
			}
			if r.Header.Get("Content-Type") != "application/json" {
				t.Errorf("expected application/json, got %s", r.Header.Get("Content-Type"))
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(expected)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))

		var result map[string]string
		err := client.Get(context.Background(), "/test", &result)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result["key"] != expected["key"] {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("error response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))

		var result map[string]string
		err := client.Get(context.Background(), "/test", &result)

		if err == nil {
			t.Fatal("expected error, got nil")
		}

		apiErr, ok := err.(*APIError)
		if !ok {
			t.Fatalf("expected APIError, got %T", err)
		}
		if apiErr.StatusCode != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", apiErr.StatusCode)
		}
	})

	t.Run("context cancellation", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		var result map[string]string
		err := client.Get(ctx, "/test", &result)

		if err == nil {
			t.Fatal("expected error due to cancelled context")
		}
	})
}

func TestClient_Post(t *testing.T) {
	t.Run("successful request with body", func(t *testing.T) {
		requestBody := map[string]string{"name": "test"}
		responseBody := map[string]string{"id": "123", "name": "test"}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST method, got %s", r.Method)
			}

			var body map[string]string
			json.NewDecoder(r.Body).Decode(&body)

			if body["name"] != requestBody["name"] {
				t.Errorf("expected body %v, got %v", requestBody, body)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseBody)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))

		var result map[string]string
		err := client.Post(context.Background(), "/test", requestBody, &result)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result["id"] != responseBody["id"] {
			t.Errorf("expected %v, got %v", responseBody, result)
		}
	})
}

func TestClient_Patch(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH method, got %s", r.Method)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"updated": "true"})
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))

		var result map[string]string
		err := client.Patch(context.Background(), "/test", map[string]string{"key": "value"}, &result)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestClient_Delete(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE method, got %s", r.Method)
			}
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))

		err := client.Delete(context.Background(), "/test")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
