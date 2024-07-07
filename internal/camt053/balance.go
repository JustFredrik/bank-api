package camt053

type Balance struct {
	Type                 BalanceType `xml:"Tp" json:"type"`
	CreditLine           *CreditLine `xml:"CdtLine" json:"creditLine,omitempty"`
	Amount               Amount      `xml:"Amt" json:"amount"`
	CreditDebitIndicator string      `xml:"CdtDbtInd" json:"creditDebitIndicator"`
	Date                 string      `xml:"Dt>Dt" json:"date"`
}

type BalanceType struct {
	CodeOrProprietary CodeOrProprietary `xml:"CdOrPrtry" json:"codeOrProprietary"`
	SubType           *string           `xml:"SubTp>Cd" json:"subType,omitempty"`
}

type CreditLine struct {
	Included bool   `xml:"Incl" json:"included"`
	Amount   Amount `xml:"Amt" json:"amount"`
}
