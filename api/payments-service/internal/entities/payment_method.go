package entities

type PaymentMethodType string

const (
	PaymentTypeAlfabank      PaymentMethodType = "alfabank"
	PaymentTypeApplePay      PaymentMethodType = "apple_pay"
	PaymentTypeB2BSberbank   PaymentMethodType = "b2b_sberbank"
	PaymentTypeBankCard      PaymentMethodType = "bank_card"
	PaymentTypeCash          PaymentMethodType = "cash"
	PaymentTypeGooglePay     PaymentMethodType = "google_pay"
	PaymentTypeInstallments  PaymentMethodType = "installments"
	PaymentTypeMobileBalance PaymentMethodType = "mobile_balance"
	PaymentTypeQiwi          PaymentMethodType = "qiwi"
	PaymentTypeSberbank      PaymentMethodType = "sberbank"
	PaymentTypeSBP           PaymentMethodType = "sbp"
	PaymentTypeTinkoffBank   PaymentMethodType = "tinkoff_bank"
	PaymentTypeWebmoney      PaymentMethodType = "webmoney"
	PaymentTypeWeChat        PaymentMethodType = "wechat"
	PaymentTypeYooMoney      PaymentMethodType = "yoo_money"
)

type PaymentMethoder interface {
}

type paymentMethod struct {
	Type  PaymentMethodType `json:"type,omitempty"`
	ID    string            `json:"id,omitempty"`
	Saved bool              `json:"saved,omitempty"`
	Title string            `json:"title,omitempty"`
}

type Alfabank struct {
	paymentMethod
	Login string `login:"login,omitempty"`
}

type ApplePay struct {
	paymentMethod
}

type B2BSberbank struct {
	paymentMethod
	PayerBankDetails PayerBankDetails `json:"payer_bank_details,omitempty"`
	PaymentPurpose   string           `json:"payment_purpose,omitempty" binding:"max=210"`
	VATData          string           `json:"vat_data,omitempty"`
}

type BankCard struct {
	paymentMethod
	Card Card `json:"card,omitempty"`
}

type Cash struct {
	paymentMethod
}

type GooglePay struct {
	paymentMethod
}

type Installments struct {
	paymentMethod
}

type MobileBalance struct {
	paymentMethod
}

type Qiwi struct {
	paymentMethod
}

type Sberbank struct {
	paymentMethod
	Card  Card   `json:"card,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type SBP struct {
	paymentMethod
}

type TinkoffBank struct {
	paymentMethod
}

type WebMoney struct {
	paymentMethod
}

type WeChat struct {
	paymentMethod
}

type YooMoney struct {
	paymentMethod
	AccountNumber string `json:"account_number,omitempty"`
}
