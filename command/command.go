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
		if err := getFlags(cmd); err != nil {
			fmt.Print(err)
			return
		}
		if flg.Driver == "" {
			fmt.Printf("Error: command needs flag with argument: -d \nUsage:\n\tapp %s [flags]n\nUse app --help to show help for status\n\n", cmd.Use)
			return
		}
		if err := doCommand(cmd, ping); err != nil {
			log.Print(err)
		}
	},
}

var cmdDB = cobra.Command{
	Use:   "list",
	Short: "list available database attributes",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getFlags(cmd); err != nil {
			fmt.Print(err)
			return
		}
		if err := requirementCheck(cmd); err != nil {
			fmt.Print(err)
			return
		}
		switch flg.Stat {
		case "db":
			if err := doCommand(cmd, listDB); err != nil {
				log.Print(err)
			}
		case "coll":
			if config["dbName"] == "" {
				fmt.Printf("Error: command needs flag with argument: -d -a --db\nUsage:\n\tapp %s [flags][flags]\n\nUse app --help to show help for status\n\n", cmd.Use)
				return
			}
			if err := doCommand(cmd, listColl); err != nil {
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
		if err := getFlags(cmd); err != nil {
			fmt.Print(err)
			return
		}
		if err := requirementCheck(cmd); err != nil {
			fmt.Print(err)
			return
		}
		if err := doCommand(cmd, infoDB); err != nil {
			log.Print(err)
		}
	},
}

var cmdStatus = cobra.Command{
	Use:   "status",
	Short: "Get status of selected connection",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getFlags(cmd); err != nil {
			fmt.Print(err)
			return
		}
		if err := requirementCheck(cmd); err != nil {
			fmt.Print(err)
			return
		}
		if err := doCommand(cmd, statusDB); err != nil {
			log.Print(err)
		}
	},
}

type attrib struct {
	Driver string
	Stat   string
	Pretty bool
	Prompt bool
}

var flg attrib

var config = map[string]string{
	"host":       "",
	"username":   "",
	"password":   "",
	"dbName":     "",
	"collection": "",
}

func init() {
	//global
	rootCmd.PersistentFlags().StringP("host", "H", "", "connection host (default- mongo:localhost:27017 / mysql:localhost:3306)") // localhost:27017 || localhost:3306
	rootCmd.PersistentFlags().StringP("username", "u", "", "database username (default- mysql: root)")
	rootCmd.PersistentFlags().String("dbName", "", "connection database name (default- mongo:xsaas_ctms / mysql:mqtt)") // xsaas_ctms || mqtt
	rootCmd.PersistentFlags().StringP("collection", "c", "", "connection database collection name")
	rootCmd.PersistentFlags().StringVarP(&flg.Driver, "driver", "d", "", "connection driver name (mongo / mysql)") // mongo || mysql
	rootCmd.PersistentFlags().StringVarP(&flg.Stat, "info", "i", "", "connection information")

	//optional
	rootCmd.PersistentFlags().BoolVarP(&flg.Pretty, "pretty", "P", false, "show pretty version of json")
	rootCmd.PersistentFlags().BoolVarP(&flg.Prompt, "password", "p", false, "call password prompt")
}

func Exec() {
	rootCmd.AddCommand(&cmdPing, &cmdInfo, &cmdStatus, &cmdDB)
	rootCmd.Execute()
}
