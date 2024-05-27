package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	httpClient "github.com/yacheru/infinity-mc.ru/backend/internal/http-client"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response/payments"
	"net/http"
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
			Items: &[]payments.Items{
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

	err := h.services.Payments.AddActivePayment(payment.ID)
	if err != nil {
		logrus.Errorf("error adding active payment: %s", err.Error())

		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "error adding active payment",
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success",
	})
	return
}

func (h *Handler) Accept(c *gin.Context) {
	//if ok := middleware.AllowedIps(c.ClientIP()); !ok {
	//	logrus.Infof("%s not found in allowed ips", c.ClientIP())
	//
	//	c.JSON(http.StatusForbidden, map[string]interface{}{
	//		"code":    http.StatusForbidden,
	//		"message": "ip not found in allowed ips",
	//	})
	//
	//	return
	//}

	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		logrus.Errorf("Invalid request body, %s", err.Error())

		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body",
		})

		return
	}

	//if err := service.GiveHronon("yacheru", "1"); err != nil {
	//	logrus.Errorf("error giving hronon: %s", err.Error())
	//
	//	c.JSON(http.StatusInternalServerError, map[string]interface{}{
	//		"code":    http.StatusInternalServerError,
	//		"message": "error giving hronon",
	//	})
	//
	//	return
	//}

	c.JSON(http.StatusOK, jsonData)

	return
}
