package command

import (
	"log"

	"github.com/isollaa/conn/helper"
	"github.com/spf13/cobra"
)

type factory func(cfg Config) *cobra.Command

var listCommand = make(map[string]factory)

func Register(list factory) {
	name := helper.GetName(helper.PACKAGE, list)
	ok := false
	for k, _ := range listCommand {
		if name != k {
			ok = true
			continue
		}
		log.Printf("Service %s already registered !", name)
	}
	if ok || len(listCommand) == 0 {
		listCommand[name] = list
	}
}

func New() map[string]factory {
	return listCommand
}

// func New(cfg Config, key ...string) []*cobra.Command {
// 	var cmds []*cobra.Command
// 	err := false
// 	if len(key) != 0 {
// 		for _, v := range key {
// 			cmd := listCommand[v]
// 			if cmd == nil {
// 				err = true
// 				break
// 			}
// 			cmds = append(cmds, cmd(cfg))
// 		}
// 		return cmds
// 	}
// 	for k, _ := range listCommand {
// 		cmd := listCommand[k]
// 		if cmd == nil {
// 			err = true
// 			break
// 		}
// 		cmds = append(cmds, cmd(cfg))
// 	}
// 	if err {
// 		log.Fatalf("driver '%s' not available.\n\nUse `app [command] --help` for more information about a command. ", key)
// 	}
// 	return cmds
// }
