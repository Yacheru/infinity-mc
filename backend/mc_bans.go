package backend

import "github.com/google/uuid"

type LbPunishments struct {
	Victim struct {
		UUID uuid.UUID `json:"uuid" db:"uuid"`
		Name *string   `json:"name" db:"lower_name"`
	} `json:"victim" `
	Reason string `json:"reason" db:"reason"`
	Time   struct {
		Start int `json:"start" db:"start"`
		End   int `json:"end" db:"end"`
	} `json:"time"`
	Operator struct {
		UUID uuid.UUID `json:"uuid" db:"operator"`
		Name string    `json:"name" db:"lower_name"`
	} `json:"operator"`
}
