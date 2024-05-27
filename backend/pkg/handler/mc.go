package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yacheru/infinity-mc.ru/backend/init/rcon"
	"log"
	"net/http"
)

func (h *Handler) Mc(c *gin.Context) {
	rcon := rcon.InitRCON()
	defer rcon.Close()

	response, err := rcon.Execute("op yacheru")
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, map[string]string{
		"response": response,
	})
}
