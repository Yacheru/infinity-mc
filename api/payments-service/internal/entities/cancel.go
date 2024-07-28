package entities

type CancellationDetails struct {
	Party  string `json:"party,omitempty"`
	Reason string `json:"reason,omitempty"`
}
