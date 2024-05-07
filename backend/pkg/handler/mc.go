package handler

import (
	"bytes"
	"encoding/json"
	"github.com/fossoreslp/go-uuid-v4"
	"github.com/gin-gonic/gin"
	"github.com/gorcon/rcon"
	"io"
	"log"
	"net/http"
)

const (
	endpoint   = "https://api.yookassa.ru/v3/payments/"
	auth       = "Basic key"
	return_url = "https://infinity-mc.ru/"
)

func (h *Handler) Mc(c *gin.Context) {
	conn, err := rcon.Dial("port", "pass")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	response, err := conn.Execute("op yacheru")
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) CreatePayment(c *gin.Context) {
	amount := c.Query("amount")
	//nickname := c.Query("nickname")
	//color := c.Query("color")
	donatType := c.Query("donat")

	value := map[string]interface{}{
		"amount": map[string]string{
			"value":    amount,
			"currency": "RUB",
		},
		"capture": true,
		"confirmation": map[string]string{
			"type":       "redirect",
			"return_url": return_url,
		},
		"description": donatType,
	}

	params, _ := json.Marshal(value)

	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(params))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	key, _ := uuid.NewString()
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", auth)
	r.Header.Set("Idempotence-Key", key)

	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer response.Body.Close()

	var result map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
