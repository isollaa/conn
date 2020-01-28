package command

import (
	"testing"

	cc "github.com/isollaa/conn/config"
	d "github.com/isollaa/conn/driver"
	"github.com/isollaa/conn/driver/mongo"
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

var svc d.CommonFeature
var c Config

func TestRequirementCheck(t *testing.T) {
	c.RequirementCheck(
		cc.DRIVER,
		cc.HOST,
		cc.PORT,
		cc.USERNAME,
		cc.PASSWORD,
		cc.DBNAME,
		cc.COLLECTION,
		STAT,
		TYPE,
		BEAUTY,
		PROMPT)
}

func TestPromptPassword(t *testing.T) {
	c.promptPassword()
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
		STAT:          "",
		TYPE:          "",
		BEAUTY:        false,
		PROMPT:        false,
	}
	c.connect(svc)
}

func TestPrint(t *testing.T) {
	TestConnect(t)
	c.DoPrint(svc)
}
