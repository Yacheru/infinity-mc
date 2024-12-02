package middleware

import (
	"jwt-auth-service/internal/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"jwt-auth-service/internal/server/http/handlers"
)

func ValidateUUID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, exists := ctx.Get("payload")
		if !exists {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "payload required")
			return
		}

		payload, ok := value.(*entities.Claims)
		if !ok {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "invalid payload")
			return
		}

		id := ctx.Param("uuid")
		if id == "" {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "uuid is required")
			return
		}

		if _, err := uuid.Parse(id); err != nil {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "uuid is invalid")
			return
		}

		if payload.Subject == id {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "you can't modify yourself")
			return
		}

		ctx.Next()
	}
}
