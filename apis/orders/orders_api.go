package orders

import (
	"context"
	"fmt"

	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// OrdersAPI represents the Orders API client
type OrdersAPI struct {
	apiClient *client.APIClient
}

// NewOrdersAPI creates a new Orders API client
func NewOrdersAPI(config *client.Configuration) *OrdersAPI {
	return &OrdersAPI{
		apiClient: client.NewAPIClient(config),
	}
}

// GetAPIClient returns the underlying API client
func (o *OrdersAPI) GetAPIClient() *client.APIClient {
	return o.apiClient
}

// GetOrders retrieves orders based on the specified criteria
func (o *OrdersAPI) GetOrders(ctx context.Context, request *GetOrdersRequest) (*GetOrdersResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	if len(request.MarketplaceIds) > 0 {
		params["marketplaceIds"] = request.MarketplaceIds
	}

	// Optional parameters
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
	if len(request.OrderStatuses) > 0 {
		params["orderStatuses"] = request.OrderStatuses
	}
	if len(request.FulfillmentChannels) > 0 {
		params["fulfillmentChannels"] = request.FulfillmentChannels
	}
	if len(request.PaymentMethods) > 0 {
		params["paymentMethods"] = request.PaymentMethods
	}
	if request.BuyerEmail != "" {
		params["buyerEmail"] = request.BuyerEmail
	}
	if request.SellerOrderId != "" {
		params["sellerOrderId"] = request.SellerOrderId
	}
	if request.MaxResultsPerPage > 0 {
		params["maxResultsPerPage"] = request.MaxResultsPerPage
	}
	if len(request.EasyShipShipmentStatuses) > 0 {
		params["easyShipShipmentStatuses"] = request.EasyShipShipmentStatuses
	}
	if len(request.ElectronicInvoiceStatuses) > 0 {
		params["electronicInvoiceStatuses"] = request.ElectronicInvoiceStatuses
	}
	if request.NextToken != "" {
		params["nextToken"] = request.NextToken
	}
	if len(request.AmazonOrderIds) > 0 {
		params["amazonOrderIds"] = request.AmazonOrderIds
	}
	if request.ActualFulfillmentSupplySourceId != "" {
		params["actualFulfillmentSupplySourceId"] = request.ActualFulfillmentSupplySourceId
	}
	if request.IsISPU != nil {
		params["isISPU"] = *request.IsISPU
	}
	if request.StoreChainStoreId != "" {
		params["storeChainStoreId"] = request.StoreChainStoreId
	}
	if request.RestrictedDataToken != "" {
		params["restrictedDataToken"] = request.RestrictedDataToken
	}

	// Build query string
	queryString := o.apiClient.BuildQueryString(params)

	// Build the full path
	path := "/orders/v0/orders"
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := o.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getOrders API: %w", err)
	}

	// Parse the response
	var result GetOrdersResponse
	if err := o.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getOrders response: %w", err)
	}

	return &result, nil
}

// GetOrdersSimple is a simplified version of GetOrders for common use cases
func (o *OrdersAPI) GetOrdersSimple(ctx context.Context, marketplaceIds []string, createdAfter string) (*GetOrdersResponse, error) {
	request := &GetOrdersRequest{
		MarketplaceIds: marketplaceIds,
		CreatedAfter:   createdAfter,
	}
	return o.GetOrders(ctx, request)
}

// GetOrdersWithStatus retrieves orders with specific status
func (o *OrdersAPI) GetOrdersWithStatus(ctx context.Context, marketplaceIds []string, orderStatuses []string) (*GetOrdersResponse, error) {
	request := &GetOrdersRequest{
		MarketplaceIds: marketplaceIds,
		OrderStatuses:  orderStatuses,
	}
	return o.GetOrders(ctx, request)
}
