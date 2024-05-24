package handler

import (
	"fmt"
	"github.com/spf13/viper"
	httpClient "github.com/yacheru/infinity-mc.ru/backend/internal/http-client"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response/payments"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePayment(c *gin.Context) {
	nickname := c.Query("nickname")
	amount := c.Query("amount")
	email := c.Query("email")
	donat := c.Query("donat")

	yooClient := httpClient.NewClient(viper.GetString("ykassa.shopid"), viper.GetString("ykassa.pass"))

	pH := httpClient.NewPaymentHandler(yooClient)

	payment, _ := pH.CreatePayment(&payments.Payment{
		Amount: &payments.Amount{
			Value:    amount,
			Currency: "RUB",
		},
		PaymentMethod: payments.PaymentMethodType("bank_card"),
		Receipt: &payments.Receipt{
			Customer: &payments.Email{
				Email: email,
			},
			Items: &[1]payments.Items{
				{
					Description: fmt.Sprintf("Услуга %s", donat),
					Amount: &payments.Amount{
						Value:    amount,
						Currency: "RUB",
					},
					VatCode:  1,
					Quantity: "1",
				},
			},
		},
		Capture: true,
		Confirmation: payments.Redirect{
			Type:      "redirect",
			ReturnURL: "https://infinity-mc.ru/",
		},
		Description: fmt.Sprintf("%s %s %s", nickname, donat, amount),
	})

	c.JSON(http.StatusOK, payment)
}

func (h *Handler) Accept(c *gin.Context) {
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Printf("Received webhook: %+v\n", jsonData)

	return
}
