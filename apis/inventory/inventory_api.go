package inventory

import (
	"context"
	"fmt"

	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// InventoryAPI represents the Inventory API client
type InventoryAPI struct {
	apiClient *client.APIClient
}

// NewInventoryAPI creates a new Inventory API client
func NewInventoryAPI(config *client.Configuration) *InventoryAPI {
	return &InventoryAPI{
		apiClient: client.NewAPIClient(config),
	}
}

// GetInventorySummaries retrieves inventory summaries
func (i *InventoryAPI) GetInventorySummaries(ctx context.Context, request *GetInventorySummariesRequest) (*GetInventorySummariesResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	params["granularityType"] = request.GranularityType
	params["granularityId"] = request.GranularityId
	params["marketplaceIds"] = request.MarketplaceIds

	// Optional parameters
	if request.SellerId != "" {
		params["sellerId"] = request.SellerId
	}
	if request.Details {
		params["details"] = request.Details
	}
	if request.StartDateTime != "" {
		params["startDateTime"] = request.StartDateTime
	}
	if len(request.SellerSkus) > 0 {
		params["sellerSkus"] = request.SellerSkus
	}
	if request.SellerSku != "" {
		params["sellerSku"] = request.SellerSku
	}
	if request.NextToken != "" {
		params["nextToken"] = request.NextToken
	}

	// Build query string
	queryString := i.apiClient.BuildQueryString(params)

	// Build the full path
	path := "/fba/inventory/v1/summaries"
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := i.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getInventorySummaries API: %w", err)
	}

	// Parse the response
	var result GetInventorySummariesResponse
	if err := i.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getInventorySummaries response: %w", err)
	}

	return &result, nil
}

// GetInventorySummariesSimple is a simplified version for common use cases
func (i *InventoryAPI) GetInventorySummariesSimple(ctx context.Context, sellerId string, marketplaceIds []string) (*GetInventorySummariesResponse, error) {
	// Use the first marketplace ID as granularity ID for simplicity
	var granularityId string
	if len(marketplaceIds) > 0 {
		granularityId = marketplaceIds[0]
	}

	request := &GetInventorySummariesRequest{
		GranularityType: "Marketplace",
		GranularityId:   granularityId,
		SellerId:        sellerId,
		MarketplaceIds:  marketplaceIds,
		Details:         true,
	}
	return i.GetInventorySummaries(ctx, request)
}
