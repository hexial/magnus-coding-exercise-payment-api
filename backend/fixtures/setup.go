// +build integration unit

package fixtures

import (
	"backend/constants"
	"backend/models"
	"backend/storage"
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

//
// Setup will create organisation, parties and payments used in testing
func Setup(ctx context.Context, organisationStorage storage.Organisation, partyStorage storage.Party, paymentStorage storage.Payment) error {

	//
	// Create organisation
	var gormOrganisation models.OrganisationDB
	gormOrganisation.ID = ctx.Value(constants.ContextOrganisationID).(string)
	err := organisationStorage.Create(ctx, gormOrganisation)
	if err != nil {
		return err
	}
	//
	// Create parties
	var gormPartyWOwens models.PartyDB
	gormPartyWOwens.ID = "a242c6c0-195f-4e7a-8951-d98159ee73a3"
	gormPartyWOwens.AccountName = "W Owens"
	gormPartyWOwens.AccountNumber = "31926819"
	gormPartyWOwens.AccountNumberCode = "BBAN"
	gormPartyWOwens.AccountType = 0
	gormPartyWOwens.Address = "1 The Beneficiary Localtown SE2"
	gormPartyWOwens.BankID = "403000"
	gormPartyWOwens.BankIDCode = "GBDSC"
	gormPartyWOwens.Name = "Wilfred Jeremiah Owens"
	err = partyStorage.Create(ctx, gormPartyWOwens)
	if err != nil {
		return err
	}
	var gormPartyEJBrownBlack models.PartyDB
	gormPartyEJBrownBlack.ID = "75bb3c47-3ce5-4306-b5a2-af0f26539e55"
	gormPartyEJBrownBlack.AccountName = "EJ Brown Black"
	gormPartyEJBrownBlack.AccountNumber = "GB29XABC10161234567801"
	gormPartyEJBrownBlack.AccountNumberCode = "IBAN"
	gormPartyEJBrownBlack.Address = "10 Debtor Crescent Sourcetown NE1"
	gormPartyEJBrownBlack.BankID = "203301"
	gormPartyEJBrownBlack.BankIDCode = "GBDSC"
	gormPartyEJBrownBlack.Name = "Emelia Jane Brown"
	err = partyStorage.Create(ctx, gormPartyEJBrownBlack)
	if err != nil {
		return err
	}
	var gormPartySponsor models.PartyDB
	gormPartySponsor.ID = "88cbeb37-7f2f-41d8-9302-3ac78026735a"
	gormPartySponsor.AccountNumber = "56781234"
	gormPartySponsor.BankID = "123123"
	gormPartySponsor.BankIDCode = "GBDSC"
	err = partyStorage.Create(ctx, gormPartySponsor)
	if err != nil {
		return err
	}
	//
	// Create payments
	for _, paymentString := range SinglePayment {
		var payment models.Payment
		var paymentDB models.PaymentDB
		err := json.Unmarshal([]byte(paymentString), &payment)
		if err != nil {
			return err
		}
		paymentDB.ID = payment.ID
		paymentDB.OrganisationID = payment.OrganisationID
		paymentDB.Amount = payment.Attributes.Amount
		paymentDB.Currency = payment.Attributes.Currency
		paymentDB.EndToEndReference = payment.Attributes.EndToEndReference
		paymentDB.NumericReference = payment.Attributes.NumericReference
		paymentDB.PaymentID = payment.Attributes.PaymentID
		paymentDB.PaymentPurpose = payment.Attributes.PaymentPurpose
		paymentDB.PaymentScheme = payment.Attributes.PaymentScheme
		paymentDB.PaymentType = payment.Attributes.PaymentType
		paymentDB.ProcessingDate = payment.Attributes.ProcessingDate
		paymentDB.Reference = payment.Attributes.Reference
		paymentDB.SchemePaymentSubType = payment.Attributes.SchemePaymentSubType
		paymentDB.SchemePaymentType = payment.Attributes.SchemePaymentType
		//
		// Charges information
		paymentDB.ChargesInformationBearerCode = payment.Attributes.ChargesInformation.BearerCode
		paymentDB.ChargesInformationReceiverChargesAmount = payment.Attributes.ChargesInformation.ReceiverChargesAmount
		paymentDB.ChargesInformationReceiverChargesCurrency = payment.Attributes.ChargesInformation.ReceiverChargesCurrency
		//
		// Fx
		paymentDB.FxContractReference = payment.Attributes.Fx.ContractReference
		paymentDB.FxExchangeRate = payment.Attributes.Fx.ExchangeRate
		paymentDB.FxOriginalAmount = payment.Attributes.Fx.OriginalAmount
		paymentDB.FxOriginalCurrency = payment.Attributes.Fx.OriginalCurrency
		//
		// Parties
		debtorPartyDB, err := partyStorage.FindByAccountNumber(ctx, payment.Attributes.DebtorParty.AccountNumber)
		if err != nil {
			return err
		}
		beneficiaryPartyDB, err := partyStorage.FindByAccountNumber(ctx, payment.Attributes.BeneficiaryParty.AccountNumber)
		if err != nil {
			return err
		}
		sponsorPartyDB, err := partyStorage.FindByAccountNumber(ctx, payment.Attributes.SponsorParty.AccountNumber)
		if err != nil {
			return err
		}

		paymentDB.DebtorParty = &debtorPartyDB
		paymentDB.BeneficiaryParty = &beneficiaryPartyDB
		paymentDB.SponsorParty = &sponsorPartyDB
		//
		// Sender Charges
		for _, senderCharge := range payment.Attributes.ChargesInformation.SenderCharges {
			var senderChargeDB models.ChargeDB
			senderChargeDB.ID = uuid.New().String()
			senderChargeDB.PaymentID = paymentDB.ID
			senderChargeDB.Amount = senderCharge.Amount
			senderChargeDB.Currency = senderCharge.Currency
			paymentDB.ChargesInformationSenderCharges = append(paymentDB.ChargesInformationSenderCharges, senderChargeDB)
		}
		//
		//
		_, err = paymentStorage.Create(ctx, paymentDB)
		if err != nil {
			return err
		}
	}
	return nil
}
