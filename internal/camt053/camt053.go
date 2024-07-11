// package camt053 models the camt053 to enable marshaling and unmarshaling of camt053 data, both JSON and XML.
package camt053

import "encoding/xml"

// Tom Payne had a good talk that gave me some pointers on how to work with encoding/xml

// Document represnts the root 'Document' tag of the camt053 XML document.
type Document struct {
	XMLName       xml.Name                `xml:"Document"`
	BankStatement BankToCustomerStatement `xml:"BkToCstmrStmt" json:"bankStatement"`
}

// BankToCustomerStatement represents the 'BkToCstmrStmt' XML tag.
type BankToCustomerStatement struct {
	//XMLName     xml.Name    `xml:"BkToCstmrStmt"`
	GroupHeader GroupHeader `xml:"GrpHdr"  json:"groupHeader"`
	Statement   Statement   `xml:"Stmt" json:"Statement"`
}

// GroupHeader represents the 'GrpHdr' XML tag.
type GroupHeader struct {
	//XMLName         xml.Name `xml:"GrpHdr"`
	MessageId         int                `xml:"MsgId" json:"messageId"`
	CreationDateTime  string             `xml:"CreDtTm" json:"creationDateTime"`
	MessageRecipient  *MessageRecipient  `xml:"MsgRcpt" json:"messageRecipient,omitempty`
	MessagePagination *MessagePagination `xml:"MsgPgntn,omitempty" json:"messagePagination,omitempty"`
}

// MessageRecipient represents the 'MsgRcpt' XML tag.
type MessageRecipient struct {
	Name string  `xml:"Nm" json:"name,omitempty"`
	Id   OtherId `xml:"Id>OrgId>Othr" json: "id"`
}

// MessagePagination represents the 'MsgPgntn' XML tag. (not implemented)
type MessagePagination struct {
}

// Statement represents the 'Stmt' XML tag.
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

// TransactionSummary represents the 'TxsSummry' XML tag.
type TransactionSummary struct {
	TotalCreditEntries *CreditDebitEntry `xml:"TtlCdtNtries" json:"totalCreditEntries,omitempty"`
	TotalDebitEntries  *CreditDebitEntry `xml:"TtlDbtNtries" json:"totalDebitEntries,omitempty"`
}

// CreditDebitEntry represents the 'TtlCdtNtries' and the 'TtlDbtNtries' XML tags.
type CreditDebitEntry struct {
	NumberOfEntries int    `xml:"NbOfNtries" json:"numberOfEntries"`
	Sum             string `xml:"Sum" json:"sum"`
}

// CodePorProprietary represents the 'CdOrPrtry' XML tag.
type CodeOrProprietary struct {
	Code        *string `xml:"Cd" json:"code"`
	Proprietary *string `xml:"Prtry" json:"proprietary,omitempty"`
}

// Amount represents the 'Amt' XML tag.
type Amount struct {
	Currency string `xml:"Ccy,attr" json:"currency"`
	Value    string `xml:",chardata" json:"value"`
}

// FromDate represents the 'FrToDt' XML tag.
type FromDate struct {
	FromDateTime string `xml:"FrDtTm" json:"fromDateTime"`
	ToDateTime   string `xml:"ToDtTm" json:"toDateTime"`
}

// OtherId represents the 'Othr' XML tag which can be found nested inside Id tags.
type OtherId struct {
	Id         uint64 `xml:"Id" json:"id"`
	SchemeName string `xml:"SchmeNm>Cd" json:"schemeName"`
}

// FinancialInstitutionId represents the 'FinInstId' XML tag.
type FinancialInstitutionId struct {
	BIC  *string `xml:"BIC" json:"bicCode,omitempty"`
	Name *string `xml:"Nm" json:"name,omitempty"`
}

//Servicer represents the 'Svcr' XML tag.
type Servicer struct {
	FinancialInstitutionId FinancialInstitutionId `xml:"FinInstnId" json:"financialInstitutionId,omitempty"`
}
