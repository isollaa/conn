package command

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

var cmdPing = cobra.Command{
	Use:   "ping",
	Short: "Check ping of selected connection",
	Run: func(cmd *cobra.Command, args []string) {
		if flg.Driver == "" {
			fmt.Printf("Error: command needs flag with argument: -d \nUsage:\n\tapp %s [flags]n\nUse app --help to show help for status\n\n", cmd.Use)
			return
		}
		err := doCommand(cmd, ping)
		if err != nil {
			log.Print(err)
		}
	},
}

var cmdDB = cobra.Command{
	Use:   "list",
	Short: "list available database attributes",
	Run: func(cmd *cobra.Command, args []string) {
		err := requirementCheck(cmd)
		if err != nil {
			fmt.Print(err)
			return
		}
		switch flg.Stat {
		case "db":
			err := doCommand(cmd, listDB)
			if err != nil {
				log.Print(err)
			}
		case "coll":
			if flg.DBName == "" {
				fmt.Printf("Error: command needs flag with argument: -d -a --db\nUsage:\n\tapp %s [flags][flags]\n\nUse app --help to show help for status\n\n", cmd.Use)
				return
			}
			err := doCommand(cmd, listColl)
			if err != nil {
				log.Print(err)
			}
		default:
			fmt.Printf("Error: flag with argument '%s' not found \n\nTry using:\n", flg.Stat)
			for k, v := range listAttribute {
				fmt.Printf("\t%s \t%s\n", k, v)
			}
			println()
		}
	},
}

var cmdInfo = cobra.Command{
	Use:   "info",
	Short: "Get information of selected connection",
	Run: func(cmd *cobra.Command, args []string) {
		err := requirementCheck(cmd)
		if err != nil {
			fmt.Print(err)
			return
		}
		err = doCommand(cmd, infoDB)
		if err != nil {
			log.Print(err)
		}
	},
}

var cmdStatus = cobra.Command{
	Use:   "status",
	Short: "Get status of selected connection",
	Run: func(cmd *cobra.Command, args []string) {
		err := requirementCheck(cmd)
		if err != nil {
			fmt.Print(err)
			return
		}
		err = doCommand(cmd, statusDB)
		if err != nil {
			log.Print(err)
		}
	},
}

type attrib struct {
	Driver     string
	Host       string
	Username   string
	Password   string
	DBName     string
	Collection string
	Stat       string
	Prompt     bool
}

var flg attrib

func init() {
	//global
	rootCmd.PersistentFlags().StringVarP(&flg.Driver, "driver", "d", "", "connection driver name (mongo / mysql)")                              // mongo || mysql
	rootCmd.PersistentFlags().StringVarP(&flg.Host, "host", "H", "", "connection host (default- mongo:localhost:27017 / mysql:localhost:3306)") // localhost:27017 || localhost:3306
	rootCmd.PersistentFlags().StringVar(&flg.DBName, "db", "", "connection database name (default- mongo:xsaas_ctms / mysql:mqtt)")             // xsaas_ctms || mqtt
	rootCmd.PersistentFlags().StringVarP(&flg.Collection, "collection", "c", "", "connection database collection name")
	rootCmd.PersistentFlags().StringVarP(&flg.Stat, "info", "i", "", "connection information")
	rootCmd.PersistentFlags().StringVarP(&flg.Username, "username", "u", "", "database username (default- mysql: root)")
	rootCmd.PersistentFlags().BoolVarP(&flg.Prompt, "password", "p", false, "call password prompt")
	//local
	// cmdInfo.PersistentFlags().StringVarP(opt["stat"], "info", "i", "", "connection information")
	// cmdStatus.PersistentFlags().StringVarP(opt["stat"], "status", "s", "", "connection status")
	// // cmdStatus.MarkFlagRequired("db")
	// // cmdStatus.MarkFlagRequired("status")
	// cmdDB.PersistentFlags().StringVarP(opt["stat"], "attribute", "a", "", "attribute db to show")
}

func Exec() {
	rootCmd.AddCommand(&cmdPing, &cmdInfo, &cmdStatus, &cmdDB)
	rootCmd.Execute()
}
