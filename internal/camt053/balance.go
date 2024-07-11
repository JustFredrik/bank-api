// package camt053 models the camt053 to enable marshaling and unmarshaling of camt053 data, both JSON and XML.
package camt053

// Balance represents the 'Bal' XML tag.
type Balance struct {
	Type                 BalanceType `xml:"Tp" json:"type"`
	CreditLine           *CreditLine `xml:"CdtLine" json:"creditLine,omitempty"`
	Amount               Amount      `xml:"Amt" json:"amount"`
	CreditDebitIndicator string      `xml:"CdtDbtInd" json:"creditDebitIndicator"`
	Date                 string      `xml:"Dt>Dt" json:"date"`
}

// BalanceType represents the 'Tp' XML tag.
type BalanceType struct {
	CodeOrProprietary CodeOrProprietary `xml:"CdOrPrtry" json:"codeOrProprietary"`
	SubType           *string           `xml:"SubTp>Cd" json:"subType,omitempty"`
}

// CreditLine represents the 'CdtLine' XML tag.
type CreditLine struct {
	Included bool   `xml:"Incl" json:"included"`
	Amount   Amount `xml:"Amt" json:"amount"`
}
