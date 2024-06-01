package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response"
	"net/http"
)

func ValidatePaymentParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		nickname := c.Query("nickname")
		price := c.Query("price")
		email := c.Query("email")
		donat := c.Query("donat")
		duration := c.Query("duration")

		if nickname == "" ||
			price == "" ||
			email == "" ||
			donat == "" ||
			duration == "" {
			logrus.Errorf("invalid request parameters")

			response.NewErrorResponse(c, http.StatusBadRequest, "invalid request parameters", "")

			return
		}

		c.Next()
		return
	}
}
