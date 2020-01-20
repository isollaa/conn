package mongo

import (
	"testing"
)

const (
	driver     = "mongo"
	host       = "localhost:27017"
	username   = ""
	password   = ""
	dbName     = "xsaas_ctms"
	collection = "relationship"
)

var m Mongo

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

func TestDBStats(t *testing.T) {
	TestConnect(t)
	defer m.Close()
	m.DbStats()
}
func TestCollStats(t *testing.T) {
	TestConnect(t)
	defer m.Close()
	m.CollStats()
}
