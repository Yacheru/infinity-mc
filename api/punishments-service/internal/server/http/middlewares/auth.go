package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"punishments-service/init/config"
	"punishments-service/internal/server/http/handlers"
	"punishments-service/internal/utils"
	"punishments-service/pkg/constants"
	"strings"
)

func Auth(cfg *config.Config) gin.HandlerFunc {
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

		payload, err := utils.ValidToken(parts[1], cfg.JWTSalt)
		if err != nil {
			handlers.NewErrorResponse(ctx, http.StatusUnauthorized, "invalid token")
			return
		}

		if payload.Role != constants.Admin {
			handlers.NewErrorResponse(ctx, http.StatusForbidden, "forbidden")
			return
		}

		ctx.Next()
	}
}
