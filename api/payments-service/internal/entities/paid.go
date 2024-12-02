package entities

type Paid struct {
	Type   string  `json:"type,omitempty"`
	Event  string  `json:"event,omitempty"`
	Object *Object `json:"object,omitempty"`
}
