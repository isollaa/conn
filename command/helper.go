package command

import (
	"errors"
	"fmt"

	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var listInfo = map[string]string{
	"build": "build info of selected server",
	"host":  "host info of selected server",
}

var listAttribute = map[string]string{
	"db":   "list databases on selected driver",
	"coll": "list collection on selected database",
}

var listStatus = map[string]string{
	"db":   "status of selected database",
	"coll": "status of selected collection",
	"disk": "status of disk space",
}

var listStatusType = map[string]string{
	"db":   "status of selected database",
	"coll": "status of selected collection",
}

func validator(flg string, list map[string]string) {
	fmt.Printf("Error: flag with argument '%s' not found \n\nTry using:\n", flg)
	for k, v := range list {
		fmt.Printf("\t%s \t%s\n", k, v)
	}
	println()
}

func requirementCheck(cmd *cobra.Command) error {
	if config["driver"] == "" || flg.Stat == "" {
		err := fmt.Sprintf("Error: command needs flag with argument: -d -i\nUsage:\n\tapp %s [flags][flags]\n\nUse app --help to show help for status\n", cmd.Use)
		return errors.New(err)
	}
	return nil
}

func promptPassword() error {
	print("Input database password : ")
	passDb, err := terminal.ReadPassword(0)
	if err != nil {
		return err
	}
	config["password"] = string(passDb)
	println()
	return nil
}

func checkAutoFill() {
	if config["host"] == "" {
		config["host"] = "localhost"
	}
	switch config["driver"] {
	case "mongo":
		if config["port"] == "" {
			config["port"] = "27017"
		}
		if config["dbName"] == "" {
			config["dbName"] = "xsaas_ctms"
		}
	case "mysql":
		if config["port"] == "" {
			config["port"] = "3306"
		}
		if config["dbName"] == "" {
			config["dbName"] = "mqtt"
		}
		if config["username"] == "" {
			config["username"] = "root"
		}
	case "postgres":
		if config["port"] == "" {
			config["port"] = "5432"
		}
		if config["dbName"] == "" {
			config["dbName"] = "postgres"
		}
		if config["username"] == "" {
			config["username"] = "postgres"
		}
		if config["password"] == "" {
			config["password"] = "12345"
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

func doPrint(result interface{}) error {
	if flg.Pretty {
		if err := printPretty(result); err != nil {
			return err
		}
		return nil
	}

	fmt.Printf("%v\n", result)
	return nil
}

func getFlags(cmd *cobra.Command) {
	for key := range config {
		v, _ := cmd.Flags().GetString(key)
		config[key] = v
	}
}
