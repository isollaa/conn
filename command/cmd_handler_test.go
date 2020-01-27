package command

import (
	"testing"

	cc "github.com/isollaa/conn/config"
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
	beauty     = true
)

var svc s.CommonFeature
var c Config

func TestRequirementCheck(t *testing.T) {
	c.requirementCheck(
		cc.DRIVER,
		cc.HOST,
		cc.PORT,
		cc.USERNAME,
		cc.PASSWORD,
		cc.DBNAME,
		cc.COLLECTION,
		cc.STAT,
		cc.TYPE,
		cc.BEAUTY,
		cc.PROMPT)
}

func TestAutoFIll(t *testing.T) {
	c.checkAutoFill()
}

func TestPromptPassword(t *testing.T) {
	c.promptPassword()
}

func TestPrintPretty(t *testing.T) {
	str := []string{"aku", "kamu", "dirimu", "bukan dirinya"}
	printPretty(str)
}

func TestConnect(t *testing.T) {
	svc = &mongo.Mongo{}
	c = Config{
		cc.DRIVER:     driver,
		cc.HOST:       host,
		cc.PORT:       port,
		cc.USERNAME:   username,
		cc.PASSWORD:   password,
		cc.DBNAME:     dbName,
		cc.COLLECTION: collection,
		cc.STAT:       "",
		cc.TYPE:       "",
		cc.BEAUTY:     false,
		cc.PROMPT:     false,
	}
	c.connect(svc)
}

func TestPing(t *testing.T) {
	TestConnect(t)
	defer svc.Close()
	ping(c, svc)
}

func TestList(t *testing.T) {
	TestConnect(t)
	defer svc.Close()
	c[cc.STAT] = DB
	c[cc.BEAUTY] = beauty
	list(c, svc)
}

func TestHostInfo(t *testing.T) {
	TestConnect(t)
	defer svc.Close()
	c[cc.STAT] = HOST
	c[cc.BEAUTY] = beauty
	infoDB(c, svc)
}

func TestBuildInfo(t *testing.T) {
	TestConnect(t)
	defer svc.Close()
	c[cc.STAT] = BUILD
	c[cc.BEAUTY] = beauty
	infoDB(c, svc)
}

func TestDBStats(t *testing.T) {
	TestConnect(t)
	defer svc.Close()
	c[cc.STAT] = DISK
	c[cc.TYPE] = DB
	c[cc.BEAUTY] = beauty
	statusDB(c, svc)
}

func TestCollStats(t *testing.T) {
	TestConnect(t)
	defer svc.Close()
	c[cc.STAT] = COLL
	c[cc.BEAUTY] = beauty
	statusDB(c, svc)
}
