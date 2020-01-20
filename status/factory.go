package status

import "log"

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
	Ping() error
	ListDB() (interface{}, error)
	ListColl() (interface{}, error)
}

type NoSQLFeature interface {
	CommonFeature
	DbStats() (interface{}, error)
	Info(info string) (interface{}, error)
	CollStats() (interface{}, error)
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
