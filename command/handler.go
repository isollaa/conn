package command

import (
	"fmt"
	"log"

	s "github.com/isollaa/conn/status"
)

func doCommand(commandFunc func(s.CommonFeature) error) {
	if err := requirementCheck(s.DRIVER); err != nil {
		log.Print("error: ", err)
		return
	}
	driver := config[s.DRIVER]
	if driver == "postgres" || driver == "mysql" {
		driver = "sql"
	}
	svc := s.New(driver.(string))
	if err := connect(svc); err != nil {
		log.Print("unable to connect: ", err)
		return
	}
	defer svc.Close()
	if err := commandFunc(svc); err != nil {
		log.Print("error due executing command: ", err)
		return
	}
	if err := doPrint(svc); err != nil {
		log.Print("unable to print: ", err)
		return
	}
}

func connect(svc s.CommonFeature) error {
	fmt.Printf("--%s\n", config[s.DRIVER])
	if flg.prompt {
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

func ping(svc s.CommonFeature) error {
	log.Printf("Pinging %s ", config[s.HOST])
	err := svc.Ping()
	if err != nil {
		return err
	}
	return nil
}

func list(svc s.CommonFeature) error {
	if err := requirementCheck(s.DBNAME); err != nil {
		return err
	}
	switch flg.stat {
	case DB:
		err := svc.ListDB()
		if err != nil {
			return err
		}
	case COLL:
		if err := requirementCheck(s.COLLECTION); err != nil {
			return err
		}
		err := svc.ListColl()
		if err != nil {
			return err
		}
	default:
		validator(flg.stat, listAttribute)
	}

	return nil
}

func infoDB(svc s.CommonFeature) error {
	if err := requirementCheck(s.DRIVER); err != nil {
		return err
	}
	var err error
	if nsvc, supported := svc.(s.NoSQLFeature); supported {
		str := fmt.Sprintf("%sInfo", flg.stat)
		valid := false
		for k := range listInfo {
			if flg.stat == k {
				err = nsvc.Info(str)
				if err != nil {
					return err
				}
				valid = true
				return nil
			}
		}
		if !valid {
			validator(flg.stat, listInfo)
		}
	} else {
		fmt.Printf("--%s : selected info not available\n", config[s.DRIVER])
	}
	return nil
}

func statusDB(svc s.CommonFeature) error {
	if err := requirementCheck(s.DRIVER); err != nil {
		return err
	}
	var err error
	valid := false
	if flg.stat == DISK {
		if flg.statType == COLL {
			if err := requirementCheck(s.COLLECTION); err != nil {
				return err
			}
		}
		for k := range listStatusType {
			if flg.statType == k {
				err = svc.DiskSpace(k)
				if err != nil {
					return err
				}
				valid = true
				return nil
			}
		}
		if !valid {
			validator(flg.statType, listStatusType)
		}
	} else {
		fmt.Printf("--%s : selected info not available\n", config[s.DRIVER])
	}
	return nil
}
