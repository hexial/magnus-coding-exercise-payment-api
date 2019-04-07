package payment

import (
	"backend/models"
	"backend/storage"
	"backend/validation"
	"context"
)

//
// Resource is handling payment resources
type Resource struct {
	PaymentStorage storage.Payment
}

//
// List will list payment resources
func (r *Resource) List(ctx context.Context) (models.Payments, error) {
	//
	// Load from storage
	paymentsDB, err := r.PaymentStorage.List(ctx)
	if err != nil {
		return models.Payments{}, err
	}

	//
	// Transform
	var payments models.Payments
	for _, paymentDB := range paymentsDB {
		//
		// Transform
		payment := transformFromDB(ctx, paymentDB)
		//
		// Validate
		err = validation.Validate.StructCtx(ctx, payment)
		if err != nil {
			return models.Payments{}, err
		}
		payments.Data = append(payments.Data, payment)
	}
	return payments, nil
}

//
// Create will create a new payment
func (r *Resource) Create(ctx context.Context, payment models.PaymentInput) (string, error) {
	//
	// Validate
	err := validation.Validate.StructCtx(ctx, payment)
	if err != nil {
		return "", err
	}
	//
	// Transform to DB
	paymentDB := transformToDB(ctx, payment)
	//
	// Create in storage
	id, err := r.PaymentStorage.Create(ctx, paymentDB)
	if err != nil {
		return "", err
	}
	//
	//
	return id, nil
}

//
// Find will retreive one payment by paymentID
func (r *Resource) Find(ctx context.Context, paymentID string) (models.Payment, error) {
	//
	// Load from storage
	paymentDB, err := r.PaymentStorage.Find(ctx, paymentID)
	if err != nil {
		return models.Payment{}, err
	}
	//
	// Transform
	payment := transformFromDB(ctx, paymentDB)
	//
	// Validate
	err = validation.Validate.StructCtx(ctx, payment)
	if err != nil {
		return models.Payment{}, err
	}
	return payment, nil
}

//
// Update will update payment with paymentID
func (r *Resource) Update(ctx context.Context, paymentID string, payment models.PaymentInput) error {
	//
	// Validate
	err := validation.Validate.StructCtx(ctx, payment)
	if err != nil {
		return err
	}
	//
	// Transform to DB
	paymentDB := transformToDB(ctx, payment)
	paymentDB.ID = paymentID
	//
	// Create in storage
	err = r.PaymentStorage.Update(ctx, paymentDB)
	if err != nil {
		return err
	}
	return nil
}

//
// Delete will delete a payment will paymentID
func (r *Resource) Delete(ctx context.Context, paymentID string) error {
	//
	// Delete from storage
	err := r.PaymentStorage.Delete(ctx, paymentID)
	if err != nil {
		return err
	}
	return nil
}
