package rcon

import (
	"fmt"
	"github.com/gorcon/rcon"
	"github.com/spf13/viper"
	"log"
)

func InitRCON() *rcon.Conn {
	ip := viper.GetString("rcon.ip")
	port := viper.GetString("rcon.port")
	ipport := fmt.Sprintf("%s:%s", ip, port)

	connect, err := rcon.Dial(ipport, viper.GetString("rcon.pass"))
	if err != nil {
		log.Fatal(err)
	}

	return connect
}
