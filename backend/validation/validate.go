package validation

import (
	"github.com/go-playground/validator"
)

//
// Validate is a initialized instance of the validator
var Validate *validator.Validate

//
// Register custom validators
func init() {
	Validate = validator.New()
	Validate.RegisterValidation("decimalNumber", DecimalNumber)
	Validate.RegisterValidation("currency", Currency)
}
