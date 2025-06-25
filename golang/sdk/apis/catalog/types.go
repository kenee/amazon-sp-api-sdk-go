package catalog

// GetCatalogItemRequest represents the request parameters for getCatalogItem
type GetCatalogItemRequest struct {
	ASIN           string   `json:"asin"`
	MarketplaceIds []string `json:"marketplaceIds"`
	IncludedData   []string `json:"includedData,omitempty"`
	Locale         string   `json:"locale,omitempty"`
}

// GetCatalogItemResponse represents the response from getCatalogItem
type GetCatalogItemResponse struct {
	ASIN            string               `json:"asin"`
	Summaries       []ItemSummary        `json:"summaries,omitempty"`
	Attributes      []ItemAttribute      `json:"attributes,omitempty"`
	Images          []ItemImage          `json:"images,omitempty"`
	Classifications []ItemClassification `json:"classifications,omitempty"`
	Dimensions      []ItemDimension      `json:"dimensions,omitempty"`
	Identifiers     []ItemIdentifier     `json:"identifiers,omitempty"`
	ProductTypes    []ItemProductType    `json:"productTypes,omitempty"`
	Relationships   []ItemRelationship   `json:"relationships,omitempty"`
	SalesRanks      []ItemSalesRank      `json:"salesRanks,omitempty"`
	VendorDetails   *ItemVendorDetails   `json:"vendorDetails,omitempty"`
}

// SearchCatalogItemsRequest represents the request parameters for searchCatalogItems
type SearchCatalogItemsRequest struct {
	MarketplaceIds    []string `json:"marketplaceIds"`
	Identifiers       []string `json:"identifiers,omitempty"`
	IdentifiersType   string   `json:"identifiersType,omitempty"`
	IncludedData      []string `json:"includedData,omitempty"`
	Locale            string   `json:"locale,omitempty"`
	SellerId          string   `json:"sellerId,omitempty"`
	Keywords          []string `json:"keywords,omitempty"`
	BrandNames        []string `json:"brandNames,omitempty"`
	ClassificationIds []string `json:"classificationIds,omitempty"`
	PageSize          int      `json:"pageSize,omitempty"`
	PageToken         string   `json:"pageToken,omitempty"`
	KeywordsLocale    string   `json:"keywordsLocale,omitempty"`
}

// SearchCatalogItemsResponse represents the response from searchCatalogItems
type SearchCatalogItemsResponse struct {
	Items      []CatalogItem `json:"items"`
	Pagination *Pagination   `json:"pagination,omitempty"`
}

// CatalogItem represents a catalog item
type CatalogItem struct {
	ASIN            string               `json:"asin"`
	Summaries       []ItemSummary        `json:"summaries,omitempty"`
	Attributes      []ItemAttribute      `json:"attributes,omitempty"`
	Images          []ItemImage          `json:"images,omitempty"`
	Classifications []ItemClassification `json:"classifications,omitempty"`
	Dimensions      []ItemDimension      `json:"dimensions,omitempty"`
	Identifiers     []ItemIdentifier     `json:"identifiers,omitempty"`
	ProductTypes    []ItemProductType    `json:"productTypes,omitempty"`
	Relationships   []ItemRelationship   `json:"relationships,omitempty"`
	SalesRanks      []ItemSalesRank      `json:"salesRanks,omitempty"`
	VendorDetails   *ItemVendorDetails   `json:"vendorDetails,omitempty"`
}

// ItemSummary represents an item summary
type ItemSummary struct {
	MarketplaceId string                 `json:"marketplaceId"`
	ASIN          string                 `json:"asin"`
	ProductType   string                 `json:"productType"`
	ItemName      string                 `json:"itemName,omitempty"`
	MainImage     *ItemImage             `json:"mainImage,omitempty"`
	Summaries     []ItemSummaryAttribute `json:"summaries,omitempty"`
}

// ItemSummaryAttribute represents an item summary attribute
type ItemSummaryAttribute struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// ItemAttribute represents an item attribute
type ItemAttribute struct {
	MarketplaceId string                 `json:"marketplaceId"`
	ASIN          string                 `json:"asin"`
	ProductType   string                 `json:"productType"`
	Attributes    map[string]interface{} `json:"attributes,omitempty"`
}

// ItemImage represents an item image
type ItemImage struct {
	Link   string `json:"link"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

// ItemClassification represents an item classification
type ItemClassification struct {
	MarketplaceId   string           `json:"marketplaceId"`
	ASIN            string           `json:"asin"`
	ProductType     string           `json:"productType"`
	Classifications []Classification `json:"classifications,omitempty"`
}

// Classification represents a classification
type Classification struct {
	ProductType           string                 `json:"productType"`
	ProductTypeId         string                 `json:"productTypeId,omitempty"`
	ProductTypeName       string                 `json:"productTypeName,omitempty"`
	ProductTypeVersion    string                 `json:"productTypeVersion,omitempty"`
	ProductTypeAttributes map[string]interface{} `json:"productTypeAttributes,omitempty"`
}

// ItemDimension represents an item dimension
type ItemDimension struct {
	MarketplaceId string                 `json:"marketplaceId"`
	ASIN          string                 `json:"asin"`
	ProductType   string                 `json:"productType"`
	Dimensions    map[string]interface{} `json:"dimensions,omitempty"`
}

// ItemIdentifier represents an item identifier
type ItemIdentifier struct {
	MarketplaceId string                 `json:"marketplaceId"`
	ASIN          string                 `json:"asin"`
	ProductType   string                 `json:"productType"`
	Identifiers   map[string]interface{} `json:"identifiers,omitempty"`
}

// ItemProductType represents an item product type
type ItemProductType struct {
	MarketplaceId         string                 `json:"marketplaceId"`
	ASIN                  string                 `json:"asin"`
	ProductType           string                 `json:"productType"`
	ProductTypeAttributes map[string]interface{} `json:"productTypeAttributes,omitempty"`
}

// ItemRelationship represents an item relationship
type ItemRelationship struct {
	MarketplaceId string         `json:"marketplaceId"`
	ASIN          string         `json:"asin"`
	ProductType   string         `json:"productType"`
	Relationships []Relationship `json:"relationships,omitempty"`
}

// Relationship represents a relationship
type Relationship struct {
	ASIN             string `json:"asin"`
	RelationshipType string `json:"relationshipType,omitempty"`
	RelationshipName string `json:"relationshipName,omitempty"`
}

// ItemSalesRank represents an item sales rank
type ItemSalesRank struct {
	MarketplaceId string      `json:"marketplaceId"`
	ASIN          string      `json:"asin"`
	ProductType   string      `json:"productType"`
	SalesRanks    []SalesRank `json:"salesRanks,omitempty"`
}

// SalesRank represents a sales rank
type SalesRank struct {
	ProductCategoryId string `json:"productCategoryId"`
	Rank              int    `json:"rank"`
	Title             string `json:"title,omitempty"`
	Link              string `json:"link,omitempty"`
}

// ItemVendorDetails represents item vendor details
type ItemVendorDetails struct {
	MarketplaceId string                 `json:"marketplaceId"`
	ASIN          string                 `json:"asin"`
	ProductType   string                 `json:"productType"`
	VendorDetails map[string]interface{} `json:"vendorDetails,omitempty"`
}

// Pagination represents pagination information
type Pagination struct {
	NextToken     string `json:"nextToken,omitempty"`
	PreviousToken string `json:"previousToken,omitempty"`
}
