package v1

// CreateDestinationRequest represents the request for creating a destination
type CreateDestinationRequest struct {
	ResourceSpecification *DestinationResourceSpecification `json:"resourceSpecification"`
	Name                  string                            `json:"name"`
}

// CreateDestinationResponse represents the response from createDestination
type CreateDestinationResponse struct {
	Payload *Destination `json:"payload,omitempty"`
	Errors  []Error      `json:"errors,omitempty"`
}

// GetDestinationsResponse represents the response from getDestinations
type GetDestinationsResponse struct {
	Payload *DestinationsPayload `json:"payload,omitempty"`
	Errors  []Error              `json:"errors,omitempty"`
}

// DestinationsPayload represents the payload in getDestinations response
type DestinationsPayload struct {
	Destinations []Destination `json:"destinations"`
}

// GetDestinationResponse represents the response from getDestination
type GetDestinationResponse struct {
	Payload *Destination `json:"payload,omitempty"`
	Errors  []Error      `json:"errors,omitempty"`
}

// DeleteDestinationResponse represents the response from deleteDestination
type DeleteDestinationResponse struct {
	Errors []Error `json:"errors,omitempty"`
}

// CreateSubscriptionRequest represents the request for creating a subscription
type CreateSubscriptionRequest struct {
	PayloadVersion      string               `json:"payloadVersion,omitempty"`
	DestinationId       string               `json:"destinationId"`
	ProcessingDirective *ProcessingDirective `json:"processingDirective,omitempty"`
}

// CreateSubscriptionResponse represents the response from createSubscription
type CreateSubscriptionResponse struct {
	Payload *Subscription `json:"payload,omitempty"`
	Errors  []Error       `json:"errors,omitempty"`
}

// GetSubscriptionResponse represents the response from getSubscription
type GetSubscriptionResponse struct {
	Payload *Subscription `json:"payload,omitempty"`
	Errors  []Error       `json:"errors,omitempty"`
}

// GetSubscriptionByIdResponse represents the response from getSubscriptionById
type GetSubscriptionByIdResponse struct {
	Payload *Subscription `json:"payload,omitempty"`
	Errors  []Error       `json:"errors,omitempty"`
}

// DeleteSubscriptionByIdResponse represents the response from deleteSubscriptionById
type DeleteSubscriptionByIdResponse struct {
	Errors []Error `json:"errors,omitempty"`
}

// Destination represents a notification destination
type Destination struct {
	Name                  string                            `json:"name"`
	DestinationId         string                            `json:"destinationId"`
	Resource              *DestinationResource              `json:"resource"`
	ResourceSpecification *DestinationResourceSpecification `json:"resourceSpecification,omitempty"`
}

// DestinationResource represents a destination resource
type DestinationResource struct {
	Sqs         *SqsResource         `json:"sqs,omitempty"`
	EventBridge *EventBridgeResource `json:"eventBridge,omitempty"`
}

// DestinationResourceSpecification represents destination resource specification
type DestinationResourceSpecification struct {
	Sqs         *SqsResourceSpecification         `json:"sqs,omitempty"`
	EventBridge *EventBridgeResourceSpecification `json:"eventBridge,omitempty"`
}

// SqsResource represents an SQS resource
type SqsResource struct {
	Arn string `json:"arn"`
}

// SqsResourceSpecification represents SQS resource specification
type SqsResourceSpecification struct {
	Arn string `json:"arn"`
}

// EventBridgeResource represents an EventBridge resource
type EventBridgeResource struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	AccountId string `json:"accountId"`
}

// EventBridgeResourceSpecification represents EventBridge resource specification
type EventBridgeResourceSpecification struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	AccountId string `json:"accountId"`
}

// Subscription represents a notification subscription
type Subscription struct {
	SubscriptionId      string               `json:"subscriptionId"`
	PayloadVersion      string               `json:"payloadVersion"`
	DestinationId       string               `json:"destinationId"`
	ProcessingDirective *ProcessingDirective `json:"processingDirective,omitempty"`
}

// ProcessingDirective represents processing directive
type ProcessingDirective struct {
	EventFilter *EventFilter `json:"eventFilter,omitempty"`
}

// EventFilter represents event filter
type EventFilter struct {
	AggregationSettings   *AggregationSettings   `json:"aggregationSettings,omitempty"`
	MarketplaceFilter     *MarketplaceFilter     `json:"marketplaceFilter,omitempty"`
	OrderChangeTypeFilter *OrderChangeTypeFilter `json:"orderChangeTypeFilter,omitempty"`
}

// AggregationSettings represents aggregation settings
type AggregationSettings struct {
	AggregationTimePeriod *AggregationTimePeriod `json:"aggregationTimePeriod,omitempty"`
}

// AggregationTimePeriod represents aggregation time period
type AggregationTimePeriod string

const (
	AggregationTimePeriodFiveMinutes AggregationTimePeriod = "FiveMinutes"
	AggregationTimePeriodTenMinutes  AggregationTimePeriod = "TenMinutes"
)

// MarketplaceFilter represents marketplace filter
type MarketplaceFilter struct {
	MarketplaceIds []string `json:"marketplaceIds,omitempty"`
}

// OrderChangeTypeFilter represents order change type filter
type OrderChangeTypeFilter struct {
	OrderChangeTypes []OrderChangeType `json:"orderChangeTypes,omitempty"`
}

// OrderChangeType represents order change type
type OrderChangeType string

const (
	OrderChangeTypeOrderStatusChange    OrderChangeType = "OrderStatusChange"
	OrderChangeTypeBuyerRequestedChange OrderChangeType = "BuyerRequestedChange"
)

// Error represents an API error
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}
