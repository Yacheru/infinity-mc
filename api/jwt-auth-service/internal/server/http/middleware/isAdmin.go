package middleware

import (
	"github.com/gin-gonic/gin"
	"jwt-auth-service/internal/entities"
	"jwt-auth-service/internal/server/http/handlers"
	"jwt-auth-service/pkg/constants"
	"net/http"
)

func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, exists := ctx.Get("payload")
		if !exists {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "payload required")
			return
		}

		claims, ok := value.(*entities.Claims)
		if !ok {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "invalid payload")
			return
		}

		if claims.Role != constants.Admin {
			handlers.NewErrorResponse(ctx, http.StatusForbidden, "forbidden")
			return
		}

		ctx.Next()
	}
}
