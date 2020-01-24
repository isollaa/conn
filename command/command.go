package command

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

var cmdPing = cobra.Command{
	Use:   "ping",
	Short: "Check ping of selected connection",
	Run: func(cmd *cobra.Command, args []string) {
		err := getFlags(cmd)
		if err != nil {
			log.Print("unable to get flag: ", err)
			return
		}
		doCommand(ping)
	},
}

var cmdDB = cobra.Command{
	Use:   "list",
	Short: "list available database attributes",
	Run: func(cmd *cobra.Command, args []string) {
		err := getFlags(cmd)
		if err != nil {
			log.Print("unable to get flag: ", err)
			return
		}
		if err := requirementCheck(STAT); err != nil {
			log.Print("error: ", err)
			return
		}
		doCommand(list)
	},
}

var cmdInfo = cobra.Command{
	Use:   "info",
	Short: "Get information of selected connection",
	Run: func(cmd *cobra.Command, args []string) {
		err := getFlags(cmd)
		if err != nil {
			log.Print("unable to get flag: ", err)
			return
		}
		if err := requirementCheck(STAT); err != nil {
			log.Print("error: ", err)
			return
		}
		doCommand(infoDB)
	},
}

var cmdStatus = cobra.Command{
	Use:   "status",
	Short: "Get status of selected connection",
	Run: func(cmd *cobra.Command, args []string) {
		err := getFlags(cmd)
		if err != nil {
			log.Print("unable to get flag: ", err)
			return
		}
		if err := requirementCheck(STAT); err != nil {
			log.Print("error: ", err)
			return
		}
		doCommand(statusDB)
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
	rootCmd.PersistentFlags().StringVarP(&flg.stat, "info", "i", "", "connection information")
	rootCmd.PersistentFlags().StringVarP(&flg.statType, "type", "t", "", "connection information type")

	//optional
	rootCmd.PersistentFlags().BoolVarP(&flg.pretty, "beautify", "b", false, "show pretty version of json")
	rootCmd.PersistentFlags().BoolVarP(&flg.prompt, "password", "p", false, "call password prompt")
}

func Exec() {
	rootCmd.AddCommand(&cmdPing, &cmdInfo, &cmdStatus, &cmdDB)
	rootCmd.Execute()
}
