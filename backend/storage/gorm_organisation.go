package storage

import (
	"backend/models"
	"context"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//
// OrganisationGORM is the GORM implementation of Organisation
type OrganisationGORM struct {
	db *gorm.DB
}

//
// Create creates a new organisation record
func (s *OrganisationGORM) Create(ctx context.Context, organisationModel models.OrganisationDB) error {
	if organisationModel.ID == "" {
		organisationModel.ID = uuid.New().String()
	}
	err := s.db.Create(&organisationModel).Error
	return err
}
