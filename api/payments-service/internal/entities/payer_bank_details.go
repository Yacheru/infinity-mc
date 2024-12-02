package entities

type PayerBankDetails struct {
	FullName   string `json:"full_name,omitempty" binding:"max=800"`
	ShortName  string `json:"short_name,omitempty" binding:"max=160"`
	Address    string `json:"address,omitempty" binding:"max=500"`
	INN        string `json:"inn,omitempty"`
	BankName   string `json:"bank_name,omitempty" binding:"min=1,max=350"`
	BankBranch string `json:"bank_branch,omitempty" binding:"min=1,max=140"`
	BankBIK    string `json:"bank_bik,omitempty"`
	Account    string `json:"account,omitempty"`
	KPP        string `json:"kpp,omitempty"`
}
