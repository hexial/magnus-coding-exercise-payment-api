package models

//
// OrganisationDB table holding organisation information
type OrganisationDB struct {
	ID       string `gorm:"type:uuid;PRIMARY_KEY;"`
	Payments []PaymentDB
}

//
// TableName returnes the name of the table
func (organisationDB *OrganisationDB) TableName() string {
	return "organisation"
}

//
// PaymentDB table holding payments
type PaymentDB struct {
	ID                   string `gorm:"type:uuid;primary_key;"`
	OrganisationID       string `gorm:"type:uuid references organisation(id);not null;"`
	BeneficiaryPartyID   string `gorm:"type:uuid references party(id);not null;"`
	BeneficiaryParty     *PartyDB
	DebtorPartyID        string `gorm:"type:uuid references party(id);not null;"`
	DebtorParty          *PartyDB
	SponsorPartyID       string `gorm:"type:uuid references party(id);not null;"`
	SponsorParty         *PartyDB
	Amount               string `gorm:"type:varchar(255);not null;"`
	Currency             string `gorm:"type:char(3);not null;"`
	EndToEndReference    string `gorm:"type:varchar(255);not null;"`
	NumericReference     string `gorm:"type:varchar(255);not null;"`
	PaymentID            string `gorm:"type:varchar(255);not null;"`
	PaymentPurpose       string `gorm:"type:varchar(255);not null;"`
	PaymentScheme        string `gorm:"type:varchar(255);not null;"`
	PaymentType          string `gorm:"type:varchar(255);not null;"`
	ProcessingDate       string `gorm:"type:varchar(255);not null;"`
	Reference            string `gorm:"type:varchar(255);not null;"`
	SchemePaymentSubType string `gorm:"type:varchar(255);not null;"`
	SchemePaymentType    string `gorm:"type:varchar(255);not null;"`
	//
	// Charges information
	ChargesInformationBearerCode              string `gorm:"type:varchar(255);not null;"`
	ChargesInformationSenderCharges           []ChargeDB
	ChargesInformationReceiverChargesAmount   string `gorm:"type:varchar(255);not null;"`
	ChargesInformationReceiverChargesCurrency string `gorm:"type:char(3);not null;"`
	//
	// Fx
	FxContractReference string `gorm:"type:varchar(255);not null;"`
	FxExchangeRate      string `gorm:"type:varchar(255);not null;"`
	FxOriginalAmount    string `gorm:"type:varchar(255);not null;"`
	FxOriginalCurrency  string `gorm:"type:char(3);not null;"`
}

//
// TableName returnes the tablename
func (paymentDB *PaymentDB) TableName() string {
	return "payment"
}

//
// PartyDB table holding parties
type PartyDB struct {
	ID                string `gorm:"type:uuid;primary_key;"`
	AccountName       string `gorm:"type:varchar(255);not null;"`
	AccountNumber     string `gorm:"type:varchar(255);not null;unique_index;"`
	AccountNumberCode string `gorm:"type:varchar(255);not null;"`
	AccountType       int32  `gorm:"type:int;not null;"`
	Address           string `gorm:"type:varchar(255);not null;"`
	BankID            string `gorm:"type:varchar(255);not null;"`
	BankIDCode        string `gorm:"type:varchar(255);not null;"`
	Name              string `gorm:"type:varchar(255);not null;"`
}

//
// TableName returnes the tablename
func (partyDB *PartyDB) TableName() string {
	return "party"
}

//
// ChargeDB table holding charges
type ChargeDB struct {
	ID        string `gorm:"type:uuid;primary_key;"`
	PaymentID string `gorm:"type:uuid references payment(id)"`
	Amount    string `gorm:"type:varchar(255);not null;"`
	Currency  string `gorm:"type:char(3);not null;"`
}

//
// TableName returnes the tablename
func (chargDB *ChargeDB) TableName() string {
	return "charge"
}
