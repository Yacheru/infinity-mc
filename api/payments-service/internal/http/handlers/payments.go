package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"payments-service/init/config"
	"payments-service/internal/http/client"
	"payments-service/internal/http/service"
	"payments-service/internal/repository/records"
	"payments-service/pkg/util/setups"
)

type PaymentsHandler struct {
	services *service.Service
}

func NewPaymentsHandler(services *service.Service) *PaymentsHandler {
	return &PaymentsHandler{
		services: services,
	}
}

func (h *PaymentsHandler) CreatePayment(c *gin.Context) {
	price := c.Query("price")
	email := c.Query("email")
	donat := c.Query("donat")
	nickname := c.Query("nickname")
	duration := c.Query("duration")

	yooClient := client.NewClient(config.ServerConfig.YKassaID, config.ServerConfig.YKassaPass)

	pH := client.NewPaymentHandler(yooClient)

	payment, err := setups.SetupPayment(pH, price, donat, nickname, duration, email)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed creating payment")

		return
	}

	NewSuccessResponse(c, http.StatusOK, "success", payment.Confirmation)
}

func (h *PaymentsHandler) Accept(c *gin.Context) {
	var paid = &records.Paid{}
	if err := c.ShouldBindJSON(&paid); err != nil {
		NewErrorResponse(c, http.StatusForbidden, "Invalid request body")

		return
	}

	// TODO: new deliver service!!!
	//if err := GiveDonat(paid.Object.Metadata.DonatType, paid.Object.Metadata.Nickname, paid.Object.Metadata.Duration); err != nil {
	//	NewErrorResponse(c, http.StatusInternalServerError, "error of donate delivery")
	//
	//	return
	//}

	//if err := h.services.Payment.CreateHistory(
	//	paid.Object.ID,
	//	paid.Object.Metadata.Nickname,
	//	paid.Object.Metadata.Price,
	//	paid.Object.Metadata.DonatType,
	//); err != nil {
	//	NewErrorResponse(c, http.StatusInternalServerError, "error creating history")
	//
	//	return
	//}

	NewSuccessResponse(c, http.StatusOK, "success", paid.Object.Metadata)
}
