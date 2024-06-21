package records

type Time struct {
	Start int `json:"start" db:"start"`
	End   int `json:"end" db:"end"`
}
