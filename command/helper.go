package command

import (
	"errors"
	"fmt"

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
	flg.Password = string(passDb)
	println()
	return nil
}

func checkAutoFill() {
	switch flg.Driver {
	case "mongo":
		if flg.Host == "" {
			flg.Host = "localhost:27017"
		}
		if flg.DBName == "" {
			flg.DBName = "xsaas_ctms"
		}
	case "mysql":
		if flg.Host == "" {
			flg.Host = "localhost:3306"
		}
		if flg.DBName == "" {
			flg.DBName = "mqtt"
		}
		if flg.Username == "" {
			flg.Username = "root"
		}
	}
}
