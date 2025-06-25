package catalog

import (
	"context"
	"fmt"

	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// CatalogAPI represents the Catalog API client
type CatalogAPI struct {
	apiClient *client.APIClient
}

// NewCatalogAPI creates a new Catalog API client
func NewCatalogAPI(config *client.Configuration) *CatalogAPI {
	return &CatalogAPI{
		apiClient: client.NewAPIClient(config),
	}
}

// GetCatalogItem retrieves a catalog item by ASIN
func (c *CatalogAPI) GetCatalogItem(ctx context.Context, request *GetCatalogItemRequest) (*GetCatalogItemResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	params["marketplaceIds"] = request.MarketplaceIds

	// Optional parameters
	if len(request.IncludedData) > 0 {
		params["includedData"] = request.IncludedData
	}
	if request.Locale != "" {
		params["locale"] = request.Locale
	}

	// Build query string
	queryString := c.apiClient.BuildQueryString(params)

	// Build the full path - 使用 2022-04-01 版本
	path := fmt.Sprintf("/catalog/2022-04-01/items/%s", request.ASIN)
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := c.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getCatalogItem API: %w", err)
	}

	// Parse the response
	var result GetCatalogItemResponse
	if err := c.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getCatalogItem response: %w", err)
	}

	return &result, nil
}

// SearchCatalogItems searches for catalog items
func (c *CatalogAPI) SearchCatalogItems(ctx context.Context, request *SearchCatalogItemsRequest) (*SearchCatalogItemsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	params["marketplaceIds"] = request.MarketplaceIds

	// Optional parameters
	if len(request.Identifiers) > 0 {
		params["identifiers"] = request.Identifiers
	}
	if request.IdentifiersType != "" {
		params["identifiersType"] = request.IdentifiersType
	}
	if len(request.IncludedData) > 0 {
		params["includedData"] = request.IncludedData
	}
	if request.Locale != "" {
		params["locale"] = request.Locale
	}
	if request.SellerId != "" {
		params["sellerId"] = request.SellerId
	}
	if len(request.Keywords) > 0 {
		params["keywords"] = request.Keywords
	}
	if len(request.BrandNames) > 0 {
		params["brandNames"] = request.BrandNames
	}
	if len(request.ClassificationIds) > 0 {
		params["classificationIds"] = request.ClassificationIds
	}
	if request.PageSize > 0 {
		params["pageSize"] = request.PageSize
	}
	if request.PageToken != "" {
		params["pageToken"] = request.PageToken
	}
	if request.KeywordsLocale != "" {
		params["keywordsLocale"] = request.KeywordsLocale
	}

	// Build query string
	queryString := c.apiClient.BuildQueryString(params)

	// Build the full path - 使用 2022-04-01 版本
	path := "/catalog/2022-04-01/items"
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := c.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call searchCatalogItems API: %w", err)
	}

	// Parse the response
	var result SearchCatalogItemsResponse
	if err := c.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse searchCatalogItems response: %w", err)
	}

	return &result, nil
}

// GetCatalogItemSimple is a simplified version of GetCatalogItem for common use cases
func (c *CatalogAPI) GetCatalogItemSimple(ctx context.Context, asin string, marketplaceIds []string) (*GetCatalogItemResponse, error) {
	request := &GetCatalogItemRequest{
		ASIN:           asin,
		MarketplaceIds: marketplaceIds,
		IncludedData:   []string{"summaries", "attributes", "images", "classifications", "dimensions", "identifiers", "productTypes", "relationships", "salesRanks", "vendorDetails"},
		Locale:         "en_US",
	}
	return c.GetCatalogItem(ctx, request)
}

// SearchCatalogItemsSimple is a simplified version of SearchCatalogItems for common use cases
func (c *CatalogAPI) SearchCatalogItemsSimple(ctx context.Context, marketplaceIds []string, keywords []string) (*SearchCatalogItemsResponse, error) {
	request := &SearchCatalogItemsRequest{
		MarketplaceIds: marketplaceIds,
		Keywords:       keywords,
		IncludedData:   []string{"summaries", "attributes"},
		Locale:         "en_US",
		PageSize:       10,
		KeywordsLocale: "en_US",
	}
	return c.SearchCatalogItems(ctx, request)
}
