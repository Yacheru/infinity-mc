package rcon

import (
	"deliver-service/init/logger"
	"deliver-service/pkg/constants"
	"fmt"
	"github.com/gorcon/rcon"
)

type Deliver interface {
	DeliverService(nickname, service, duration string) error
}

type RCON struct {
	*rcon.Conn
}

func NewRCON(conn *rcon.Conn) Deliver {
	return &RCON{
		conn,
	}
}

func (r *RCON) DeliverService(nickname, service, duration string) error {
	command := fmt.Sprintf("lp user %s parent addtemp %s %smo", nickname, service, duration)
	log := fmt.Sprintf("service %s delivered for %s with duration %s month", service, nickname, duration)

	_, err := r.Execute(command)
	_, _ = r.Execute("say " + log)
	if err != nil {
		return err
	}

	logger.Info(log, constants.LoggerCategoryRCON)

	return nil
}
