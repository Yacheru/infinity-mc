package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AllowedIps() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		if ip != "185.71.76.0/27" && ip != "185.71.77.0/27" &&
			ip != "77.75.153.0/25" && ip != "77.75.156.11" &&
			ip != "77.75.156.35" && ip != "77.75.154.128/25" &&
			ip != "2a02:5180::/32" {

			logrus.Errorf("%s ip not allowed", c.ClientIP())

			return
		}

		c.Next()
	}
}
