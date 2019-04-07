package models

//
// AttributesInput Attributes of the payment to update/create
type AttributesInput struct {
	Amount               string             `json:"amount,omitempty" validate:"required,decimalNumber=2"`
	BeneficiaryParty     PartyInput         `json:"beneficiary_party,omitempty" validate:"required"`
	ChargesInformation   ChargesInformation `json:"charges_information,omitempty" validate:"required"`
	Currency             string             `json:"currency,omitempty" validate:"required,currency"`
	DebtorParty          PartyInput         `json:"debtor_party,omitempty" validate:"required"`
	EndToEndReference    string             `json:"end_to_end_reference,omitempty" validate:"required"`
	Fx                   Fx                 `json:"fx,omitempty" validate:"required"`
	NumericReference     string             `json:"numeric_reference,omitempty" validate:"required"`
	PaymentID            string             `json:"payment_id,omitempty" validate:"required"`
	PaymentPurpose       string             `json:"payment_purpose,omitempty" validate:"required"`
	PaymentScheme        string             `json:"payment_scheme,omitempty" validate:"required"`
	PaymentType          string             `json:"payment_type,omitempty" validate:"required"`
	ProcessingDate       string             `json:"processing_date,omitempty" validate:"required"`
	Reference            string             `json:"reference,omitempty" validate:"required"`
	SchemePaymentSubType string             `json:"scheme_payment_sub_type,omitempty" validate:"required"`
	SchemePaymentType    string             `json:"scheme_payment_type,omitempty" validate:"required"`
	SponsorParty         PartyInput         `json:"sponsor_party,omitempty" validate:"required"`
}

//
// PaymentInput The payment to update/create
type PaymentInput struct {
	Attributes AttributesInput `json:"attributes,omitempty" validate:"required"`
	Type       string          `json:"type,omitempty" validate:"required,eq=Payment"`
	Version    *int32          `json:"version" validate:"required,eq=0"`
}

//
// PartyInput The account number of a party
type PartyInput struct {
	AccountNumber string `json:"account_number,omitempty" validate:"required"`
}
