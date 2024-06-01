package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response"
	"net/http"
	"strconv"
)

func ValidatePunishmentsParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			logrus.Errorf("invalid limit parameter, %s", err.Error())

			response.NewErrorResponse(c, http.StatusBadRequest, "invalid limit parameter", err.Error())

			return
		}

		allowedTypes := map[string]bool{
			"bans":  true,
			"mutes": true,
			"warns": true,
		}

		if _, ok := allowedTypes[c.Query("type")]; !ok {
			logrus.Error("invalid or missing punishment type parameter")

			response.NewErrorResponse(c, http.StatusBadRequest, "invalid or missing punishment type parameter", "")

			return
		}

		c.Next()
		return
	}
}
