package validation

import "github.com/go-playground/validator"

//
// IsValidationError checks the type of error. Returnes true if it is an validator.ValidationErrors
func IsValidationError(err error) bool {
	_, ok := err.(validator.ValidationErrors)
	return ok
}
