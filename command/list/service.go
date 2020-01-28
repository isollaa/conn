package list

import (
	"log"

	c "github.com/isollaa/conn/command"
	"github.com/isollaa/conn/config"
	"github.com/isollaa/conn/driver"
	"github.com/spf13/cobra"
)

var listAttributes = map[string]string{
	c.DB:   "list databases on selected driver",
	c.COLL: "list collection on selected database",
}

func list(cfg c.Config, svc driver.CommonFeature) error {
	switch cfg[c.STAT] {
	case c.DB:
		err := svc.ListDB()
		if err != nil {
			return err
		}
	case c.COLL:
		if err := cfg.RequirementCheck(config.COLLECTION); err != nil {
			return err
		}
		err := svc.ListColl()
		if err != nil {
			return err
		}
	default:
		c.Validator(cfg[c.STAT].(string), listAttributes)
	}

	return nil
}

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list available database attributes",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := c.DoConfig(cmd)
			if err := cfg.RequirementCheck(config.DRIVER, config.DBNAME, c.STAT); err != nil {
				log.Print("error: ", err)
				return
			}
			cfg.DoCommand(list)
		},
	}
}
