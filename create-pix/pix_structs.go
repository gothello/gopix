package pix

import (
	"time"

	"github.com/google/uuid"
)

type InputPix struct {
	ID               string
	Amount           float64
	Description      string
	TimeOfExpiration time.Duration
	UrlNotify        string
	Email            string
}

type OutputPix struct {
	ID                    string
	IDExternalTransaction int64
	CreateAt              string
	ExpiresAt             string
	Status                string
	Type                  string
	Amount                float64
	Ticket                string
	QrCode                string
	QrCodeBase64          string
}

func NewPix(amount float64, desc string, time time.Duration, url, email string) *InputPix {
	return &InputPix{
		ID:               uuid.New().String(),
		Amount:           amount,
		Description:      desc,
		TimeOfExpiration: time,
		UrlNotify:        url,
		Email:            email,
	}
}

type RefundData struct {
	ID        int   `json:"id"`
	PaymentID int64 `json:"payment_id"`
	Amount    int   `json:"amount"`
	Metadata  struct {
	} `json:"metadata"`
	Source struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"source"`
	DateCreated           string        `json:"date_created"`
	UniqueSequenceNumber  interface{}   `json:"unique_sequence_number"`
	RefundMode            string        `json:"refund_mode"`
	AdjustmentAmount      int           `json:"adjustment_amount"`
	Status                string        `json:"status"`
	Reason                interface{}   `json:"reason"`
	Labels                []interface{} `json:"labels"`
	AmountRefundedToPayer int           `json:"amount_refunded_to_payer"`
	AdditionalData        interface{}   `json:"additional_data"`
	PartitionDetails      []interface{} `json:"partition_details"`
}

type ResponseMP struct {
	ID                 int64       `json:"id"`
	DateCreated        string      `json:"date_created"`
	DateApproved       interface{} `json:"date_approved"`
	DateLastUpdated    string      `json:"date_last_updated"`
	DateOfExpiration   string      `json:"date_of_expiration"`
	MoneyReleaseDate   interface{} `json:"money_release_date"`
	MoneyReleaseStatus interface{} `json:"money_release_status"`
	OperationType      string      `json:"operation_type"`
	IssuerID           interface{} `json:"issuer_id"`
	PaymentMethodID    string      `json:"payment_method_id"`
	PaymentTypeID      string      `json:"payment_type_id"`
	PaymentMethod      struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"payment_method"`
	Status             string      `json:"status"`
	StatusDetail       string      `json:"status_detail"`
	CurrencyID         string      `json:"currency_id"`
	Description        string      `json:"description"`
	LiveMode           bool        `json:"live_mode"`
	SponsorID          interface{} `json:"sponsor_id"`
	AuthorizationCode  interface{} `json:"authorization_code"`
	MoneyReleaseSchema interface{} `json:"money_release_schema"`
	TaxesAmount        int         `json:"taxes_amount"`
	CounterCurrency    interface{} `json:"counter_currency"`
	BrandID            interface{} `json:"brand_id"`
	ShippingAmount     int         `json:"shipping_amount"`
	BuildVersion       string      `json:"build_version"`
	PosID              interface{} `json:"pos_id"`
	StoreID            interface{} `json:"store_id"`
	IntegratorID       interface{} `json:"integrator_id"`
	PlatformID         interface{} `json:"platform_id"`
	CorporationID      interface{} `json:"corporation_id"`
	CollectorID        int         `json:"collector_id"`
	Payer              struct {
		Type           interface{} `json:"type"`
		ID             string      `json:"id"`
		OperatorID     interface{} `json:"operator_id"`
		Email          interface{} `json:"email"`
		Identification struct {
			Type   interface{} `json:"type"`
			Number interface{} `json:"number"`
		} `json:"identification"`
		Phone struct {
			AreaCode  interface{} `json:"area_code"`
			Number    interface{} `json:"number"`
			Extension interface{} `json:"extension"`
		} `json:"phone"`
		FirstName  interface{} `json:"first_name"`
		LastName   interface{} `json:"last_name"`
		EntityType interface{} `json:"entity_type"`
	} `json:"payer"`
	MarketplaceOwner interface{} `json:"marketplace_owner"`
	Metadata         struct {
	} `json:"metadata"`
	AdditionalInfo struct {
		AvailableBalance   interface{} `json:"available_balance"`
		NsuProcessadora    interface{} `json:"nsu_processadora"`
		AuthenticationCode interface{} `json:"authentication_code"`
	} `json:"additional_info"`
	Order struct {
	} `json:"order"`
	ExternalReference         interface{} `json:"external_reference"`
	TransactionAmount         int         `json:"transaction_amount"`
	TransactionAmountRefunded int         `json:"transaction_amount_refunded"`
	CouponAmount              int         `json:"coupon_amount"`
	DifferentialPricingID     interface{} `json:"differential_pricing_id"`
	FinancingGroup            interface{} `json:"financing_group"`
	DeductionSchema           interface{} `json:"deduction_schema"`
	CallbackURL               interface{} `json:"callback_url"`
	Installments              int         `json:"installments"`
	TransactionDetails        struct {
		PaymentMethodReferenceID interface{} `json:"payment_method_reference_id"`
		NetReceivedAmount        int         `json:"net_received_amount"`
		TotalPaidAmount          int         `json:"total_paid_amount"`
		OverpaidAmount           int         `json:"overpaid_amount"`
		ExternalResourceURL      interface{} `json:"external_resource_url"`
		InstallmentAmount        int         `json:"installment_amount"`
		FinancialInstitution     interface{} `json:"financial_institution"`
		PayableDeferralPeriod    interface{} `json:"payable_deferral_period"`
		AcquirerReference        interface{} `json:"acquirer_reference"`
		BankTransferID           interface{} `json:"bank_transfer_id"`
		TransactionID            interface{} `json:"transaction_id"`
	} `json:"transaction_details"`
	FeeDetails     []interface{} `json:"fee_details"`
	ChargesDetails []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		Accounts struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"accounts"`
		ClientID    int    `json:"client_id"`
		DateCreated string `json:"date_created"`
		LastUpdated string `json:"last_updated"`
		Amounts     struct {
			Original float64 `json:"original"`
			Refunded int     `json:"refunded"`
		} `json:"amounts"`
		Metadata struct {
		} `json:"metadata"`
		ReserveID     interface{}   `json:"reserve_id"`
		RefundCharges []interface{} `json:"refund_charges"`
	} `json:"charges_details"`
	Captured            bool        `json:"captured"`
	BinaryMode          bool        `json:"binary_mode"`
	CallForAuthorizeID  interface{} `json:"call_for_authorize_id"`
	StatementDescriptor interface{} `json:"statement_descriptor"`
	Card                struct {
	} `json:"card"`
	NotificationURL        string        `json:"notification_url"`
	Refunds                []interface{} `json:"refunds"`
	ProcessingMode         string        `json:"processing_mode"`
	MerchantAccountID      interface{}   `json:"merchant_account_id"`
	MerchantNumber         interface{}   `json:"merchant_number"`
	AcquirerReconciliation []interface{} `json:"acquirer_reconciliation"`
	PointOfInteraction     struct {
		Type         string `json:"type"`
		BusinessInfo struct {
			Unit    string `json:"unit"`
			SubUnit string `json:"sub_unit"`
		} `json:"business_info"`
		Location struct {
			StateID interface{} `json:"state_id"`
			Source  interface{} `json:"source"`
		} `json:"location"`
		ApplicationData struct {
			Name    interface{} `json:"name"`
			Version interface{} `json:"version"`
		} `json:"application_data"`
		TransactionData struct {
			QrCode               string      `json:"qr_code"`
			BankTransferID       interface{} `json:"bank_transfer_id"`
			TransactionID        interface{} `json:"transaction_id"`
			E2EID                interface{} `json:"e2e_id"`
			FinancialInstitution interface{} `json:"financial_institution"`
			TicketURL            string      `json:"ticket_url"`
			BankInfo             struct {
				Payer struct {
					AccountID         interface{} `json:"account_id"`
					ID                interface{} `json:"id"`
					LongName          interface{} `json:"long_name"`
					ExternalAccountID interface{} `json:"external_account_id"`
				} `json:"payer"`
				Collector struct {
					AccountID         interface{} `json:"account_id"`
					LongName          interface{} `json:"long_name"`
					AccountHolderName string      `json:"account_holder_name"`
					TransferAccountID interface{} `json:"transfer_account_id"`
				} `json:"collector"`
				IsSameBankAccountOwner interface{} `json:"is_same_bank_account_owner"`
				OriginBankID           interface{} `json:"origin_bank_id"`
				OriginWalletID         interface{} `json:"origin_wallet_id"`
			} `json:"bank_info"`
			QrCodeBase64 string `json:"qr_code_base64"`
		} `json:"transaction_data"`
	} `json:"point_of_interaction"`
}
