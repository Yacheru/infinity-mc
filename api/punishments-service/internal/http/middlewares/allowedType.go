package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AllowedType() gin.HandlerFunc {
	return func(c *gin.Context) {
		qType := c.Query("type")

		if qType != "bans" && qType != "mutes" && qType != "warns" {
			logrus.Error("invalid or missing punishment type parameter")

			return
		}

		c.Next()
	}
}
