package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response"
	"net/http"
)

func userIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || auth != viper.GetString("api.pass") {
			logrus.Errorf("Authorization header is invalid")

			response.NewErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "Authorization header is invalid")

			return
		}

		c.Next()
		return
	}
}
