package entities

type Settlement struct {
	Amount Amount `json:"amount,omitempty"`
	Type   string `json:"type,omitempty"`
}
