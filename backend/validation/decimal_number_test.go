// +build unit

package validation

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDecimalNumber(t *testing.T) {
	Convey("Given we are validating numbers with 2 decimals", t, func() {
		d := 2
		Convey("When the system receives 1000", func() {
			s := "1000"
			Convey("Then validation should fail", func() {
				So(decimalNumberString(s, d), ShouldBeFalse)
			})
		})
		Convey("When the system receives 1000.0", func() {
			s := "1000.0"
			Convey("Then validation should fail", func() {
				So(decimalNumberString(s, d), ShouldBeFalse)
			})
		})
		Convey("When the system receives 1000.00", func() {
			s := "1000.00"
			Convey("Then validation should be succesfull", func() {
				So(decimalNumberString(s, d), ShouldBeTrue)
			})
			Convey("When the system receives 1000.000", func() {
				s := "1000.000"
				Convey("Then validation should fail", func() {
					So(decimalNumberString(s, d), ShouldBeFalse)
				})
			})
			Convey("When the system receives 1 000.00", func() {
				s := "1 000.00"
				Convey("Then validation should fail", func() {
					So(decimalNumberString(s, d), ShouldBeFalse)
				})
			})
			Convey("When the system receives 1000,00", func() {
				s := "1000,00"
				Convey("Then validation should fail", func() {
					So(decimalNumberString(s, d), ShouldBeFalse)
				})
			})
			Convey("When the system receives ' 1000.00'", func() {
				s := " 1000.00"
				Convey("Then validation should fail", func() {
					So(decimalNumberString(s, d), ShouldBeFalse)
				})
			})
			Convey("When the system receives '1000.00 '", func() {
				s := "1000.00 "
				Convey("Then validation should fail", func() {
					So(decimalNumberString(s, d), ShouldBeFalse)
				})
			})
		})
	})
}
