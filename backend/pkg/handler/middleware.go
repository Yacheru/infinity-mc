package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/yacheru/infinity-mc.ru/backend/configs"
	"log"
	"net/http"
)

func (h *Handler) userIdentity(c *gin.Context) {
	if err := configs.InitConfig(); err != nil {
		log.Fatalf("Error reading config.json file, %s", err.Error())
	}

	auth := c.GetHeader("Authorization")
	if auth == "" || auth != viper.GetString("api.pass") {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":        http.StatusUnauthorized,
			"description": "Unauthorized",
		})
		c.Abort()
	}
}
