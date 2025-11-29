package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestToken_IsExpired(t *testing.T) {
	tests := []struct {
		name      string
		expiresAt time.Time
		expected  bool
	}{
		{
			name:      "expired token",
			expiresAt: time.Now().Add(-1 * time.Hour),
			expected:  true,
		},
		{
			name:      "valid token",
			expiresAt: time.Now().Add(1 * time.Hour),
			expected:  false,
		},
		{
			name:      "just expired",
			expiresAt: time.Now().Add(-1 * time.Second),
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := &Token{
				AccessToken: "test-token",
				TokenType:   "Bearer",
				ExpiresAt:   tt.expiresAt,
			}

			result := token.IsExpired()
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAuthenticate(t *testing.T) {
	t.Run("successful authentication", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST method, got %s", r.Method)
			}

			if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
				t.Errorf("expected form content type, got %s", r.Header.Get("Content-Type"))
			}

			r.ParseForm()
			if r.Form.Get("grant_type") != "client_credentials" {
				t.Errorf("expected client_credentials grant type")
			}
			if r.Form.Get("client_id") != "test-client-id" {
				t.Errorf("expected test-client-id, got %s", r.Form.Get("client_id"))
			}
			if r.Form.Get("client_secret") != "test-client-secret" {
				t.Errorf("expected test-client-secret, got %s", r.Form.Get("client_secret"))
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(TokenResponse{
				AccessToken: "access-token-123",
				TokenType:   "Bearer",
				ExpiresIn:   3600,
			})
		}))
		defer server.Close()

		originalAuthURL := AuthURL
		defer func() {
			// Note: Cannot modify const, this test uses the real AuthURL
			// In production, we'd use dependency injection
			_ = originalAuthURL
		}()

		// For this test, we'll create a custom authenticate function
		token, err := authenticateWithURL(context.Background(), server.URL, "test-client-id", "test-client-secret")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if token.AccessToken != "access-token-123" {
			t.Errorf("expected access-token-123, got %s", token.AccessToken)
		}
		if token.TokenType != "Bearer" {
			t.Errorf("expected Bearer, got %s", token.TokenType)
		}
		if token.IsExpired() {
			t.Error("token should not be expired")
		}
	})

	t.Run("authentication failure", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "invalid_client"}`))
		}))
		defer server.Close()

		_, err := authenticateWithURL(context.Background(), server.URL, "bad-id", "bad-secret")
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		apiErr, ok := err.(*APIError)
		if !ok {
			t.Fatalf("expected APIError, got %T", err)
		}
		if apiErr.StatusCode != http.StatusUnauthorized {
			t.Errorf("expected status 401, got %d", apiErr.StatusCode)
		}
	})
}

// authenticateWithURL is a helper for testing with custom auth URL
func authenticateWithURL(ctx context.Context, authURL, clientID, clientSecret string) (*Token, error) {
	return authenticateInternal(ctx, authURL, clientID, clientSecret)
}

func TestSaveAndLoadToken(t *testing.T) {
	tempDir := t.TempDir()
	tokenPath := filepath.Join(tempDir, "token.json")

	originalToken := &Token{
		AccessToken: "test-access-token",
		TokenType:   "Bearer",
		ExpiresAt:   time.Now().Add(1 * time.Hour).Truncate(time.Second),
	}

	t.Run("save token", func(t *testing.T) {
		err := SaveToken(originalToken, tokenPath)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if _, err := os.Stat(tokenPath); os.IsNotExist(err) {
			t.Fatal("token file was not created")
		}
	})

	t.Run("load token", func(t *testing.T) {
		loadedToken, err := LoadToken(tokenPath)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if loadedToken.AccessToken != originalToken.AccessToken {
			t.Errorf("expected %s, got %s", originalToken.AccessToken, loadedToken.AccessToken)
		}
		if loadedToken.TokenType != originalToken.TokenType {
			t.Errorf("expected %s, got %s", originalToken.TokenType, loadedToken.TokenType)
		}
		if !loadedToken.ExpiresAt.Equal(originalToken.ExpiresAt) {
			t.Errorf("expected %v, got %v", originalToken.ExpiresAt, loadedToken.ExpiresAt)
		}
	})

	t.Run("load non-existent token", func(t *testing.T) {
		_, err := LoadToken(filepath.Join(tempDir, "non-existent.json"))
		if err == nil {
			t.Fatal("expected error for non-existent file")
		}
	})
}
