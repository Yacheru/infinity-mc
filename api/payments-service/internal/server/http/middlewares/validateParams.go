package middlewares

import (
	"net/http"
	"net/mail"
	"payments-service/internal/server/http/handlers"
	"payments-service/pkg/util/constants"

	"github.com/gin-gonic/gin"
)

func ValidatePaymentParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		nickname := c.Query("nickname")
		price := c.Query("price")
		service := c.Query("donat")
		duration := c.Query("duration")

		_, err := mail.ParseAddress(c.Query("email"))

		if err != nil {
			handlers.NewErrorResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		if service != constants.Nickname && service != constants.Badge && service != constants.Hronon {
			handlers.NewErrorResponse(c, http.StatusBadRequest, constants.ErrService.Error())

			return
		}

		if nickname == "" || price == "" || duration == "" {
			handlers.NewErrorResponse(c, http.StatusBadRequest, constants.ErrReqParams.Error())

			return
		}

		c.Next()
	}
}
