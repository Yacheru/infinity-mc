package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response"
	"net/http"
)

func AllowedIps() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedIps := map[string]bool{
			"185.71.76.0/27":   true,
			"185.71.77.0/27":   true,
			"77.75.153.0/25":   true,
			"77.75.156.11":     true,
			"77.75.156.35":     true,
			"77.75.154.128/25": true,
			"2a02:5180::/32":   true,
		}

		if _, ok := allowedIps[c.ClientIP()]; !ok {
			logrus.Errorf("%s ip not allowed", c.ClientIP())
			message := fmt.Sprintf("%s ip not allowed", c.ClientIP())

			response.NewErrorResponse(c, http.StatusForbidden, message, "")

			return
		}

		c.Next()

		return
	}
}
