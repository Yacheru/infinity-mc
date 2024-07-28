package entities

type Recipient struct {
	AccountId string `json:"account_id,omitempty"`
	GatewayId string `json:"gateway_id,omitempty"`
}
