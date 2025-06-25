package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// APIClient represents the base API client
type APIClient struct {
	config     *Configuration
	httpClient *http.Client
}

// NewAPIClient creates a new API client
func NewAPIClient(config *Configuration) *APIClient {
	httpClient := config.GetHTTPClient()
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	return &APIClient{
		config:     config,
		httpClient: httpClient,
	}
}

// CallAPI makes an API call
func (c *APIClient) CallAPI(ctx context.Context, method, path string, body interface{}, headers map[string]string) (*http.Response, error) {
	// Build the full URL
	fullURL := c.config.GetHost() + path

	// Create request body
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, method, fullURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set default headers
	req.Header.Set("User-Agent", c.config.GetUserAgent())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Sign the request with LWA authorization
	if err := c.config.SignRequest(req); err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}

	// Make the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	return resp, nil
}

// CallAPIWithQueryParams makes an API call with query parameters
func (c *APIClient) CallAPIWithQueryParams(ctx context.Context, method, path string, body interface{}, headers map[string]string, queryParams map[string]string) (*http.Response, error) {
	// Build query string
	if len(queryParams) > 0 {
		values := url.Values{}
		for key, value := range queryParams {
			values.Set(key, value)
		}
		path += "?" + values.Encode()
	}

	return c.CallAPI(ctx, method, path, body, headers)
}

// CallAPIWithArrayQueryParams makes an API call with array query parameters
func (c *APIClient) CallAPIWithArrayQueryParams(ctx context.Context, method, path string, body interface{}, headers map[string]string, queryParams map[string][]string) (*http.Response, error) {
	// Build query string with array parameters
	if len(queryParams) > 0 {
		values := url.Values{}
		for key, valuesList := range queryParams {
			for _, value := range valuesList {
				values.Add(key, value)
			}
		}
		path += "?" + values.Encode()
	}

	return c.CallAPI(ctx, method, path, body, headers)
}

// ProcessResponse processes the API response
func (c *APIClient) ProcessResponse(resp *http.Response, target interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check if response is successful
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse response body
	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}

// BuildQueryString builds a query string from parameters
func (c *APIClient) BuildQueryString(params map[string]interface{}) string {
	values := url.Values{}
	for key, value := range params {
		switch v := value.(type) {
		case string:
			values.Set(key, v)
		case []string:
			for _, item := range v {
				values.Add(key, item)
			}
		case bool:
			if c.config.GetBooleanFormatForQuery() == "int" {
				if v {
					values.Set(key, "1")
				} else {
					values.Set(key, "0")
				}
			} else {
				values.Set(key, fmt.Sprintf("%t", v))
			}
		case int, int32, int64:
			values.Set(key, fmt.Sprintf("%d", v))
		case float32, float64:
			values.Set(key, fmt.Sprintf("%f", v))
		default:
			values.Set(key, fmt.Sprintf("%v", v))
		}
	}
	return values.Encode()
}

// GetConfig returns the configuration
func (c *APIClient) GetConfig() *Configuration {
	return c.config
}

// SetConfig sets the configuration
func (c *APIClient) SetConfig(config *Configuration) {
	c.config = config
}
