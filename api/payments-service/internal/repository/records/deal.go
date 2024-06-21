package records

type Deal struct {
	ID          string       `json:"id,omitempty" binding:"min=36,max=50"`
	Settlements []Settlement `json:"settlements,omitempty"`
}
