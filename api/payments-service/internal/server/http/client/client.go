package client

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"payments-service/pkg/util/constants"
)

type Client struct {
	client    *http.Client
	account   string
	secretKey string
}

func NewClient(account, secretKey string) *Client {
	return &Client{
		client:    new(http.Client),
		account:   account,
		secretKey: secretKey,
	}
}

func (c *Client) makeRequest(method string, endpoint string, body []byte) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", constants.YooKassa, endpoint)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
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
