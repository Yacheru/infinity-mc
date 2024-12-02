package utils

import (
	"fmt"

	"deliver-service/init/config"
	r "deliver-service/init/rcon"
	"github.com/gorcon/rcon"
)

func SetupRCONConnection() (*rcon.Conn, error) {
	RCONConfig := r.ConfigRCON{
		Address:  fmt.Sprintf("%s:%d", config.ServerConfig.RCONIp, config.ServerConfig.RCONPort),
		Password: config.ServerConfig.RCONPass,
	}

	conn, err := RCONConfig.InitRCON()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
