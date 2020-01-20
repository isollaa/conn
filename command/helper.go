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
	"dbstats":   "status of selected database",
	"collstats": "status of selected collection",
}

func requirementCheck(cmd *cobra.Command) error {
	if flg.Driver == "" || flg.Stat == "" {
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
	switch flg.Driver {
	case "mongo":
		if config["host"] == "" {
			config["host"] = "localhost:27017"
		}
		if config["dbName"] == "" {
			config["dbName"] = "xsaas_ctms"
		}
	case "mysql":
		if config["host"] == "" {
			config["host"] = "localhost:3306"
		}
		if config["dbName"] == "" {
			config["dbName"] = "mqtt"
		}
		if config["username"] == "" {
			config["username"] = "root"
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

func getFlags(cmd *cobra.Command) error {
	for key := range config {
		v, err := cmd.Flags().GetString(key)
		if err != nil {
			return err
		}
		config[key] = v
	}
	return nil
}
