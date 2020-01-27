package command

import (
	"fmt"

	cc "github.com/isollaa/conn/config"
	s "github.com/isollaa/conn/status"
	"golang.org/x/crypto/ssh/terminal"
)

func (c Config) connect(svc s.CommonFeature) error {
	fmt.Printf("--%s\n", c[cc.DRIVER])
	if c[cc.PROMPT].(bool) {
		if err := c.promptPassword(); err != nil {
			return err
		}
	}
	c.checkAutoFill()
	err := svc.Connect(
		cc.Config{
			cc.DRIVER:     c[cc.DRIVER],
			cc.HOST:       c[cc.HOST],
			cc.PORT:       c[cc.PORT],
			cc.USERNAME:   c[cc.USERNAME],
			cc.PASSWORD:   c[cc.PASSWORD],
			cc.DBNAME:     c[cc.DBNAME],
			cc.COLLECTION: c[cc.COLLECTION],
		})
	if err != nil {
		return err
	}
	return nil
}

func (c Config) requirementCheck(arg ...string) error {
	length := len(arg)
	for k, v := range arg {
		flag := ""
		if c[v] == "" || c[v] == false {
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
			case cc.STAT:
				flag = "-s"
			case cc.TYPE:
				flag = "-t"
			case cc.BEAUTY:
				flag = "-b"
			case cc.PROMPT:
				flag = "-p"
			}
			if k == length-1 {
				return fmt.Errorf("Command needs flag with argument: %s `%s`", flag, v)
			}
			fmt.Printf("Command needs flag with argument: %s `%s`", flag, v)
		}
	}

	return nil
}

func (c Config) checkAutoFill() {
	switch c[cc.DRIVER] {
	case "mongo":
		if c[cc.PORT] == 0 {
			c[cc.PORT] = 27017
		}
		if c[cc.DBNAME] == "" {
			c[cc.DBNAME] = "xsaas_ctms"
		}
	case "mysql":
		if c[cc.PORT] == 0 {
			c[cc.PORT] = 3306
		}
		if c[cc.DBNAME] == "" {
			c[cc.DBNAME] = "mqtt"
		}
		if c[cc.USERNAME] == "" {
			c[cc.USERNAME] = "root"
		}
	case "postgres":
		if c[cc.PORT] == 0 {
			c[cc.PORT] = 5432
		}
		if c[cc.DBNAME] == "" {
			c[cc.DBNAME] = "postgres"
		}
		if c[cc.USERNAME] == "" {
			c[cc.USERNAME] = "postgres"
		}
		if c[cc.PASSWORD] == "" {
			c[cc.PASSWORD] = "12345"
		}
	}
}

func (c Config) promptPassword() error {
	print("Input database password : ")
	passDb, err := terminal.ReadPassword(0)
	if err != nil {
		return err
	}
	c[cc.PASSWORD] = string(passDb)
	println()
	return nil
}
