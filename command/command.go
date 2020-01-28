package command

import (
	"github.com/spf13/cobra"
)

func initFlag(cmd *cobra.Command) {
	//global
	cmd.PersistentFlags().StringP("driver", "d", "", "connection driver name (mongo / mysql / postgres)")
	cmd.PersistentFlags().StringP("host", "H", "localhost", "connection host ")
	cmd.PersistentFlags().IntP("port", "P", 0, "connection port (default - mongo:27017 / mysql:3306 / postgres:5432)")
	cmd.PersistentFlags().StringP("username", "u", "", "database username (default - mysql:root / postgres:postgres)")
	cmd.PersistentFlags().String("dbname", "", "connection database name (default - mongo:xsaas_ctms / mysql:mqtt)")
	cmd.PersistentFlags().StringP("collection", "c", "", "connection database collection name")
	cmd.PersistentFlags().StringP("stat", "s", "", "connection information")
	cmd.PersistentFlags().StringP("type", "t", "", "connection information type")
	//optional
	cmd.PersistentFlags().BoolP("beauty", "b", false, "show pretty version of json")
	cmd.PersistentFlags().BoolP("prompt", "p", false, "call password prompt")
}

func Exec() {
	rootCmd := &cobra.Command{Use: "app"}
	initFlag(rootCmd)
	for _, v := range New() {
		rootCmd.AddCommand(v())
	}
	rootCmd.Execute()
}
