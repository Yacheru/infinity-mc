package rcon

import (
	"deliver-service/init/logger"
	"github.com/gorcon/rcon"
	"github.com/sirupsen/logrus"

	"deliver-service/pkg/util/constants"
)

type ConfigRCON struct {
	Address  string
	Password string
}

func (c *ConfigRCON) InitRCON() (*rcon.Conn, error) {
	connect, err := rcon.Dial(c.Address, c.Password)
	if err != nil {
		logger.FatalF("err rcon dial: %s", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryRCON}, err)

		return nil, constants.ErrDialRCON
	}

	return connect, nil
}
