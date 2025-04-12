package models

import "time"

type Purchase struct {
	OrderID             string             `json:"order_id"`
	OrderRef            string             `json:"order_ref"`
	OrderStatus         string             `json:"order_status"`
	ProductType         string             `json:"product_type"`
	PaymentMethod       string             `json:"payment_method"`
	StoreID             string             `json:"store_id"`
	PaymentMerchantID   int                `json:"payment_merchant_id"`
	Installments        int                `json:"installments"`
	CardType            string             `json:"card_type"`
	CardLast4Digits     string             `json:"card_last4digits"`
	CardRejectionReason *string            `json:"card_rejection_reason"`
	BoletoURL           *string            `json:"boleto_URL"`
	BoletoBarcode       *string            `json:"boleto_barcode"`
	BoletoExpiryDate    *string            `json:"boleto_expiry_date"`
	PixCode             *string            `json:"pix_code"`
	PixExpiration       *string            `json:"pix_expiration"`
	SaleType            string             `json:"sale_type"`
	RefundedAt          *time.Time         `json:"refunded_at"`
	WebhookEventType    string             `json:"webhook_event_type"`
	Product             Product            `json:"Product"`
	Customer            Customer           `json:"Customer"`
	Commissions         Commissions        `json:"Commissions"`
	TrackingParameters  TrackingParameters `json:"TrackingParameters"`
	Subscription        Subscription       `json:"Subscription"`
	SubscriptionID      string             `json:"subscription_id"`
	AccessURL           *string            `json:"access_url"`
}

type Product struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
}

type Customer struct {
	FullName     string `json:"full_name"`
	FirstName    string `json:"first_name"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	CNPJ         string `json:"cnpj"`
	IP           string `json:"ip"`
	Instagram    string `json:"instagram"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zipcode      string `json:"zipcode"`
}

type Commissions struct {
	ChargeAmount             int                 `json:"charge_amount"`
	ProductBasePrice         int                 `json:"product_base_price"`
	ProductBasePriceCurrency string              `json:"product_base_price_currency"`
	KiwifyFee                int                 `json:"kiwify_fee"`
	KiwifyFeeCurrency        string              `json:"kiwify_fee_currency"`
	SettlementAmount         int                 `json:"settlement_amount"`
	SettlementAmountCurrency string              `json:"settlement_amount_currency"`
	SaleTaxRate              int                 `json:"sale_tax_rate"`
	SaleTaxAmount            int                 `json:"sale_tax_amount"`
	CommissionedStores       []CommissionedStore `json:"commissioned_stores"`
	Currency                 string              `json:"currency"`
	MyCommission             int                 `json:"my_commission"`
	FundsStatus              *string             `json:"funds_status"`
	EstimatedDepositDate     *string             `json:"estimated_deposit_date"`
	DepositDate              *string             `json:"deposit_date"`
}

type CommissionedStore struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	CustomName  string  `json:"custom_name"`
	Email       string  `json:"email"`
	Value       string  `json:"value"`
	AffiliateID *string `json:"affiliate_id,omitempty"`
}

type TrackingParameters struct {
	Src         *string `json:"src"`
	Sck         *string `json:"sck"`
	UtmSource   *string `json:"utm_source"`
	UtmMedium   *string `json:"utm_medium"`
	UtmCampaign *string `json:"utm_campaign"`
	UtmContent  *string `json:"utm_content"`
	UtmTerm     *string `json:"utm_term"`
	S1          *string `json:"s1"`
	S2          *string `json:"s2"`
	S3          *string `json:"s3"`
}

type Subscription struct {
	ID          string  `json:"id"`
	StartDate   string  `json:"start_date"`
	NextPayment string  `json:"next_payment"`
	Status      string  `json:"status"`
	Plan        Plan    `json:"plan"`
	Charges     Charges `json:"charges"`
}

type Plan struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Frequency  string `json:"frequency"`
	QtyCharges int    `json:"qty_charges"`
}

type Charges struct {
	Completed []CompletedCharge `json:"completed"`
	Future    []FutureCharge    `json:"future"`
}

type CompletedCharge struct {
	OrderID         string `json:"order_id"`
	Amount          int    `json:"amount"`
	Status          string `json:"status"`
	Installments    int    `json:"installments"`
	CardType        string `json:"card_type"`
	CardLastDigits  string `json:"card_last_digits"`
	CardFirstDigits string `json:"card_first_digits"`
	CreatedAt       string `json:"created_at"`
}

type FutureCharge struct {
	ChargeDate string `json:"charge_date"`
}
