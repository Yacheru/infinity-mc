package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/yacheru/infinity-mc.ru/backend"
	"github.com/yacheru/infinity-mc.ru/backend/configs"
	"io"
	"log"
	"net/http"
)

const (
	endpoint  = "https://api.yookassa.ru/v3/payments/"
	returnUrl = "https://infinity-mc.ru/"
)

// SetHeaders Устанавливаем заголовки для api запроса
func SetHeaders(r *http.Request) *http.Request {
	if err := configs.InitConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err.Error())
	}

	key := uuid.NewString()

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", viper.GetString("ykassa.pass"))
	r.Header.Set("Idempotence-Key", key)

	return r
}

func (h *Handler) CreatePayment(c *gin.Context) {
	nickname := c.Query("nickname")
	email := c.Query("email")
	amount := c.Query("amount")
	donatType := c.Query("donat")
	description := fmt.Sprintf("Услуга: %s\nНикнейм: %s\nПочта: %s", donatType, nickname, email)

	payment := backend.PaymentStruct{
		Amount: backend.AmountType{
			Value:    amount,
			Currency: "RUB",
		},
		Receipt: struct {
			Customer struct {
				Email string `json:"email"`
			} `json:"customer"`
			Items [1]struct {
				Description string             `json:"description"`
				Amount      backend.AmountType `json:"amount"`
				VatCode     int                `json:"vat_code"`
				Quantity    string             `json:"quantity"`
			} `json:"items"`
		}{
			Customer: struct {
				Email string `json:"email"`
			}{
				Email: email,
			},
			Items: [1]struct {
				Description string             `json:"description"`
				Amount      backend.AmountType `json:"amount"`
				VatCode     int                `json:"vat_code"`
				Quantity    string             `json:"quantity"`
			}{
				{
					Description: donatType,
					Amount: backend.AmountType{
						Value:    amount,
						Currency: "RUB",
					},
					VatCode:  1,
					Quantity: "1",
				},
			},
		},
		Capture: true,
		Confirmation: struct {
			Type      string `json:"type"`
			ReturnUrl string `json:"return_url"`
		}{
			Type:      "redirect",
			ReturnUrl: returnUrl,
		},
		Description: description,
	}

	params, _ := json.Marshal(payment)

	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(params))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	r = SetHeaders(r)

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
