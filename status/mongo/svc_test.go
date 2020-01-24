package mongo

import (
	s "github.com/isollaa/conn/status"
	"testing"
)

const (
	driver     = "mongo"
	host       = "localhost"
	port       = 27017
	username   = ""
	password   = ""
	dbName     = "xsaas_ctms"
	collection = "project_form"
)

var m Mongo

func TestConnect(t *testing.T) {
	m.Connect(s.Config{
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

func TestDiskSpace(t *testing.T) {
	TestConnect(t)
	defer m.Close()
	m.DiskSpace("db")
}

func TestInfo(t *testing.T) {
	TestConnect(t)
	defer m.Close()
	m.Info("hostInfo")
}
