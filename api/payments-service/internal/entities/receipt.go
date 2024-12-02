package entities

type Receipt struct {
	Customer *Email   `json:"customer,omitempty"`
	Items    *[]Items `json:"items,omitempty"`
}

type Items struct {
	Description string  `json:"description,omitempty"`
	Amount      *Amount `json:"amount,omitempty"`
	VatCode     int     `json:"vat_code,omitempty"`
	Quantity    string  `json:"quantity,omitempty"`
}

type Email struct {
	Email string `json:"email,omitempty"`
}
