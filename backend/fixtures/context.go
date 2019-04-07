package fixtures

import (
	"backend/constants"
	"context"
)

//
// NewContext generates a new context with MyOrganisationID as organisation
func NewContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.ContextOrganisationID, MyOrganisationID)
	return ctx
}
