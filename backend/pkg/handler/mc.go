package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorcon/rcon"
	"github.com/spf13/viper"
	"github.com/yacheru/infinity-mc.ru/backend/configs"
	"log"
	"net/http"
)

func InitRCON() *rcon.Conn {
	if err := configs.InitConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err.Error())
	}

	ip := viper.GetString("rcon.ip")
	port := viper.GetString("rcon.port")
	ipport := fmt.Sprintf("%s:%s", ip, port)

	conn, err := rcon.Dial(ipport, viper.GetString("rcon.pass"))
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func (h *Handler) Mc(c *gin.Context) {
	rcon := InitRCON()
	defer rcon.Close()

	response, err := rcon.Execute("op yacheru")
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, response)
}
