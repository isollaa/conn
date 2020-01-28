package ping

import (
	"log"

	c "github.com/isollaa/conn/command"
	"github.com/isollaa/conn/config"
	"github.com/isollaa/conn/driver"
	"github.com/spf13/cobra"
)

func ping(cfg c.Config, svc driver.CommonFeature) error {
	log.Printf("Pinging %s ", cfg[config.HOST])
	err := svc.Ping()
	if err != nil {
		return err
	}
	return nil
}

func Command(cfg c.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "Check ping of selected connection",
		Run: func(cmd *cobra.Command, args []string) {
			cfg.SetFlag(cmd)
			if err := cfg.RequirementCheck(config.DRIVER); err != nil {
				log.Print("error: ", err)
				return
			}
			cfg.DoCommand(ping)
		},
	}
}
