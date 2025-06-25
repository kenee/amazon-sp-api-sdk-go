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

// GetOrders retrieves orders based on the specified criteria
func (o *OrdersAPI) GetOrders(ctx context.Context, request *GetOrdersRequest) (*GetOrdersResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})

	// Required parameters
	if len(request.MarketplaceIds) > 0 {
		params["MarketplaceIds"] = request.MarketplaceIds
	}

	// Optional parameters
	if request.CreatedAfter != "" {
		params["CreatedAfter"] = request.CreatedAfter
	}
	if request.CreatedBefore != "" {
		params["CreatedBefore"] = request.CreatedBefore
	}
	if request.LastUpdatedAfter != "" {
		params["LastUpdatedAfter"] = request.LastUpdatedAfter
	}
	if request.LastUpdatedBefore != "" {
		params["LastUpdatedBefore"] = request.LastUpdatedBefore
	}
	if len(request.OrderStatuses) > 0 {
		params["OrderStatuses"] = request.OrderStatuses
	}
	if len(request.FulfillmentChannels) > 0 {
		params["FulfillmentChannels"] = request.FulfillmentChannels
	}
	if len(request.PaymentMethods) > 0 {
		params["PaymentMethods"] = request.PaymentMethods
	}
	if request.BuyerEmail != "" {
		params["BuyerEmail"] = request.BuyerEmail
	}
	if request.SellerOrderId != "" {
		params["SellerOrderId"] = request.SellerOrderId
	}
	if request.MaxResultsPerPage > 0 {
		params["MaxResultsPerPage"] = request.MaxResultsPerPage
	}
	if len(request.EasyShipShipmentStatuses) > 0 {
		params["EasyShipShipmentStatuses"] = request.EasyShipShipmentStatuses
	}
	if len(request.ElectronicInvoiceStatuses) > 0 {
		params["ElectronicInvoiceStatuses"] = request.ElectronicInvoiceStatuses
	}
	if request.NextToken != "" {
		params["NextToken"] = request.NextToken
	}
	if len(request.AmazonOrderIds) > 0 {
		params["AmazonOrderIds"] = request.AmazonOrderIds
	}
	if request.ActualFulfillmentSupplySourceId != "" {
		params["ActualFulfillmentSupplySourceId"] = request.ActualFulfillmentSupplySourceId
	}
	if request.IsISPU != nil {
		params["IsISPU"] = *request.IsISPU
	}
	if request.StoreChainStoreId != "" {
		params["StoreChainStoreId"] = request.StoreChainStoreId
	}
	if request.RestrictedDataToken != "" {
		params["RestrictedDataToken"] = request.RestrictedDataToken
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
