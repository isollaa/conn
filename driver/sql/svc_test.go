package sql

import (
	"testing"

	cc "github.com/isollaa/conn/config"
)

const (
	driver     = /* "mysql"  */ "postgres"
	host       = "localhost"
	port       = /* "3306" */ "5432"
	username   = /* "root" */ "postgres"
	password   = /* "" */ "12345"
	dbName     = "mqtt"
	collection = "listclient"
)

var m SQL

func TestConnect(t *testing.T) {
	m.Connect(cc.Config{
		"driver":     driver,
		"host":       host,
		"port":       port,
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

func TestDiscSpace(t *testing.T) {
	TestConnect(t)
	defer m.Close()
	m.DiskSpace("coll")
}
