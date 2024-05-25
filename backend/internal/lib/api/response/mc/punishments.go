package mc

type LbPunishments struct {
	Victim   *Victim   `json:"victim,omitempty"`
	Reason   string    `json:"reason,omitempty"`
	Time     *Time     `json:"time,omitempty"`
	Operator *Operator `json:"operator,omitempty"`
}
