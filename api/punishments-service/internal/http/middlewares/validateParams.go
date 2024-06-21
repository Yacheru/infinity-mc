package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"punishments-service/internal/http/handlers"
	"strconv"
)

func ValidateParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			handlers.NewErrorResponse(c, http.StatusBadRequest, "invalid or missing limit parameter")
			c.Abort()

			return
		}

		qType := c.Query("type")

		if qType != "bans" && qType != "mutes" && qType != "warns" {
			handlers.NewErrorResponse(c, http.StatusBadRequest, "invalid or missing type parameter")
			c.Abort()

			return
		}

		c.Next()
	}
}
