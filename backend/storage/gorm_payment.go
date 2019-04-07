package storage

import (
	"backend/models"
	"context"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

//
// PaymentGORM is the GORM implementation of Payment
type PaymentGORM struct {
	db    *gorm.DB
	party Party
}

//
// List will select all payments for supplied organisationID
func (s *PaymentGORM) List(ctx context.Context) ([]models.PaymentDB, error) {
	organisationID := ctx.Value("OrganisationID").(string)
	//
	// Load from database
	var paymentsDB []models.PaymentDB
	err := s.db.Set("gorm:auto_preload", true).Where("organisation_id = ?", organisationID).Find(&paymentsDB).Error
	if err != nil {
		return paymentsDB, err
	}
	//
	// Populate SenderCharges
	for i := range paymentsDB {
		err = s.db.Where("payment_id = ?", paymentsDB[i].ID).Find(&paymentsDB[i].ChargesInformationSenderCharges).Error
		if err != nil {
			return paymentsDB, err
		}
	}
	return paymentsDB, nil
}

//
// Find will select one payment for supplied organisation and payment ID
func (s *PaymentGORM) Find(ctx context.Context, paymentID string) (models.PaymentDB, error) {
	organisationID := ctx.Value("OrganisationID").(string)
	//
	// Load from database
	var paymentDB models.PaymentDB
	err := s.db.Set("gorm:auto_preload", true).Where("organisation_id = ? and id = ?", organisationID, paymentID).First(&paymentDB).Error
	if err != nil {
		return paymentDB, err
	}
	//
	// Populate SenderCharges
	err = s.db.Where("payment_id = ?", paymentDB.ID).Find(&paymentDB.ChargesInformationSenderCharges).Error
	if err != nil {
		return paymentDB, err
	}
	return paymentDB, nil
}

//
// Create will create one new payment record
func (s *PaymentGORM) Create(ctx context.Context, paymentDB models.PaymentDB) (string, error) {
	//
	//
	var err error
	//
	// Load party information
	beneficiearyParty, err := s.party.FindByAccountNumber(ctx, paymentDB.BeneficiaryParty.AccountNumber)
	if err != nil {
		log.Error().Msgf("[PaymentGORM.GORM.Create] Unable to find beneficiary party. Reason: %v", err)
		return "", err
	}
	debtorParty, err := s.party.FindByAccountNumber(ctx, paymentDB.DebtorParty.AccountNumber)
	if err != nil {
		log.Error().Msgf("[PaymentGORM.GORM.Create] Unable to find debtor party. Reason: %v", err)
		return "", err
	}
	sponsorParty, err := s.party.FindByAccountNumber(ctx, paymentDB.SponsorParty.AccountNumber)
	if err != nil {
		log.Error().Msgf("[PaymentGORM.GORM.Create] Unable to find sponsor party. Reason: %v", err)
		return "", err
	}
	//
	// Clear out some foreign keys that should not be updated
	paymentDB.DebtorParty = nil
	paymentDB.BeneficiaryParty = nil
	paymentDB.SponsorParty = nil
	//
	// Database IDs
	paymentDB.OrganisationID = ctx.Value("OrganisationID").(string)
	if paymentDB.ID == "" {
		paymentDB.ID = uuid.New().String()
	}
	paymentDB.BeneficiaryPartyID = beneficiearyParty.ID
	paymentDB.DebtorPartyID = debtorParty.ID
	paymentDB.SponsorPartyID = sponsorParty.ID
	for i := range paymentDB.ChargesInformationSenderCharges {
		paymentDB.ChargesInformationSenderCharges[i].ID = uuid.New().String()
		paymentDB.ChargesInformationSenderCharges[i].PaymentID = paymentDB.ID
	}
	log.Info().Msgf("[PaymentGORM.Create] Data: %+v", paymentDB)
	//
	//
	tx := s.db.Begin()
	defer tx.Rollback()
	//
	// Create payment
	err = tx.Create(&paymentDB).Error
	if err != nil {
		log.Error().Msgf("[PaymentGORM.Create] Unable to execute tx.Create(&paymentmodelDB). Reason: %v", err)
		return "", err
	}
	for _, c := range paymentDB.ChargesInformationSenderCharges {
		err = tx.Create(&c).Error
		if err != nil {
			log.Error().Msgf("[PaymentGORM.Create] Unable to execute tx.Create(&c). Reason: %v", err)
			return "", err
		}
	}
	err = tx.Commit().Error
	if err != nil {
		return "", err
	}
	return paymentDB.ID, nil
}

//
// Update will update one existing payment record
func (s *PaymentGORM) Update(ctx context.Context, p models.PaymentDB) error {
	organisationID := ctx.Value("OrganisationID").(string)
	//
	// Find src
	var o models.PaymentDB
	err := s.db.Where("organisation_id = ? and id = ?", organisationID, p.ID).First(&o).Error
	if err != nil {
		return err
	}
	//
	// Load party information
	beneficiearyParty, err := s.party.FindByAccountNumber(ctx, p.BeneficiaryParty.AccountNumber)
	if err != nil {
		log.Error().Msgf("[PaymentGORM.Update] Unable to find beneficiary party. Reason: %v", err)
		return err
	}
	debtorParty, err := s.party.FindByAccountNumber(ctx, p.DebtorParty.AccountNumber)
	if err != nil {
		log.Error().Msgf("[PaymentGORM.Update] Unable to find debtor party. Reason: %v", err)
		return err
	}
	sponsorParty, err := s.party.FindByAccountNumber(ctx, p.SponsorParty.AccountNumber)
	if err != nil {
		log.Error().Msgf("[PaymentGORM.Update] Unable to find sponsor party. Reason: %v", err)
		return err
	}
	//
	// Clear out some foreign keys that should not be updated
	p.DebtorParty = nil
	p.BeneficiaryParty = nil
	p.SponsorParty = nil
	p.BeneficiaryPartyID = beneficiearyParty.ID
	p.DebtorPartyID = debtorParty.ID
	p.SponsorPartyID = sponsorParty.ID
	for i := range p.ChargesInformationSenderCharges {
		p.ChargesInformationSenderCharges[i].ID = uuid.New().String()
		p.ChargesInformationSenderCharges[i].PaymentID = p.ID
	}
	//
	//
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Rollback()
	err = tx.Model(&o).Update(p).Error
	if err != nil {
		return err
	}
	err = tx.Exec("delete from charge where payment_id = ?", p.ID).Error
	if err != nil {
		return err
	}
	for _, c := range p.ChargesInformationSenderCharges {
		err = tx.Save(&c).Error
		if err != nil {
			return err
		}
	}
	err = tx.Commit().Error
	if err != nil {
		return err
	}
	return nil
}

//
// Delete will delete one payment record for supplied organisation and payment ID
func (s *PaymentGORM) Delete(ctx context.Context, paymentID string) error {
	organisationID := ctx.Value("OrganisationID").(string)
	//
	// Begin transaction
	tx := s.db.Begin()
	defer tx.Rollback()
	//
	// Delete charges
	err := tx.Exec("delete from charge where payment_id = ?", paymentID).Error
	if err != nil {
		return err
	}
	//
	// Delete payment
	query := tx.Exec("delete from payment where id = ? and organisation_id = ?", paymentID, organisationID)
	rowsAffected, err := query.RowsAffected, query.Error
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else if rowsAffected > 1 {
		return fmt.Errorf("To many records affected")
	}
	//
	// Commit transaction
	err = tx.Commit().Error
	if err != nil {
		return err
	}
	return nil
}
