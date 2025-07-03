package v1

import (
	"context"
	"fmt"

	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// NotificationsAPI represents the Notifications API client
type NotificationsAPI struct {
	apiClient *client.APIClient
}

// NewNotificationsAPI creates a new Notifications API client
func NewNotificationsAPI(config *client.Configuration) *NotificationsAPI {
	return &NotificationsAPI{
		apiClient: client.NewAPIClient(config),
	}
}

// CreateDestination creates a destination for receiving notifications
func (n *NotificationsAPI) CreateDestination(ctx context.Context, request *CreateDestinationRequest) (*CreateDestinationResponse, error) {
	// Build the full path
	path := "/notifications/v1/destinations"

	// Make the API call
	resp, err := n.apiClient.CallAPI(ctx, "POST", path, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call createDestination API: %w", err)
	}

	// Parse the response
	var result CreateDestinationResponse
	if err := n.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse createDestination response: %w", err)
	}

	return &result, nil
}

// GetDestinations retrieves all destinations
func (n *NotificationsAPI) GetDestinations(ctx context.Context) (*GetDestinationsResponse, error) {
	// Build the full path
	path := "/notifications/v1/destinations"

	// Make the API call
	resp, err := n.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getDestinations API: %w", err)
	}

	// Parse the response
	var result GetDestinationsResponse
	if err := n.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getDestinations response: %w", err)
	}

	return &result, nil
}

// GetDestination retrieves a specific destination by ID
func (n *NotificationsAPI) GetDestination(ctx context.Context, destinationId string) (*GetDestinationResponse, error) {
	// Build the full path
	path := fmt.Sprintf("/notifications/v1/destinations/%s", destinationId)

	// Make the API call
	resp, err := n.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getDestination API: %w", err)
	}

	// Parse the response
	var result GetDestinationResponse
	if err := n.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getDestination response: %w", err)
	}

	return &result, nil
}

// DeleteDestination deletes a destination by ID
func (n *NotificationsAPI) DeleteDestination(ctx context.Context, destinationId string) (*DeleteDestinationResponse, error) {
	// Build the full path
	path := fmt.Sprintf("/notifications/v1/destinations/%s", destinationId)

	// Make the API call
	resp, err := n.apiClient.CallAPI(ctx, "DELETE", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call deleteDestination API: %w", err)
	}

	// Parse the response
	var result DeleteDestinationResponse
	if err := n.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse deleteDestination response: %w", err)
	}

	return &result, nil
}

// CreateSubscription creates a subscription for a notification type
func (n *NotificationsAPI) CreateSubscription(ctx context.Context, notificationType string, request *CreateSubscriptionRequest) (*CreateSubscriptionResponse, error) {
	// Build the full path
	path := fmt.Sprintf("/notifications/v1/subscriptions/%s", notificationType)

	// Make the API call
	resp, err := n.apiClient.CallAPI(ctx, "POST", path, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call createSubscription API: %w", err)
	}

	// Parse the response
	var result CreateSubscriptionResponse
	if err := n.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse createSubscription response: %w", err)
	}

	return &result, nil
}

// GetSubscription retrieves a subscription for a notification type
func (n *NotificationsAPI) GetSubscription(ctx context.Context, notificationType string, payloadVersion string) (*GetSubscriptionResponse, error) {
	// Build query parameters
	params := make(map[string]interface{})
	if payloadVersion != "" {
		params["payloadVersion"] = payloadVersion
	}

	// Build query string
	queryString := n.apiClient.BuildQueryString(params)

	// Build the full path
	path := fmt.Sprintf("/notifications/v1/subscriptions/%s", notificationType)
	if queryString != "" {
		path += "?" + queryString
	}

	// Make the API call
	resp, err := n.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getSubscription API: %w", err)
	}

	// Parse the response
	var result GetSubscriptionResponse
	if err := n.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getSubscription response: %w", err)
	}

	return &result, nil
}

// GetSubscriptionById retrieves a specific subscription by ID
func (n *NotificationsAPI) GetSubscriptionById(ctx context.Context, subscriptionId string, notificationType string) (*GetSubscriptionByIdResponse, error) {
	// Build the full path
	path := fmt.Sprintf("/notifications/v1/subscriptions/%s/%s", notificationType, subscriptionId)

	// Make the API call
	resp, err := n.apiClient.CallAPI(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call getSubscriptionById API: %w", err)
	}

	// Parse the response
	var result GetSubscriptionByIdResponse
	if err := n.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse getSubscriptionById response: %w", err)
	}

	return &result, nil
}

// DeleteSubscriptionById deletes a subscription by ID
func (n *NotificationsAPI) DeleteSubscriptionById(ctx context.Context, subscriptionId string, notificationType string) (*DeleteSubscriptionByIdResponse, error) {
	// Build the full path
	path := fmt.Sprintf("/notifications/v1/subscriptions/%s/%s", notificationType, subscriptionId)

	// Make the API call
	resp, err := n.apiClient.CallAPI(ctx, "DELETE", path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call deleteSubscriptionById API: %w", err)
	}

	// Parse the response
	var result DeleteSubscriptionByIdResponse
	if err := n.apiClient.ProcessResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse deleteSubscriptionById response: %w", err)
	}

	return &result, nil
}

// Helper methods for common use cases

// CreateSqsDestination creates an SQS destination
func (n *NotificationsAPI) CreateSqsDestination(ctx context.Context, name string, sqsArn string) (*CreateDestinationResponse, error) {
	request := &CreateDestinationRequest{
		Name: name,
		ResourceSpecification: &DestinationResourceSpecification{
			Sqs: &SqsResourceSpecification{
				Arn: sqsArn,
			},
		},
	}
	return n.CreateDestination(ctx, request)
}

// CreateEventBridgeDestination creates an EventBridge destination
func (n *NotificationsAPI) CreateEventBridgeDestination(ctx context.Context, name string, eventBridgeName string, region string, accountId string) (*CreateDestinationResponse, error) {
	request := &CreateDestinationRequest{
		Name: name,
		ResourceSpecification: &DestinationResourceSpecification{
			EventBridge: &EventBridgeResourceSpecification{
				Name:      eventBridgeName,
				Region:    region,
				AccountId: accountId,
			},
		},
	}
	return n.CreateDestination(ctx, request)
}

// CreateOrderChangeSubscription creates a subscription for order change notifications
func (n *NotificationsAPI) CreateOrderChangeSubscription(ctx context.Context, destinationId string, marketplaceIds []string) (*CreateSubscriptionResponse, error) {
	request := &CreateSubscriptionRequest{
		DestinationId: destinationId,
		ProcessingDirective: &ProcessingDirective{
			EventFilter: &EventFilter{
				MarketplaceFilter: &MarketplaceFilter{
					MarketplaceIds: marketplaceIds,
				},
			},
		},
	}
	return n.CreateSubscription(ctx, "ORDERS_CHANGE", request)
}
