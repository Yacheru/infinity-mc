package setups

import (
	"fmt"

	"github.com/gorcon/rcon"
	"github.com/sirupsen/logrus"

	"deliver-service/init/config"
	"deliver-service/init/logger"
	"deliver-service/pkg/util/constants"

	r "deliver-service/init/rcon"
)

func SetupRCONConnection() (*rcon.Conn, error) {
	rcfg := r.ConfigRCON{
		Address:  fmt.Sprintf("%s:%d", config.ServerConfig.RCONIp, config.ServerConfig.RCONPort),
		Password: config.ServerConfig.RCONPass,
	}

	conn, err := rcfg.InitRCON()
	if err != nil {
		logger.FatalF("error setup rcon: %s", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryRCON}, err)

		return nil, err
	}

	return conn, nil
}
