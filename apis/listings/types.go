package listings

// GetListingsItemRequest represents the request parameters for getListingsItem
type GetListingsItemRequest struct {
	SellerId       string   `json:"sellerId"`
	SKU            string   `json:"sku"`
	MarketplaceIds []string `json:"marketplaceIds"`
	IssueLocale    string   `json:"issueLocale,omitempty"`
	IncludedData   []string `json:"includedData,omitempty"`
}

// GetListingsItemResponse represents the response from getListingsItem
type GetListingsItemResponse struct {
	SKU                     string                    `json:"sku"`
	Summaries               []ListingSummary          `json:"summaries,omitempty"`
	Attributes              []ListingAttribute        `json:"attributes,omitempty"`
	Issues                  []ListingIssue            `json:"issues,omitempty"`
	Offers                  []ListingOffer            `json:"offers,omitempty"`
	FulfillmentAvailability []FulfillmentAvailability `json:"fulfillmentAvailability,omitempty"`
	Procurement             []ListingProcurement      `json:"procurement,omitempty"`
}

// ListingSummary represents a listing summary
type ListingSummary struct {
	MarketplaceId string                    `json:"marketplaceId"`
	ASIN          string                    `json:"asin"`
	ProductType   string                    `json:"productType"`
	ItemName      string                    `json:"itemName,omitempty"`
	MainImage     *ListingImage             `json:"mainImage,omitempty"`
	Summaries     []ListingSummaryAttribute `json:"summaries,omitempty"`
}

// ListingSummaryAttribute represents a listing summary attribute
type ListingSummaryAttribute struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// ListingAttribute represents a listing attribute
type ListingAttribute struct {
	MarketplaceId string                 `json:"marketplaceId"`
	ASIN          string                 `json:"asin"`
	ProductType   string                 `json:"productType"`
	Attributes    map[string]interface{} `json:"attributes,omitempty"`
}

// ListingImage represents a listing image
type ListingImage struct {
	Link   string `json:"link"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

// ListingIssue represents a listing issue
type ListingIssue struct {
	Code          string `json:"code"`
	Message       string `json:"message"`
	Severity      string `json:"severity"`
	AttributeName string `json:"attributeName,omitempty"`
}

// ListingOffer represents a listing offer
type ListingOffer struct {
	MarketplaceId string  `json:"marketplaceId"`
	ASIN          string  `json:"asin"`
	ProductType   string  `json:"productType"`
	Offers        []Offer `json:"offers,omitempty"`
}

// Offer represents an offer
type Offer struct {
	MarketplaceId        string                `json:"marketplaceId"`
	ASIN                 string                `json:"asin"`
	ProductType          string                `json:"productType"`
	OfferType            string                `json:"offerType,omitempty"`
	Price                *Money                `json:"price,omitempty"`
	Shipping             *Shipping             `json:"shipping,omitempty"`
	Points               *Points               `json:"points,omitempty"`
	SellerFeedbackRating *SellerFeedbackRating `json:"sellerFeedbackRating,omitempty"`
	ShippingTime         *ShippingTime         `json:"shippingTime,omitempty"`
	PrimeInformation     *PrimeInformation     `json:"primeInformation,omitempty"`
	B2B                  *B2B                  `json:"b2b,omitempty"`
}

// Money represents a monetary amount
type Money struct {
	CurrencyCode string `json:"currencyCode"`
	Amount       string `json:"amount"`
}

// Shipping represents shipping information
type Shipping struct {
	CurrencyCode string `json:"currencyCode"`
	Amount       string `json:"amount"`
}

// Points represents points information
type Points struct {
	PointsNumber        int    `json:"pointsNumber,omitempty"`
	PointsMonetaryValue *Money `json:"pointsMonetaryValue,omitempty"`
}

// SellerFeedbackRating represents seller feedback rating
type SellerFeedbackRating struct {
	SellerPositiveFeedbackRating float64 `json:"sellerPositiveFeedbackRating,omitempty"`
	FeedbackCount                int     `json:"feedbackCount,omitempty"`
}

// ShippingTime represents shipping time information
type ShippingTime struct {
	MinimumHours     int    `json:"minimumHours,omitempty"`
	MaximumHours     int    `json:"maximumHours,omitempty"`
	AvailableDate    string `json:"availableDate,omitempty"`
	AvailabilityType string `json:"availabilityType,omitempty"`
}

// PrimeInformation represents prime information
type PrimeInformation struct {
	IsPrime         bool `json:"isPrime,omitempty"`
	IsNationalPrime bool `json:"isNationalPrime,omitempty"`
}

// B2B represents B2B information
type B2B struct {
	IsPrime bool `json:"isPrime,omitempty"`
	IsB2B   bool `json:"isB2B,omitempty"`
}

// FulfillmentAvailability represents fulfillment availability
type FulfillmentAvailability struct {
	MarketplaceId           string         `json:"marketplaceId"`
	ASIN                    string         `json:"asin"`
	ProductType             string         `json:"productType"`
	FulfillmentAvailability []Availability `json:"fulfillmentAvailability,omitempty"`
}

// Availability represents availability information
type Availability struct {
	FulfillmentChannel string `json:"fulfillmentChannel,omitempty"`
	Quantity           int    `json:"quantity,omitempty"`
}

// ListingProcurement represents listing procurement
type ListingProcurement struct {
	MarketplaceId string                 `json:"marketplaceId"`
	ASIN          string                 `json:"asin"`
	ProductType   string                 `json:"productType"`
	Procurement   map[string]interface{} `json:"procurement,omitempty"`
}
