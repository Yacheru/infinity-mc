package entities

type AuthorizationDetails struct {
	RRN          string `json:"rrn,omitempty"`
	AuthCode     string `json:"auth_code,omitempty"`
	ThreeDSecure struct {
		Applied bool `json:"applied,omitempty"`
	} `json:"three_d_secure,omitempty"`
}
