package mc

type LbPunishments struct {
	ID       int      `json:"id" db:"-"`
	Victim   Victim   `json:"victim" db:"victim"`
	Reason   string   `json:"reason" db:"reason"`
	Time     Time     `json:"time" db:"time"`
	Operator Operator `json:"operator" db:"operator"`
}
