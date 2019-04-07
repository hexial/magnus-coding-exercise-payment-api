// +build unit

package validation

import (
	"backend/fixtures"
	"backend/models"
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func decodePaymentString(input string) (models.Payment, error) {
	var err error
	var m models.Payment
	err = json.Unmarshal([]byte(input), &m)
	if err != nil {
		return m, err
	}
	return m, Validate.Struct(m)
}

func TestValidatePayment(t *testing.T) {
	Convey("Given we are using swagger", t, func() {
		for i, s := range fixtures.SinglePayment {
			Convey(fmt.Sprintf("When the system receives singlePayment #%d", i), func() {
				Convey("Then decoding should be successfull", func() {
					p, err := decodePaymentString(s)
					So(err, ShouldBeNil)
					So(p.Type, ShouldEqual, "Payment")
				})
			})
		}
		Convey("When the system receives 'xyz'", func() {
			Convey("Then decoding should be failing", func() {
				_, err := decodePaymentString("xyz")
				So(err.Error(), ShouldEqual, "invalid character 'x' looking for beginning of value")
			})
		})
		Convey("When the system receives payment with incorrect Attributes/Amount", func() {
			Convey("Then decoding should fail", func() {
				_, err := decodePaymentString(fixtures.IncorrectPaymentAttributesAmount)
				So(err, ShouldNotBeNil)
				So(IsValidationError(err), ShouldBeTrue)
				So(err.Error(), ShouldEqual, "Key: 'Payment.Attributes.Amount' Error:Field validation for 'Amount' failed on the 'decimalNumber' tag")
			})
		})
		Convey("When the system receives payment with incorrect currency", func() {
			Convey("Then decoding should fail", func() {
				_, err := decodePaymentString(fixtures.InvalidPaymentWrongCurrency)
				So(err, ShouldNotBeNil)
				So(IsValidationError(err), ShouldBeTrue)
				So(err.Error(), ShouldEqual, "Key: 'Payment.Attributes.ChargesInformation.SenderCharges[0].Currency' Error:Field validation for 'Currency' failed on the 'currency' tag")
			})
		})
		Convey("When the system receives payment with incorrect charge amount", func() {
			Convey("Then decoding should fail", func() {
				_, err := decodePaymentString(fixtures.InvalidPaymentWrongChargeAmount)
				So(err, ShouldNotBeNil)
				So(IsValidationError(err), ShouldBeTrue)
				So(err.Error(), ShouldEqual, "Key: 'Payment.Attributes.ChargesInformation.ReceiverChargesAmount' Error:Field validation for 'ReceiverChargesAmount' failed on the 'decimalNumber' tag")
			})
		})
		Convey("When the system receives payment with incorrect amount", func() {
			Convey("Then decoding should fail", func() {
				_, err := decodePaymentString(fixtures.InvalidPaymentWrongAmount)
				So(err, ShouldNotBeNil)
				So(IsValidationError(err), ShouldBeTrue)
				So(err.Error(), ShouldEqual, "Key: 'Payment.Attributes.ChargesInformation.SenderCharges[0].Amount' Error:Field validation for 'Amount' failed on the 'decimalNumber' tag")
			})
		})
		Convey("When the system receives payment with incorrect type", func() {
			Convey("Then decoding should fail", func() {
				_, err := decodePaymentString(fixtures.InvalidPaymentWrongType)
				So(err, ShouldNotBeNil)
				So(IsValidationError(err), ShouldBeTrue)
				So(err.Error(), ShouldEqual, "Key: 'Payment.Type' Error:Field validation for 'Type' failed on the 'eq' tag")
			})
		})
	})
}
