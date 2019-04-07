package payment

import (
	"backend/models"
	"context"
)

func transformFromDB(ctx context.Context, input models.PaymentDB) models.Payment {
	//
	// Transform
	var output models.Payment
	output.Type = "Payment"
	version := int32(0)
	output.Version = &version
	output.ID = input.ID
	output.OrganisationID = input.OrganisationID
	output.Attributes.Amount = input.Amount
	output.Attributes.Currency = input.Currency
	output.Attributes.EndToEndReference = input.EndToEndReference
	output.Attributes.NumericReference = input.NumericReference
	output.Attributes.PaymentID = input.PaymentID
	output.Attributes.PaymentPurpose = input.PaymentPurpose
	output.Attributes.PaymentScheme = input.PaymentScheme
	output.Attributes.PaymentType = input.PaymentType
	output.Attributes.ProcessingDate = input.ProcessingDate
	output.Attributes.Reference = input.Reference
	output.Attributes.SchemePaymentType = input.SchemePaymentType
	output.Attributes.SchemePaymentSubType = input.SchemePaymentSubType
	//
	// Charges information
	output.Attributes.ChargesInformation.BearerCode = input.ChargesInformationBearerCode
	output.Attributes.ChargesInformation.ReceiverChargesAmount = input.ChargesInformationReceiverChargesAmount
	output.Attributes.ChargesInformation.ReceiverChargesCurrency = input.ChargesInformationReceiverChargesCurrency
	for _, charge := range input.ChargesInformationSenderCharges {
		var c models.Charge
		c.Amount = charge.Amount
		c.Currency = charge.Currency
		output.Attributes.ChargesInformation.SenderCharges = append(output.Attributes.ChargesInformation.SenderCharges, c)
	}
	//
	// Fx
	output.Attributes.Fx.ContractReference = input.FxContractReference
	output.Attributes.Fx.ExchangeRate = input.FxExchangeRate
	output.Attributes.Fx.OriginalAmount = input.FxOriginalAmount
	output.Attributes.Fx.OriginalCurrency = input.FxOriginalCurrency
	//
	// Sponsor Party
	output.Attributes.SponsorParty.AccountNumber = input.SponsorParty.AccountNumber
	output.Attributes.SponsorParty.BankID = input.SponsorParty.BankID
	output.Attributes.SponsorParty.BankIDCode = input.SponsorParty.BankIDCode
	//
	// Debtor Party
	output.Attributes.DebtorParty.AccountName = input.DebtorParty.AccountName
	output.Attributes.DebtorParty.BankID = input.DebtorParty.BankID
	output.Attributes.DebtorParty.BankIDCode = input.DebtorParty.BankIDCode
	output.Attributes.DebtorParty.AccountNumber = input.DebtorParty.AccountNumber
	output.Attributes.DebtorParty.AccountNumberCode = input.DebtorParty.AccountNumberCode
	output.Attributes.DebtorParty.Address = input.DebtorParty.Address
	output.Attributes.DebtorParty.Name = input.DebtorParty.Name
	//
	// Beneficiary Party
	output.Attributes.BeneficiaryParty.AccountName = input.BeneficiaryParty.AccountName
	output.Attributes.BeneficiaryParty.BankID = input.BeneficiaryParty.BankID
	output.Attributes.BeneficiaryParty.BankIDCode = input.BeneficiaryParty.BankIDCode
	output.Attributes.BeneficiaryParty.AccountType = &input.BeneficiaryParty.AccountType
	output.Attributes.BeneficiaryParty.AccountNumber = input.BeneficiaryParty.AccountNumber
	output.Attributes.BeneficiaryParty.AccountNumberCode = input.BeneficiaryParty.AccountNumberCode
	output.Attributes.BeneficiaryParty.Address = input.BeneficiaryParty.Address
	output.Attributes.BeneficiaryParty.Name = input.BeneficiaryParty.Name
	//
	//
	return output
}

func transformToDB(ctx context.Context, input models.PaymentInput) models.PaymentDB {
	//
	// Transform
	var output models.PaymentDB
	output.Amount = input.Attributes.Amount
	output.Currency = input.Attributes.Currency
	output.EndToEndReference = input.Attributes.EndToEndReference
	output.NumericReference = input.Attributes.NumericReference
	output.PaymentID = input.Attributes.PaymentID
	output.PaymentPurpose = input.Attributes.PaymentPurpose
	output.PaymentScheme = input.Attributes.PaymentScheme
	output.PaymentType = input.Attributes.PaymentType
	output.ProcessingDate = input.Attributes.ProcessingDate
	output.Reference = input.Attributes.Reference
	output.SchemePaymentType = input.Attributes.SchemePaymentType
	output.SchemePaymentSubType = input.Attributes.SchemePaymentSubType
	//
	// Charges information
	output.ChargesInformationBearerCode = input.Attributes.ChargesInformation.BearerCode
	output.ChargesInformationReceiverChargesAmount = input.Attributes.ChargesInformation.ReceiverChargesAmount
	output.ChargesInformationReceiverChargesCurrency = input.Attributes.ChargesInformation.ReceiverChargesCurrency
	for _, charge := range input.Attributes.ChargesInformation.SenderCharges {
		var c models.ChargeDB
		c.Amount = charge.Amount
		c.Currency = charge.Currency
		output.ChargesInformationSenderCharges = append(output.ChargesInformationSenderCharges, c)
	}
	//
	// Fx
	output.FxContractReference = input.Attributes.Fx.ContractReference
	output.FxExchangeRate = input.Attributes.Fx.ExchangeRate
	output.FxOriginalAmount = input.Attributes.Fx.OriginalAmount
	output.FxOriginalCurrency = input.Attributes.Fx.OriginalCurrency
	//
	// Sponsor Party
	output.SponsorParty = new(models.PartyDB)
	output.SponsorParty.AccountNumber = input.Attributes.SponsorParty.AccountNumber
	//
	// Debtor Party
	output.DebtorParty = new(models.PartyDB)
	output.DebtorParty.AccountNumber = input.Attributes.DebtorParty.AccountNumber
	//
	// Beneficiary Party
	output.BeneficiaryParty = new(models.PartyDB)
	output.BeneficiaryParty.AccountNumber = input.Attributes.BeneficiaryParty.AccountNumber
	//
	//
	return output
}
