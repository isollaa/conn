package info

import (
	"fmt"
	"log"

	c "github.com/isollaa/conn/command"
	"github.com/isollaa/conn/config"
	"github.com/isollaa/conn/driver"
	"github.com/spf13/cobra"
)

const (
	SERVER = "server"
	BUILD  = "build"
)

var listInfo = map[string]string{
	SERVER:      "server info of selected driver",
	BUILD:       "build info of selected driver",
	config.HOST: "host info of selected driver",
}

func infoDB(cfg c.Config, svc driver.CommonFeature) error {
	if nsvc, supported := svc.(driver.NoSQLFeature); supported {
		str := fmt.Sprintf("%sInfo", cfg[c.STAT])
		valid := false
		for k := range listInfo {
			if cfg[c.STAT] == k {
				err := nsvc.Info(str)
				if err != nil {
					return err
				}
				valid = true
				return nil
			}
		}
		if !valid {
			c.Validator(cfg[c.STAT].(string), listInfo)
		}
	} else {
		fmt.Printf("--%s : selected info not available\n", cfg[config.DRIVER])
	}
	return nil
}

func Command(cfg c.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Get information of selected connection",
		Run: func(cmd *cobra.Command, args []string) {
			cfg.SetFlag(cmd)
			if err := cfg.RequirementCheck(config.DRIVER, c.STAT); err != nil {
				log.Print("error: ", err)
				return
			}
			cfg.DoCommand(infoDB)
		},
	}
}
