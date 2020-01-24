package command

import (
	"errors"
	"fmt"

	"github.com/hokaccha/go-prettyjson"
	s "github.com/isollaa/conn/status"
	m "github.com/isollaa/conn/status/mongo"
	"github.com/isollaa/conn/status/sql"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var listInfo = map[string]string{
	BUILD: "build info of selected server",
	HOST:  "host info of selected server",
}

var listAttribute = map[string]string{
	DB:   "list databases on selected driver",
	COLL: "list collection on selected database",
}

var listStatus = map[string]string{
	DISK: "status of disk space (postgres only)",
}

var listStatusType = map[string]string{
	DB:   "status of selected database",
	COLL: "status of selected collection",
}

func validator(flg string, list map[string]string) {
	fmt.Printf("Error: flag with argument '%s' not found \n\nTry using:\n", flg)
	for k, v := range list {
		fmt.Printf("\t%s \t%s\n", k, v)
	}
	println()
}

func requirementCheck(arg ...string) error {
	length := len(arg)
	for k, v := range arg {
		err := configCheck(k, length, v)
		if err != nil {
			return err
		}
		err = flagCheck(k, length, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func flagCheck(key, length int, value string) error {
	flag := ""
	switch value {
	case STAT:
		if flg.stat == "" {
			flag = "-i"
		}
	case STATTYPE:
		if flg.statType == "" {
			flag = "-t"
		}
	case PRETTY:
		if !flg.pretty {
			flag = "-b"
		}
	case PROMPT:
		if !flg.prompt {
			flag = "-p"
		}
	}
	if flag != "" {
		if key == length-1 {
			return fmt.Errorf("Command needs flag with argument: %s `%s`", flag, value)
		}
		fmt.Printf("Command needs flag with argument: %s `%s`", flag, value)
	}
	return nil
}

func configCheck(key, length int, value string) error {
	flag := ""
	if config[value] == "" {
		switch value {
		case s.DRIVER:
			flag = "-d"
		case s.HOST:
			flag = "-H"
		case s.PORT:
			flag = "-P"
		case s.USERNAME:
			flag = "-u"
		case s.PASSWORD:
			flag = "-p"
		case s.DBNAME:
			flag = "--dbname"
		case s.COLLECTION:
			flag = "-c"
		}
		if key == length-1 {
			return fmt.Errorf("Command needs flag with argument: %s `%s`", flag, value)
		}
		fmt.Printf("Command needs flag with argument: %s `%s`", flag, value)
	}
	return nil
}

func promptPassword() error {
	print("Input database password : ")
	passDb, err := terminal.ReadPassword(0)
	if err != nil {
		return err
	}
	config[s.PASSWORD] = string(passDb)
	println()
	return nil
}

func checkAutoFill() {
	switch config[s.DRIVER] {
	case "mongo":
		if config[s.PORT] == 0 {
			config[s.PORT] = 27017
		}
		if config[s.DBNAME] == "" {
			config[s.DBNAME] = "xsaas_ctms"
		}
	case "mysql":
		if config[s.PORT] == 0 {
			config[s.PORT] = 3306
		}
		if config[s.DBNAME] == "" {
			config[s.DBNAME] = "mqtt"
		}
		if config[s.USERNAME] == "" {
			config[s.USERNAME] = "root"
		}
	case "postgres":
		if config[s.PORT] == 0 {
			config[s.PORT] = 5432
		}
		if config[s.DBNAME] == "" {
			config[s.DBNAME] = "postgres"
		}
		if config[s.USERNAME] == "" {
			config[s.USERNAME] = "postgres"
		}
		if config[s.PASSWORD] == "" {
			config[s.PASSWORD] = "12345"
		}
	}
}

func printPretty(result interface{}) error {
	v, err := prettyjson.Marshal(result)
	if err != nil {
		return err
	}
	fmt.Println(string(v))
	return nil
}

func doPrint(svc s.CommonFeature) error {
	var res interface{}
	if ss, ok := svc.(*m.Mongo); ok {
		res = ss.Result
	} else {
		ss := svc.(*sql.SQL)
		res = ss.Result
	}

	if flg.pretty {
		if err := printPretty(res); err != nil {
			return err
		}
		return nil
	}

	fmt.Printf("%v\n", res)
	return nil
}

func getFlags(cmd *cobra.Command) error {
	for key := range config {
		if key == s.PASSWORD {
			continue
		}
		v, err := cmd.Flags().GetString(key)
		if err != nil {
			v, err := cmd.Flags().GetInt(key)
			if err != nil {
				v, err := cmd.Flags().GetFloat64(key)
				if err != nil {
					v, err := cmd.Flags().GetBool(key)
					if err != nil {
						return errors.New(fmt.Sprintf("flag %s doesn't exist", key))
					}
					config[key] = v
					continue
				}
				config[key] = v
				continue
			}
			config[key] = v
			continue
		}
		config[key] = v
		continue
	}
	return nil
}
