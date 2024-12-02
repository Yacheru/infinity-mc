package entities

type PaymentService struct {
	Price    string `json:"price"`
	Email    string `json:"email"`
	Service  string `json:"service"`
	Nickname string `json:"nickname"`
	Duration int64  `json:"duration"`
}
