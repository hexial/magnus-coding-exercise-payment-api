package validation

import (
	"github.com/go-playground/validator"
)

//
// currencies is a list of accepted currencies
var currencies = []string{"USD", "GBP"}

//
// Currency is a custom validator that validates currency.
// Validation is done by calling currencyString
// It returns true or false
func Currency(fl validator.FieldLevel) bool {
	return currencyString(fl.Field().String())
}

//
// A custom validator the validates currency.
// It returns true or false
func currencyString(input string) bool {
	for _, c := range currencies {
		if c == input {
			return true
		}
	}
	return false
}
