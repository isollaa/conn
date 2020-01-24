package status

import (
	"log"
)

type CommonFeature interface {
	Connect(Config) error
	Close()
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
	name := getPackageName(list)
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
		log.Fatalf("driver '%s' not available", key)
	}
	return run()
}
