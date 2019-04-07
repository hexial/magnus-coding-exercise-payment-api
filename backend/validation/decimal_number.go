package validation

import (
	"fmt"
	"regexp"

	"strconv"

	"github.com/go-playground/validator"
)

//
// DecimalNumber is a custom validator for decimal numbers in strings
// Validation is done by calling decimalNumberString
// It returns true or false
func DecimalNumber(fl validator.FieldLevel) bool {
	decimals, err := strconv.Atoi(fl.Param())
	if err != nil {
		panic("Requires decimals as parameter")
	}
	return decimalNumberString(fl.Field().String(), decimals)
}

//
// A custom validator for decimal numbers in strings
// It validates by running a RegEx on the input string
// It returns true or false
func decimalNumberString(input string, decimals int) bool {
	match, err := regexp.MatchString(fmt.Sprintf("^\\d+\\.\\d{%d}$", decimals), input)
	if err != nil {
		panic("regexp failed")
	}
	return match
}
