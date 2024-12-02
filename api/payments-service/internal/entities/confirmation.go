package entities

type ConfirmationType string

const (
	TypeEmbedded          ConfirmationType = "embedded"
	TypeExternal          ConfirmationType = "external"
	TypeMobileApplication ConfirmationType = "mobile_application"
	TypeQR                ConfirmationType = "qr"
	TypeRedirect          ConfirmationType = "redirect"
)

type Confirmer interface {
}

type Embedded struct {
	Type              ConfirmationType `json:"type,omitempty"`
	ConfirmationToken string           `json:"confirmation_token,omitempty"`
}

type External struct {
	Type ConfirmationType `json:"type,omitempty"`
}

type MobileApplication struct {
	Type            ConfirmationType `json:"type,omitempty"`
	ConfirmationURL string           `json:"confirmation_url,omitempty"`
}

type QR struct {
	Type             ConfirmationType `json:"type,omitempty"`
	ConfirmationData string           `json:"confirmation_data,omitempty"`
}

type Redirect struct {
	Type            ConfirmationType `json:"type,omitempty"`
	ConfirmationURL string           `json:"confirmation_url,omitempty"`
	ReturnURL       string           `json:"return_url,omitempty" binding:"max=2048"`
	Enforce         bool             `json:"enforce,omitempty"`
}
