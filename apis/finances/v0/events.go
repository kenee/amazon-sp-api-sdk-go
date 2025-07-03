package v0

// ShipmentEvent represents a shipment event
type ShipmentEvent struct {
	AmazonOrderId           string            `json:"amazonOrderId,omitempty"`
	SellerOrderId           string            `json:"sellerOrderId,omitempty"`
	MarketplaceName         string            `json:"marketplaceName,omitempty"`
	OrderChargeList         []ChargeComponent `json:"orderChargeList,omitempty"`
	OrderChargeAdjustments  []ChargeComponent `json:"orderChargeAdjustments,omitempty"`
	ShipmentFeeList         []FeeComponent    `json:"shipmentFeeList,omitempty"`
	ShipmentFeeAdjustments  []FeeComponent    `json:"shipmentFeeAdjustments,omitempty"`
	OrderFeeList            []FeeComponent    `json:"orderFeeList,omitempty"`
	OrderFeeAdjustments     []FeeComponent    `json:"orderFeeAdjustments,omitempty"`
	DirectPaymentList       []DirectPayment   `json:"directPaymentList,omitempty"`
	PostedDate              string            `json:"postedDate,omitempty"`
	ShipmentItemList        []ShipmentItem    `json:"shipmentItemList,omitempty"`
	ShipmentItemAdjustments []ShipmentItem    `json:"shipmentItemAdjustments,omitempty"`
}

// ShipmentItem represents a shipment item
type ShipmentItem struct {
	SellerSKU               string                 `json:"sellerSKU,omitempty"`
	OrderItemId             string                 `json:"orderItemId,omitempty"`
	OrderAdjustmentItemId   string                 `json:"orderAdjustmentItemId,omitempty"`
	QuantityShipped         int                    `json:"quantityShipped,omitempty"`
	ItemChargeList          []ChargeComponent      `json:"itemChargeList,omitempty"`
	ItemChargeAdjustments   []ChargeComponent      `json:"itemChargeAdjustments,omitempty"`
	ItemFeeList             []FeeComponent         `json:"itemFeeList,omitempty"`
	ItemFeeAdjustments      []FeeComponent         `json:"itemFeeAdjustments,omitempty"`
	ItemTaxWithheldList     []TaxWithheldComponent `json:"itemTaxWithheldList,omitempty"`
	PromotionList           []Promotion            `json:"promotionList,omitempty"`
	PromotionAdjustmentList []Promotion            `json:"promotionAdjustmentList,omitempty"`
	CostOfPointsGranted     *Currency              `json:"costOfPointsGranted,omitempty"`
	CostOfPointsReturned    *Currency              `json:"costOfPointsReturned,omitempty"`
}

// ChargeComponent represents a charge component
type ChargeComponent struct {
	ChargeType   string    `json:"chargeType,omitempty"`
	ChargeAmount *Currency `json:"chargeAmount,omitempty"`
}

// FeeComponent represents a fee component
type FeeComponent struct {
	FeeType   string    `json:"feeType,omitempty"`
	FeeAmount *Currency `json:"feeAmount,omitempty"`
}

// TaxWithheldComponent represents a tax withheld component
type TaxWithheldComponent struct {
	TaxCollectionModel string            `json:"taxCollectionModel,omitempty"`
	TaxesWithheld      []ChargeComponent `json:"taxesWithheld,omitempty"`
}

// Promotion represents a promotion
type Promotion struct {
	PromotionType   string    `json:"promotionType,omitempty"`
	PromotionId     string    `json:"promotionId,omitempty"`
	PromotionAmount *Currency `json:"promotionAmount,omitempty"`
}

// DirectPayment represents a direct payment
type DirectPayment struct {
	DirectPaymentType   string    `json:"directPaymentType,omitempty"`
	DirectPaymentAmount *Currency `json:"directPaymentAmount,omitempty"`
}

// PayWithAmazonEvent represents a Pay with Amazon event
type PayWithAmazonEvent struct {
	SellerOrderId         string           `json:"sellerOrderId,omitempty"`
	TransactionPostedDate string           `json:"transactionPostedDate,omitempty"`
	BusinessObjectType    string           `json:"businessObjectType,omitempty"`
	SalesChannel          string           `json:"salesChannel,omitempty"`
	Charge                *ChargeComponent `json:"charge,omitempty"`
	FeeList               []FeeComponent   `json:"feeList,omitempty"`
	PaymentAmountType     string           `json:"paymentAmountType,omitempty"`
	AmountDescription     string           `json:"amountDescription,omitempty"`
	FulfillmentChannel    string           `json:"fulfillmentChannel,omitempty"`
	StoreName             string           `json:"storeName,omitempty"`
}

// SolutionProviderCreditEvent represents a solution provider credit event
type SolutionProviderCreditEvent struct {
	ProviderTransactionType string    `json:"providerTransactionType,omitempty"`
	SellerOrderId           string    `json:"sellerOrderId,omitempty"`
	MarketplaceId           string    `json:"marketplaceId,omitempty"`
	MarketplaceCountryCode  string    `json:"marketplaceCountryCode,omitempty"`
	SellerId                string    `json:"sellerId,omitempty"`
	SellerStoreName         string    `json:"sellerStoreName,omitempty"`
	ProviderId              string    `json:"providerId,omitempty"`
	ProviderStoreName       string    `json:"providerStoreName,omitempty"`
	TransactionAmount       *Currency `json:"transactionAmount,omitempty"`
	TransactionCreationDate string    `json:"transactionCreationDate,omitempty"`
}

// RetrochargeEvent represents a retrocharge event
type RetrochargeEvent struct {
	RetrochargeEventType       string                 `json:"retrochargeEventType,omitempty"`
	AmazonOrderId              string                 `json:"amazonOrderId,omitempty"`
	PostedDate                 string                 `json:"postedDate,omitempty"`
	BaseTax                    *Currency              `json:"baseTax,omitempty"`
	ShippingTax                *Currency              `json:"shippingTax,omitempty"`
	MarketplaceName            string                 `json:"marketplaceName,omitempty"`
	RetrochargeTaxWithheldList []TaxWithheldComponent `json:"retrochargeTaxWithheldList,omitempty"`
}

// RentalTransactionEvent represents a rental transaction event
type RentalTransactionEvent struct {
	AmazonOrderId         string                 `json:"amazonOrderId,omitempty"`
	RentalEventType       string                 `json:"rentalEventType,omitempty"`
	ExtensionLength       int                    `json:"extensionLength,omitempty"`
	PostedDate            string                 `json:"postedDate,omitempty"`
	RentalChargeList      []ChargeComponent      `json:"rentalChargeList,omitempty"`
	RentalFeeList         []FeeComponent         `json:"rentalFeeList,omitempty"`
	MarketplaceName       string                 `json:"marketplaceName,omitempty"`
	RentalInitialValue    *Currency              `json:"rentalInitialValue,omitempty"`
	RentalReimbursement   *Currency              `json:"rentalReimbursement,omitempty"`
	RentalTaxWithheldList []TaxWithheldComponent `json:"rentalTaxWithheldList,omitempty"`
}

// ProductAdsPaymentEvent represents a product ads payment event
type ProductAdsPaymentEvent struct {
	PostedDate       string    `json:"postedDate,omitempty"`
	TransactionType  string    `json:"transactionType,omitempty"`
	InvoiceId        string    `json:"invoiceId,omitempty"`
	BaseValue        *Currency `json:"baseValue,omitempty"`
	TaxValue         *Currency `json:"taxValue,omitempty"`
	TransactionValue *Currency `json:"transactionValue,omitempty"`
}

// ServiceFeeEvent represents a service fee event
type ServiceFeeEvent struct {
	AmazonOrderId  string         `json:"amazonOrderId,omitempty"`
	FeeReason      string         `json:"feeReason,omitempty"`
	FeeList        []FeeComponent `json:"feeList,omitempty"`
	SellerSKU      string         `json:"sellerSKU,omitempty"`
	FnSKU          string         `json:"fnSKU,omitempty"`
	FeeDescription string         `json:"feeDescription,omitempty"`
	ASIN           string         `json:"asin,omitempty"`
}

// SellerDealPaymentEvent represents a seller deal payment event
type SellerDealPaymentEvent struct {
	PostedDate      string    `json:"postedDate,omitempty"`
	DealId          string    `json:"dealId,omitempty"`
	DealDescription string    `json:"dealDescription,omitempty"`
	EventType       string    `json:"eventType,omitempty"`
	FeeType         string    `json:"feeType,omitempty"`
	FeeAmount       *Currency `json:"feeAmount,omitempty"`
	TaxAmount       *Currency `json:"taxAmount,omitempty"`
	TotalAmount     *Currency `json:"totalAmount,omitempty"`
}

// DebtRecoveryEvent represents a debt recovery event
type DebtRecoveryEvent struct {
	DebtRecoveryType     string             `json:"debtRecoveryType,omitempty"`
	RecoveryAmount       *Currency          `json:"recoveryAmount,omitempty"`
	OverPaymentCredit    *Currency          `json:"overPaymentCredit,omitempty"`
	DebtRecoveryItemList []DebtRecoveryItem `json:"debtRecoveryItemList,omitempty"`
	ChargeInstrumentList []ChargeInstrument `json:"chargeInstrumentList,omitempty"`
}

// DebtRecoveryItem represents a debt recovery item
type DebtRecoveryItem struct {
	RecoveryAmount *Currency `json:"recoveryAmount,omitempty"`
	OriginalAmount *Currency `json:"originalAmount,omitempty"`
	GroupBeginDate string    `json:"groupBeginDate,omitempty"`
	GroupEndDate   string    `json:"groupEndDate,omitempty"`
}

// ChargeInstrument represents a charge instrument
type ChargeInstrument struct {
	Description string    `json:"description,omitempty"`
	Tail        string    `json:"tail,omitempty"`
	Amount      *Currency `json:"amount,omitempty"`
}

// LoanServicingEvent represents a loan servicing event
type LoanServicingEvent struct {
	LoanAmount              *Currency `json:"loanAmount,omitempty"`
	SourceBusinessEventType string    `json:"sourceBusinessEventType,omitempty"`
}

// AdjustmentEvent represents an adjustment event
type AdjustmentEvent struct {
	AdjustmentType     string           `json:"adjustmentType,omitempty"`
	PostedDate         string           `json:"postedDate,omitempty"`
	AdjustmentAmount   *Currency        `json:"adjustmentAmount,omitempty"`
	AdjustmentItemList []AdjustmentItem `json:"adjustmentItemList,omitempty"`
}

// AdjustmentItem represents an adjustment item
type AdjustmentItem struct {
	Quantity           string    `json:"quantity,omitempty"`
	PerUnitAmount      *Currency `json:"perUnitAmount,omitempty"`
	TotalAmount        *Currency `json:"totalAmount,omitempty"`
	SellerSKU          string    `json:"sellerSKU,omitempty"`
	FnSKU              string    `json:"fnSKU,omitempty"`
	ProductDescription string    `json:"productDescription,omitempty"`
	ASIN               string    `json:"asin,omitempty"`
}

// SAFETReimbursementEvent represents a SAFET reimbursement event
type SAFETReimbursementEvent struct {
	PostedDate                 string                   `json:"postedDate,omitempty"`
	SAFETClaimId               string                   `json:"safetClaimId,omitempty"`
	ReimbursedAmount           *Currency                `json:"reimbursedAmount,omitempty"`
	SAFETReimbursementItemList []SAFETReimbursementItem `json:"safetReimbursementItemList,omitempty"`
}

// SAFETReimbursementItem represents a SAFET reimbursement item
type SAFETReimbursementItem struct {
	ItemChargeList     []ChargeComponent `json:"itemChargeList,omitempty"`
	ProductDescription string            `json:"productDescription,omitempty"`
	Quantity           string            `json:"quantity,omitempty"`
}

// SellerReviewEnrollmentPaymentEvent represents a seller review enrollment payment event
type SellerReviewEnrollmentPaymentEvent struct {
	PostedDate           string        `json:"postedDate,omitempty"`
	EnrollmentId         string        `json:"enrollmentId,omitempty"`
	ParentASIN           string        `json:"parentASIN,omitempty"`
	FeeComponent         *FeeComponent `json:"feeComponent,omitempty"`
	ChargeAmount         *Currency     `json:"chargeAmount,omitempty"`
	ChargeAmountRefunded *Currency     `json:"chargeAmountRefunded,omitempty"`
}

// FBALiquidationEvent represents an FBA liquidation event
type FBALiquidationEvent struct {
	PostedDate             string    `json:"postedDate,omitempty"`
	OriginalRemovalOrderId string    `json:"originalRemovalOrderId,omitempty"`
	LiquidationProceeds    *Currency `json:"liquidationProceeds,omitempty"`
	LiquidationFeeAmount   *Currency `json:"liquidationFeeAmount,omitempty"`
}

// CouponPaymentEvent represents a coupon payment event
type CouponPaymentEvent struct {
	PostedDate              string        `json:"postedDate,omitempty"`
	CouponId                string        `json:"couponId,omitempty"`
	SellerCouponDescription string        `json:"sellerCouponDescription,omitempty"`
	ClipOrRedemptionCount   int           `json:"clipOrRedemptionCount,omitempty"`
	PaymentEventId          string        `json:"paymentEventId,omitempty"`
	FeeComponent            *FeeComponent `json:"feeComponent,omitempty"`
	ChargeAmount            *Currency     `json:"chargeAmount,omitempty"`
	TotalAmount             *Currency     `json:"totalAmount,omitempty"`
}

// ImagingServicesFeeEvent represents an imaging services fee event
type ImagingServicesFeeEvent struct {
	ImagingRequestBillingItemID string         `json:"imagingRequestBillingItemID,omitempty"`
	ASIN                        string         `json:"asin,omitempty"`
	PostedDate                  string         `json:"postedDate,omitempty"`
	FeeList                     []FeeComponent `json:"feeList,omitempty"`
}

// NetworkComminglingTransactionEvent represents a network commingling transaction event
type NetworkComminglingTransactionEvent struct {
	PostedDate         string    `json:"postedDate,omitempty"`
	NetCoTransactionID string    `json:"netCoTransactionID,omitempty"`
	SwapReason         string    `json:"swapReason,omitempty"`
	ASIN               string    `json:"asin,omitempty"`
	MarketplaceId      string    `json:"marketplaceId,omitempty"`
	TaxExclusiveAmount *Currency `json:"taxExclusiveAmount,omitempty"`
	TaxAmount          *Currency `json:"taxAmount,omitempty"`
}

// AffordabilityExpenseEvent represents an affordability expense event
type AffordabilityExpenseEvent struct {
	AmazonOrderId   string    `json:"amazonOrderId,omitempty"`
	PostedDate      string    `json:"postedDate,omitempty"`
	MarketplaceId   string    `json:"marketplaceId,omitempty"`
	TransactionType string    `json:"transactionType,omitempty"`
	BaseExpense     *Currency `json:"baseExpense,omitempty"`
	TaxTypeCGST     *Currency `json:"taxTypeCGST,omitempty"`
	TaxTypeSGST     *Currency `json:"taxTypeSGST,omitempty"`
	TaxTypeIGST     *Currency `json:"taxTypeIGST,omitempty"`
	TotalExpense    *Currency `json:"totalExpense,omitempty"`
}

// RemovalShipmentEvent represents a removal shipment event
type RemovalShipmentEvent struct {
	PostedDate              string                `json:"postedDate,omitempty"`
	MerchantOrderId         string                `json:"merchantOrderId,omitempty"`
	OrderId                 string                `json:"orderId,omitempty"`
	TransactionType         string                `json:"transactionType,omitempty"`
	RemovalShipmentItemList []RemovalShipmentItem `json:"removalShipmentItemList,omitempty"`
}

// RemovalShipmentItem represents a removal shipment item
type RemovalShipmentItem struct {
	SellerSKU           string                 `json:"sellerSKU,omitempty"`
	FnSKU               string                 `json:"fnSKU,omitempty"`
	Quantity            int                    `json:"quantity,omitempty"`
	Revenue             *Currency              `json:"revenue,omitempty"`
	ItemFeeList         []FeeComponent         `json:"itemFeeList,omitempty"`
	ItemFeeAdjustments  []FeeComponent         `json:"itemFeeAdjustments,omitempty"`
	ItemTaxWithheldList []TaxWithheldComponent `json:"itemTaxWithheldList,omitempty"`
	ProductDescription  string                 `json:"productDescription,omitempty"`
	ASIN                string                 `json:"asin,omitempty"`
}

// RemovalShipmentAdjustmentEvent represents a removal shipment adjustment event
type RemovalShipmentAdjustmentEvent struct {
	PostedDate         string                          `json:"postedDate,omitempty"`
	AdjustmentEventId  string                          `json:"adjustmentEventId,omitempty"`
	MerchantOrderId    string                          `json:"merchantOrderId,omitempty"`
	OrderId            string                          `json:"orderId,omitempty"`
	TransactionType    string                          `json:"transactionType,omitempty"`
	AdjustmentItemList []RemovalShipmentItemAdjustment `json:"adjustmentItemList,omitempty"`
}

// RemovalShipmentItemAdjustment represents a removal shipment item adjustment
type RemovalShipmentItemAdjustment struct {
	SellerSKU                  string                 `json:"sellerSKU,omitempty"`
	FnSKU                      string                 `json:"fnSKU,omitempty"`
	Quantity                   int                    `json:"quantity,omitempty"`
	RevenueAdjustment          *Currency              `json:"revenueAdjustment,omitempty"`
	TaxAmountAdjustment        *Currency              `json:"taxAmountAdjustment,omitempty"`
	TaxWithheldAdjustment      *Currency              `json:"taxWithheldAdjustment,omitempty"`
	ItemFeeAdjustments         []FeeComponent         `json:"itemFeeAdjustments,omitempty"`
	ItemTaxWithheldAdjustments []TaxWithheldComponent `json:"itemTaxWithheldAdjustments,omitempty"`
	ProductDescription         string                 `json:"productDescription,omitempty"`
	ASIN                       string                 `json:"asin,omitempty"`
}

// TrialShipmentEvent represents a trial shipment event
type TrialShipmentEvent struct {
	AmazonOrderId         string         `json:"amazonOrderId,omitempty"`
	FinancialEventGroupId string         `json:"financialEventGroupId,omitempty"`
	PostedDate            string         `json:"postedDate,omitempty"`
	SKU                   string         `json:"sku,omitempty"`
	FeeList               []FeeComponent `json:"feeList,omitempty"`
}

// TDSReimbursementEvent represents a TDS reimbursement event
type TDSReimbursementEvent struct {
	PostedDate       string    `json:"postedDate,omitempty"`
	TdsOrderId       string    `json:"tdsOrderId,omitempty"`
	ReimbursedAmount *Currency `json:"reimbursedAmount,omitempty"`
}

// AdhocDisbursementEvent represents an adhoc disbursement event
type AdhocDisbursementEvent struct {
	PostedDate           string    `json:"postedDate,omitempty"`
	TransactionId        string    `json:"transactionId,omitempty"`
	DisbursementAmount   *Currency `json:"disbursementAmount,omitempty"`
	DisbursementCurrency string    `json:"disbursementCurrency,omitempty"`
}

// TaxWithholdingEvent represents a tax withholding event
type TaxWithholdingEvent struct {
	PostedDate           string                `json:"postedDate,omitempty"`
	BaseAmount           *Currency             `json:"baseAmount,omitempty"`
	WithheldAmount       *Currency             `json:"withheldAmount,omitempty"`
	TaxWithholdingPeriod *TaxWithholdingPeriod `json:"taxWithholdingPeriod,omitempty"`
}

// TaxWithholdingPeriod represents a tax withholding period
type TaxWithholdingPeriod struct {
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
}

// ChargeRefundEvent represents a charge refund event
type ChargeRefundEvent struct {
	PostedDate               string                    `json:"postedDate,omitempty"`
	ReasonCode               string                    `json:"reasonCode,omitempty"`
	ReasonCodeDescription    string                    `json:"reasonCodeDescription,omitempty"`
	ChargeRefundTransactions []ChargeRefundTransaction `json:"chargeRefundTransactions,omitempty"`
}

// ChargeRefundTransaction represents a charge refund transaction
type ChargeRefundTransaction struct {
	ChargeAmount *Currency `json:"chargeAmount,omitempty"`
	ChargeType   string    `json:"chargeType,omitempty"`
}

// FailedAdhocDisbursementEvent represents a failed adhoc disbursement event
type FailedAdhocDisbursementEvent struct {
	FundsTransfersType string    `json:"fundsTransfersType,omitempty"`
	TransferId         string    `json:"transferId,omitempty"`
	DisbursementId     string    `json:"disbursementId,omitempty"`
	ProcessingStatus   string    `json:"processingStatus,omitempty"`
	TransferAmount     *Currency `json:"transferAmount,omitempty"`
	PostedDate         string    `json:"postedDate,omitempty"`
	TransferFailReason string    `json:"transferFailReason,omitempty"`
	TransferFailCode   string    `json:"transferFailCode,omitempty"`
}

// ValueAddedServiceChargeEvent represents a value added service charge event
type ValueAddedServiceChargeEvent struct {
	TransactionType   string    `json:"transactionType,omitempty"`
	PostedDate        string    `json:"postedDate,omitempty"`
	Description       string    `json:"description,omitempty"`
	TransactionAmount *Currency `json:"transactionAmount,omitempty"`
}

// CapacityReservationBillingEvent represents a capacity reservation billing event
type CapacityReservationBillingEvent struct {
	TransactionType   string    `json:"transactionType,omitempty"`
	PostedDate        string    `json:"postedDate,omitempty"`
	Description       string    `json:"description,omitempty"`
	TransactionAmount *Currency `json:"transactionAmount,omitempty"`
}
