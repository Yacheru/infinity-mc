package entities

type Card struct {
	First6        string `json:"first6,omitempty"`
	Last4         string `json:"last4,omitempty"`
	ExpiryYear    string `json:"expiry_year,omitempty"`
	ExpiryMonth   string `json:"expiry_month,omitempty"`
	CardType      string `json:"card_type,omitempty"`
	IssuerCountry string `json:"issuer_country,omitempty"`
	IssuerName    string `json:"issuer_name,omitempty"`
	Source        string `json:"source,omitempty"`
}
