package middleware

import (
	"github.com/gin-gonic/gin"
	"jwt-auth-service/internal/jwt"
	"jwt-auth-service/internal/server/http/handlers"
	"net/http"
	"strings"
)

func Auth(m jwt.TokenManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("authorization")
		if auth == "" {
			handlers.NewErrorResponse(ctx, http.StatusUnauthorized, "missing authorization header")
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			handlers.NewErrorResponse(ctx, http.StatusUnauthorized, "invalid token")
			return
		}

		payload, err := m.ValidToken(parts[1])
		if err != nil {
			handlers.NewErrorResponse(ctx, http.StatusUnauthorized, "invalid token")
			return
		}

		ctx.Set("payload", payload)

		ctx.Next()
	}
}
