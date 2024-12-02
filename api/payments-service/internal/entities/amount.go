package entities

type Amount struct {
	Value    string `json:"value,omitempty"`
	Currency string `json:"currency,omitempty"`
}
