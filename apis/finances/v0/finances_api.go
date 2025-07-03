package v0

import (
	"context"
	"fmt"

	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// FinancesAPI represents the Finances API client
type FinancesAPI struct {
	apiClient *client.APIClient
}

// NewFinancesAPI creates a new Finances API client
func NewFinancesAPI(config *client.Configuration) *FinancesAPI {
	return &FinancesAPI{
		apiClient: client.NewAPIClient(config),
	}
}

// ListFinancialEventGroups retrieves financial event groups
func (f *FinancesAPI) ListFinancialEventGroups(ctx context.Context, request *ListFinancialEventGroupsRequest) (*ListFinancialEventGroupsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Optional parameters
	if request.MaxResultsPerPage > 0 {
		params["MaxResultsPerPage"] = request.MaxResultsPerPage
	}
	if request.FinancialEventGroupStartedBefore != "" {
		params["FinancialEventGroupStartedBefore"] = request.FinancialEventGroupStartedBefore
	}
	if request.FinancialEventGroupStartedAfter != "" {
		params["FinancialEventGroupStartedAfter"] = request.FinancialEventGroupStartedAfter
	}
	if request.NextToken != "" {
		params["NextToken"] = request.NextToken
	}

	// Build query string
	queryString := f.apiClient.BuildQueryString(params)

	// Build the full path
	path := "/finances/v0/financialEventGroups"
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call listFinancialEventGroups API: %w", err)
	}

	// Parse the response
	var result ListFinancialEventGroupsResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse listFinancialEventGroups response: %w", err)
	}

	return &result, nil
}

// ListFinancialEvents retrieves financial events
func (f *FinancesAPI) ListFinancialEvents(ctx context.Context, request *ListFinancialEventsRequest) (*ListFinancialEventsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Optional parameters
	if request.MaxResultsPerPage > 0 {
		params["MaxResultsPerPage"] = request.MaxResultsPerPage
	}
	if request.PostedAfter != "" {
		params["PostedAfter"] = request.PostedAfter
	}
	if request.PostedBefore != "" {
		params["PostedBefore"] = request.PostedBefore
	}
	if request.NextToken != "" {
		params["NextToken"] = request.NextToken
	}

	// Build query string
	queryString := f.apiClient.BuildQueryString(params)

	// Build the full path
	path := "/finances/v0/financialEvents"
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call listFinancialEvents API: %w", err)
	}

	// Parse the response
	var result ListFinancialEventsResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse listFinancialEvents response: %w", err)
	}

	return &result, nil
}

// ListFinancialEventsByGroupId retrieves financial events by group ID
func (f *FinancesAPI) ListFinancialEventsByGroupId(ctx context.Context, eventGroupId string, request *ListFinancialEventsByGroupIdRequest) (*ListFinancialEventsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Optional parameters
	if request.MaxResultsPerPage > 0 {
		params["MaxResultsPerPage"] = request.MaxResultsPerPage
	}
	if request.PostedAfter != "" {
		params["PostedAfter"] = request.PostedAfter
	}
	if request.PostedBefore != "" {
		params["PostedBefore"] = request.PostedBefore
	}
	if request.NextToken != "" {
		params["NextToken"] = request.NextToken
	}

	// Build query string
	queryString := f.apiClient.BuildQueryString(params)

	// Build the full path
	path := fmt.Sprintf("/finances/v0/financialEventGroups/%s/financialEvents", eventGroupId)
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call listFinancialEventsByGroupId API: %w", err)
	}

	// Parse the response
	var result ListFinancialEventsResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse listFinancialEventsByGroupId response: %w", err)
	}

	return &result, nil
}

// ListFinancialEventsByOrderId retrieves financial events by order ID
func (f *FinancesAPI) ListFinancialEventsByOrderId(ctx context.Context, orderId string, request *ListFinancialEventsByOrderIdRequest) (*ListFinancialEventsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Optional parameters
	if request.MaxResultsPerPage > 0 {
		params["MaxResultsPerPage"] = request.MaxResultsPerPage
	}
	if request.NextToken != "" {
		params["NextToken"] = request.NextToken
	}

	// Build query string
	queryString := f.apiClient.BuildQueryString(params)

	// Build the full path
	path := fmt.Sprintf("/finances/v0/orders/%s/financialEvents", orderId)
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call listFinancialEventsByOrderId API: %w", err)
	}

	// Parse the response
	var result ListFinancialEventsResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse listFinancialEventsByOrderId response: %w", err)
	}

	return &result, nil
}

// Helper methods for common use cases

// ListFinancialEventGroupsSimple is a simplified version for common use cases
func (f *FinancesAPI) ListFinancialEventGroupsSimple(ctx context.Context, maxResults int) (*ListFinancialEventGroupsResponse, error) {
	request := &ListFinancialEventGroupsRequest{
		MaxResultsPerPage: maxResults,
	}
	return f.ListFinancialEventGroups(ctx, request)
}

// ListFinancialEventsSimple is a simplified version for common use cases
func (f *FinancesAPI) ListFinancialEventsSimple(ctx context.Context, maxResults int) (*ListFinancialEventsResponse, error) {
	request := &ListFinancialEventsRequest{
		MaxResultsPerPage: maxResults,
	}
	return f.ListFinancialEvents(ctx, request)
}

// ListFinancialEventsByGroupIdSimple is a simplified version for common use cases
func (f *FinancesAPI) ListFinancialEventsByGroupIdSimple(ctx context.Context, eventGroupId string, maxResults int) (*ListFinancialEventsResponse, error) {
	request := &ListFinancialEventsByGroupIdRequest{
		MaxResultsPerPage: maxResults,
	}
	return f.ListFinancialEventsByGroupId(ctx, eventGroupId, request)
}

// ListFinancialEventsByOrderIdSimple is a simplified version for common use cases
func (f *FinancesAPI) ListFinancialEventsByOrderIdSimple(ctx context.Context, orderId string, maxResults int) (*ListFinancialEventsResponse, error) {
	request := &ListFinancialEventsByOrderIdRequest{
		MaxResultsPerPage: maxResults,
	}
	return f.ListFinancialEventsByOrderId(ctx, orderId, request)
}
