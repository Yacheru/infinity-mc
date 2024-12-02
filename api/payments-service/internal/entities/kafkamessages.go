package entities

type KafkaMcMessage struct {
	Duration int64  `json:"duration"`
	Nickname string `json:"nickname"`
	Service  string `json:"service"`
}
