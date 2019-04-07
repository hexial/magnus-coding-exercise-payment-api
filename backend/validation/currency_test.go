// +build unit

package validation

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCurrency(t *testing.T) {
	Convey("Given we are validating currencies", t, func() {
		Convey("When the system receives GBP", func() {
			Convey("Then validaton should be successfull", func() {
				So(currencyString("GBP"), ShouldBeTrue)
			})
		})
		Convey("When the system receives USD", func() {
			Convey("Then validaton should be successfull", func() {
				So(currencyString("USD"), ShouldBeTrue)
			})
		})
		Convey("When the system receives US", func() {
			Convey("Then validaton should NOT be successfull", func() {
				So(currencyString("US"), ShouldBeFalse)
			})
		})
		Convey("When the system receives BP", func() {
			Convey("Then validaton should NOT be successfull", func() {
				So(currencyString("BP"), ShouldBeFalse)
			})
		})
	})
}
