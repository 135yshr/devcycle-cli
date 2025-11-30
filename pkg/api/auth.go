package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// TokenResponse represents the raw OAuth2 token response from the DevCycle authentication endpoint.
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// Token represents an authenticated session with DevCycle.
// It includes the access token, token type, and expiration time.
type Token struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresAt   time.Time `json:"expires_at"`
}

// IsExpired reports whether the token has expired based on the current time.
func (t *Token) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}

// Authenticate obtains an OAuth2 access token from DevCycle using client credentials.
// The clientID and clientSecret can be obtained from the DevCycle dashboard.
// Returns a Token containing the access token and expiration information.
func Authenticate(ctx context.Context, clientID, clientSecret string) (*Token, error) {
	return authenticateInternal(ctx, AuthURL, clientID, clientSecret)
}

func authenticateInternal(ctx context.Context, authURL, clientID, clientSecret string) (*Token, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("audience", "https://api.devcycle.com/")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, authURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create auth request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: DefaultTimeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("auth request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read auth response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(body),
		}
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w", err)
	}

	token := &Token{
		AccessToken: tokenResp.AccessToken,
		TokenType:   tokenResp.TokenType,
		ExpiresAt:   time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second),
	}

	return token, nil
}

// SaveToken writes the token to a file at the specified path.
// The token is stored as JSON with indented formatting.
func SaveToken(token *Token, path string) error {
	data, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal token: %w", err)
	}

	var buf bytes.Buffer
	buf.Write(data)

	return writeFile(path, buf.Bytes())
}

// LoadToken reads a token from a file at the specified path.
// Returns an error if the file does not exist or contains invalid JSON.
func LoadToken(path string) (*Token, error) {
	data, err := readFile(path)
	if err != nil {
		return nil, err
	}

	var token Token
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, fmt.Errorf("failed to parse token file: %w", err)
	}

	return &token, nil
}
