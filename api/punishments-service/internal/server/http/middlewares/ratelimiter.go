package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"punishments-service/internal/server/http/handlers"
	"punishments-service/pkg/rate"
)

func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limiter := rate.Limiter(ctx.RemoteIP())

		if !limiter.Allow() {
			handlers.NewErrorResponse(ctx, http.StatusTooManyRequests, "too many requests")
			return
		}

		ctx.Next()
	}
}
