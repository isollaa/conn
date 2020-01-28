package command

import (
	"fmt"
	"log"

	cc "github.com/isollaa/conn/config"
	d "github.com/isollaa/conn/driver"
	m "github.com/isollaa/conn/driver/mongo"
	"github.com/isollaa/conn/driver/sql"
	"github.com/isollaa/conn/helper"
	"golang.org/x/crypto/ssh/terminal"
)

func (c Config) connect(svc d.CommonFeature) error {
	fmt.Printf("--%s\n", c[cc.DRIVER])
	if c[PROMPT].(bool) {
		if err := c.promptPassword(); err != nil {
			return err
		}
	}
	cfg := cc.Config{
		cc.DRIVER:     c[cc.DRIVER],
		cc.HOST:       c[cc.HOST],
		cc.PORT:       c[cc.PORT],
		cc.USERNAME:   c[cc.USERNAME],
		cc.PASSWORD:   c[cc.PASSWORD],
		cc.DBNAME:     c[cc.DBNAME],
		cc.COLLECTION: c[cc.COLLECTION],
	}
	svc.AutoFill(cfg)
	err := svc.Connect(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (c Config) DoCommand(commandFunc func(Config, d.CommonFeature) error) {
	driver := c[cc.DRIVER]
	if driver == "postgres" || driver == "mysql" {
		driver = "sql"
	}
	svc := d.New(driver.(string))
	if err := c.connect(svc); err != nil {
		log.Print("unable to connect: ", err)
		return
	}
	defer svc.Close()
	if err := commandFunc(c, svc); err != nil {
		log.Print("error due executing command: ", err)
		return
	}
	if err := c.DoPrint(svc); err != nil {
		log.Print("unable to print: ", err)
		return
	}
}

func (c Config) RequirementCheck(arg ...string) error {
	length := len(arg)
	for k, v := range arg {
		flag := requirementCase(v)
		msg := ""
		switch c[v] {
		case "", 0:
			msg = fmt.Sprintf("Command needs flag with argument: %s `%s`\n", flag, v)
		case false:
			msg = fmt.Sprintf("Command needs flag: %s\n", flag)
		}
		if k == length-1 && msg != "" {
			return fmt.Errorf(msg)
		}
		// log.Print("error: ", msg)
	}
	return nil
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

func (c Config) DoPrint(svc d.CommonFeature) error {
	var res interface{}
	if ss, ok := svc.(*m.Mongo); ok {
		res = ss.Result
	} else {
		ss := svc.(*sql.SQL)
		res = ss.Result
	}

	if c[BEAUTY].(bool) {
		if err := helper.PrintPretty(res); err != nil {
			return err
		}
		return nil
	}

	fmt.Printf("%v\n", res)
	return nil
}

func Validator(flg string, list map[string]string) error {
	fmt.Printf("Error: flag with argument '%s' not found \n\nTry using:\n", flg)
	for k, v := range list {
		fmt.Printf("\t%s \t%s\n", k, v)
	}
	println()
	return nil
}
