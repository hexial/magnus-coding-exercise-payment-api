package main

import (
	"backend/constants"
	"backend/fixtures"
	"backend/payment"
	"backend/storage"
	"os"
	"time"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//
// FakeAuthentication is just population values to the context
/// This would normally be doing when parsing a JWT-token or similar
func FakeAuthentication(c *gin.Context) {
	c.Set(constants.ContextOrganisationID, fixtures.MyOrganisationID)
}

func main() {
	//
	// Zerolog
	zerolog.TimeFieldFormat = ""
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	//
	// Startup info
	log.Info().Msg("Magnus Coding Exercise Payment API - Backend")
	//
	// Setup storage, resource, handle
	_, _, paymentStorage := storage.NewGORM()
	defer storage.Close()
	paymentResource := payment.Resource{PaymentStorage: paymentStorage}
	paymentHandle := &payment.Handler{PaymentResource: paymentResource}

	//
	// GIN router
	router := gin.New()
	router.Use(ginzerolog.Logger("gin"))
	router.Use(FakeAuthentication)
	router.GET("/api/v1/payments", paymentHandle.List)
	router.POST("/api/v1/payments", paymentHandle.Create)
	router.GET("/api/v1/payments/:paymentID", paymentHandle.Find)
	router.PUT("/api/v1/payments/:paymentID", paymentHandle.Update)
	router.DELETE("/api/v1/payments/:paymentID", paymentHandle.Delete)
	router.Run()
}
