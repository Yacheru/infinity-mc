package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yacheru/infinity-mc.ru/backend/init/rcon"
	"log"
	"net/http"
)

func (h *Handler) Mc(c *gin.Context) {
	rconn := rcon.InitRCON()
	defer rconn.Close()

	response, err := rconn.Execute("op yacheru")
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, map[string]string{
		"response": response,
	})
}
