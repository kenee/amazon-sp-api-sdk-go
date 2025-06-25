package orders

// GetOrdersRequest represents the request parameters for getOrders
type GetOrdersRequest struct {
	MarketplaceIds                  []string `json:"marketplaceIds"`
	CreatedAfter                    string   `json:"createdAfter,omitempty"`
	CreatedBefore                   string   `json:"createdBefore,omitempty"`
	LastUpdatedAfter                string   `json:"lastUpdatedAfter,omitempty"`
	LastUpdatedBefore               string   `json:"lastUpdatedBefore,omitempty"`
	OrderStatuses                   []string `json:"orderStatuses,omitempty"`
	FulfillmentChannels             []string `json:"fulfillmentChannels,omitempty"`
	PaymentMethods                  []string `json:"paymentMethods,omitempty"`
	BuyerEmail                      string   `json:"buyerEmail,omitempty"`
	SellerOrderId                   string   `json:"sellerOrderId,omitempty"`
	MaxResultsPerPage               int      `json:"maxResultsPerPage,omitempty"`
	EasyShipShipmentStatuses        []string `json:"easyShipShipmentStatuses,omitempty"`
	ElectronicInvoiceStatuses       []string `json:"electronicInvoiceStatuses,omitempty"`
	NextToken                       string   `json:"nextToken,omitempty"`
	AmazonOrderIds                  []string `json:"amazonOrderIds,omitempty"`
	ActualFulfillmentSupplySourceId string   `json:"actualFulfillmentSupplySourceId,omitempty"`
	IsISPU                          *bool    `json:"isISPU,omitempty"`
	StoreChainStoreId               string   `json:"storeChainStoreId,omitempty"`
	RestrictedDataToken             string   `json:"restrictedDataToken,omitempty"`
}

// GetOrdersResponse represents the response from getOrders
type GetOrdersResponse struct {
	Payload *OrdersPayload `json:"payload,omitempty"`
}

// OrdersPayload represents the payload in the getOrders response
type OrdersPayload struct {
	Orders        []Order `json:"orders"`
	NextToken     string  `json:"nextToken,omitempty"`
	CreatedBefore string  `json:"createdBefore,omitempty"`
}

// Order represents an order
type Order struct {
	AmazonOrderId                  string                     `json:"amazonOrderId"`
	SellerOrderId                  string                     `json:"sellerOrderId,omitempty"`
	PurchaseDate                   string                     `json:"purchaseDate"`
	LastUpdateDate                 string                     `json:"lastUpdateDate"`
	OrderStatus                    string                     `json:"orderStatus"`
	FulfillmentChannel             string                     `json:"fulfillmentChannel,omitempty"`
	SalesChannel                   string                     `json:"salesChannel,omitempty"`
	OrderChannel                   string                     `json:"orderChannel,omitempty"`
	ShipServiceLevel               string                     `json:"shipServiceLevel,omitempty"`
	OrderTotal                     *Money                     `json:"orderTotal,omitempty"`
	NumberOfItemsShipped           int                        `json:"numberOfItemsShipped,omitempty"`
	NumberOfItemsUnshipped         int                        `json:"numberOfItemsUnshipped,omitempty"`
	PaymentMethod                  string                     `json:"paymentMethod,omitempty"`
	MarketplaceId                  string                     `json:"marketplaceId,omitempty"`
	ShipmentServiceLevelCategory   string                     `json:"shipmentServiceLevelCategory,omitempty"`
	OrderType                      string                     `json:"orderType,omitempty"`
	EarliestShipDate               string                     `json:"earliestShipDate,omitempty"`
	LatestShipDate                 string                     `json:"latestShipDate,omitempty"`
	EarliestDeliveryDate           string                     `json:"earliestDeliveryDate,omitempty"`
	LatestDeliveryDate             string                     `json:"latestDeliveryDate,omitempty"`
	IsBusinessOrder                bool                       `json:"isBusinessOrder,omitempty"`
	IsPrime                        bool                       `json:"isPrime,omitempty"`
	IsGlobalExpressEnabled         bool                       `json:"isGlobalExpressEnabled,omitempty"`
	IsReplacementOrder             bool                       `json:"isReplacementOrder,omitempty"`
	IsAccessPointOrder             bool                       `json:"isAccessPointOrder,omitempty"`
	IsISPU                         bool                       `json:"isISPU,omitempty"`
	IsPremiumOrder                 bool                       `json:"isPremiumOrder,omitempty"`
	IsPrimeAccess                  bool                       `json:"isPrimeAccess,omitempty"`
	IsPendingVulnerabilityScan     bool                       `json:"isPendingVulnerabilityScan,omitempty"`
	IsEstimatedShipDateSet         bool                       `json:"isEstimatedShipDateSet,omitempty"`
	IsSoldByAB                     bool                       `json:"isSoldByAB,omitempty"`
	DefaultShipFromLocationAddress *Address                   `json:"defaultShipFromLocationAddress,omitempty"`
	FulfillmentInstruction         *FulfillmentInstruction    `json:"fulfillmentInstruction,omitempty"`
	IsISPUAccessPointOrder         bool                       `json:"isISPUAccessPointOrder,omitempty"`
	MarketplaceTaxInfo             *MarketplaceTaxInfo        `json:"marketplaceTaxInfo,omitempty"`
	SellerDisplayName              string                     `json:"sellerDisplayName,omitempty"`
	ShippingAddress                *Address                   `json:"shippingAddress,omitempty"`
	BuyerInfo                      *BuyerInfo                 `json:"buyerInfo,omitempty"`
	AutomatedShippingSettings      *AutomatedShippingSettings `json:"automatedShippingSettings,omitempty"`
	HasRegulatedItems              bool                       `json:"hasRegulatedItems,omitempty"`
	ElectronicInvoiceStatus        string                     `json:"electronicInvoiceStatus,omitempty"`
	ItemApprovalTypes              []string                   `json:"itemApprovalTypes,omitempty"`
	ItemApprovalStatus             []string                   `json:"itemApprovalStatus,omitempty"`
	RegulatedInformation           *RegulatedInformation      `json:"regulatedInformation,omitempty"`
	PointsGranted                  *PointsGranted             `json:"pointsGranted,omitempty"`
}

// Money represents a monetary amount
type Money struct {
	CurrencyCode string `json:"currencyCode"`
	Amount       string `json:"amount"`
}

// Address represents an address
type Address struct {
	Name          string `json:"name,omitempty"`
	AddressLine1  string `json:"addressLine1,omitempty"`
	AddressLine2  string `json:"addressLine2,omitempty"`
	AddressLine3  string `json:"addressLine3,omitempty"`
	City          string `json:"city,omitempty"`
	County        string `json:"county,omitempty"`
	District      string `json:"district,omitempty"`
	StateOrRegion string `json:"stateOrRegion,omitempty"`
	Municipality  string `json:"municipality,omitempty"`
	PostalCode    string `json:"postalCode,omitempty"`
	CountryCode   string `json:"countryCode,omitempty"`
	Phone         string `json:"phone,omitempty"`
	AddressType   string `json:"addressType,omitempty"`
}

// FulfillmentInstruction represents fulfillment instructions
type FulfillmentInstruction struct {
	FulfillmentSupplySourceId string `json:"fulfillmentSupplySourceId,omitempty"`
}

// MarketplaceTaxInfo represents marketplace tax information
type MarketplaceTaxInfo struct {
	TaxClassifications []TaxClassification `json:"taxClassifications,omitempty"`
}

// TaxClassification represents tax classification
type TaxClassification struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// BuyerInfo represents buyer information
type BuyerInfo struct {
	BuyerEmail          string        `json:"buyerEmail,omitempty"`
	BuyerName           string        `json:"buyerName,omitempty"`
	BuyerCounty         string        `json:"buyerCounty,omitempty"`
	BuyerTaxInfo        *BuyerTaxInfo `json:"buyerTaxInfo,omitempty"`
	PurchaseOrderNumber string        `json:"purchaseOrderNumber,omitempty"`
}

// BuyerTaxInfo represents buyer tax information
type BuyerTaxInfo struct {
	CompanyLegalName   string              `json:"companyLegalName,omitempty"`
	TaxingRegion       string              `json:"taxingRegion,omitempty"`
	TaxClassifications []TaxClassification `json:"taxClassifications,omitempty"`
}

// AutomatedShippingSettings represents automated shipping settings
type AutomatedShippingSettings struct {
	HasAutomatedShippingSettings bool   `json:"hasAutomatedShippingSettings,omitempty"`
	AutomatedCarrier             string `json:"automatedCarrier,omitempty"`
	AutomatedShipMethod          string `json:"automatedShipMethod,omitempty"`
}

// RegulatedInformation represents regulated information
type RegulatedInformation struct {
	Fields []RegulatedInformationField `json:"fields,omitempty"`
}

// RegulatedInformationField represents a regulated information field
type RegulatedInformationField struct {
	FieldId    string `json:"fieldId,omitempty"`
	FieldLabel string `json:"fieldLabel,omitempty"`
	FieldType  string `json:"fieldType,omitempty"`
	FieldValue string `json:"fieldValue,omitempty"`
}

// PointsGranted represents points granted
type PointsGranted struct {
	PointsNumber        int    `json:"pointsNumber,omitempty"`
	PointsMonetaryValue *Money `json:"pointsMonetaryValue,omitempty"`
}
