package rcon

import (
	"github.com/gorcon/rcon"
)

type ConfigRCON struct {
	Address  string
	Password string
}

func (c *ConfigRCON) InitRCON() (*rcon.Conn, error) {
	connect, err := rcon.Dial(c.Address, c.Password)
	if err != nil {
		return nil, err
	}

	return connect, nil
}
