package driver

import (
	"log"

	cc "github.com/isollaa/conn/config"
	"github.com/isollaa/conn/helper"
)

type Initial interface {
	AutoFill(cc.Config)
	Connect(cc.Config) error
	Close()
}

type CommonFeature interface {
	Initial
	Ping() error
	ListDB() error
	ListColl() error
	DiskSpace(info string) error
}

type NoSQLFeature interface {
	Info(info string) error
}

type factory func() CommonFeature

var listFactory = make(map[string]factory)

//auto register service by its package name
func Register(list factory) {
	name := helper.GetPackageName(list)
	if list == nil {
		log.Panicf("Service %s does not exist.", name)
	}
	_, registered := listFactory[name]
	if registered {
		log.Fatalf("Service %s already registered. Ignoring.", name)
	}
	listFactory[name] = list
}

// fill parameter using selected service package name
func New(key string) CommonFeature {
	run := listFactory[key]
	if run == nil {
		log.Fatalf("driver '%s' not available.\n\nUse `app [command] --help` for more information about a command. ", key)
	}
	return run()
}
