package mc

import "github.com/google/uuid"

type Operator struct {
	UUID uuid.UUID `json:"uuid,omitempty"`
	Name string    `json:"name,omitempty"`
}
