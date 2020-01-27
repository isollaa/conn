package command

import (
	"fmt"
	"log"

	cc "github.com/isollaa/conn/config"
	s "github.com/isollaa/conn/status"
)

func ping(c Config, svc s.CommonFeature) error {
	log.Printf("Pinging %s ", c[cc.HOST])
	err := svc.Ping()
	if err != nil {
		return err
	}
	return nil
}

func list(c Config, svc s.CommonFeature) error {
	if err := c.requirementCheck(cc.DBNAME); err != nil {
		return err
	}
	switch c[cc.STAT] {
	case DB:
		err := svc.ListDB()
		if err != nil {
			return err
		}
	case COLL:
		if err := c.requirementCheck(cc.COLLECTION); err != nil {
			return err
		}
		err := svc.ListColl()
		if err != nil {
			return err
		}
	default:
		validator(c[cc.STAT].(string), listAttribute)
	}

	return nil
}

func infoDB(c Config, svc s.CommonFeature) error {
	if err := c.requirementCheck(cc.DRIVER); err != nil {
		return err
	}
	if nsvc, supported := svc.(s.NoSQLFeature); supported {
		str := fmt.Sprintf("%sInfo", c[cc.STAT])
		valid := false
		for k := range listInfo {
			if c[cc.STAT] == k {
				err := nsvc.Info(str)
				if err != nil {
					return err
				}
				valid = true
				return nil
			}
		}
		if !valid {
			validator(c[cc.STAT].(string), listInfo)
		}
	} else {
		fmt.Printf("--%s : selected info not available\n", c[cc.DRIVER])
	}
	return nil
}

func statusDB(c Config, svc s.CommonFeature) error {
	if err := c.requirementCheck(cc.DRIVER); err != nil {
		return err
	}
	valid := false
	if c[cc.STAT] == DISK {
		if c[cc.DRIVER] == "postgres" {
			if err := c.requirementCheck(cc.PROMPT); err != nil {
				return err
			}
		}
		if c[cc.TYPE] == COLL {
			if err := c.requirementCheck(cc.COLLECTION); err != nil {
				return err
			}
		}
		for k := range listStatusType {
			if c[cc.TYPE] == k {
				err := svc.DiskSpace(k)
				if err != nil {
					return err
				}
				valid = true
				return nil
			}
		}
		if !valid {
			validator(c[cc.TYPE].(string), listStatusType)
		}
	} else {
		fmt.Printf("--%s : selected info not available\n", c[cc.DRIVER])
	}
	return nil
}
