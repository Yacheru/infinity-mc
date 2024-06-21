package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"payments-service/internal/http/handlers"
)

func ValidatePaymentParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		nickname := c.Query("nickname")
		price := c.Query("price")
		email := c.Query("email")
		donat := c.Query("donat")
		duration := c.Query("duration")

		if nickname == "" || price == "" || email == "" || donat == "" || duration == "" {
			handlers.NewErrorResponse(c, http.StatusBadRequest, "invalid request parameters")

			return
		}

		c.Next()
	}
}
