package backend

type PaymentStruct struct {
	Amount  AmountType `json:"amount"`
	Receipt struct {
		Customer struct {
			Email string `json:"email"`
		} `json:"customer"`
		Items [1]struct {
			Description string     `json:"description"`
			Amount      AmountType `json:"amount"`
			VatCode     int        `json:"vat_code"`
			Quantity    string     `json:"quantity"`
		} `json:"items"`
	} `json:"receipt"`
	Capture      bool `json:"capture"`
	Confirmation struct {
		Type      string `json:"type"`
		ReturnUrl string `json:"return_url"`
	} `json:"confirmation"`
	Description string `json:"description"`
}

type AmountType struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}
