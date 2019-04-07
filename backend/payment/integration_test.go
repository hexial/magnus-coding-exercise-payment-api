// +build integration

package payment

import (
	"backend/fixtures"
	"backend/models"
	"backend/storage"
	"backend/validation"
	"os"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//
// Setup the environment for integration testing
// Connect to database
// Mockup some data
func TestMain(m *testing.M) {
	//
	// Setup storage
	organisationStorage, partyStorage, paymentStorage := storage.NewGORM()
	defer storage.Close()
	//
	// Create context for fixtures
	ctx := fixtures.NewContext()
	err := fixtures.Setup(ctx, organisationStorage, partyStorage, paymentStorage)
	if err != nil {
		panic(err)
	}
	//
	//
	os.Exit(m.Run())
}

//
// Run integration tests
func TestGetPayment(t *testing.T) {
	Convey("Given we are getting payments using the HTTP REST API", t, func() {
		Convey("When we're calling the /api/v1/payments/09fe827a-b3c2-4437-b999-6c0e780c0983", func() {
			uri := "/api/v1/payments/09fe827a-b3c2-4437-b999-6c0e780c0983"
			Convey("Then http status code should be 2000", func() {
				var r models.Payment
				statusCode, err := testHttpGetBody(uri, &r)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 200)
				So(validation.Validate.Struct(r), ShouldBeNil)
			})
		})
		Convey("When we're calling the /api/v1/payments/00000000-0000-0000-0000-000000000000", func() {
			uri := "/api/v1/payments/00000000-0000-0000-0000-000000000000"
			Convey("Then http status code should be 404", func() {
				statusCode, err := testHttpGetStatusCode(uri)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 404)
			})
		})
	})
}

//
// TestCreatePayment is doing a fullstack test of creating a payment
func TestCreatePaymen(t *testing.T) {
	Convey("Given we are creating payments", t, func() {
		Convey("When we call the /api/v1/payments", func() {
			uri := "/api/v1/payments"
			Convey("With a correct json input the status code should be 201", func() {
				input := `{ "type": "Payment", "version": 0, "attributes": { "amount": "100.21", "beneficiary_party": { "account_number": "31926819" }, "charges_information": { "bearer_code": "SHAR", "sender_charges": [ { "amount": "5.00", "currency": "GBP" }, { "amount": "10.00", "currency": "USD" } ], "receiver_charges_amount": "1.00", "receiver_charges_currency": "USD" }, "currency": "GBP", "debtor_party": { "account_number": "GB29XABC10161234567801" }, "end_to_end_reference": "Wil piano Jan", "fx": { "contract_reference": "FX123", "exchange_rate": "2.00000", "original_amount": "200.42", "original_currency": "USD" }, "numeric_reference": "1002001", "payment_id": "123456789012345678", "payment_purpose": "Paying for goods/services", "payment_scheme": "FPS", "payment_type": "Credit", "processing_date": "2017-01-18", "reference": "Payment for Em's piano lessons", "scheme_payment_sub_type": "InternetBanking", "scheme_payment_type": "ImmediatePayment", "sponsor_party": { "account_number": "56781234" } } }`
				statusCode, err := testHttpPost(uri, strings.NewReader(input), nil)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 201)
			})
			Convey("With an invalid json the status code should be 400", func() {
				input := `xyz`
				statusCode, err := testHttpPost(uri, strings.NewReader(input), nil)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 400)
			})
			Convey("With an incorrect json input the status code should be 400", func() {
				input := `{ "type": "Payment", "version": 0, "attributes": { "___amount": "100.21", "beneficiary_party": { "account_number": "31926819" }, "charges_information": { "bearer_code": "SHAR", "sender_charges": [ { "amount": "5.00", "currency": "GBP" }, { "amount": "10.00", "currency": "USD" } ], "receiver_charges_amount": "1.00", "receiver_charges_currency": "USD" }, "currency": "GBP", "debtor_party": { "account_number": "GB29XABC10161234567801" }, "end_to_end_reference": "Wil piano Jan", "fx": { "contract_reference": "FX123", "exchange_rate": "2.00000", "original_amount": "200.42", "original_currency": "USD" }, "numeric_reference": "1002001", "payment_id": "123456789012345678", "payment_purpose": "Paying for goods/services", "payment_scheme": "FPS", "payment_type": "Credit", "processing_date": "2017-01-18", "reference": "Payment for Em's piano lessons", "scheme_payment_sub_type": "InternetBanking", "scheme_payment_type": "ImmediatePayment", "sponsor_party": { "account_number": "56781234" } } }`
				statusCode, err := testHttpPost(uri, strings.NewReader(input), nil)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 400)
			})
		})
	})
}

//
// TestDeletePayment is doing a fullstack test of deleting a payment
func TestDeletePayment(t *testing.T) {
	Convey("Given we are deleting payments", t, func() {
		Convey("When we call the /api/v1/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", func() {
			uri := "/api/v1/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"
			Convey("The status code should be 200", func() {
				statusCode, err := testHttpDeleteStatusCode(uri)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 200)
			})
			Convey("Deleting same payment again, the status code should be 404", func() {
				statusCode, err := testHttpDeleteStatusCode(uri)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 404)
			})
		})
	})
}

//
// TestUpdatePayment is doing a fullstack test of updating a payment
func TestUpdatePayment(t *testing.T) {
	Convey("Given we are updating payments", t, func() {
		Convey("Using the API", func() {
			Convey("After the update the reference should be 'Payment for Em's guitar lessons'", func() {
				var uri string
				//
				// Base URI
				uri = "/api/v1/payments"
				//
				// Create
				var response models.JSONAPISuccessObject
				inputCreate := `{ "type": "Payment", "version": 0, "attributes": { "amount": "100.21", "beneficiary_party": { "account_number": "31926819" }, "charges_information": { "bearer_code": "SHAR", "sender_charges": [ { "amount": "5.00", "currency": "GBP" }, { "amount": "10.00", "currency": "USD" } ], "receiver_charges_amount": "1.00", "receiver_charges_currency": "USD" }, "currency": "GBP", "debtor_party": { "account_number": "GB29XABC10161234567801" }, "end_to_end_reference": "Wil piano Jan", "fx": { "contract_reference": "FX123", "exchange_rate": "2.00000", "original_amount": "200.42", "original_currency": "USD" }, "numeric_reference": "1002001", "payment_id": "123456789012345678", "payment_purpose": "Paying for goods/services", "payment_scheme": "FPS", "payment_type": "Credit", "processing_date": "2017-01-18", "reference": "Payment for Em's piano lessons", "scheme_payment_sub_type": "InternetBanking", "scheme_payment_type": "ImmediatePayment", "sponsor_party": { "account_number": "56781234" } } }`
				statusCode, err := testHttpPost(uri, strings.NewReader(inputCreate), &response)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 201)
				//
				// New URI
				uri = "/api/v1/payments/" + response.ID
				//
				// Update
				inputUpdate := `{ "type": "Payment", "version": 0, "attributes": { "amount": "100.21", "beneficiary_party": { "account_number": "31926819" }, "charges_information": { "bearer_code": "SHAR", "sender_charges": [ { "amount": "5.00", "currency": "GBP" }, { "amount": "10.00", "currency": "USD" } ], "receiver_charges_amount": "1.00", "receiver_charges_currency": "USD" }, "currency": "GBP", "debtor_party": { "account_number": "GB29XABC10161234567801" }, "end_to_end_reference": "Wil piano Jan", "fx": { "contract_reference": "FX123", "exchange_rate": "2.00000", "original_amount": "200.42", "original_currency": "USD" }, "numeric_reference": "1002001", "payment_id": "123456789012345678", "payment_purpose": "Paying for goods/services", "payment_scheme": "FPS", "payment_type": "Credit", "processing_date": "2017-01-18", "reference": "Payment for Em's guitar lessons", "scheme_payment_sub_type": "InternetBanking", "scheme_payment_type": "ImmediatePayment", "sponsor_party": { "account_number": "56781234" } } }`
				statusCode, err = testHttpPut(uri, strings.NewReader(inputUpdate), nil)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 202)
				//
				//
				var payment models.Payment
				statusCode, err = testHttpGetBody(uri, &payment)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 200)
				So(payment.Attributes.Reference, ShouldEqual, "Payment for Em's guitar lessons")
			})
			Convey("After the update the sender charge for GBP should be 1500.00", func() {
				var uri string
				//
				// Base URI
				uri = "/api/v1/payments"
				//
				// Create
				var response models.JSONAPISuccessObject
				inputCreate := `{ "type": "Payment", "version": 0, "attributes": { "amount": "100.21", "beneficiary_party": { "account_number": "31926819" }, "charges_information": { "bearer_code": "SHAR", "sender_charges": [ { "amount": "5.00", "currency": "GBP" }, { "amount": "10.00", "currency": "USD" } ], "receiver_charges_amount": "1.00", "receiver_charges_currency": "USD" }, "currency": "GBP", "debtor_party": { "account_number": "GB29XABC10161234567801" }, "end_to_end_reference": "Wil piano Jan", "fx": { "contract_reference": "FX123", "exchange_rate": "2.00000", "original_amount": "200.42", "original_currency": "USD" }, "numeric_reference": "1002001", "payment_id": "123456789012345678", "payment_purpose": "Paying for goods/services", "payment_scheme": "FPS", "payment_type": "Credit", "processing_date": "2017-01-18", "reference": "Payment for Em's piano lessons", "scheme_payment_sub_type": "InternetBanking", "scheme_payment_type": "ImmediatePayment", "sponsor_party": { "account_number": "56781234" } } }`
				statusCode, err := testHttpPost(uri, strings.NewReader(inputCreate), &response)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 201)
				//
				// New URI
				uri = "/api/v1/payments/" + response.ID
				//
				// Update
				inputUpdate := `{ "type": "Payment", "version": 0, "attributes": { "amount": "100.21", "beneficiary_party": { "account_number": "31926819" }, "charges_information": { "bearer_code": "SHAR", "sender_charges": [ { "amount": "1500.00", "currency": "GBP" }, { "amount": "10.00", "currency": "USD" } ], "receiver_charges_amount": "1.00", "receiver_charges_currency": "USD" }, "currency": "GBP", "debtor_party": { "account_number": "GB29XABC10161234567801" }, "end_to_end_reference": "Wil piano Jan", "fx": { "contract_reference": "FX123", "exchange_rate": "2.00000", "original_amount": "200.42", "original_currency": "USD" }, "numeric_reference": "1002001", "payment_id": "123456789012345678", "payment_purpose": "Paying for goods/services", "payment_scheme": "FPS", "payment_type": "Credit", "processing_date": "2017-01-18", "reference": "Payment for Em's guitar lessons", "scheme_payment_sub_type": "InternetBanking", "scheme_payment_type": "ImmediatePayment", "sponsor_party": { "account_number": "56781234" } } }`
				statusCode, err = testHttpPut(uri, strings.NewReader(inputUpdate), nil)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 202)
				//
				//
				var payment models.Payment
				statusCode, err = testHttpGetBody(uri, &payment)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 200)
				index := 0
				if payment.Attributes.ChargesInformation.SenderCharges[index].Currency != "GBP" {
					index = 1
				}
				So(payment.Attributes.ChargesInformation.SenderCharges[index].Amount, ShouldEqual, "1500.00")
			})
			Convey("With an invalid json the status code should be 400", func() {
				input := `xyz`
				statusCode, err := testHttpPut("/api/v1/payments/00000000-0000-0000-0000-000000000000", strings.NewReader(input), nil)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 400)
			})
			Convey("With an incorrect json input the status code should be 400", func() {
				input := `{ "type": "Payment", "version": 0, "attributes": { "___amount": "100.21", "beneficiary_party": { "account_number": "31926819" }, "charges_information": { "bearer_code": "SHAR", "sender_charges": [ { "amount": "5.00", "currency": "GBP" }, { "amount": "10.00", "currency": "USD" } ], "receiver_charges_amount": "1.00", "receiver_charges_currency": "USD" }, "currency": "GBP", "debtor_party": { "account_number": "GB29XABC10161234567801" }, "end_to_end_reference": "Wil piano Jan", "fx": { "contract_reference": "FX123", "exchange_rate": "2.00000", "original_amount": "200.42", "original_currency": "USD" }, "numeric_reference": "1002001", "payment_id": "123456789012345678", "payment_purpose": "Paying for goods/services", "payment_scheme": "FPS", "payment_type": "Credit", "processing_date": "2017-01-18", "reference": "Payment for Em's piano lessons", "scheme_payment_sub_type": "InternetBanking", "scheme_payment_type": "ImmediatePayment", "sponsor_party": { "account_number": "56781234" } } }`
				statusCode, err := testHttpPut("/api/v1/payments/00000000-0000-0000-0000-000000000000", strings.NewReader(input), nil)
				So(err, ShouldBeNil)
				So(statusCode, ShouldEqual, 400)
			})
		})
	})
}
