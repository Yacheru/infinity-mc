package middlewares

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"

	"payments-service/internal/server/http/handlers"
	"payments-service/pkg/constants"
)

func ValidatePaymentParams() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nickname := ctx.Query("nickname")
		price := ctx.Query("price")
		service := ctx.Query("donat")
		duration := ctx.Query("duration")

		_, err := mail.ParseAddress(ctx.Query("email"))

		if err != nil {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		if service != constants.Nickname && service != constants.Badge && service != constants.Hronon {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, constants.ErrService.Error())
			return
		}
		if nickname == "" || price == "" || duration == "" {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, constants.ErrReqParams.Error())
			return
		}

		ctx.Next()
	}
}
