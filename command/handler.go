package command

import (
	"fmt"
	"log"

	"github.com/isollaa/conn/status"
	"github.com/spf13/cobra"
)

func doCommand(cmd *cobra.Command, commandFunc func(status.CommonFeature) error) error {
	println(cmd.Short, "\n")
	svc := status.New(flg.Driver)
	err := connect(svc)
	if err != nil {
		return err
	}
	defer svc.Close()
	err = commandFunc(svc)
	if err != nil {
		return err
	}
	return nil
}

func connect(svc status.CommonFeature) error {
	println("--", flg.Driver)
	if flg.Prompt {
		err := promptPassword()
		if err != nil {
			return err
		}
	}
	checkAutoFill()
	// err := svc.Connect(&status.Config{
	// 	Host:     flg.Host,
	// 	Username: flg.Username,
	// 	Password: flg.Password,
	// 	Attribute: &status.Attribute{
	// 		DBName:     flg.DBName,
	// 		Collection: flg.Collection,
	// 	},
	// })
	err := svc.Connect(map[string]string{
		"host":       flg.Host,
		"username":   flg.Username,
		"password":   flg.Password,
		"dbName":     flg.DBName,
		"collection": flg.Collection,
	})
	if err != nil {
		return err
	}
	return nil
}

func ping(svc status.CommonFeature) error {
	log.Printf("Pinging %s ", flg.Host)
	err := svc.Ping()
	if err != nil {
		return err
	}
	return nil
}

func listDB(svc status.CommonFeature) error {
	err := svc.ListDB()
	if err != nil {
		return err
	}
	return nil
}

func listColl(svc status.CommonFeature) error {
	err := svc.ListColl()
	if err != nil {
		return err
	}
	return nil
}

func infoDB(svc status.CommonFeature) error {
	if nsvc, supported := svc.(status.NoSQLFeature); supported {
		str := fmt.Sprintf("%sInfo", flg.Stat)
		valid := false
		for k := range listInfo {
			if flg.Stat == k {
				err := nsvc.Info(str)
				if err != nil {
					return err
				}
				valid = true
				break
			}
		}
		if !valid {
			fmt.Printf("Error: flag with argument '%s' not found \n\nTry using:\n", flg.Stat)
			for k, v := range listInfo {
				fmt.Printf("\t%s \t%s\n", k, v)
			}
			println()
		}
	} else {
		fmt.Printf("--%s : info not available", flg.Driver)
	}
	return nil
}

func statusDB(svc status.CommonFeature) error {
	if nsvc, supported := svc.(status.NoSQLFeature); supported {
		var err error
		if flg.Stat != "" {
			switch flg.Stat {
			case "collstats":
				err = nsvc.CollStats()
			case "dbstats":
				err = nsvc.DbStats()
			default:
				fmt.Printf("Error: flag with argument '%s' not found \n\nTry using:\n", flg.Stat)
				for k, v := range listStatus {
					fmt.Printf("\t%s \t%s\n", k, v)
				}
				println()
			}
		}
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("--%s : status not available", flg.Driver)
	}
	return nil
}
