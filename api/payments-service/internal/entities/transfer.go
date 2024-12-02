package entities

type Transfer struct {
	AccountID         string      `json:"account_id,omitempty"`
	Amount            Amount      `json:"amount,omitempty"`
	Status            Status      `json:"status,omitempty"`
	PlatformFeeAmount Amount      `json:"platform_fee_amount,omitempty"`
	Description       string      `json:"description,omitempty" binding:"max=128"`
	Metadata          interface{} `json:"metadata,omitempty"`
}
