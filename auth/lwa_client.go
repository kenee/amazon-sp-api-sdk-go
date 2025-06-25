package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// LWAClient handles Login with Amazon (LWA) authentication
type LWAClient struct {
	client   *http.Client
	endpoint string
	cache    LWAAccessTokenCache
}

// NewLWAClient creates a new LWA client
func NewLWAClient(endpoint string) *LWAClient {
	return &LWAClient{
		client:   &http.Client{Timeout: 30 * time.Second},
		endpoint: endpoint,
	}
}

// SetCache sets the token cache
func (l *LWAClient) SetCache(cache LWAAccessTokenCache) {
	l.cache = cache
}

// GetAccessToken retrieves an access token, using cache if available
func (l *LWAClient) GetAccessToken(meta *LWAAccessTokenRequestMeta) (string, error) {
	if l.cache != nil {
		return l.getAccessTokenFromCache(meta)
	}
	return l.getAccessTokenFromEndpoint(meta)
}

// getAccessTokenFromCache attempts to get token from cache
func (l *LWAClient) getAccessTokenFromCache(meta *LWAAccessTokenRequestMeta) (string, error) {
	key := l.generateCacheKey(meta)
	if token, found := l.cache.Get(key); found {
		return token, nil
	}

	// If not in cache, get from endpoint and cache it
	token, err := l.getAccessTokenFromEndpoint(meta)
	if err != nil {
		return "", err
	}

	// Cache the token with TTL based on expires_in
	// We'll set a default TTL of 1 hour if we can't determine it
	ttl := time.Hour
	l.cache.Set(key, token, ttl)

	return token, nil
}

// getAccessTokenFromEndpoint retrieves token directly from LWA endpoint
func (l *LWAClient) getAccessTokenFromEndpoint(meta *LWAAccessTokenRequestMeta) (string, error) {
	requestBody, err := json.Marshal(meta)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", l.endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := l.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unsuccessful LWA token exchange: status=%d, body=%s", resp.StatusCode, string(bodyBytes))
	}

	var tokenResp LWAAccessTokenResponse
	if err := json.Unmarshal(bodyBytes, &tokenResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("response did not have required access_token")
	}

	return tokenResp.AccessToken, nil
}

// generateCacheKey creates a unique cache key for the request
func (l *LWAClient) generateCacheKey(meta *LWAAccessTokenRequestMeta) string {
	// Create a simple key based on client_id and scope
	key := fmt.Sprintf("%s:%s", meta.ClientID, meta.Scope)
	return key
}

// SetClient allows setting a custom HTTP client
func (l *LWAClient) SetClient(client *http.Client) {
	l.client = client
}

// GetEndpoint returns the LWA endpoint
func (l *LWAClient) GetEndpoint() string {
	return l.endpoint
}
