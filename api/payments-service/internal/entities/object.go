package entities

import "time"

type Object struct {
	ID                   string               `json:"id,omitempty"`
	Status               string               `json:"status,omitempty"`
	Paid                 bool                 `json:"paid,omitempty"`
	Amount               *Amount              `json:"amount,omitempty"`
	AuthorizationDetails AuthorizationDetails `json:"authorization_details,omitempty"`
	CreatedAt            *time.Time           `json:"created_at,omitempty"`
	CapturedAt           *time.Time           `json:"captured_at,omitempty"`
	Description          string               `json:"description,omitempty"`
	Metadata             *Metadata            `json:"metadata,omitempty"`
	PaymentMethod        PaymentMethoder      `json:"payment_method,omitempty"`
	IncomeAmount         *Amount              `json:"income_amount,omitempty"`
	ReceiptRegistration  Status               `json:"receipt_registration,omitempty"`
	Recipient            *Recipient           `json:"recipient,omitempty"`
	Refundable           bool                 `json:"refundable,omitempty"`
	RefundedAmount       *Amount              `json:"refunded_amount,omitempty"`
	Test                 bool                 `json:"test,omitempty"`
}
