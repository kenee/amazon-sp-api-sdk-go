package inventory

import "time"

// GetInventorySummariesRequest represents the request for getting inventory summaries
type GetInventorySummariesRequest struct {
	GranularityType string   `json:"granularityType"`
	GranularityId   string   `json:"granularityId"`
	MarketplaceIds  []string `json:"marketplaceIds"`
	SellerId        string   `json:"sellerId,omitempty"`
	Details         bool     `json:"details,omitempty"`
	StartDateTime   string   `json:"startDateTime,omitempty"`
	SellerSkus      []string `json:"sellerSkus,omitempty"`
	SellerSku       string   `json:"sellerSku,omitempty"`
	NextToken       string   `json:"nextToken,omitempty"`
}

// GetInventorySummariesResponse represents the response from getting inventory summaries
type GetInventorySummariesResponse struct {
	InventorySummaries []InventorySummary `json:"inventorySummaries"`
	NextToken          string             `json:"nextToken,omitempty"`
}

// InventorySummary represents an inventory summary
type InventorySummary struct {
	Asin             string           `json:"asin"`
	FnSku            string           `json:"fnSku,omitempty"`
	SellerSku        string           `json:"sellerSku,omitempty"`
	Condition        string           `json:"condition"`
	InventoryDetails InventoryDetails `json:"inventoryDetails"`
	LastUpdatedTime  time.Time        `json:"lastUpdatedTime"`
	ProductName      string           `json:"productName,omitempty"`
	TotalQuantity    int              `json:"totalQuantity"`
}

// InventoryDetails represents inventory details
type InventoryDetails struct {
	FulfillableQuantity      int                            `json:"fulfillableQuantity"`
	InboundWorkingQuantity   int                            `json:"inboundWorkingQuantity"`
	InboundShippedQuantity   int                            `json:"inboundShippedQuantity"`
	InboundReceivingQuantity int                            `json:"inboundReceivingQuantity"`
	ReservedQuantity         InventoryReservedQuantity      `json:"reservedQuantity"`
	ResearchingQuantity      InventoryResearchingQuantity   `json:"researchingQuantity"`
	UnfulfillableQuantity    InventoryUnfulfillableQuantity `json:"unfulfillableQuantity"`
}

// InventoryReservedQuantity represents reserved quantity
type InventoryReservedQuantity struct {
	TotalReservedQuantity        int `json:"totalReservedQuantity"`
	PendingCustomerOrderQuantity int `json:"pendingCustomerOrderQuantity"`
	PendingTransshipmentQuantity int `json:"pendingTransshipmentQuantity"`
	FcProcessingQuantity         int `json:"fcProcessingQuantity"`
}

// InventoryResearchingQuantity represents researching quantity
type InventoryResearchingQuantity struct {
	TotalResearchingQuantity     int                                     `json:"totalResearchingQuantity"`
	ResearchingQuantityBreakdown []InventoryResearchingQuantityBreakdown `json:"researchingQuantityBreakdown"`
}

// InventoryResearchingQuantityBreakdown represents researching quantity breakdown
type InventoryResearchingQuantityBreakdown struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

// InventoryUnfulfillableQuantity represents unfulfillable quantity
type InventoryUnfulfillableQuantity struct {
	TotalUnfulfillableQuantity int `json:"totalUnfulfillableQuantity"`
	CustomerDamagedQuantity    int `json:"customerDamagedQuantity"`
	WarehouseDamagedQuantity   int `json:"warehouseDamagedQuantity"`
	DistributorDamagedQuantity int `json:"distributorDamagedQuantity"`
	CarrierDamagedQuantity     int `json:"carrierDamagedQuantity"`
	DefectiveQuantity          int `json:"defectiveQuantity"`
	ExpiredQuantity            int `json:"expiredQuantity"`
}
