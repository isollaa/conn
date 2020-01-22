package sql

import "testing"

const (
	driver     = "mysql"
	host       = "localhost:3306"
	username   = "root"
	password   = ""
	dbName     = "mqtt"
	collection = "relationship"
)

var m SQL

func TestConnect(t *testing.T) {
	m.Connect(map[string]string{
		"host":       host,
		"username":   username,
		"password":   password,
		"dbName":     dbName,
		"collection": collection,
	})
}

func TestPing(t *testing.T) {
	TestConnect(t)
	defer m.Close()
	m.Ping()
}

func TestListDB(t *testing.T) {
	TestConnect(t)
	defer m.Close()
	m.ListDB()
}

func TestListColl(t *testing.T) {
	TestConnect(t)
	defer m.Close()
	m.ListColl()
}
