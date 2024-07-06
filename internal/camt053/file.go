package camt053

import "encoding/xml"

// Tom Payne gave me some pointers on how to work with encoding/xml

type Document struct {
	XMLName       xml.Name                `xml:"Document"`
	BankStatement BankToCustomerStatement `xml:"BkToCstmrStmt" json:"bankStatement"`
}

type BankToCustomerStatement struct {
	//XMLName     xml.Name    `xml:"BkToCstmrStmt"`
	GroupHeader GroupHeader `xml:"GrpHdr"  json:"groupHeader"`
	Statement   Statement   `xml:"Stmt" json:"Statement"`
}

type GroupHeader struct {
	//XMLName         xml.Name `xml:"GrpHdr"`
	MessageId         int                `xml:"MsgId" json:"messageId"`
	CreationDateTime  string             `xml:"CreDtTm" json:"creationDateTime"`
	MessageRecipient  *MessageRecipient  `xml:"MsgRcpt" json:"messageRecipient,omitempty`
	MessagePagination *MessagePagination `xml:"MsgPgntn,omitempty" json:"messagePagination,omitempty"`
}

type MessageRecipient struct {
	Name string  `xml:"Nm" json:"name,omitempty"`
	Id   OtherId `xml:"Id>OrgId>Othr" json: "id"`
}

type MessagePagination struct {
}

type Statement struct {
	//XMLName                  xml.Name `xml:"Stmt"`
	Id                       string              `xml:"Id" json:"id"`
	ElectronicSequenceNumber *int                `xml:"ElctrncSeqNb" json:"electronicSequenceNumber"`
	LegalSequenceNumber      *int                `xml:"LglSeqNb" json:"legalSequenceNumber"`
	CreationDateTime         string              `xml:"CreDtTm" json:"createdDateTime"`
	FromDate                 *FromDate           `xml:"FrToDt" json:"fromDate"`
	Account                  Account             `xml:"Acct" json:"account"`
	Balances                 []Balance           `xml:"Bal" json:"balances"`
	Entries                  *[]Entry            `xml:"Ntry" json:"entries"`
	TransactionSummary       *TransactionSummary `xml:"TxsSummry" json:"transactionSummary"`
}

type TransactionSummary struct {
}

type Balance struct {
	Type                 BalanceType `xml:"Tp" json:"type"`
	CreditLine           *CreditLine `xml:"CdtLine" json:"creditLine,omitempty"`
	Amount               Amount      `xml:"Amt" json:"amount"`
	CreditDebitIndicator string      `xml:"CdtDbtInd" json:"creditDebitIndicator"`
	Date                 string      `xml:"Dt>Dt" json:"date"`
}

type CreditLine struct {
	Included bool   `xml:"Incl" json:"included"`
	Amount   Amount `xml:"Amt" json:"amount"`
}

type BalanceType struct {
	CodeOrProprietary CodeOrProprietary `xml:"CdOrPrtry" json:"codeOrProprietary"`
	SubType           *string           `xml:"SubTp>Cd" json:"subType,omitempty"`
}

type CodeOrProprietary struct {
	Code        *string `xml:"Cd" json:"code"`
	Proprietary *string `xml:"Prtry" json:"proprietary,omitempty"`
}

type Amount struct {
	Currency string `xml:"Ccy,attr"`
	Value    string `xml:",chardata"`
}

type FromDate struct {
	FromDateTime string `xml:"FrDtTm" json:"fromDateTime"`
	ToDateTime   string `xml:"ToDtTm" json:"toDateTime"`
}

type Account struct {
	//XMLName                  xml.Name `xml:"Acct"`
	Id       AccountId     `xml:"Id" json:"id"`
	Currency *string       `xml:"Ccy" json:"currency,omitempty"`
	Owner    *AccountOwner `xml:"Ownr" json:"owner,omitempty"`
	Servicer *Servicer     `xml:"Svcr" json:"servicer,omitempty"`
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

type OtherId struct {
	Id         int    `xml:"Id" json:"id"`
	SchemeName string `xml:"SchmeNm>Cd" json:"schemeName"`
}

type SchemeName struct {
	Code string `xml:"Cd" json:"code"`
}

type FinancialInstitutionId struct {
	BIC  string `xml:"BIC" json:"bicCode,omitempty"`
	Name string `xml:"Nm" json:"name,omitempty"`
}

type Entry struct {
	Reference            string `xml:"NtryRef"`
	Amount               string `xml:"Amt"`
	CreditDebitIndicator string `xml:"CdtDbtInd"`
}

type Servicer struct {
	FinancialInstitutionId FinancialInstitutionId `xml:"FinInstnId" json:"financialInstitutionId,omitempty"`
}
