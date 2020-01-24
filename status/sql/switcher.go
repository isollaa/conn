package sql

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	s "github.com/isollaa/conn/status"
	_ "github.com/lib/pq"
)

type SQLConf map[string]interface{}

func (c SQLConf) GetSource() string {
	source := ""
	switch c[s.DRIVER] {
	case "mysql":
		source = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c[s.USERNAME], c[s.PASSWORD], c[s.HOST], c[s.PORT], c[s.DBNAME])
	case "postgres":
		source = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", c[s.HOST], c[s.PORT], c[s.USERNAME], c[s.PASSWORD], c[s.DBNAME])
	}
	return source
}

func (c SQLConf) GetQueryDB() string {
	query := ""
	switch c[s.DRIVER] {
	case "mysql":
		query = "SHOW DATABASES"
	case "postgres":
		query = "SELECT datname FROM pg_database WHERE datistemplate = false"
	}
	return query
}

func (c SQLConf) GetQueryTable() string {
	query := ""
	switch c[s.DRIVER] {
	case "mysql":
		query = "SHOW TABLES"
	case "postgres":
		query = "SELECT table_schema,table_name FROM information_schema.tables ORDER BY table_schema,table_name"
	}
	return query
}

func (c SQLConf) GetDiskSpace(info string) (map[string]string, error) {
	v := map[string]string{}
	switch c[s.DRIVER] {
	case "mysql":
		switch info {
		case "db":
			return v, errors.New("disk status not available")
		case "coll":
			v["title"] = fmt.Sprintf("Table - %s", c[s.COLLECTION])
			v["query"] = fmt.Sprintf("SELECT (data_length+index_length)/power(1024,1) FROM information_schema.tables WHERE table_schema='%s' and table_name='%s'", c[s.DBNAME], c[s.COLLECTION])
		}
	case "postgres":
		switch info {
		case "db":
			info = "pg_database_size"
			v["title"] = "DB - " + c[s.DBNAME].(string)
		case "coll":
			info = "pg_total_relation_size"
			c[s.DBNAME] = c[s.COLLECTION].(string)
			v["title"] = fmt.Sprintf("Table - %s", c[s.COLLECTION])
		}
		v["query"] = fmt.Sprintf("SELECT pg_size_pretty(%s('%s'))", info, c[s.DBNAME])
	}
	return v, nil
}
