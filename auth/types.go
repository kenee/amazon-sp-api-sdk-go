package auth

import (
	"time"
)

// LWAAccessTokenRequestMeta represents the metadata for LWA access token request
type LWAAccessTokenRequestMeta struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope,omitempty"`
}

// LWAAccessTokenResponse represents the response from LWA token endpoint
type LWAAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope,omitempty"`
}

// LWAAuthorizationCredentials represents the credentials for LWA authorization
type LWAAuthorizationCredentials struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
	Endpoint     string
	Scopes       []string
}

// LWAAccessTokenCache represents the interface for caching access tokens
type LWAAccessTokenCache interface {
	Get(key string) (string, bool)
	Set(key, value string, ttl time.Duration)
	Delete(key string)
}

// RateLimitConfiguration represents rate limiting configuration
type RateLimitConfiguration struct {
	ID     string `json:"id"`
	Policy string `json:"policy"`
	Limit  int    `json:"limit"`
	Rate   struct {
		Interval string `json:"interval"`
		Amount   int    `json:"amount"`
	} `json:"rate"`
}

// ScopeConstants defines common scope constants
const (
	ScopeNotifications = "sellingpartnerapi::notifications"
	ScopeMigration     = "sellingpartnerapi::migration"
	ScopeOrders        = "sellingpartnerapi::orders"
	ScopeInventory     = "sellingpartnerapi::inventory"
	ScopeProducts      = "sellingpartnerapi::products"
	ScopePricing       = "sellingpartnerapi::pricing"
	ScopeFeeds         = "sellingpartnerapi::feeds"
	ScopeReports       = "sellingpartnerapi::reports"
	ScopeFinances      = "sellingpartnerapi::finances"
	ScopeShipping      = "sellingpartnerapi::shipping"
	ScopeMessaging     = "sellingpartnerapi::messaging"
	ScopeSolicitations = "sellingpartnerapi::solicitations"
)
