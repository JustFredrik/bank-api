package camt053

type Account struct {
	//XMLName                  xml.Name `xml:"Acct"`
	Id       AccountId     `xml:"Id" json:"id"`
	Currency *string       `xml:"Ccy" json:"currency,omitempty"`
	Owner    *AccountOwner `xml:"Ownr" json:"owner,omitempty"`
	Servicer *Servicer     `xml:"Svcr" json:"servicer,omitempty"`
}

func (acc Account) GetId() uint64 {
	return acc.Id.Other.Id
}

type AccountOwner struct {
	Name     string    `xml:"Nm" json:"name"`
	Id       AccountId `xml:"Id>OrgId>Othr"`
	Servicer Servicer  `xml:"Svcr" json:"servicer,omitempty"`
}

type AccountId struct {
	IBAN  *string  `xml:"IBAN" json:"IBAN,omitempty"`
	Other *OtherId `xml:"Othr" json:"other,omitempty"`
}
