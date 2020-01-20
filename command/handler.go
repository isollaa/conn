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
	if err := connect(svc); err != nil {
		return err
	}
	defer svc.Close()
	if err := commandFunc(svc); err != nil {
		return err
	}
	return nil
}

func connect(svc status.CommonFeature) error {
	println("--", flg.Driver)
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

func ping(svc status.CommonFeature) error {
	log.Printf("Pinging %s ", config["host"])
	if err := svc.Ping(); err != nil {
		return err
	}

	return nil
}

func listDB(svc status.CommonFeature) error {
	v, err := svc.ListDB()
	if err != nil {
		return err
	}
	if err = doPrint(v); err != nil {
		return err
	}

	return nil
}

func listColl(svc status.CommonFeature) error {
	v, err := svc.ListColl()
	if err != nil {
		return err
	}
	if err = doPrint(v); err != nil {
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
				v, err := nsvc.Info(str)
				if err != nil {
					return err
				}
				if err = doPrint(v); err != nil {
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
		if flg.Stat != "" {
			var (
				err error
				v   interface{}
			)
			switch flg.Stat {
			case "collstats":
				if v, err = nsvc.CollStats(); err != nil {
					return err
				}
			case "dbstats":
				if v, err = nsvc.DbStats(); err != nil {
					return err
				}
			default:
				fmt.Printf("Error: flag with argument '%s' not found \n\nTry using:\n", flg.Stat)
				for k, v := range listStatus {
					fmt.Printf("\t%s \t%s\n", k, v)
				}
				println()
			}
			if err = doPrint(v); err != nil {
				return err
			}
		}
	} else {
		fmt.Printf("--%s : status not available", flg.Driver)
	}
	return nil
}
