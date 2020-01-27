package command

import (
	"fmt"

	"github.com/hokaccha/go-prettyjson"
	cc "github.com/isollaa/conn/config"
	s "github.com/isollaa/conn/status"
	m "github.com/isollaa/conn/status/mongo"
	"github.com/isollaa/conn/status/sql"
)

var listInfo = map[string]string{
	SERVER: "server info of selected driver",
	BUILD:  "build info of selected driver",
	HOST:   "host info of selected driver",
}

var listAttribute = map[string]string{
	DB:   "list databases on selected driver",
	COLL: "list collection on selected database",
}

var listStatus = map[string]string{
	DISK: "status of disk space (postgres only)",
}

var listStatusType = map[string]string{
	DB:   "status of selected database",
	COLL: "status of selected collection",
}

func validator(flg string, list map[string]string) {
	fmt.Printf("Error: flag with argument '%s' not found \n\nTry using:\n", flg)
	for k, v := range list {
		fmt.Printf("\t%s \t%s\n", k, v)
	}
	println()
}

func printPretty(result interface{}) error {
	v, err := prettyjson.Marshal(result)
	if err != nil {
		return err
	}
	fmt.Println(string(v))
	return nil
}

func doPrint(c Config, svc s.CommonFeature) error {
	var res interface{}
	if ss, ok := svc.(*m.Mongo); ok {
		res = ss.Result
	} else {
		ss := svc.(*sql.SQL)
		res = ss.Result
	}

	if c[cc.BEAUTY].(bool) {
		if err := printPretty(res); err != nil {
			return err
		}
		return nil
	}

	fmt.Printf("%v\n", res)
	return nil
}
