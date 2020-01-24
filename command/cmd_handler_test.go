package command

import (
	"testing"

	s "github.com/isollaa/conn/status"
	"github.com/isollaa/conn/status/mongo"
)

const (
	driver     = "mongo"
	host       = "localhost"
	port       = 27017
	username   = ""
	password   = ""
	dbName     = "xsaas_ctms"
	collection = "project_form"
	pretty     = true
)

var str = []string{"aku", "kamu", "dirimu", "bukan dirinya"}
var svc s.CommonFeature

func TestPromptPassword(t *testing.T) {
	promptPassword()
}

func TestAutoFIll(t *testing.T) {
	checkAutoFill()
}

func TestPrintPretty(t *testing.T) {
	printPretty(str)
}

func TestConnect(t *testing.T) {
	svc = &mongo.Mongo{}
	config = s.Config{
		"driver":     driver,
		"host":       host,
		"port":       port,
		"username":   username,
		"password":   password,
		"dbname":     dbName,
		"collection": collection,
	}
	connect(svc)
}

func TestPing(t *testing.T) {
	TestConnect(t)
	defer svc.Close()
	ping(svc)
}

func TestList(t *testing.T) {
	flg.stat = DB
	flg.pretty = pretty
	TestConnect(t)
	defer svc.Close()
	list(svc)
}

func TestHostInfo(t *testing.T) {
	flg.stat = HOST
	flg.pretty = pretty
	TestConnect(t)
	defer svc.Close()
	infoDB(svc)
}

func TestBuildInfo(t *testing.T) {
	flg.stat = HOST
	flg.pretty = pretty
	TestConnect(t)
	defer svc.Close()
	infoDB(svc)
}

func TestDBStats(t *testing.T) {
	flg.stat = DISK
	flg.statType = DB
	flg.pretty = pretty
	TestConnect(t)
	defer svc.Close()
	statusDB(svc)
}

func TestCollStats(t *testing.T) {
	flg.stat = COLL
	flg.pretty = pretty
	TestConnect(t)
	defer svc.Close()
	statusDB(svc)
}
