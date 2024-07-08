package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"

	"payments-service/init/config"
	"payments-service/internal/http/client"
	"payments-service/internal/kafka/producer"
	"payments-service/internal/records"
	"payments-service/pkg/util/setups"
)

type PaymentsHandler struct {
	producer *producer.KafkaProducer
}

func NewPaymentsHandler(producer *producer.KafkaProducer) *PaymentsHandler {
	return &PaymentsHandler{
		producer: producer,
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
	if err := c.ShouldBindJSON(paid); err != nil {
		NewErrorResponse(c, http.StatusForbidden, "Invalid request body")

		return
	}

	message, err := json.Marshal(&records.MC{
		Nickname: paid.Object.Metadata.Nickname,
		Duration: paid.Object.Metadata.Duration,
		Service:  paid.Object.Metadata.Service,
	})

	if err != nil {
		NewErrorResponse(c, http.StatusForbidden, "Invalid request body")

		return
	}

	err = h.producer.PrepareMessage(message)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to prepare message")

		return
	}

	NewSuccessResponse(c, http.StatusOK, "success", paid.Object.Metadata)
}
