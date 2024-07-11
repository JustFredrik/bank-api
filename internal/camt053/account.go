// package camt053 models the camt053 to enable marshaling and unmarshaling of camt053 data, both JSON and XML.
package camt053

// Account represents the 'Acct' XML tag.
type Account struct {
	//XMLName                  xml.Name `xml:"Acct"`
	Id       AccountId     `xml:"Id" json:"id"`
	Currency *string       `xml:"Ccy" json:"currency,omitempty"`
	Owner    *AccountOwner `xml:"Ownr" json:"owner,omitempty"`
	Servicer *Servicer     `xml:"Svcr" json:"servicer,omitempty"`
}

// GetId returns the accounts id.
func (acc Account) GetId() uint64 {
	return acc.Id.Other.Id
}

// AccountOwner represents the 'Ownr' XML tag.
type AccountOwner struct {
	Name     string    `xml:"Nm" json:"name"`
	Id       AccountId `xml:"Id>OrgId>Othr"`
	Servicer Servicer  `xml:"Svcr" json:"servicer,omitempty"`
}

// Account represents the 'Id>Othr' XML tag.
type AccountId struct {
	IBAN  *string  `xml:"IBAN" json:"IBAN,omitempty"`
	Other *OtherId `xml:"Othr" json:"other,omitempty"`
}
