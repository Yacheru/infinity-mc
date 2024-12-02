package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"punishments-service/internal/server/http/handlers"
	"strconv"
)

func ValidateParams() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "invalid or missing limit parameter")
			return
		}

		category := ctx.Query("type")

		if category != "bans" && category != "mutes" && category != "warns" {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "invalid or missing type parameter")
			return
		}

		ctx.Next()
	}
}
