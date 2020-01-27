package command

import (
	"fmt"
	"log"

	cc "github.com/isollaa/conn/config"
	s "github.com/isollaa/conn/status"
	"github.com/spf13/cobra"
)

const (
	SERVER = "server"
	BUILD  = "build"
	HOST   = "host"
	DB     = "db"
	COLL   = "coll"
	DISK   = "disk"
)

type Config cc.Config

func doConfig(cmd *cobra.Command) Config {
	c := Config{
		cc.DRIVER:     "",
		cc.HOST:       "",
		cc.PORT:       0,
		cc.USERNAME:   "",
		cc.PASSWORD:   "",
		cc.DBNAME:     "",
		cc.COLLECTION: "",
		cc.STAT:       "",
		cc.TYPE:       "",
		cc.BEAUTY:     false,
		cc.PROMPT:     false,
	}
	c.SetFlag(cmd)
	return c
}

func (c Config) SetFlag(cmd *cobra.Command) {
	for key := range c {
		if key == cc.PASSWORD {
			continue
		}
		if v, err := cmd.Flags().GetString(key); err == nil {
			c[key] = v
			continue
		}
		if v, err := cmd.Flags().GetInt(key); err == nil {
			c[key] = v
			continue
		}
		if v, err := cmd.Flags().GetFloat64(key); err == nil {
			c[key] = v
			continue
		}
		if v, err := cmd.Flags().GetBool(key); err == nil {
			c[key] = v
			continue
		}
		fmt.Printf("flag %s doesn't exist", key)
	}
}

func (c Config) doCommand(commandFunc func(Config, s.CommonFeature) error) {
	if err := c.requirementCheck(cc.DRIVER); err != nil {
		log.Print("error: ", err)
		return
	}
	driver := c[cc.DRIVER]
	if driver == "postgres" || driver == "mysql" {
		driver = "sql"
	}
	svc := s.New(driver.(string))
	if err := c.connect(svc); err != nil {
		log.Print("unable to connect: ", err)
		return
	}
	defer svc.Close()
	if err := commandFunc(c, svc); err != nil {
		log.Print("error due executing command: ", err)
		return
	}
	if err := doPrint(c, svc); err != nil {
		log.Print("unable to print: ", err)
		return
	}
}
