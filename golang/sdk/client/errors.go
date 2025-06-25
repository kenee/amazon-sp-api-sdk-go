package client

import (
	"fmt"
	"net/http"
)

// APIError represents an API error
type APIError struct {
	StatusCode int
	Message    string
	Body       string
	Headers    http.Header
}

// Error implements the error interface
func (e *APIError) Error() string {
	return fmt.Sprintf("API Error %d: %s", e.StatusCode, e.Message)
}

// RateLimitExceededError represents a rate limit exceeded error
type RateLimitExceededError struct {
	Message    string
	RetryAfter int
}

// Error implements the error interface
func (e *RateLimitExceededError) Error() string {
	return fmt.Sprintf("Rate limit exceeded: %s", e.Message)
}

// ConfigurationError represents a configuration error
type ConfigurationError struct {
	Message string
}

// Error implements the error interface
func (e *ConfigurationError) Error() string {
	return fmt.Sprintf("Configuration error: %s", e.Message)
}

// AuthenticationError represents an authentication error
type AuthenticationError struct {
	Message string
}

// Error implements the error interface
func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("Authentication error: %s", e.Message)
}

// Common error messages
const (
	ErrMissingCredentials = "missing required credentials"
	ErrInvalidHost        = "invalid host configuration"
	ErrInvalidEndpoint    = "invalid endpoint configuration"
	ErrRateLimitExceeded  = "rate limit exceeded"
	ErrUnauthorized       = "unauthorized access"
	ErrForbidden          = "forbidden access"
	ErrNotFound           = "resource not found"
	ErrInternalServer     = "internal server error"
	ErrBadRequest         = "bad request"
)
