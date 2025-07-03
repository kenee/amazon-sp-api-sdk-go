package v2020_07_01

import (
	"context"
	"fmt"

	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// FulfillmentOutboundAPI represents the Fulfillment Outbound API client
type FulfillmentOutboundAPI struct {
	apiClient *client.APIClient
}

// NewFulfillmentOutboundAPI creates a new Fulfillment Outbound API client
func NewFulfillmentOutboundAPI(config *client.Configuration) *FulfillmentOutboundAPI {
	return &FulfillmentOutboundAPI{
		apiClient: client.NewAPIClient(config),
	}
}

// CancelFulfillmentOrder cancels a fulfillment order
func (api *FulfillmentOutboundAPI) CancelFulfillmentOrder(ctx context.Context, sellerFulfillmentOrderId string) (*CancelFulfillmentOrderResponse, error) {
	path := fmt.Sprintf("/fba/outbound/2020-07-01/fulfillmentOrders/%s/cancel", sellerFulfillmentOrderId)

	resp, err := api.apiClient.CallAPI(ctx, "PUT", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call cancelFulfillmentOrder API: %w", err)
	}

	var response CancelFulfillmentOrderResponse
	if err := api.apiClient.ProcessResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse cancelFulfillmentOrder response: %w", err)
	}

	return &response, nil
}

// CreateFulfillmentOrder creates a new fulfillment order
func (api *FulfillmentOutboundAPI) CreateFulfillmentOrder(ctx context.Context, request *CreateFulfillmentOrderRequest) (*CreateFulfillmentOrderResponse, error) {
	path := "/fba/outbound/2020-07-01/fulfillmentOrders"

	resp, err := api.apiClient.CallAPI(ctx, "POST", path, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call createFulfillmentOrder API: %w", err)
	}

	var response CreateFulfillmentOrderResponse
	if err := api.apiClient.ProcessResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse createFulfillmentOrder response: %w", err)
	}

	return &response, nil
}

// CreateFulfillmentReturn creates a fulfillment return
func (api *FulfillmentOutboundAPI) CreateFulfillmentReturn(ctx context.Context, sellerFulfillmentOrderId string, request *CreateFulfillmentReturnRequest) (*CreateFulfillmentReturnResponse, error) {
	path := fmt.Sprintf("/fba/outbound/2020-07-01/fulfillmentOrders/%s/return", sellerFulfillmentOrderId)

	resp, err := api.apiClient.CallAPI(ctx, "PUT", path, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call createFulfillmentReturn API: %w", err)
	}

	var response CreateFulfillmentReturnResponse
	if err := api.apiClient.ProcessResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse createFulfillmentReturn response: %w", err)
	}

	return &response, nil
}

// GetFulfillmentOrder gets a fulfillment order by ID
func (api *FulfillmentOutboundAPI) GetFulfillmentOrder(ctx context.Context, sellerFulfillmentOrderId string) (*GetFulfillmentOrderResponse, error) {
	path := fmt.Sprintf("/fba/outbound/2020-07-01/fulfillmentOrders/%s", sellerFulfillmentOrderId)

	resp, err := api.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getFulfillmentOrder API: %w", err)
	}

	var response GetFulfillmentOrderResponse
	if err := api.apiClient.ProcessResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse getFulfillmentOrder response: %w", err)
	}

	return &response, nil
}

// GetFulfillmentPreview gets a fulfillment preview
func (api *FulfillmentOutboundAPI) GetFulfillmentPreview(ctx context.Context, request *GetFulfillmentPreviewRequest) (*GetFulfillmentPreviewResponse, error) {
	path := "/fba/outbound/2020-07-01/fulfillmentOrders/preview"

	resp, err := api.apiClient.CallAPI(ctx, "POST", path, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getFulfillmentPreview API: %w", err)
	}

	var response GetFulfillmentPreviewResponse
	if err := api.apiClient.ProcessResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse getFulfillmentPreview response: %w", err)
	}

	return &response, nil
}

// ListAllFulfillmentOrders lists all fulfillment orders
func (api *FulfillmentOutboundAPI) ListAllFulfillmentOrders(ctx context.Context, queryStartTime *string, nextToken *string) (*ListAllFulfillmentOrdersResponse, error) {
	path := "/fba/outbound/2020-07-01/fulfillmentOrders"

	// Build query parameters
	queryParams := make(map[string]string)
	if queryStartTime != nil {
		queryParams["queryStartTime"] = *queryStartTime
	}
	if nextToken != nil {
		queryParams["nextToken"] = *nextToken
	}

	resp, err := api.apiClient.CallAPIWithQueryParams(ctx, "GET", path, nil, nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to call listAllFulfillmentOrders API: %w", err)
	}

	var response ListAllFulfillmentOrdersResponse
	if err := api.apiClient.ProcessResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse listAllFulfillmentOrders response: %w", err)
	}

	return &response, nil
}

// UpdateFulfillmentOrder updates a fulfillment order
func (api *FulfillmentOutboundAPI) UpdateFulfillmentOrder(ctx context.Context, sellerFulfillmentOrderId string, request *UpdateFulfillmentOrderRequest) (*UpdateFulfillmentOrderResponse, error) {
	path := fmt.Sprintf("/fba/outbound/2020-07-01/fulfillmentOrders/%s", sellerFulfillmentOrderId)

	resp, err := api.apiClient.CallAPI(ctx, "PATCH", path, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call updateFulfillmentOrder API: %w", err)
	}

	var response UpdateFulfillmentOrderResponse
	if err := api.apiClient.ProcessResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse updateFulfillmentOrder response: %w", err)
	}

	return &response, nil
}

// GetPackageTrackingDetails gets package tracking details
func (api *FulfillmentOutboundAPI) GetPackageTrackingDetails(ctx context.Context, packageNumber string) (*GetPackageTrackingDetailsResponse, error) {
	path := fmt.Sprintf("/fba/outbound/2020-07-01/tracking/%s", packageNumber)

	resp, err := api.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getPackageTrackingDetails API: %w", err)
	}

	var response GetPackageTrackingDetailsResponse
	if err := api.apiClient.ProcessResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse getPackageTrackingDetails response: %w", err)
	}

	return &response, nil
}

// ListReturnReasonCodes lists return reason codes
func (api *FulfillmentOutboundAPI) ListReturnReasonCodes(ctx context.Context, sellerSku string, language *string) (*ListReturnReasonCodesResponse, error) {
	path := "/fba/outbound/2020-07-01/returnReasonCodes"

	// Build query parameters
	queryParams := make(map[string]string)
	queryParams["sellerSku"] = sellerSku
	if language != nil {
		queryParams["language"] = *language
	}

	resp, err := api.apiClient.CallAPIWithQueryParams(ctx, "GET", path, nil, nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to call listReturnReasonCodes API: %w", err)
	}

	var response ListReturnReasonCodesResponse
	if err := api.apiClient.ProcessResponse(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to parse listReturnReasonCodes response: %w", err)
	}

	return &response, nil
}

// GetAPIClient returns the underlying API client
func (api *FulfillmentOutboundAPI) GetAPIClient() *client.APIClient {
	return api.apiClient
}
