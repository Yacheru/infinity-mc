package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"websocket-chat-service/internal/server/rest/handlers"
)

func ValidatePagination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limit := ctx.DefaultQuery("limit", "50")
		offset := ctx.DefaultQuery("offset", "0")

		if _, err := strconv.Atoi(limit); err != nil {
			handlers.ErrorResponse(ctx, http.StatusBadRequest, "limit must be a positive integer")
			return
		}

		if _, err := strconv.Atoi(offset); err != nil {
			handlers.ErrorResponse(ctx, http.StatusBadRequest, "offset must be a positive integer")
			return
		}

		ctx.Next()
	}
}
