package register

import (
	"github.com/isollaa/conn/command"
	"github.com/isollaa/conn/command/info"
	"github.com/isollaa/conn/command/list"
	"github.com/isollaa/conn/command/ping"
	"github.com/isollaa/conn/command/status"
	"github.com/isollaa/conn/driver"
	"github.com/isollaa/conn/driver/mongo"
	"github.com/isollaa/conn/driver/sql"
)

func initCommand() {
	command.Register(ping.Command)
	command.Register(list.Command)
	command.Register(info.Command)
	command.Register(status.Command)
}

func initDriver() {
	driver.Register(mongo.New)
	driver.Register(sql.New)
}

func init() {
	initDriver()
	initCommand()
}
