package client

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/kenee/amazon-sp-api-sdk-go/auth"
)

// Configuration represents the SDK configuration
type Configuration struct {
	Host                   string
	UserAgent              string
	Debug                  bool
	DebugFile              string
	TempFolderPath         string
	LwaAuthSigner          *auth.LWAAuthorizationSigner
	HTTPClient             *http.Client
	BooleanFormatForQuery  string
	RateLimiterEnabled     bool
	RateLimitConfiguration map[string]interface{}
}

// NewConfiguration creates a new configuration
func NewConfiguration() *Configuration {
	return &Configuration{
		Host:                   "https://sellingpartnerapi-na.amazon.com",
		UserAgent:              fmt.Sprintf("amazon-selling-partner-api-sdk/1.0.0/Go (%s; %s)", runtime.GOOS, runtime.GOARCH),
		Debug:                  false,
		DebugFile:              "",
		TempFolderPath:         "",
		HTTPClient:             &http.Client{},
		BooleanFormatForQuery:  "string",
		RateLimiterEnabled:     true,
		RateLimitConfiguration: make(map[string]interface{}),
	}
}

// NewConfigurationWithCredentials creates a new configuration with LWA credentials
func NewConfigurationWithCredentials(credentials *auth.LWAAuthorizationCredentials) *Configuration {
	config := NewConfiguration()
	config.LwaAuthSigner = auth.NewLWAAuthorizationSigner(credentials)
	return config
}

// NewConfigurationWithCredentialsAndCache creates a new configuration with LWA credentials and cache
func NewConfigurationWithCredentialsAndCache(credentials *auth.LWAAuthorizationCredentials, cache auth.LWAAccessTokenCache) *Configuration {
	config := NewConfiguration()
	config.LwaAuthSigner = auth.NewLWAAuthorizationSignerWithCache(credentials, cache)
	return config
}

// NewConfigurationFromMap creates a new configuration from a map
func NewConfigurationFromMap(configMap map[string]interface{}) (*Configuration, error) {
	requiredFields := []string{"clientId", "clientSecret", "endpoint", "region"}
	missingFields := []string{}

	for _, field := range requiredFields {
		if _, exists := configMap[field]; !exists {
			missingFields = append(missingFields, field)
		}
	}

	if len(missingFields) > 0 {
		return nil, fmt.Errorf("missing required configuration fields: %s", strings.Join(missingFields, ", "))
	}

	credentials := &auth.LWAAuthorizationCredentials{
		ClientID:     configMap["clientId"].(string),
		ClientSecret: configMap["clientSecret"].(string),
		Endpoint:     configMap["endpoint"].(string),
		RefreshToken: getStringFromMap(configMap, "refreshToken"),
		Scopes:       getStringSliceFromMap(configMap, "scopes"),
	}

	config := NewConfigurationWithCredentials(credentials)

	// Set region-specific host
	if region, exists := configMap["region"]; exists {
		config.SetHostByRegion(region.(string))
	}

	// Set cache if provided
	if cache, exists := configMap["lwaTokenCache"]; exists {
		if tokenCache, ok := cache.(auth.LWAAccessTokenCache); ok {
			config.LwaAuthSigner.SetCache(tokenCache)
		}
	}

	return config, nil
}

// SetHost sets the API host
func (c *Configuration) SetHost(host string) {
	c.Host = host
}

// GetHost returns the API host
func (c *Configuration) GetHost() string {
	return c.Host
}

// SetHostByRegion sets the host based on region
func (c *Configuration) SetHostByRegion(region string) {
	switch strings.ToLower(region) {
	case "na":
		c.Host = "https://sellingpartnerapi-na.amazon.com"
	case "eu":
		c.Host = "https://sellingpartnerapi-eu.amazon.com"
	case "fe":
		c.Host = "https://sellingpartnerapi-fe.amazon.com"
	default:
		c.Host = "https://sellingpartnerapi-na.amazon.com"
	}
}

// SetUserAgent sets the user agent
func (c *Configuration) SetUserAgent(userAgent string) {
	c.UserAgent = userAgent
}

// GetUserAgent returns the user agent
func (c *Configuration) GetUserAgent() string {
	return c.UserAgent
}

// SetDebug sets debug mode
func (c *Configuration) SetDebug(debug bool) {
	c.Debug = debug
}

// IsDebug returns debug mode status
func (c *Configuration) IsDebug() bool {
	return c.Debug
}

// SetHTTPClient sets the HTTP client
func (c *Configuration) SetHTTPClient(client *http.Client) {
	c.HTTPClient = client
}

// GetHTTPClient returns the HTTP client
func (c *Configuration) GetHTTPClient() *http.Client {
	return c.HTTPClient
}

// SetBooleanFormatForQuery sets the boolean format for query strings
func (c *Configuration) SetBooleanFormatForQuery(format string) {
	c.BooleanFormatForQuery = format
}

// GetBooleanFormatForQuery returns the boolean format for query strings
func (c *Configuration) GetBooleanFormatForQuery() string {
	return c.BooleanFormatForQuery
}

// SetRateLimiterEnabled sets whether rate limiter is enabled
func (c *Configuration) SetRateLimiterEnabled(enabled bool) {
	c.RateLimiterEnabled = enabled
}

// IsRateLimiterEnabled returns whether rate limiter is enabled
func (c *Configuration) IsRateLimiterEnabled() bool {
	return c.RateLimiterEnabled
}

// SignRequest signs a request with LWA authorization
func (c *Configuration) SignRequest(req *http.Request) error {
	if c.LwaAuthSigner == nil {
		return fmt.Errorf("LWA authorization signer not configured")
	}
	return c.LwaAuthSigner.Sign(req)
}

// EnableCache enables token caching
func (c *Configuration) EnableCache(cache auth.LWAAccessTokenCache) {
	if c.LwaAuthSigner != nil {
		c.LwaAuthSigner.SetCache(cache)
	}
}

// DisableCache disables token caching
func (c *Configuration) DisableCache() {
	if c.LwaAuthSigner != nil {
		c.LwaAuthSigner.SetCache(nil)
	}
}

// Helper functions
func getStringFromMap(m map[string]interface{}, key string) string {
	if val, exists := m[key]; exists {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func getStringSliceFromMap(m map[string]interface{}, key string) []string {
	if val, exists := m[key]; exists {
		if slice, ok := val.([]string); ok {
			return slice
		}
		if slice, ok := val.([]interface{}); ok {
			result := make([]string, len(slice))
			for i, v := range slice {
				if str, ok := v.(string); ok {
					result[i] = str
				}
			}
			return result
		}
	}
	return nil
}
