package entities

type Metadata struct {
	Nickname string `json:"nickname"`
	Price    string `json:"price,omitempty"`
	Service  string `json:"service"`
	Duration string `json:"duration"`
}
