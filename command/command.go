package command

import (
	"log"

	cc "github.com/isollaa/conn/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

var cmdPing = cobra.Command{
	Use:   "ping",
	Short: "Check ping of selected connection",
	Run: func(cmd *cobra.Command, args []string) {
		c := doConfig(cmd)
		c.doCommand(ping)
	},
}

var cmdDB = cobra.Command{
	Use:   "list",
	Short: "list available database attributes",
	Run: func(cmd *cobra.Command, args []string) {
		c := doConfig(cmd)
		if err := c.requirementCheck(cc.STAT); err != nil {
			log.Print("error: ", err)
			return
		}
		c.doCommand(list)
	},
}

var cmdInfo = cobra.Command{
	Use:   "info",
	Short: "Get information of selected connection",
	Run: func(cmd *cobra.Command, args []string) {
		c := doConfig(cmd)
		if err := c.requirementCheck(cc.STAT); err != nil {
			log.Print("error: ", err)
			return
		}
		c.doCommand(infoDB)
	},
}

var cmdStatus = cobra.Command{
	Use:   "status",
	Short: "Get status of selected connection",
	Run: func(cmd *cobra.Command, args []string) {
		c := doConfig(cmd)
		if err := c.requirementCheck(cc.STAT); err != nil {
			log.Print("error: ", err)
			return
		}
		c.doCommand(statusDB)
	},
}

func init() {
	//global
	rootCmd.PersistentFlags().StringP("driver", "d", "", "connection driver name (mongo / mysql / postgres)")
	rootCmd.PersistentFlags().StringP("host", "H", "localhost", "connection host (default- localhost)")
	rootCmd.PersistentFlags().IntP("port", "P", 0, "connection port (default- mongo:27017 / mysql:3306 / postgres:5432)")
	rootCmd.PersistentFlags().StringP("username", "u", "", "database username (default- mysql: root / postgres: postgres)")
	rootCmd.PersistentFlags().String("dbname", "", "connection database name (default- mongo:xsaas_ctms / mysql:mqtt)")
	rootCmd.PersistentFlags().StringP("collection", "c", "", "connection database collection name")
	rootCmd.PersistentFlags().StringP("stat", "s", "", "connection information")
	rootCmd.PersistentFlags().StringP("type", "t", "", "connection information type")

	//optional
	rootCmd.PersistentFlags().BoolP("beauty", "b", false, "show pretty version of json")
	rootCmd.PersistentFlags().BoolP("prompt", "p", false, "call password prompt")
}

func Exec() {
	rootCmd.AddCommand(&cmdPing, &cmdInfo, &cmdStatus, &cmdDB)
	rootCmd.Execute()
}
