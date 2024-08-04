package rcon

import (
	"deliver-service/init/logger"
	"deliver-service/pkg/util/constants"
	"fmt"
	"github.com/gorcon/rcon"
	"github.com/sirupsen/logrus"
)

type Deliver interface {
	DeliverService(nickname, service, duration string) error
}

type RCON struct {
	rcon *rcon.Conn
}

func NewRCON(rcon *rcon.Conn) Deliver {
	return &RCON{
		rcon: rcon,
	}
}

func (r *RCON) DeliverService(nickname, service, duration string) error {
	command := fmt.Sprintf("lp user %s parent addtemp %s %smo", nickname, service, duration)
	log := fmt.Sprintf("service %s delivered for %s with duration %s month", service, nickname, duration)

	_, err := r.rcon.Execute(command)
	_, _ = r.rcon.Execute("say " + log)
	if err != nil {
		return err
	}

	logger.Info(log, logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryRCON})

	return nil
}
