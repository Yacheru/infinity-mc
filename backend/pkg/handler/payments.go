package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	httpClient "github.com/yacheru/infinity-mc.ru/backend/internal/http-client"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/middleware"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response/payments"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/service"
	"net/http"
)

func (h *Handler) CreatePayment(c *gin.Context) {
	nickname, price, email, donat, duration, ok := middleware.ValidateParams(
		c.Query("nickname"), c.Query("price"), c.Query("email"), c.Query("donat"), c.Query("duration"),
	)
	if !ok {
		logrus.Infof("invalid request parameters")

		response.NewErrorResponse(c, http.StatusBadRequest, "invalid request parameters", "")

		return
	}

	yooClient := httpClient.NewClient(viper.GetString("ykassa.shopid"), viper.GetString("ykassa.pass"))

	pH := httpClient.NewPaymentHandler(yooClient)

	payment, err := pH.CreatePayment(&payments.Payment{
		Amount: &payments.Amount{
			Value:    price,
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
						Value:    price,
						Currency: "RUB",
					},
					VatCode:  1,
					Quantity: "1",
				},
			},
		},
		Metadata: &payments.Metadata{
			Nickname:  nickname,
			DonatType: donat,
			Price:     price,
			Duration:  duration,
		},
		Capture: true,
		Confirmation: payments.Redirect{
			Type:      "redirect",
			ReturnURL: "https://infinity-mc.ru/",
		},
		Description: fmt.Sprintf("Ваш никнейм: %s, Услуга: %s, Стоимость: %s", nickname, donat, price),
	})
	if err != nil {
		logrus.Errorf("Failed to create payment: %s", err.Error())

		response.NewErrorResponse(c, http.StatusInternalServerError, "Failed to create payment", err.Error())

		return
	}

	err = h.services.Payments.AddActivePayment(payment.ID)
	if err != nil {
		logrus.Errorf("error adding active payment: %s", err.Error())

		response.NewErrorResponse(c, http.StatusInternalServerError, "error adding active payment", err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":         http.StatusOK,
		"message":      "success",
		"confirmation": payment.Confirmation,
	})

	return
}

func (h *Handler) Accept(c *gin.Context) {
	if ok := middleware.AllowedIps(c.ClientIP()); !ok {
		logrus.Infof("%s not found in allowed ips", c.ClientIP())

		response.NewErrorResponse(c, http.StatusForbidden, "ip not found in allowed ips", "no provided")

		return
	}

	var paid = &payments.Paid{}
	if err := c.ShouldBindJSON(&paid); err != nil {
		logrus.Errorf("Invalid request body, %s", err.Error())

		response.NewErrorResponse(c, http.StatusForbidden, "Invalid request body", err.Error())

		return
	}

	if err := service.GiveDonat(paid.Object.Metadata.DonatType, paid.Object.Metadata.Nickname, paid.Object.Metadata.Duration); err != nil {
		logrus.Errorf("error giving hronon: %s", err.Error())

		response.NewErrorResponse(c, http.StatusInternalServerError, "error of donate delivery", err.Error())

		return
	}

	if err := h.services.Payments.CreateHistory(
		paid.Object.ID,
		paid.Object.Metadata.Nickname,
		paid.Object.Metadata.Price,
		paid.Object.Metadata.DonatType,
	); err != nil {
		logrus.Errorf("error creating history: %s", err.Error())

		response.NewErrorResponse(c, http.StatusInternalServerError, "error creating history", err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":     http.StatusOK,
		"message":  "success",
		"metadata": paid.Object.Metadata,
	})

	return
}
