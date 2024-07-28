package entities

type Settlement struct {
	Type   string `json:"type,omitempty"`
	Amount Amount `json:"amount,omitempty"`
}
