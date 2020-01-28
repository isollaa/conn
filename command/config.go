package command

import (
	"fmt"

	cc "github.com/isollaa/conn/config"
	"github.com/spf13/cobra"
)

const (
	DB   = "db"
	COLL = "coll"

	STAT   = "stat"
	TYPE   = "type"
	BEAUTY = "beauty"
	PROMPT = "prompt"
)

type Config cc.Config

func SetConfig(cmd *cobra.Command) Config {
	c := Config{
		cc.DRIVER:     "",
		cc.HOST:       "",
		cc.PORT:       0,
		cc.USERNAME:   "",
		cc.PASSWORD:   "",
		cc.DBNAME:     "",
		cc.COLLECTION: "",
		STAT:          "",
		TYPE:          "",
		BEAUTY:        false,
		PROMPT:        false,
	}
	return c
}

func requirementCase(v string) string {
	flag := ""
	switch v {
	case cc.DRIVER:
		flag = "-d"
	case cc.HOST:
		flag = "-H"
	case cc.PORT:
		flag = "-P"
	case cc.USERNAME:
		flag = "-u"
	case cc.PASSWORD:
		flag = "-p"
	case cc.DBNAME:
		flag = "--dbname"
	case cc.COLLECTION:
		flag = "-c"
	case STAT:
		flag = "-s"
	case TYPE:
		flag = "-t"
	case BEAUTY:
		flag = "-b"
	case PROMPT:
		flag = "-p"
	}
	return flag
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
		fmt.Printf("flag %s doesn't exist\n", key)
	}
}
