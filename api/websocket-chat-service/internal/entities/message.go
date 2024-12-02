package entities

type Message struct {
	Message string  `json:"message"`
	Player  Player  `json:"player"`
	SentAt  *string `json:"sent_at,omitempty"`
}

type Player struct {
	UUID     string `json:"uuid"`
	Username string `json:"Username"`
}
