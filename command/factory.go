package command

import (
	"log"

	"github.com/isollaa/conn/helper"
	"github.com/spf13/cobra"
)

type factory func() *cobra.Command

var listCommand []factory

func Register(list factory) {
	name := helper.GetPackageName(list)
	ok := false
	for _, v := range listCommand {
		vName := helper.GetPackageName(v)
		if name != vName {
			ok = true
			continue
		}
		log.Printf("Service %s already registered !", name)
	}
	if ok || len(listCommand) == 0 {
		listCommand = append(listCommand, list)
	}
}

func New() []factory {
	return listCommand
}
