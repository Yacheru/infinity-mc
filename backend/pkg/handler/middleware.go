package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func (h *Handler) userIdentity(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" || auth != viper.GetString("api.pass") {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":        http.StatusUnauthorized,
			"description": "Unauthorized",
		})
		c.Abort()
	}
}
