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
