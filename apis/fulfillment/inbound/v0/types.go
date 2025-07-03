package v0

// Request types

// GetShipmentsRequest represents the request for getting shipments
type GetShipmentsRequest struct {
	QueryType          string   `json:"queryType,omitempty"`
	MarketplaceId      string   `json:"marketplaceId,omitempty"`
	ShipmentStatusList []string `json:"shipmentStatusList,omitempty"`
	ShipmentIdList     []string `json:"shipmentIdList,omitempty"`
	LastUpdatedAfter   string   `json:"lastUpdatedAfter,omitempty"`
	LastUpdatedBefore  string   `json:"lastUpdatedBefore,omitempty"`
	NextToken          string   `json:"nextToken,omitempty"`
}

// GetShipmentItemsRequest represents the request for getting shipment items
type GetShipmentItemsRequest struct {
	QueryType         string `json:"queryType,omitempty"`
	MarketplaceId     string `json:"marketplaceId,omitempty"`
	ShipmentId        string `json:"shipmentId,omitempty"`
	LastUpdatedAfter  string `json:"lastUpdatedAfter,omitempty"`
	LastUpdatedBefore string `json:"lastUpdatedBefore,omitempty"`
	NextToken         string `json:"nextToken,omitempty"`
}

// GetShipmentItemsByShipmentIdRequest represents the request for getting shipment items by shipment ID
type GetShipmentItemsByShipmentIdRequest struct {
	MarketplaceId string `json:"marketplaceId,omitempty"`
	NextToken     string `json:"nextToken,omitempty"`
}

// GetPrepInstructionsRequest represents the request for getting prep instructions
type GetPrepInstructionsRequest struct {
	ShipToCountryCode string   `json:"shipToCountryCode,omitempty"`
	SellerSKUList     []string `json:"sellerSKUList,omitempty"`
	ASINList          []string `json:"asinList,omitempty"`
}

// GetLabelsRequest represents the request for getting labels
type GetLabelsRequest struct {
	PageType             string   `json:"pageType,omitempty"`
	LabelType            string   `json:"labelType,omitempty"`
	NumberOfPackages     int      `json:"numberOfPackages,omitempty"`
	PackageLabelsToPrint []string `json:"packageLabelsToPrint,omitempty"`
	NumberOfPallets      int      `json:"numberOfPallets,omitempty"`
	PageSize             int      `json:"pageSize,omitempty"`
	PageStartIndex       int      `json:"pageStartIndex,omitempty"`
}

// Response types

// GetShipmentsResponse represents the response for getting shipments
type GetShipmentsResponse struct {
	Payload *GetShipmentsPayload `json:"payload,omitempty"`
	Errors  []Error              `json:"errors,omitempty"`
}

// GetShipmentsPayload represents the payload for getting shipments
type GetShipmentsPayload struct {
	ShipmentData []ShipmentData `json:"shipmentData,omitempty"`
	NextToken    string         `json:"nextToken,omitempty"`
}

// ShipmentData represents shipment data
type ShipmentData struct {
	ShipmentId                     string                `json:"shipmentId,omitempty"`
	ShipmentName                   string                `json:"shipmentName,omitempty"`
	ShipFromAddress                Address               `json:"shipFromAddress,omitempty"`
	DestinationFulfillmentCenterId string                `json:"destinationFulfillmentCenterId,omitempty"`
	ShipmentStatus                 string                `json:"shipmentStatus,omitempty"`
	LabelPrepType                  string                `json:"labelPrepType,omitempty"`
	BoxContentsSource              string                `json:"boxContentsSource,omitempty"`
	TrackingId                     string                `json:"trackingId,omitempty"`
	CreatedDate                    string                `json:"createdDate,omitempty"`
	LastUpdatedDate                string                `json:"lastUpdatedDate,omitempty"`
	EstimatedArrivalDate           string                `json:"estimatedArrivalDate,omitempty"`
	ShippingId                     string                `json:"shippingId,omitempty"`
	ShippingAddress                Address               `json:"shippingAddress,omitempty"`
	InboundShipmentHeader          InboundShipmentHeader `json:"inboundShipmentHeader,omitempty"`
	InboundShipmentItems           []InboundShipmentItem `json:"inboundShipmentItems,omitempty"`
}

// Address represents an address
type Address struct {
	Name                string `json:"name,omitempty"`
	AddressLine1        string `json:"addressLine1,omitempty"`
	AddressLine2        string `json:"addressLine2,omitempty"`
	City                string `json:"city,omitempty"`
	StateOrProvinceCode string `json:"stateOrProvinceCode,omitempty"`
	PostalCode          string `json:"postalCode,omitempty"`
	CountryCode         string `json:"countryCode,omitempty"`
	Phone               string `json:"phone,omitempty"`
}

// InboundShipmentHeader represents inbound shipment header
type InboundShipmentHeader struct {
	ShipmentName                   string  `json:"shipmentName,omitempty"`
	ShipFromAddress                Address `json:"shipFromAddress,omitempty"`
	DestinationFulfillmentCenterId string  `json:"destinationFulfillmentCenterId,omitempty"`
	AreCasesRequired               bool    `json:"areCasesRequired,omitempty"`
	ShipmentStatus                 string  `json:"shipmentStatus,omitempty"`
	LabelPrepType                  string  `json:"labelPrepType,omitempty"`
	IntendedBoxContentsSource      string  `json:"intendedBoxContentsSource,omitempty"`
}

// InboundShipmentItem represents inbound shipment item
type InboundShipmentItem struct {
	ShipmentId            string        `json:"shipmentId,omitempty"`
	SellerSKU             string        `json:"sellerSKU,omitempty"`
	FulfillmentNetworkSKU string        `json:"fulfillmentNetworkSKU,omitempty"`
	QuantityShipped       int           `json:"quantityShipped,omitempty"`
	QuantityReceived      int           `json:"quantityReceived,omitempty"`
	QuantityInCase        int           `json:"quantityInCase,omitempty"`
	PrepDetailsList       []PrepDetails `json:"prepDetailsList,omitempty"`
}

// PrepDetails represents prep details
type PrepDetails struct {
	PrepInstruction string `json:"prepInstruction,omitempty"`
	PrepOwner       string `json:"prepOwner,omitempty"`
}

// GetShipmentItemsResponse represents the response for getting shipment items
type GetShipmentItemsResponse struct {
	Payload *GetShipmentItemsPayload `json:"payload,omitempty"`
	Errors  []Error                  `json:"errors,omitempty"`
}

// GetShipmentItemsPayload represents the payload for getting shipment items
type GetShipmentItemsPayload struct {
	ItemData  []InboundShipmentItem `json:"itemData,omitempty"`
	NextToken string                `json:"nextToken,omitempty"`
}

// GetPrepInstructionsResponse represents the response for getting prep instructions
type GetPrepInstructionsResponse struct {
	Payload *GetPrepInstructionsPayload `json:"payload,omitempty"`
	Errors  []Error                     `json:"errors,omitempty"`
}

// GetPrepInstructionsPayload represents the payload for getting prep instructions
type GetPrepInstructionsPayload struct {
	PrepInstructions []PrepInstruction `json:"prepInstructions,omitempty"`
}

// PrepInstruction represents prep instruction
type PrepInstruction struct {
	SellerSKU           string                `json:"sellerSKU,omitempty"`
	ASIN                string                `json:"asin,omitempty"`
	PrepGuidance        string                `json:"prepGuidance,omitempty"`
	PrepInstructionList []PrepInstructionItem `json:"prepInstructionList,omitempty"`
}

// PrepInstructionItem represents prep instruction item
type PrepInstructionItem struct {
	PrepInstruction string `json:"prepInstruction,omitempty"`
	PrepOwner       string `json:"prepOwner,omitempty"`
}

// GetLabelsResponse represents the response for getting labels
type GetLabelsResponse struct {
	Payload *GetLabelsPayload `json:"payload,omitempty"`
	Errors  []Error           `json:"errors,omitempty"`
}

// GetLabelsPayload represents the payload for getting labels
type GetLabelsPayload struct {
	DownloadURL string `json:"downloadURL,omitempty"`
}

// GetBillOfLadingResponse represents the response for getting bill of lading
type GetBillOfLadingResponse struct {
	Payload *GetBillOfLadingPayload `json:"payload,omitempty"`
	Errors  []Error                 `json:"errors,omitempty"`
}

// GetBillOfLadingPayload represents the payload for getting bill of lading
type GetBillOfLadingPayload struct {
	DownloadURL string `json:"downloadURL,omitempty"`
}

// Error represents an API error
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}
