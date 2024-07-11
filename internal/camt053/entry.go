// package camt053 models the camt053 to enable marshaling and unmarshaling of camt053 data, both JSON and XML.
package camt053

// Entry represents the 'Ntry' XML tag.
type Entry struct {
	Reference            *string `xml:"NtryRef" json:"reference"`
	URLReference         *string `xml:"-" json:"urlReference"` // Not part of camt053, used as resource ref in API.
	Amount               Amount  `xml:"Amt" json:"amount"`
	CreditDebitIndicator string  `xml:"CdtDbtInd" json:"creditDebitIndicator"`
	//RvslInd optional
	Status              string              `xml:"Sts" json:"status"`
	BookingDate         *string             `xml:"BookgDt>Dt" json:"bookingDate,omitempty"`
	ValueDate           *string             `xml:"ValDt>Dt" json:"valueDate,omitempty"`
	AccountServicerRef  *string             `xml:"AcctSvcrRef" json:"accountServicerRef,omitempty"`
	BankTransactionCode BankTransactionCode `xml:"BkTxCd" json:"bankTransactionCode,omitempty"`
	AmountDetails       *AmountDetails      `xml:"AmtDtls" json:"amountDetails,omitempty"`
	Charges             *[]Charge           `xml:"Chrgs" json:"charges,omitempty"`
	//TechInptChanl, optional
	EntryDetails *[]EntryDetail `xml:"NtryDtls" json:"entryDetails,omitempty"`
}

// BankTransactionCode represents the 'BkTxCd' XML tag.
type BankTransactionCode struct {
	Domain          *BankTransactionCodeDomain      `xml:"Domn" json:"domain,omitempty"`
	ProprietaryCode *BankTransactionProprietaryCode `xml:"Prtry" json:"proprietaryCode"`
}

// BankTransactionCodeDomain represents the 'Domn' XML tag.
type BankTransactionCodeDomain struct {
	Code   string                    `xml:"Cd" json:"code"`
	Family BankTransactionCodeFamily `xml:"Fmly" json:"family"`
}

// BankTransactionCodeFamily represents the 'Fmly' XML tag.
type BankTransactionCodeFamily struct {
	Code          string `xml:"Cd" json:"code"`
	SubFamilyCode string `xml:"SubFmlyCd" json:"subFamilyCode"`
}

// BanTransactionProprietaryCode represents the 'Prtry' XML tag inside bank transaction code sections.
type BankTransactionProprietaryCode struct {
	Code   string `xml:"Cd" json:"code"`
	Issuer string `xml:"Issr" json:"issuer"`
}

// AmountDetails represents the 'AmtDtls' XML tag.
type AmountDetails struct {
	InstructedAmount  *AmountAndCurrencyExchangeDetails `xml:"InstdAmt" json:"instructedAmount,omitempty"`
	TransactionAmount *AmountAndCurrencyExchangeDetails `xml:"TxAmt" json:"transactionAmount,omitempty"`
	//CntrValAmt, optional
	//AnncdPstngAmt, optional
	ProprietaryAmount *ProprietaryAmount `xml:"PrtryAmt" json:"proprietaryAmount,omitempty"`
}

// AmountAndCurrencyExchangeDetails represents the 'InstdAmt' and 'TxAmt' XML tags.
type AmountAndCurrencyExchangeDetails struct {
	Amount Amount `xml:"Amt" json:"amount"`
	//CcyXchg (Currency Exchange), optional
}

// ProprietaryAmount represents the 'PrtryAmt' XML tag.
type ProprietaryAmount struct {
	Type   string `xml:"Tp" json:"type"`
	Amount Amount `xml:"Amt" json:"amount"`
}

// Charge represents the 'Chrgs' XML tag. (not part of test data set)
type Charge struct { // Charges are specified in 2.172 and 2.152 in SEB MIG camt052-053-054v2 spec
	TotalChargesAndTaxAmount *string `xml:"TtlChrgsAndTaxAmt" json:"totalChargesAndTaxAmount,omitempty"`
	Amount                   Amount  `xml:"Amt" json:"amount"`
	Type                     *string `xml:"Tp>Cd" json:"type"`
	Rate                     *string `xml:"Rate" json:"rate"`
	// Br, optional
	// Pty, optional
	Tax *string `xml:"Tax" json:"tax,omitempty"`
}

// EntryDetail represents the 'NtryDtls' XML tag.
type EntryDetail struct {
	Batch              *BatchInformation    `xml:"Btch" json:"batch,omitempty"`
	TransactionDetails *[]TransactionDetail `xml:"TxDtls" json:"transactionDetails,omitempty"`
}

// BatchInformation represents the 'Btch' XML tag.
type BatchInformation struct {
	MessageId            string `xml:"MsgId" json:"messageId"`
	PaymentInformationId string `xml:"PmtInfId" json:"paymentInformationId"`
	//NbOfTxs, optional
	//TtlAmt, optional
	//CdtDbtInd, optional
}

// TransactionDetail represents the 'TxDtls' XML tag.
type TransactionDetail struct {
	References    *TransactionReferences `xml:"Refs" json:"references"`
	AmountDetails *AmountDetails         `xml:"AmtDtls" json:"amountDetails"`
	//Related Parties
	// Related Dates
	RemittanceInformation *RemittanceInformation `xml:"RmtInf" json:"remittanceInformation"`
}

// TransactionReferences represents the 'Refs' XML tag in a TransactionDetail.
type TransactionReferences struct {
	MessageId                *string `xml:"MsgId" json:"messageId,omitempty"`
	AccountServicerReference *string `xml:"AcctSvcrRef" json:"accountServicerReference,omitempty"`
	PaymentInformationId     *string `xml:"PmtInfId" json:"paymentInformationId,omitempty"`
	InstructionId            *string `xml:"InstrId" json:"instructionId,omitempty"`
	EndToEndId               *string `xml:"EndToEndId" json:"endToEndId"`
	// TxId, optional
	// MndtId (MandateId), optional
	// ChqNb (ChequeNumber), optional
	// ClrSysRef, optional
	// Prtry, optional
}

// RemittanceInformation represents the 'RmtInf' XML tag.
type RemittanceInformation struct {
	Unstructured *[]string                          `xml:"Ustrd" json:"unstructured,omitempty"`
	Structured   *[]StructuredRemittanceInformation `xml:"Strd" json:"structured,omitempty"`
}

// StructuredRemittanceInformation represents the 'Strd' XML tag.
type StructuredRemittanceInformation struct {
	ReferredDocumentInformation  *ReferredDocumentInformation  `xml:"RfrdDocInf" json:"referredDocumentInformation,omitempty"`
	ReferredDocumentAmount       *ReferredDocumentAmount       `xml:"RfrdDocAmt" json:"referredDocumentAmount,omitempty"`
	CreditorReferenceInformation *CreditorReferenceInformation `xml:"CdtrRefInf" json:"creditorReferenceInformation"`
}

// CreditorReferenceInformation represents the 'CdtrRefInf' XML tag.
type CreditorReferenceInformation struct {
	Type      *CodeOrProprietary `xml:"Tp>CdOrPrtry" json:"type,omitempty"`
	Reference *string            `xml:"Ref" json:"reference,omitempty"`
}

// RefferedDocumentInformation represents the 'RfrdDocInf' XML tag.
type ReferredDocumentInformation struct {
	Type   CodeOrProprietary `xml:"Tp>CdOrPrtry" json:"type,omitempty"`
	Number string            `xml:"Nb" json:"number,omitempty"`
}

// ReferredDocumentAmount represents the 'RfrdDocAmt' XML tag.
type ReferredDocumentAmount struct {
	RemittedAmount   *Amount `xml:"RmtdAmt" json:"remittedAmount,omitempty"`
	DuePayableAmount *Amount `xml:"DuePyblAmt" json:"duePayableAmount,omitempty"`
}
