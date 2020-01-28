package status

import (
	"log"

	c "github.com/isollaa/conn/command"
	"github.com/isollaa/conn/config"
	"github.com/isollaa/conn/driver"
	"github.com/spf13/cobra"
)

const DISK = "disk"

var listStatus = map[string]string{
	DISK: "status of disk space",
}

var listStatusType = map[string]string{
	c.DB:   "status of selected database",
	c.COLL: "status of selected collection",
}

func statusDB(cfg c.Config, svc driver.CommonFeature) error {
	valid := false
	if cfg[c.STAT] == DISK {
		if err := cfg.RequirementCheck(c.TYPE); err != nil {
			return err
		}
		if cfg[config.DRIVER] == "postgres" {
			if err := cfg.RequirementCheck(c.PROMPT); err != nil {
				return err
			}
		}
		if cfg[c.TYPE] == c.COLL {
			if err := cfg.RequirementCheck(config.COLLECTION); err != nil {
				return err
			}
		}
		for k := range listStatusType {
			if cfg[c.TYPE] == k {
				err := svc.DiskSpace(k)
				if err != nil {
					return err
				}
				valid = true
				return nil
			}
		}
		if !valid {
			c.Validator(cfg[c.TYPE].(string), listStatusType)
		}
	} else {
		c.Validator(cfg[c.STAT].(string), listStatus)
	}
	return nil
}

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Get status of selected connection",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := c.DoConfig(cmd)
			if err := cfg.RequirementCheck(config.DRIVER, c.STAT); err != nil {
				log.Print("error: ", err)
				return
			}
			cfg.DoCommand(statusDB)
		},
	}
}
