package rcon

import (
	"fmt"
	"github.com/gorcon/rcon"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitRCON() *rcon.Conn {
	ip := viper.GetString("rcon.ip")
	port := viper.GetString("rcon.port")
	ipport := fmt.Sprintf("%s:%s", ip, port)

	connect, err := rcon.Dial(ipport, viper.GetString("rcon.pass"))
	if err != nil {
		logrus.Fatalf("Error creating rcon connection, %s", err.Error())
	}

	return connect
}
