package payment

import (
	"backend/models"
	"backend/validation"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

func handleError(ctx *gin.Context, err error, forceHTTPStatusCode int) {
	var o models.JSONAPIErrorObject
	o.Code = err.Error()
	if forceHTTPStatusCode > 0 {
		o.Status = forceHTTPStatusCode
	} else {
		if gorm.IsRecordNotFoundError(err) {
			o.Status = http.StatusNotFound
		} else if validation.IsValidationError(err) {
			o.Status = http.StatusBadRequest
		} else {
			o.Status = http.StatusInternalServerError
		}
	}
	ctx.AbortWithStatusJSON(o.Status, o)
	log.Error().Msgf("%v", err)
}

//
// Handler is handling gin http request
type Handler struct {
	PaymentResource Resource
}

//
// List returns a list of payments
// /api/v1/payment
func (h *Handler) List(ctx *gin.Context) {
	payments, err := h.PaymentResource.List(ctx)
	if err != nil {
		handleError(ctx, err, 0)
	} else {
		//
		// Populate Links/Self
		payments.Links.Self = fmt.Sprintf("http://%s%s", ctx.Request.Host, ctx.Request.URL.String())
		//
		// Return JSON object
		ctx.JSON(http.StatusOK, payments)
	}
}

//
// Find retreives an existing payment
// /api/v1/payment/:paymentID", payment.Find)
func (h *Handler) Find(ctx *gin.Context) {
	paymentID := ctx.Param("paymentID")
	payment, err := h.PaymentResource.Find(ctx, paymentID)
	if err != nil {
		handleError(ctx, err, 0)
	} else {
		//
		// Populate Links/Self
		payment.Links.Self = fmt.Sprintf("http://%s%s", ctx.Request.Host, ctx.Request.URL.String())
		//
		// Return JSON object
		ctx.JSON(http.StatusOK, payment)
	}
}

//
// Create creats a new payment
// /api/v1/payment", payment.Create)
func (h *Handler) Create(ctx *gin.Context) {
	var payment models.PaymentInput
	err := json.NewDecoder(ctx.Request.Body).Decode(&payment)
	if err != nil {
		handleError(ctx, err, http.StatusBadRequest)
	} else {
		id, err := h.PaymentResource.Create(ctx, payment)
		if err != nil {
			handleError(ctx, err, 0)
		} else {
			ctx.JSON(http.StatusCreated, models.JSONAPISuccessObject{Status: http.StatusCreated, ID: id})
		}
	}
}

//
// Update updates an existing payment
// /api/v1/payment/:paymentID", payment.Update)
func (h *Handler) Update(ctx *gin.Context) {
	paymentID := ctx.Param("paymentID")
	var payment models.PaymentInput
	defer ctx.Request.Body.Close()
	err := json.NewDecoder(ctx.Request.Body).Decode(&payment)
	if err != nil {
		handleError(ctx, err, http.StatusBadRequest)
	} else {
		err := h.PaymentResource.Update(ctx, paymentID, payment)
		if err != nil {
			handleError(ctx, err, 0)
		} else {
			ctx.JSON(http.StatusAccepted, models.JSONAPISuccessObject{Status: http.StatusAccepted, ID: paymentID})
		}
	}
}

//
// Delete deletes an existing payment
// /api/v1/payment/:paymentID", payment.Delete)
func (h *Handler) Delete(ctx *gin.Context) {
	paymentID := ctx.Param("paymentID")
	err := h.PaymentResource.Delete(ctx, paymentID)
	if err != nil {
		handleError(ctx, err, 0)
	} else {
		ctx.JSON(http.StatusOK, models.JSONAPISuccessObject{Status: http.StatusOK, ID: paymentID})
	}
}
