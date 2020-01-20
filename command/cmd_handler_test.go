package command

import (
	"github.com/isollaa/conn/status"
	"github.com/isollaa/conn/status/mongo"
	"testing"
)

const (
	driver     = "mongo"
	host       = "localhost:27017"
	username   = ""
	password   = ""
	dbName     = "xsaas_ctms"
	collection = "project_form"
	pretty     = true
)

var str = []string{"aku", "kamu", "dirimu", "bukan dirinya"}
var svc status.CommonFeature

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
	err := svc.Connect(map[string]string{
		"host":       host,
		"username":   username,
		"password":   password,
		"dbName":     dbName,
		"collection": collection,
	})
	if err != nil {
		t.Log(err)
	}
}

func TestPing(t *testing.T) {
	TestConnect(t)
	defer svc.Close()
	err := ping(svc)
	if err != nil {
		t.Log(err)
	}
}

func TestListDB(t *testing.T) {
	flg.Pretty = pretty
	TestConnect(t)
	defer svc.Close()
	listDB(svc)
}

func TestListColl(t *testing.T) {
	flg.Pretty = pretty
	TestConnect(t)
	defer svc.Close()
	listColl(svc)
}

func TestHostInfo(t *testing.T) {
	flg.Stat = "host"
	flg.Pretty = pretty
	TestConnect(t)
	defer svc.Close()
	infoDB(svc)
}

func TestBuildInfo(t *testing.T) {
	flg.Stat = "build"
	flg.Pretty = pretty
	TestConnect(t)
	defer svc.Close()
	infoDB(svc)
}

func TestDBStats(t *testing.T) {
	flg.Stat = "dbstats"
	flg.Pretty = pretty
	TestConnect(t)
	defer svc.Close()
	statusDB(svc)
}

func TestCollStats(t *testing.T) {
	flg.Stat = "collstats"
	flg.Pretty = pretty
	TestConnect(t)
	defer svc.Close()
	statusDB(svc)
}
