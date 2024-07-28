package entities

type Status string

const (
	Pending           Status = "pending"
	WaitingForCapture Status = "waiting_for_capture"
	Succeeded         Status = "succeeded"
	Canceled          Status = "canceled"
)
