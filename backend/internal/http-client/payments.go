package http_client

import (
	"encoding/json"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response/payments"
	"io"
	"net/http"
)

const (
	createPayment = "payments"
)

type PaymentHandler struct {
	client *Client
}

func NewPaymentHandler(client *Client) *PaymentHandler {
	return &PaymentHandler{client: client}
}

func (p *PaymentHandler) CreatePayment(payment *payments.Payment) (*payments.Payment, error) {
	paymentJSON, err := json.MarshalIndent(payment, "", "\t")
	if err != nil {
		return nil, err
	}

	resp, err := p.client.makeRequest(http.MethodPost, createPayment, paymentJSON)
	if err != nil {
		return nil, err
	}

	paymentResponse, err := p.parsePaymentResponse(resp)
	if err != nil {
		return nil, err
	}

	return paymentResponse, nil
}

func (p *PaymentHandler) parsePaymentResponse(resp *http.Response) (*payments.Payment, error) {
	var responseBytes []byte
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	paymentResponse := payments.Payment{}
	err = json.Unmarshal(responseBytes, &paymentResponse)
	if err != nil {
		return nil, err
	}
	return &paymentResponse, nil
}
