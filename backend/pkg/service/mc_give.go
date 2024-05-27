package service

import (
	"fmt"
	"github.com/yacheru/infinity-mc.ru/backend/init/rcon"
)

func GiveHronon(nickname, latency string) error {
	rcon := rcon.InitRCON()
	defer rcon.Close()

	command := fmt.Sprintf("lp user %s parent addtemp Hronon %smo", nickname, latency)

	_, err := rcon.Execute(command)

	if err != nil {
		return err
	}
	return nil
}
