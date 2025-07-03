package v2020_07_01

// 这里只定义主要结构体，字段可后续补充完善

type CancelFulfillmentOrderResponse struct {
	Payload *FulfillmentOrder `json:"payload,omitempty"`
	Errors  []Error           `json:"errors,omitempty"`
}

type CreateFulfillmentOrderRequest struct {
	MarketplaceId            string                       `json:"marketplaceId"`
	SellerFulfillmentOrderId string                       `json:"sellerFulfillmentOrderId"`
	DisplayableOrderId       string                       `json:"displayableOrderId"`
	DisplayableOrderDate     string                       `json:"displayableOrderDate"`
	DisplayableOrderComment  string                       `json:"displayableOrderComment"`
	ShippingSpeedCategory    string                       `json:"shippingSpeedCategory"`
	DestinationAddress       *Address                     `json:"destinationAddress"`
	FulfillmentAction        string                       `json:"fulfillmentAction,omitempty"`
	FulfillmentPolicy        string                       `json:"fulfillmentPolicy,omitempty"`
	FulfillmentMethod        string                       `json:"fulfillmentMethod,omitempty"`
	CODSettings              *CODSettings                 `json:"codSettings,omitempty"`
	ShipFromAddress          *Address                     `json:"shipFromAddress,omitempty"`
	NotificationEmails       []string                     `json:"notificationEmails,omitempty"`
	Items                    []CreateFulfillmentOrderItem `json:"items"`
	FeatureConstraints       []FeatureSettings            `json:"featureConstraints,omitempty"`
}

type CreateFulfillmentOrderResponse struct {
	Payload *FulfillmentOrder `json:"payload,omitempty"`
	Errors  []Error           `json:"errors,omitempty"`
}

type CreateFulfillmentReturnRequest struct {
	Items []CreateReturnItem `json:"items"`
}

type CreateFulfillmentReturnResponse struct {
	Payload *CreateFulfillmentReturnResult `json:"payload,omitempty"`
	Errors  []Error                        `json:"errors,omitempty"`
}

type GetFulfillmentOrderResponse struct {
	Payload *FulfillmentOrder `json:"payload,omitempty"`
	Errors  []Error           `json:"errors,omitempty"`
}

type GetFulfillmentPreviewRequest struct {
	MarketplaceId                string                      `json:"marketplaceId"`
	Address                      *Address                    `json:"address"`
	Items                        []GetFulfillmentPreviewItem `json:"items"`
	ShippingSpeedCategories      []string                    `json:"shippingSpeedCategories,omitempty"`
	IncludeCODFulfillmentPreview bool                        `json:"includeCODFulfillmentPreview,omitempty"`
	IncludeDeliveryWindows       bool                        `json:"includeDeliveryWindows,omitempty"`
	FeatureConstraints           []FeatureSettings           `json:"featureConstraints,omitempty"`
}

type GetFulfillmentPreviewResponse struct {
	Payload *GetFulfillmentPreviewResult `json:"payload,omitempty"`
	Errors  []Error                      `json:"errors,omitempty"`
}

type ListAllFulfillmentOrdersResponse struct {
	Payload *ListAllFulfillmentOrdersResult `json:"payload,omitempty"`
	Errors  []Error                         `json:"errors,omitempty"`
}

type UpdateFulfillmentOrderRequest struct {
	DisplayableOrderId      string                       `json:"displayableOrderId"`
	DisplayableOrderDate    string                       `json:"displayableOrderDate"`
	DisplayableOrderComment string                       `json:"displayableOrderComment"`
	Items                   []UpdateFulfillmentOrderItem `json:"items"`
	NotificationEmails      []string                     `json:"notificationEmails,omitempty"`
}

type UpdateFulfillmentOrderResponse struct {
	Payload *FulfillmentOrder `json:"payload,omitempty"`
	Errors  []Error           `json:"errors,omitempty"`
}

type GetPackageTrackingDetailsResponse struct {
	Payload *PackageTrackingDetails `json:"payload,omitempty"`
	Errors  []Error                 `json:"errors,omitempty"`
}

type ListReturnReasonCodesResponse struct {
	Payload *ListReturnReasonCodesResult `json:"payload,omitempty"`
	Errors  []Error                      `json:"errors,omitempty"`
}

// 下面是部分嵌套结构体定义（可后续补充）
type FulfillmentOrder struct{}
type Error struct{}
type Address struct{}
type CODSettings struct{}
type CreateFulfillmentOrderItem struct{}
type FeatureSettings struct{}
type CreateReturnItem struct{}
type CreateFulfillmentReturnResult struct{}
type GetFulfillmentPreviewItem struct{}
type GetFulfillmentPreviewResult struct{}
type ListAllFulfillmentOrdersResult struct{}
type UpdateFulfillmentOrderItem struct{}
type PackageTrackingDetails struct{}
type ListReturnReasonCodesResult struct{}
