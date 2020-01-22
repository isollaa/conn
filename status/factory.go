package status

import (
	"log"

	"github.com/globalsign/mgo/bson"
)

// type Attribute struct {
// 	DBName     string
// 	Collection string
// }

// type Config struct {
// 	Host     string
// 	Username string
// 	Password string
// 	*Attribute
// }

type CommonFeature interface {
	// Connect(*Config) error
	Connect(map[string]string) error
	Close()
	Ping() (string, error)
	ListDB() ([]string, error)
	ListColl() ([]string, error)
}

type NoSQLFeature interface {
	DbStats() (bson.M, error)
	Info(info string) (bson.M, error)
	CollStats() (bson.M, error)
}

type SQLFeature interface {
	DiskSpace(info string) (string, error)
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
