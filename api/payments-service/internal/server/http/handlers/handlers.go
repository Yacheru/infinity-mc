package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payments-service/internal/service"

	"payments-service/internal/entities"
)

type PaymentsHandler struct {
	service service.PaymentsService
}

func NewPaymentsHandler(service service.PaymentsService) *PaymentsHandler {
	return &PaymentsHandler{
		service: service,
	}
}

func (h *PaymentsHandler) CreatePayment(ctx *gin.Context) {
	var payment = new(entities.PaymentService)
	if err := ctx.ShouldBindJSON(payment); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	p, err := h.service.CreatePayment(ctx.Request.Context(), payment)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Failed creating payment")
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "success", p.Confirmation)
	return
}

func (h *PaymentsHandler) AcceptPayment(ctx *gin.Context) {
	var paid = new(entities.Paid)
	if err := ctx.ShouldBindJSON(paid); err != nil {
		NewErrorResponse(ctx, http.StatusForbidden, "Invalid request body")
		return
	}

	err := h.service.AcceptPayment(ctx.Request.Context(), paid)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "ok", nil)
	return
}
