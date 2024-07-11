package camt053

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

type BankTransactionCode struct {
	Domain          *BankTransactionCodeDomain      `xml:"Domn" json:"domain,omitempty"`
	ProprietaryCode *BankTransactionProprietaryCode `xml:"Prtry" json:"proprietaryCode"`
}

type BankTransactionCodeDomain struct {
	Code   string                    `xml:"Cd" json:"code"`
	Family BankTransactionCodeFamily `xml:"Fmly" json:"family"`
}

type BankTransactionCodeFamily struct {
	Code          string `xml:"Cd" json:"code"`
	SubFamilyCode string `xml:"SubFmlyCd" json:"subFamilyCode"`
}

type BankTransactionProprietaryCode struct {
	Code   string `xml:"Cd" json:"code"`
	Issuer string `xml:"Issr" json:"issuer"`
}

type AmountDetails struct {
	InstructedAmount  *AmountAndCurrencyExchangeDetails `xml:"InstdAmt" json:"instructedAmount,omitempty"`
	TransactionAmount *AmountAndCurrencyExchangeDetails `xml:"TxAmt" json:"transactionAmount,omitempty"`
	//CntrValAmt, optional
	//AnncdPstngAmt, optional
	ProprietaryAmount *ProprietaryAmount `xml:"PrtryAmt" json:"proprietaryAmount,omitempty"`
}

type AmountAndCurrencyExchangeDetails struct {
	Amount Amount `xml:"Amt" json:"amount"`
	//CcyXchg (Currency Exchange), optional
}

type ProprietaryAmount struct {
	Type   string `xml:"Tp" json:"type"`
	Amount Amount `xml:"Amt" json:"amount"`
}

type Charge struct { // Charges are specified in 2.172 and 2.152 in SEB MIG camt052-053-054v2 spec
	TotalChargesAndTaxAmount *string `xml:"TtlChrgsAndTaxAmt" json:"totalChargesAndTaxAmount,omitempty"`
	Amount                   Amount  `xml:"Amt" json:"amount"`
	Type                     *string `xml:"Tp>Cd" json:"type"`
	Rate                     *string `xml:"Rate" json:"rate"`
	// Br, optional
	// Pty, optional
	Tax *string `xml:"Tax" json:"tax,omitempty"`
}

type EntryDetail struct {
	Batch              *BatchInformation    `xml:"Btch" json:"batch,omitempty"`
	TransactionDetails *[]TransactionDetail `xml:"TxDtls" json:"transactionDetails,omitempty"`
}

type BatchInformation struct {
	MessageId            string `xml:"MsgId" json:"messageId"`
	PaymentInformationId string `xml:"PmtInfId" json:"paymentInformationId"`
	//NbOfTxs, optional
	//TtlAmt, optional
	//CdtDbtInd, optional
}

type TransactionDetail struct {
	References    *TransactionReferences `xml:"Refs" json:"references"`
	AmountDetails *AmountDetails         `xml:"AmtDtls" json:"amountDetails"`
	//Related Parties
	// Related Dates
	RemittanceInformation *RemittanceInformation `xml:"RmtInf" json:"remittanceInformation"`
}

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

type CreditorReferenceInformation struct {
	Type      *CodeOrProprietary `xml:"Tp>CdOrPrtry" json:"type,omitempty"`
	Reference *string            `xml:"Ref" json:"reference,omitempty"`
}

type ReferredDocumentInformation struct {
	Type   CodeOrProprietary `xml:"Tp>CdOrPrtry" json:"type,omitempty"`
	Number string            `xml:"Nb" json:"number,omitempty"`
}

type ReferredDocumentAmount struct {
	RemittedAmount   *Amount `xml:"RmtdAmt" json:"remittedAmount,omitempty"`
	DuePayableAmount *Amount `xml:"DuePyblAmt" json:"duePayableAmount,omitempty"`
}
