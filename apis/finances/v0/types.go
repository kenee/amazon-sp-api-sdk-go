package v0

// Response types

// ListFinancialEventsResponse represents the response for listing financial events
type ListFinancialEventsResponse struct {
	Payload *ListFinancialEventsPayload `json:"payload,omitempty"`
	Errors  []Error                     `json:"errors,omitempty"`
}

// ListFinancialEventsPayload represents the payload for listing financial events
type ListFinancialEventsPayload struct {
	NextToken       string           `json:"nextToken,omitempty"`
	FinancialEvents *FinancialEvents `json:"financialEvents,omitempty"`
}

// ListFinancialEventGroupsResponse represents the response for listing financial event groups
type ListFinancialEventGroupsResponse struct {
	Payload *ListFinancialEventGroupsPayload `json:"payload,omitempty"`
	Errors  []Error                          `json:"errors,omitempty"`
}

// ListFinancialEventGroupsPayload represents the payload for listing financial event groups
type ListFinancialEventGroupsPayload struct {
	NextToken               string                `json:"nextToken,omitempty"`
	FinancialEventGroupList []FinancialEventGroup `json:"financialEventGroupList,omitempty"`
}

// FinancialEvents represents all types of financial events
type FinancialEvents struct {
	ShipmentEventList                      []ShipmentEvent                      `json:"shipmentEventList,omitempty"`
	ShipmentSettleEventList                []ShipmentEvent                      `json:"shipmentSettleEventList,omitempty"`
	RefundEventList                        []ShipmentEvent                      `json:"refundEventList,omitempty"`
	GuaranteeClaimEventList                []ShipmentEvent                      `json:"guaranteeClaimEventList,omitempty"`
	ChargebackEventList                    []ShipmentEvent                      `json:"chargebackEventList,omitempty"`
	PayWithAmazonEventList                 []PayWithAmazonEvent                 `json:"payWithAmazonEventList,omitempty"`
	ServiceProviderCreditEventList         []SolutionProviderCreditEvent        `json:"serviceProviderCreditEventList,omitempty"`
	RetrochargeEventList                   []RetrochargeEvent                   `json:"retrochargeEventList,omitempty"`
	RentalTransactionEventList             []RentalTransactionEvent             `json:"rentalTransactionEventList,omitempty"`
	ProductAdsPaymentEventList             []ProductAdsPaymentEvent             `json:"productAdsPaymentEventList,omitempty"`
	ServiceFeeEventList                    []ServiceFeeEvent                    `json:"serviceFeeEventList,omitempty"`
	SellerDealPaymentEventList             []SellerDealPaymentEvent             `json:"sellerDealPaymentEventList,omitempty"`
	DebtRecoveryEventList                  []DebtRecoveryEvent                  `json:"debtRecoveryEventList,omitempty"`
	LoanServicingEventList                 []LoanServicingEvent                 `json:"loanServicingEventList,omitempty"`
	AdjustmentEventList                    []AdjustmentEvent                    `json:"adjustmentEventList,omitempty"`
	SAFETReimbursementEventList            []SAFETReimbursementEvent            `json:"safetReimbursementEventList,omitempty"`
	SellerReviewEnrollmentPaymentEventList []SellerReviewEnrollmentPaymentEvent `json:"sellerReviewEnrollmentPaymentEventList,omitempty"`
	FBALiquidationEventList                []FBALiquidationEvent                `json:"fbaLiquidationEventList,omitempty"`
	CouponPaymentEventList                 []CouponPaymentEvent                 `json:"couponPaymentEventList,omitempty"`
	ImagingServicesFeeEventList            []ImagingServicesFeeEvent            `json:"imagingServicesFeeEventList,omitempty"`
	NetworkComminglingTransactionEventList []NetworkComminglingTransactionEvent `json:"networkComminglingTransactionEventList,omitempty"`
	AffordabilityExpenseEventList          []AffordabilityExpenseEvent          `json:"affordabilityExpenseEventList,omitempty"`
	AffordabilityExpenseReversalEventList  []AffordabilityExpenseEvent          `json:"affordabilityExpenseReversalEventList,omitempty"`
	RemovalShipmentEventList               []RemovalShipmentEvent               `json:"removalShipmentEventList,omitempty"`
	RemovalShipmentAdjustmentEventList     []RemovalShipmentAdjustmentEvent     `json:"removalShipmentAdjustmentEventList,omitempty"`
	TrialShipmentEventList                 []TrialShipmentEvent                 `json:"trialShipmentEventList,omitempty"`
	TDSReimbursementEventList              []TDSReimbursementEvent              `json:"tdsReimbursementEventList,omitempty"`
	AdhocDisbursementEventList             []AdhocDisbursementEvent             `json:"adhocDisbursementEventList,omitempty"`
	TaxWithholdingEventList                []TaxWithholdingEvent                `json:"taxWithholdingEventList,omitempty"`
	ChargeRefundEventList                  []ChargeRefundEvent                  `json:"chargeRefundEventList,omitempty"`
	FailedAdhocDisbursementEventList       []FailedAdhocDisbursementEvent       `json:"failedAdhocDisbursementEventList,omitempty"`
	ValueAddedServiceChargeEventList       []ValueAddedServiceChargeEvent       `json:"valueAddedServiceChargeEventList,omitempty"`
	CapacityReservationBillingEventList    []CapacityReservationBillingEvent    `json:"capacityReservationBillingEventList,omitempty"`
}

// FinancialEventGroup represents a financial event group
type FinancialEventGroup struct {
	FinancialEventGroupId    string    `json:"financialEventGroupId,omitempty"`
	ProcessingStatus         string    `json:"processingStatus,omitempty"`
	FundTransferStatus       string    `json:"fundTransferStatus,omitempty"`
	OriginalTotal            *Currency `json:"originalTotal,omitempty"`
	ConvertedTotal           *Currency `json:"convertedTotal,omitempty"`
	FundTransferDate         string    `json:"fundTransferDate,omitempty"`
	TraceId                  string    `json:"traceId,omitempty"`
	AccountTail              string    `json:"accountTail,omitempty"`
	BeginningBalance         *Currency `json:"beginningBalance,omitempty"`
	FinancialEventGroupStart string    `json:"financialEventGroupStart,omitempty"`
	FinancialEventGroupEnd   string    `json:"financialEventGroupEnd,omitempty"`
}

// Currency represents a currency amount
type Currency struct {
	CurrencyCode string `json:"currencyCode,omitempty"`
	Amount       string `json:"amount,omitempty"`
}

// Error represents an API error
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Request types for API calls

// ListFinancialEventGroupsRequest represents the request for listing financial event groups
type ListFinancialEventGroupsRequest struct {
	MaxResultsPerPage                int    `json:"maxResultsPerPage,omitempty"`
	FinancialEventGroupStartedBefore string `json:"financialEventGroupStartedBefore,omitempty"`
	FinancialEventGroupStartedAfter  string `json:"financialEventGroupStartedAfter,omitempty"`
	NextToken                        string `json:"nextToken,omitempty"`
}

// ListFinancialEventsRequest represents the request for listing financial events
type ListFinancialEventsRequest struct {
	MaxResultsPerPage int    `json:"maxResultsPerPage,omitempty"`
	PostedAfter       string `json:"postedAfter,omitempty"`
	PostedBefore      string `json:"postedBefore,omitempty"`
	NextToken         string `json:"nextToken,omitempty"`
}

// ListFinancialEventsByGroupIdRequest represents the request for listing financial events by group ID
type ListFinancialEventsByGroupIdRequest struct {
	MaxResultsPerPage int    `json:"maxResultsPerPage,omitempty"`
	PostedAfter       string `json:"postedAfter,omitempty"`
	PostedBefore      string `json:"postedBefore,omitempty"`
	NextToken         string `json:"nextToken,omitempty"`
}

// ListFinancialEventsByOrderIdRequest represents the request for listing financial events by order ID
type ListFinancialEventsByOrderIdRequest struct {
	MaxResultsPerPage int    `json:"maxResultsPerPage,omitempty"`
	NextToken         string `json:"nextToken,omitempty"`
}
