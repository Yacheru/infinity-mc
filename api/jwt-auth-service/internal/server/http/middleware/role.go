package middleware

import (
	"github.com/gin-gonic/gin"
	"jwt-auth-service/internal/server/http/handlers"
	"jwt-auth-service/pkg/constants"
	"net/http"
)

func ValidateRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.Query("role")
		if role == "" {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "missing role")
			return
		}

		if role != constants.Admin && role != constants.Player {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, "invalid role")
			return
		}
		
		ctx.Next()
	}
}
