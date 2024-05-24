package http_client

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

const (
	yooKassaURL = "https://api.yookassa.ru/v3/"
)

type Client struct {
	client    http.Client
	account   string
	secretKey string
}

func NewClient(account, secretKey string) *Client {
	return &Client{
		client:    http.Client{},
		account:   account,
		secretKey: secretKey,
	}
}

func (c *Client) makeRequest(method string, endpoint string, body []byte) (*http.Response, error) {
	uri := fmt.Sprintf("%s%s", yooKassaURL, endpoint)

	req, err := http.NewRequest(method, uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Idempotence-Key", uuid.NewString())
	}

	req.SetBasicAuth(c.account, c.secretKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
