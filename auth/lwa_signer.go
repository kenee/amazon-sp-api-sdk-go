package auth

import (
	"fmt"
	"net/http"
	"strings"
)

// LWAAuthorizationSigner handles signing requests with LWA access tokens
type LWAAuthorizationSigner struct {
	credentials *LWAAuthorizationCredentials
	lwaClient   *LWAClient
}

// NewLWAAuthorizationSigner creates a new LWA authorization signer
func NewLWAAuthorizationSigner(credentials *LWAAuthorizationCredentials) *LWAAuthorizationSigner {
	lwaClient := NewLWAClient(credentials.Endpoint)
	return &LWAAuthorizationSigner{
		credentials: credentials,
		lwaClient:   lwaClient,
	}
}

// NewLWAAuthorizationSignerWithCache creates a new LWA authorization signer with cache
func NewLWAAuthorizationSignerWithCache(credentials *LWAAuthorizationCredentials, cache LWAAccessTokenCache) *LWAAuthorizationSigner {
	lwaClient := NewLWAClient(credentials.Endpoint)
	lwaClient.SetCache(cache)
	return &LWAAuthorizationSigner{
		credentials: credentials,
		lwaClient:   lwaClient,
	}
}

// Sign adds the authorization header to the request
func (l *LWAAuthorizationSigner) Sign(req *http.Request) error {
	accessToken, err := l.getAccessToken()
	if err != nil {
		return fmt.Errorf("failed to get access token: %w", err)
	}

	// Add the access token to the request headers
	req.Header.Set("x-amz-access-token", accessToken)

	return nil
}

// getAccessToken retrieves an access token using the credentials
func (l *LWAAuthorizationSigner) getAccessToken() (string, error) {
	meta := &LWAAccessTokenRequestMeta{
		GrantType:    "refresh_token",
		RefreshToken: l.credentials.RefreshToken,
		ClientID:     l.credentials.ClientID,
		ClientSecret: l.credentials.ClientSecret,
	}

	if len(l.credentials.Scopes) > 0 {
		meta.Scope = strings.Join(l.credentials.Scopes, " ")
	}

	return l.lwaClient.GetAccessToken(meta)
}

// GetLwaClient returns the LWA client
func (l *LWAAuthorizationSigner) GetLwaClient() *LWAClient {
	return l.lwaClient
}

// SetCache sets the token cache
func (l *LWAAuthorizationSigner) SetCache(cache LWAAccessTokenCache) {
	l.lwaClient.SetCache(cache)
}
