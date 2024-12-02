package entities

import "github.com/google/uuid"

type Operator struct {
	UUID uuid.UUID `json:"uuid,omitempty" db:"uuid"`
	Name string    `json:"name,omitempty" db:"name"`
}
