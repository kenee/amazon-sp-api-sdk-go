package client

import (
	"fmt"
	"net/http"
	"strings"
)

// APIError represents an API error
type APIError struct {
	StatusCode int
	Message    string
	Body       string
	Headers    http.Header
	RequestID  string
}

// Error implements the error interface
func (e *APIError) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("API Error %d (RequestID: %s): %s", e.StatusCode, e.RequestID, e.Message)
	}
	return fmt.Sprintf("API Error %d: %s", e.StatusCode, e.Message)
}

// IsRetryable returns true if the error is retryable
func (e *APIError) IsRetryable() bool {
	// 5xx errors are retryable
	if e.StatusCode >= 500 && e.StatusCode < 600 {
		return true
	}
	// 429 (Too Many Requests) is retryable
	if e.StatusCode == 429 {
		return true
	}
	// 408 (Request Timeout) is retryable
	if e.StatusCode == 408 {
		return true
	}
	return false
}

// RateLimitExceededError represents a rate limit exceeded error
type RateLimitExceededError struct {
	Message    string
	RetryAfter int
	RequestID  string
}

// Error implements the error interface
func (e *RateLimitExceededError) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("Rate limit exceeded (RequestID: %s): %s", e.RequestID, e.Message)
	}
	return fmt.Sprintf("Rate limit exceeded: %s", e.Message)
}

// ConfigurationError represents a configuration error
type ConfigurationError struct {
	Message string
	Field   string
}

// Error implements the error interface
func (e *ConfigurationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("Configuration error in field '%s': %s", e.Field, e.Message)
	}
	return fmt.Sprintf("Configuration error: %s", e.Message)
}

// AuthenticationError represents an authentication error
type AuthenticationError struct {
	Message   string
	RequestID string
}

// Error implements the error interface
func (e *AuthenticationError) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("Authentication error (RequestID: %s): %s", e.RequestID, e.Message)
	}
	return fmt.Sprintf("Authentication error: %s", e.Message)
}

// ValidationError represents a validation error
type ValidationError struct {
	Message      string
	Field        string
	Value        interface{}
	RequestID    string
	IsSandboxEnv bool
	Suggestions  []string
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	var sb strings.Builder

	// 基础错误信息
	if e.RequestID != "" {
		sb.WriteString(fmt.Sprintf("Validation error (RequestID: %s): ", e.RequestID))
	} else {
		sb.WriteString("Validation error: ")
	}

	// 字段信息
	if e.Field != "" {
		sb.WriteString(fmt.Sprintf("field '%s' with value '%v' - ", e.Field, e.Value))
	} else {
		sb.WriteString(fmt.Sprintf("field '' with value '%v' - ", e.Value))
	}

	sb.WriteString(e.Message)

	// 沙盒环境特殊提示
	if e.IsSandboxEnv {
		sb.WriteString(" (This may be a sandbox environment limitation)")
	}

	// 建议
	if len(e.Suggestions) > 0 {
		sb.WriteString("\nSuggestions:")
		for _, suggestion := range e.Suggestions {
			sb.WriteString(fmt.Sprintf("\n  - %s", suggestion))
		}
	}

	return sb.String()
}

// IsSandboxError returns true if this is likely a sandbox environment error
func (e *ValidationError) IsSandboxError() bool {
	return e.IsSandboxEnv
}

// NetworkError represents a network-related error
type NetworkError struct {
	Message string
	Cause   error
}

// Error implements the error interface
func (e *NetworkError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("Network error: %s (caused by: %v)", e.Message, e.Cause)
	}
	return fmt.Sprintf("Network error: %s", e.Message)
}

// Unwrap returns the underlying error
func (e *NetworkError) Unwrap() error {
	return e.Cause
}

// Common error messages
const (
	ErrMissingCredentials   = "missing required credentials"
	ErrInvalidHost          = "invalid host configuration"
	ErrInvalidEndpoint      = "invalid endpoint configuration"
	ErrRateLimitExceeded    = "rate limit exceeded"
	ErrUnauthorized         = "unauthorized access"
	ErrForbidden            = "forbidden access"
	ErrNotFound             = "resource not found"
	ErrInternalServer       = "internal server error"
	ErrBadRequest           = "bad request"
	ErrTimeout              = "request timeout"
	ErrConnectionFailed     = "connection failed"
	ErrInvalidResponse      = "invalid response format"
	ErrMissingRequiredField = "missing required field"
	ErrInvalidFieldValue    = "invalid field value"
	ErrSandboxLimitation    = "sandbox environment limitation"
)

// ParseAPIError parses an API error from response
func ParseAPIError(statusCode int, body string, headers http.Header) error {
	requestID := headers.Get("x-amzn-RequestId")

	// 检测是否为沙盒环境
	isSandboxEnv := isSandboxEnvironment(headers)

	// 如果无法从头部检测，但错误模式符合沙盒特征，也标记为沙盒
	if !isSandboxEnv && statusCode == 400 {
		if strings.Contains(body, `field '' with value '<nil>'`) {
			isSandboxEnv = true
		}
	}

	// Create base API error
	apiError := &APIError{
		StatusCode: statusCode,
		Body:       body,
		Headers:    headers,
		RequestID:  requestID,
	}

	// Set specific error message based on status code
	switch statusCode {
	case 400:
		apiError.Message = ErrBadRequest

		// 尝试从响应体中提取更详细的错误信息
		field, value, message := extractValidationDetails(body)

		suggestions := []string{
			"Check that all required parameters are provided",
			"Verify parameter formats (dates, IDs, etc.)",
			"Ensure marketplace IDs are correct",
			"Check seller ID and permissions",
		}

		// 如果是沙盒环境，添加特殊建议
		if isSandboxEnv {
			suggestions = append(suggestions,
				"This may be a sandbox environment limitation",
				"Try testing in production environment",
				"Check if the requested data exists in sandbox",
			)
		}

		return &ValidationError{
			Message:      message,
			Field:        field,
			Value:        value,
			RequestID:    requestID,
			IsSandboxEnv: isSandboxEnv,
			Suggestions:  suggestions,
		}
	case 401:
		apiError.Message = ErrUnauthorized
		return &AuthenticationError{
			Message:   "Authentication failed - check your credentials",
			RequestID: requestID,
		}
	case 403:
		apiError.Message = ErrForbidden
		return &AuthenticationError{
			Message:   "Access forbidden - check your permissions",
			RequestID: requestID,
		}
	case 404:
		apiError.Message = ErrNotFound
		return apiError
	case 408:
		apiError.Message = ErrTimeout
		return apiError
	case 429:
		apiError.Message = ErrRateLimitExceeded
		retryAfter := 0
		if retryAfterStr := headers.Get("Retry-After"); retryAfterStr != "" {
			fmt.Sscanf(retryAfterStr, "%d", &retryAfter)
		}
		return &RateLimitExceededError{
			Message:    "Rate limit exceeded - please retry later",
			RetryAfter: retryAfter,
			RequestID:  requestID,
		}
	case 500, 502, 503, 504:
		apiError.Message = ErrInternalServer
		return apiError
	default:
		apiError.Message = fmt.Sprintf("Unexpected error with status code %d", statusCode)
		return apiError
	}
}

// isSandboxEnvironment checks if the request is going to sandbox environment
func isSandboxEnvironment(headers http.Header) bool {
	// 检查常见的沙盒环境标识
	sandboxIndicators := []string{
		"sandbox",
		"test",
		"dev",
	}

	// 检查各种可能包含环境信息的头部
	for _, header := range []string{"x-amz-cf-id", "x-amz-id-2", "server", "via", "x-amzn-requestid"} {
		value := strings.ToLower(headers.Get(header))
		for _, indicator := range sandboxIndicators {
			if strings.Contains(value, indicator) {
				return true
			}
		}
	}

	// 检查环境变量中的端点信息
	if host := headers.Get("host"); host != "" {
		if strings.Contains(strings.ToLower(host), "sandbox") {
			return true
		}
	}

	// 如果无法从头部确定，但错误模式符合沙盒环境特征，也标记为沙盒
	// 沙盒环境通常返回 field '' with value '<nil>' 的错误
	return false
}

// extractValidationDetails attempts to extract field and value information from error response
func extractValidationDetails(body string) (field, value, message string) {
	// 默认值
	field = ""
	value = "<nil>"
	message = "Bad request - check your input parameters"

	// 尝试解析JSON响应
	if strings.Contains(body, `"errors"`) {
		// 这里可以添加JSON解析逻辑来提取更详细的错误信息
		// 由于Amazon SP-API的错误响应格式可能变化，我们保持简单的字符串匹配
		if strings.Contains(body, `"InvalidInput"`) {
			message = "Invalid input parameters - check your request format"
		} else if strings.Contains(body, `"MissingRequiredField"`) {
			message = "Missing required field - check your request parameters"
		} else if strings.Contains(body, `"InvalidFieldValue"`) {
			message = "Invalid field value - check your parameter values"
		}
	}

	// 尝试从错误消息中提取字段名
	if strings.Contains(body, "field") {
		// 简单的字符串匹配来提取字段名
		// 这里可以根据实际的错误响应格式进行优化
	}

	return field, value, message
}
