package listings

import (
	"context"
	"fmt"

	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// ListingsAPI represents the Listings API client
type ListingsAPI struct {
	apiClient *client.APIClient
}

// NewListingsAPI creates a new Listings API client
func NewListingsAPI(config *client.Configuration) *ListingsAPI {
	return &ListingsAPI{
		apiClient: client.NewAPIClient(config),
	}
}

// GetListingsItem retrieves a listings item by SKU
func (l *ListingsAPI) GetListingsItem(ctx context.Context, request *GetListingsItemRequest) (*GetListingsItemResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	params["marketplaceIds"] = request.MarketplaceIds

	// Optional parameters
	if request.IssueLocale != "" {
		params["issueLocale"] = request.IssueLocale
	}
	if len(request.IncludedData) > 0 {
		params["includedData"] = request.IncludedData
	}

	// Build query string
	queryString := l.apiClient.BuildQueryString(params)

	// Build the full path
	path := fmt.Sprintf("/listings/2021-08-01/items/%s/%s", request.SellerId, request.SKU)
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := l.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getListingsItem API: %w", err)
	}

	// Parse the response
	var result GetListingsItemResponse
	if err := l.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getListingsItem response: %w", err)
	}

	return &result, nil
}

// GetListingsItemSimple is a simplified version of GetListingsItem for common use cases
func (l *ListingsAPI) GetListingsItemSimple(ctx context.Context, sellerId, sku string, marketplaceIds []string) (*GetListingsItemResponse, error) {
	request := &GetListingsItemRequest{
		SellerId:       sellerId,
		SKU:            sku,
		MarketplaceIds: marketplaceIds,
		IssueLocale:    "en_US",
		IncludedData:   []string{"summaries", "offers", "fulfillmentAvailability", "issues"},
	}
	return l.GetListingsItem(ctx, request)
}

// GetListingsRestrictions retrieves listing restrictions for an ASIN
func (l *ListingsAPI) GetListingsRestrictions(ctx context.Context, request *GetListingsRestrictionsRequest) (*GetListingsRestrictionsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	params["asin"] = request.ASIN
	params["sellerId"] = request.SellerId
	params["marketplaceIds"] = request.MarketplaceIds

	// Optional parameters
	if request.ConditionType != "" {
		params["conditionType"] = request.ConditionType
	}
	if request.IssueLocale != "" {
		params["reasonLocale"] = request.IssueLocale
	}

	// Build query string
	queryString := l.apiClient.BuildQueryString(params)

	// Build the full path - 参考PHP SDK的正确路径
	path := "/listings/2021-08-01/restrictions"
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := l.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getListingsRestrictions API: %w", err)
	}

	// Parse the response
	var result GetListingsRestrictionsResponse
	if err := l.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getListingsRestrictions response: %w", err)
	}

	return &result, nil
}

// PutListingsItem creates or updates a listings item
func (l *ListingsAPI) PutListingsItem(ctx context.Context, request *PutListingsItemRequest) (*PutListingsItemResponse, error) {
	// Build the full path
	path := fmt.Sprintf("/listings/2021-08-01/items/%s/%s", request.SellerId, request.SKU)

	// Make the API call
	resp, err := l.apiClient.CallAPI(ctx, "PUT", path, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call putListingsItem API: %w", err)
	}

	// Parse the response
	var result PutListingsItemResponse
	if err := l.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse putListingsItem response: %w", err)
	}

	return &result, nil
}

// PatchListingsItem patches a listings item
func (l *ListingsAPI) PatchListingsItem(ctx context.Context, request *PatchListingsItemRequest) (*PatchListingsItemResponse, error) {
	// Build the full path
	path := fmt.Sprintf("/listings/2021-08-01/items/%s/%s", request.SellerId, request.SKU)

	// Make the API call
	resp, err := l.apiClient.CallAPI(ctx, "PATCH", path, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call patchListingsItem API: %w", err)
	}

	// Parse the response
	var result PatchListingsItemResponse
	if err := l.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse patchListingsItem response: %w", err)
	}

	return &result, nil
}

// SearchListingsItems searches for listings items
func (l *ListingsAPI) SearchListingsItems(ctx context.Context, request *SearchListingsItemsRequest) (*SearchListingsItemsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	params["marketplaceIds"] = request.MarketplaceIds

	// Optional parameters
	if request.IssueLocale != "" {
		params["issueLocale"] = request.IssueLocale
	}
	if len(request.IncludedData) > 0 {
		params["includedData"] = request.IncludedData
	}
	if len(request.Identifiers) > 0 {
		params["identifiers"] = request.Identifiers
	}
	if request.IdentifiersType != "" {
		params["identifiersType"] = request.IdentifiersType
	}
	if request.VariationParentSku != "" {
		params["variationParentSku"] = request.VariationParentSku
	}
	if request.PackageHierarchySku != "" {
		params["packageHierarchySku"] = request.PackageHierarchySku
	}
	if request.CreatedAfter != "" {
		params["createdAfter"] = request.CreatedAfter
	}
	if request.CreatedBefore != "" {
		params["createdBefore"] = request.CreatedBefore
	}
	if request.LastUpdatedAfter != "" {
		params["lastUpdatedAfter"] = request.LastUpdatedAfter
	}
	if request.LastUpdatedBefore != "" {
		params["lastUpdatedBefore"] = request.LastUpdatedBefore
	}
	if len(request.WithIssueSeverity) > 0 {
		params["withIssueSeverity"] = request.WithIssueSeverity
	}
	if len(request.WithStatus) > 0 {
		params["withStatus"] = request.WithStatus
	}
	if len(request.WithoutStatus) > 0 {
		params["withoutStatus"] = request.WithoutStatus
	}
	if request.SortBy != "" {
		params["sortBy"] = request.SortBy
	}
	if request.SortOrder != "" {
		params["sortOrder"] = request.SortOrder
	}
	if request.PageSize > 0 {
		params["pageSize"] = request.PageSize
	}
	if request.PageToken != "" {
		params["pageToken"] = request.PageToken
	}

	// Build query string
	queryString := l.apiClient.BuildQueryString(params)

	// Build the full path - 参考PHP SDK的正确路径
	path := fmt.Sprintf("/listings/2021-08-01/items/%s", request.SellerId)
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := l.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call searchListingsItems API: %w", err)
	}

	// Parse the response
	var result SearchListingsItemsResponse
	if err := l.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse searchListingsItems response: %w", err)
	}

	return &result, nil
}
