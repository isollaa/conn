package command

import (
	"fmt"
	"log"

	"github.com/isollaa/conn/status"
	"github.com/spf13/cobra"
)

func doCommand(cmd *cobra.Command, commandFunc func(status.CommonFeature) (interface{}, error)) error {
	println(cmd.Short, "\n")
	driver := config["driver"]
	if driver == "postgres" || driver == "mysql" {
		driver = "sql"
	}
	svc := status.New(driver)
	if err := connect(svc); err != nil {
		return err
	}
	defer svc.Close()
	v, err := commandFunc(svc)
	if err != nil || v == nil {
		return err
	}
	if err = doPrint(v); err != nil {
		return err
	}
	return nil
}

func connect(svc status.CommonFeature) error {
	println("--", config["driver"])
	if flg.Prompt {
		if err := promptPassword(); err != nil {
			return err
		}
	}
	checkAutoFill()
	if err := svc.Connect(config); err != nil {
		return err
	}
	return nil
}

func ping(svc status.CommonFeature) (interface{}, error) {
	log.Printf("Pinging %s ", config["host"])
	v, err := svc.Ping()
	if err != nil {
		return v, err
	}

	return v, nil
}

func listDB(svc status.CommonFeature) (interface{}, error) {
	v, err := svc.ListDB()
	if err != nil {
		return v, err
	}

	return v, nil
}

func listColl(svc status.CommonFeature) (interface{}, error) {
	v, err := svc.ListColl()
	if err != nil {
		return v, err
	}

	return v, nil
}

func infoDB(svc status.CommonFeature) (interface{}, error) {
	var v interface{}
	var err error
	if nsvc, supported := svc.(status.NoSQLFeature); supported {
		str := fmt.Sprintf("%sInfo", flg.Stat)
		valid := false
		for k := range listInfo {
			if flg.Stat == k {
				v, err = nsvc.Info(str)
				if err != nil {
					return v, err
				}
				valid = true
				break
			}
		}
		if !valid {
			validator(flg.Stat, listInfo)
		}
	} else {
		fmt.Printf("--%s : selected info not available\n", config["driver"])
	}
	return v, nil
}

func statusDB(svc status.CommonFeature) (interface{}, error) {
	var v interface{}
	var err error
	nsvc, noSQL := svc.(status.NoSQLFeature)
	switch flg.Stat {
	case "coll":
		if noSQL {
			v, err = nsvc.CollStats()
			if err != nil {
				return v, err
			}
		}
	case "db":
		if noSQL {
			v, err = nsvc.DbStats()
			if err != nil {
				return v, err
			}
		}
	case "disk":
		if ssvc, sql := svc.(status.SQLFeature); sql {
			if flg.StatType == "" {
				flg.StatType = "db"
			}
			valid := false
			for k := range listStatusType {
				if flg.StatType == k {
					v, err = ssvc.DiskSpace(flg.StatType)
					if err != nil {
						return v, err
					}
					valid = true
					break
				}
			}
			if !valid {
				validator(flg.StatType, listStatusType)
			}
		}
	default:
		validator(flg.Stat, listStatus)
	}
	if v == nil {
		fmt.Printf("--%s : selected status not available\n", config["driver"])
		return v, nil
	}
	return v, nil
}
