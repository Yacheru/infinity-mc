package mc

import "github.com/google/uuid"

type Victim struct {
	UUID uuid.UUID `json:"uuid" db:"uuid"`
	Name string    `json:"name" db:"name"`
}
