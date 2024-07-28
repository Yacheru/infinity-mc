package entities

import "time"

type Payment struct {
	ID                   string                `json:"id,omitempty"`
	Status               Status                `json:"status,omitempty"`
	Amount               *Amount               `json:"amount,omitempty"`
	Capture              bool                  `json:"capture,omitempty"`
	Description          string                `json:"description,omitempty" binding:"max=128"`
	Recipient            *Recipient            `json:"recipient,omitempty"`
	Receipt              *Receipt              `json:"receipt,omitempty"`
	PaymentMethod        PaymentMethoder       `json:"payment_method,omitempty"`
	CapturedAt           *time.Time            `json:"captured_at,omitempty"`
	CreatedAt            *time.Time            `json:"created_at,omitempty"`
	ExpiresAt            *time.Time            `json:"expires_at,omitempty"`
	Confirmation         Confirmer             `json:"confirmation,omitempty"`
	Test                 bool                  `json:"test,omitempty"`
	RefundedAmount       *Amount               `json:"refunded_amount,omitempty"`
	Paid                 bool                  `json:"paid,omitempty"`
	Refundable           bool                  `json:"refundable,omitempty"`
	ReceiptRegistration  Status                `json:"receipt_registration,omitempty"`
	Metadata             *Metadata             `json:"metadata,omitempty"`
	CancellationDetails  *CancellationDetails  `json:"cancellation_details,omitempty"`
	AuthorizationDetails *AuthorizationDetails `json:"authorization_details,omitempty"`
	Transfers            []Transfer            `json:"transfers,omitempty"`
	Deal                 *Deal                 `json:"deal,omitempty"`
	MerchantCustomerID   string                `json:"merchant_customer_id,omitempty" binding:"max=200"`
}
