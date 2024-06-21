package client

import (
	"encoding/json"
	"io"
	"net/http"

	"payments-service/internal/repository/records"
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

func (p *PaymentHandler) CreatePayment(payment *records.Payment) (*records.Payment, error) {
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

func (p *PaymentHandler) parsePaymentResponse(resp *http.Response) (*records.Payment, error) {
	var responseBytes []byte
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	paymentResponse := records.Payment{}
	err = json.Unmarshal(responseBytes, &paymentResponse)
	if err != nil {
		return nil, err
	}
	return &paymentResponse, nil
}
