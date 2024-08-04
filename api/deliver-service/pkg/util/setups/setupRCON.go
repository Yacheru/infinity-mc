package setups

import (
	"fmt"

	"deliver-service/init/config"
	r "deliver-service/init/rcon"
	"github.com/gorcon/rcon"
)

func SetupRCONConnection() (*rcon.Conn, error) {
	rconConfig := r.ConfigRCON{
		Address:  fmt.Sprintf("%s:%d", config.ServerConfig.RCONIp, config.ServerConfig.RCONPort),
		Password: config.ServerConfig.RCONPass,
	}

	conn, err := rconConfig.InitRCON()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
