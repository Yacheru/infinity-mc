package records

type Metadata struct {
	Nickname  string `json:"nickname,omitempty"`
	Price     string `json:"amount,omitempty"`
	DonatType string `json:"donat,omitempty"`
	Duration  string `json:"duration,omitempty"`
}
