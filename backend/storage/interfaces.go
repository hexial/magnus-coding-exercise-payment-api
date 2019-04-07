package storage

import (
	"backend/models"
	"context"
)

//
// Organisation handle storing of organisation data
type Organisation interface {
	//
	// Create creates a new organisation record
	Create(ctx context.Context, organisationModel models.OrganisationDB) error
}

//
// Party handle storing of party data
type Party interface {
	//
	// Create creates a new party
	Create(ctx context.Context, party models.PartyDB) error
	//
	// FindByAccountNumber selects a party record with account number supplied
	FindByAccountNumber(ctx context.Context, accountNumber string) (models.PartyDB, error)
}

//
// Payment handle storing of payment data
type Payment interface {
	//
	// List will select all payments for supplied organisationID
	List(ctx context.Context) ([]models.PaymentDB, error)
	//
	// Find will select one payment for supplied organisation and payment ID
	Find(ctx context.Context, paymentID string) (models.PaymentDB, error)
	//
	// Create will create one new payment record
	Create(ctx context.Context, paymentDB models.PaymentDB) (string, error)
	//
	// Update will update one existing payment record
	Update(ctx context.Context, p models.PaymentDB) error
	//
	// Delete will delete one payment record for supplied organisation and payment ID
	Delete(ctx context.Context, paymentID string) error
}
