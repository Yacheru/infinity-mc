package middleware

import (
	"github.com/gin-gonic/gin"
	"jwt-auth-service/internal/server/http/handlers"
	"net/http"
	"strconv"
)

func ValidateVerifCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Query("code")

		if code == "" {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "code is required")
			return
		}

		intCode, err := strconv.Atoi(code)
		if err != nil {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "code must be integer")
			return
		}

		if intCode < 1000 || intCode > 9999 {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "code must be between 1000 and 9999")
			return
		}

		ctx.Next()
	}
}
