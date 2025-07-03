package v0

import (
	"context"
	"fmt"

	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// FulfillmentInboundAPI represents the Fulfillment Inbound API client
type FulfillmentInboundAPI struct {
	apiClient *client.APIClient
}

// NewFulfillmentInboundAPI creates a new Fulfillment Inbound API client
func NewFulfillmentInboundAPI(config *client.Configuration) *FulfillmentInboundAPI {
	return &FulfillmentInboundAPI{
		apiClient: client.NewAPIClient(config),
	}
}

// GetShipments retrieves shipment information
func (f *FulfillmentInboundAPI) GetShipments(ctx context.Context, request *GetShipmentsRequest) (*GetShipmentsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	if request.QueryType == "" {
		return nil, fmt.Errorf("queryType is required")
	}
	params["QueryType"] = request.QueryType

	if request.MarketplaceId == "" {
		return nil, fmt.Errorf("marketplaceId is required")
	}
	params["MarketplaceId"] = request.MarketplaceId

	// Optional parameters
	if request.ShipmentStatusList != nil && len(request.ShipmentStatusList) > 0 {
		params["ShipmentStatusList"] = request.ShipmentStatusList
	}
	if request.ShipmentIdList != nil && len(request.ShipmentIdList) > 0 {
		params["ShipmentIdList"] = request.ShipmentIdList
	}
	if request.LastUpdatedAfter != "" {
		params["LastUpdatedAfter"] = request.LastUpdatedAfter
	}
	if request.LastUpdatedBefore != "" {
		params["LastUpdatedBefore"] = request.LastUpdatedBefore
	}
	if request.NextToken != "" {
		params["NextToken"] = request.NextToken
	}

	// Build query string
	queryString := f.apiClient.BuildQueryString(params)

	// Build the full path
	path := "/fba/inbound/v0/shipments"
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getShipments API: %w", err)
	}

	// Parse the response
	var result GetShipmentsResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getShipments response: %w", err)
	}

	return &result, nil
}

// GetShipmentItems retrieves shipment items
func (f *FulfillmentInboundAPI) GetShipmentItems(ctx context.Context, request *GetShipmentItemsRequest) (*GetShipmentItemsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	if request.QueryType == "" {
		return nil, fmt.Errorf("queryType is required")
	}
	params["QueryType"] = request.QueryType

	if request.MarketplaceId == "" {
		return nil, fmt.Errorf("marketplaceId is required")
	}
	params["MarketplaceId"] = request.MarketplaceId

	// Optional parameters
	if request.ShipmentId != "" {
		params["ShipmentId"] = request.ShipmentId
	}
	if request.LastUpdatedAfter != "" {
		params["LastUpdatedAfter"] = request.LastUpdatedAfter
	}
	if request.LastUpdatedBefore != "" {
		params["LastUpdatedBefore"] = request.LastUpdatedBefore
	}
	if request.NextToken != "" {
		params["NextToken"] = request.NextToken
	}

	// Build query string
	queryString := f.apiClient.BuildQueryString(params)

	// Build the full path
	path := "/fba/inbound/v0/shipmentItems"
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getShipmentItems API: %w", err)
	}

	// Parse the response
	var result GetShipmentItemsResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getShipmentItems response: %w", err)
	}

	return &result, nil
}

// GetShipmentItemsByShipmentId retrieves shipment items by shipment ID
func (f *FulfillmentInboundAPI) GetShipmentItemsByShipmentId(ctx context.Context, shipmentId string, request *GetShipmentItemsByShipmentIdRequest) (*GetShipmentItemsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	if shipmentId == "" {
		return nil, fmt.Errorf("shipmentId is required")
	}

	// Optional parameters
	if request.MarketplaceId != "" {
		params["MarketplaceId"] = request.MarketplaceId
	}
	if request.NextToken != "" {
		params["NextToken"] = request.NextToken
	}

	// Build query string
	queryString := f.apiClient.BuildQueryString(params)

	// Build the full path
	path := fmt.Sprintf("/fba/inbound/v0/shipments/%s/items", shipmentId)
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getShipmentItemsByShipmentId API: %w", err)
	}

	// Parse the response
	var result GetShipmentItemsResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getShipmentItemsByShipmentId response: %w", err)
	}

	return &result, nil
}

// GetPrepInstructions retrieves prep instructions
func (f *FulfillmentInboundAPI) GetPrepInstructions(ctx context.Context, request *GetPrepInstructionsRequest) (*GetPrepInstructionsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	if request.ShipToCountryCode == "" {
		return nil, fmt.Errorf("shipToCountryCode is required")
	}
	params["ShipToCountryCode"] = request.ShipToCountryCode

	// Optional parameters
	if request.SellerSKUList != nil && len(request.SellerSKUList) > 0 {
		params["SellerSKUList"] = request.SellerSKUList
	}
	if request.ASINList != nil && len(request.ASINList) > 0 {
		params["ASINList"] = request.ASINList
	}

	// Build query string
	queryString := f.apiClient.BuildQueryString(params)

	// Build the full path
	path := "/fba/inbound/v0/prepInstructions"
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getPrepInstructions API: %w", err)
	}

	// Parse the response
	var result GetPrepInstructionsResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getPrepInstructions response: %w", err)
	}

	return &result, nil
}

// GetLabels retrieves labels for a shipment
func (f *FulfillmentInboundAPI) GetLabels(ctx context.Context, shipmentId string, request *GetLabelsRequest) (*GetLabelsResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	if shipmentId == "" {
		return nil, fmt.Errorf("shipmentId is required")
	}

	if request.PageType == "" {
		return nil, fmt.Errorf("pageType is required")
	}
	params["PageType"] = request.PageType

	if request.LabelType == "" {
		return nil, fmt.Errorf("labelType is required")
	}
	params["LabelType"] = request.LabelType

	// Optional parameters
	if request.NumberOfPackages > 0 {
		params["NumberOfPackages"] = request.NumberOfPackages
	}
	if request.PackageLabelsToPrint != nil && len(request.PackageLabelsToPrint) > 0 {
		params["PackageLabelsToPrint"] = request.PackageLabelsToPrint
	}
	if request.NumberOfPallets > 0 {
		params["NumberOfPallets"] = request.NumberOfPallets
	}
	if request.PageSize > 0 {
		params["PageSize"] = request.PageSize
	}
	if request.PageStartIndex > 0 {
		params["PageStartIndex"] = request.PageStartIndex
	}

	// Build query string
	queryString := f.apiClient.BuildQueryString(params)

	// Build the full path
	path := fmt.Sprintf("/fba/inbound/v0/shipments/%s/labels", shipmentId)
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getLabels API: %w", err)
	}

	// Parse the response
	var result GetLabelsResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getLabels response: %w", err)
	}

	return &result, nil
}

// GetBillOfLading retrieves bill of lading for a shipment
func (f *FulfillmentInboundAPI) GetBillOfLading(ctx context.Context, shipmentId string) (*GetBillOfLadingResponse, error) {
	// Required parameters
	if shipmentId == "" {
		return nil, fmt.Errorf("shipmentId is required")
	}

	// Build the full path
	path := fmt.Sprintf("/fba/inbound/v0/shipments/%s/billOfLading", shipmentId)

	// Make the API call
	resp, err := f.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getBillOfLading API: %w", err)
	}

	// Parse the response
	var result GetBillOfLadingResponse
	if err := f.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getBillOfLading response: %w", err)
	}

	return &result, nil
}
