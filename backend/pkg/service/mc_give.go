package service

import (
	"fmt"
	"github.com/yacheru/infinity-mc.ru/backend/init/rcon"
)

func GiveDonat(donatType, nickname, duration string) error {
	rcon := rcon.InitRCON()
	defer rcon.Close()

	command := fmt.Sprintf("lp user %s parent addtemp %s %smo", nickname, donatType, duration)

	_, err := rcon.Execute(command)

	if err != nil {
		return err
	}
	return nil
}
