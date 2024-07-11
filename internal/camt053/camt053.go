package camt053

import "encoding/xml"

// Tom Payne had a good talk that gave me some pointers on how to work with encoding/xml

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
	ElectronicSequenceNumber *int                `xml:"ElctrncSeqNb" json:"electronicSequenceNumber,omitempty"`
	LegalSequenceNumber      *int                `xml:"LglSeqNb" json:"legalSequenceNumber,omitempty"`
	CreationDateTime         string              `xml:"CreDtTm" json:"createdDateTime"`
	FromDate                 *FromDate           `xml:"FrToDt" json:"fromDate,omitempty"`
	Account                  Account             `xml:"Acct" json:"account"`
	Balances                 []Balance           `xml:"Bal" json:"balances"`
	Entries                  *[]Entry            `xml:"Ntry" json:"entries,omitempty"`
	TransactionSummary       *TransactionSummary `xml:"TxsSummry" json:"transactionSummary,omitempty"`
}

type TransactionSummary struct {
	TotalCreditEntries *CreditDebitEntry `xml:"TtlCdtNtries" json:"totalCreditEntries,omitempty"`
	TotalDebitEntries  *CreditDebitEntry `xml:"TtlDbtNtries" json:"totalDebitEntries,omitempty"`
}

type CreditDebitEntry struct {
	NumberOfEntries int    `xml:"NbOfNtries" json:"numberOfEntries"`
	Sum             string `xml:"Sum" json:"sum"`
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

type OtherId struct {
	Id         uint64 `xml:"Id" json:"id"`
	SchemeName string `xml:"SchmeNm>Cd" json:"schemeName"`
}

type FinancialInstitutionId struct {
	BIC  *string `xml:"BIC" json:"bicCode,omitempty"`
	Name *string `xml:"Nm" json:"name,omitempty"`
}

type Servicer struct {
	FinancialInstitutionId FinancialInstitutionId `xml:"FinInstnId" json:"financialInstitutionId,omitempty"`
}
