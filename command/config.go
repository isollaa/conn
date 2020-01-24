package command

import s "github.com/isollaa/conn/status"

const (
	BUILD = "build"
	HOST  = "host"
	DB    = "db"
	COLL  = "coll"
	DISK  = "disk"

	STAT     = "stat"
	STATTYPE = "statType"
	PRETTY   = "pretty"
	PROMPT   = "prompt"
)

type attrib struct {
	stat     string
	statType string
	pretty   bool
	prompt   bool
}

var flg attrib

var config = s.Config{
	s.DRIVER:     "",
	s.HOST:       "",
	s.PORT:       0,
	s.USERNAME:   "",
	s.PASSWORD:   "",
	s.DBNAME:     "",
	s.COLLECTION: "",
}
