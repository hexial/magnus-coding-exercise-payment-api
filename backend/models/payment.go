package models

//
// Charge single charge
type Charge struct {
	Amount   string `json:"amount,omitempty" validate:"required,decimalNumber=2"`
	Currency string `json:"currency,omitempty" validate:"required,currency"`
}

//
// Links Links to endpoints
type Links struct {
	Self string `json:"self,omitempty"`
}

//
// DebtorParty the sender
type DebtorParty struct {
	AccountName       string `json:"account_name,omitempty" validate:"required"`
	AccountNumber     string `json:"account_number,omitempty" validate:"required"`
	AccountNumberCode string `json:"account_number_code,omitempty" validate:"required"`
	Address           string `json:"address,omitempty" validate:"required"`
	BankID            string `json:"bank_id,omitempty" validate:"required"`
	BankIDCode        string `json:"bank_id_code,omitempty" validate:"required"`
	Name              string `json:"name,omitempty" validate:"required"`
}

//
// SponsorParty the sponsor
type SponsorParty struct {
	AccountNumber string `json:"account_number,omitempty" validate:"required"`
	BankID        string `json:"bank_id,omitempty" validate:"required"`
	BankIDCode    string `json:"bank_id_code,omitempty" validate:"required"`
}

//
// BeneficiaryParty the recipient
type BeneficiaryParty struct {
	AccountName       string `json:"account_name,omitempty" validate:"required"`
	AccountNumber     string `json:"account_number,omitempty" validate:"required"`
	AccountNumberCode string `json:"account_number_code,omitempty" validate:"required"`
	AccountType       *int32 `json:"account_type" validate:"required"`
	Address           string `json:"address,omitempty" validate:"required"`
	BankID            string `json:"bank_id,omitempty" validate:"required"`
	BankIDCode        string `json:"bank_id_code,omitempty" validate:"required"`
	Name              string `json:"name,omitempty" validate:"required"`
}

//
// ChargesInformation the charges for the payment
type ChargesInformation struct {
	BearerCode              string   `json:"bearer_code,omitempty" validate:"required"`
	SenderCharges           []Charge `json:"sender_charges,omitempty" validate:"required,dive"`
	ReceiverChargesAmount   string   `json:"receiver_charges_amount,omitempty" validate:"required,decimalNumber=2"`
	ReceiverChargesCurrency string   `json:"receiver_charges_currency,omitempty" validate:"required,currency"`
}

//
// Fx foreign exchange of a payment
type Fx struct {
	ContractReference string `json:"contract_reference,omitempty" validate:"required"`
	ExchangeRate      string `json:"exchange_rate,omitempty" validate:"required,decimalNumber=5"`
	OriginalAmount    string `json:"original_amount,omitempty" validate:"required,decimalNumber=2"`
	OriginalCurrency  string `json:"original_currency,omitempty" validate:"required,currency"`
}

//
// Attributes the attributes of a payment
type Attributes struct {
	Amount               string             `json:"amount,omitempty" validate:"required,decimalNumber=2"`
	BeneficiaryParty     BeneficiaryParty   `json:"beneficiary_party,omitempty" validate:"required"`
	ChargesInformation   ChargesInformation `json:"charges_information,omitempty" validate:"required"`
	Currency             string             `json:"currency,omitempty" validate:"required,currency"`
	DebtorParty          DebtorParty        `json:"debtor_party,omitempty" validate:"required"`
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
	SponsorParty         SponsorParty       `json:"sponsor_party,omitempty" validate:"required"`
}

//
// Payment is a single payment
type Payment struct {
	Type           string     `json:"type,omitempty" validate:"required,eq=Payment"`
	ID             string     `json:"id,omitempty" validate:"required"`
	Version        *int32     `json:"version" validate:"required,eq=0"`
	OrganisationID string     `json:"organisation_id,omitempty" validate:"required"`
	Attributes     Attributes `json:"attributes,omitempty" validate:"required"`
	Links          Links      `json:"links,omitempty"`
}

//
// Payments is a list of payments
type Payments struct {
	Data  []Payment `json:"data,omitempty" validate:"required,dive"`
	Links Links     `json:"links,omitempty"`
}
