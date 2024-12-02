package entities

type Message struct {
	Nickname string `json:"nickname" binding:"required" validate:"required"`
	Duration string `json:"duration" binding:"required" validate:"required"`
	Service  string `json:"service" binding:"required" validate:"required"`
}
