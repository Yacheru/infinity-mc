package entities

type Deal struct {
	Settlements []Settlement `json:"settlements,omitempty"`
	ID          string       `json:"id,omitempty" binding:"min=36,max=50"`
}
