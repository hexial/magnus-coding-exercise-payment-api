package storage

import (
	"backend/models"
	"context"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"github.com/rs/zerolog/log"
)

//
// PartyGORM is the GORM implementation of Party
type PartyGORM struct {
	db *gorm.DB
}

//
// Create creates a new party
func (s *PartyGORM) Create(ctx context.Context, party models.PartyDB) error {
	if party.ID == "" {
		party.ID = uuid.New().String()
	}
	err := s.db.Create(&party).Error
	return err
}

//
// FindByAccountNumber selects a party record with account number supplied
func (s *PartyGORM) FindByAccountNumber(ctx context.Context, accountNumber string) (models.PartyDB, error) {
	var p models.PartyDB
	err := s.db.Where("account_number = ?", accountNumber).First(&p).Error
	log.Info().Msgf("[PartyFindByAccountNumber] %s : %+v", accountNumber, p)
	if err != nil {
		return p, err
	}
	return p, nil
}
