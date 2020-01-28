package sql

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	cc "github.com/isollaa/conn/config"
	_ "github.com/lib/pq"
)

type SQLConf cc.Config

func (c SQLConf) GetSource() string {
	source := ""
	switch c[cc.DRIVER] {
	case "mysql":
		source = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c[cc.USERNAME], c[cc.PASSWORD], c[cc.HOST], c[cc.PORT], c[cc.DBNAME])
	case "postgres":
		source = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", c[cc.HOST], c[cc.PORT], c[cc.USERNAME], c[cc.PASSWORD], c[cc.DBNAME])
	}
	return source
}

func (c SQLConf) GetQueryDB() string {
	query := ""
	switch c[cc.DRIVER] {
	case "mysql":
		query = "SHOW DATABASES"
	case "postgres":
		query = "SELECT datname FROM pg_database WHERE datistemplate = false"
	}
	return query
}

func (c SQLConf) GetQueryTable() string {
	query := ""
	switch c[cc.DRIVER] {
	case "mysql":
		query = "SHOW TABLES"
	case "postgres":
		query = "SELECT table_schema,table_name FROM information_schema.tables ORDER BY table_schema,table_name"
	}
	return query
}

func (c SQLConf) GetDiskSpace(info string) (map[string]string, error) {
	v := map[string]string{}
	switch c[cc.DRIVER] {
	case "mysql":
		switch info {
		case "db":
			// v["title"] = fmt.Sprintf("Table - %s", c[cc.DBNAME])
			// v["query"] = fmt.Sprintf("SELECT table_schema AS 'Db Name', Round( Sum( data_length + index_length ) / 1024 / 1024, 3) AS 'Db Size (MB)', Round( Sum( data_free ) / 1024 / 1024, 3 ) AS 'Free Space (MB)' FROM information_schema.tables GROUP BY table_schema")
			return v, errors.New("disk status not available")
		case "coll":
			v["title"] = fmt.Sprintf("Table - %s", c[cc.COLLECTION])
			v["query"] = fmt.Sprintf("SELECT (data_length+index_length)/power(1024,1) FROM information_schema.tables WHERE table_schema='%s' and table_name='%s'", c[cc.DBNAME], c[cc.COLLECTION])
		}
	case "postgres":
		switch info {
		case "db":
			info = "pg_database_size"
			v["title"] = "DB - " + c[cc.DBNAME].(string)
		case "coll":
			info = "pg_total_relation_size"
			c[cc.DBNAME] = c[cc.COLLECTION].(string)
			v["title"] = fmt.Sprintf("Table - %s", c[cc.COLLECTION])
		}
		v["query"] = fmt.Sprintf("SELECT pg_size_pretty(%s('%s'))", info, c[cc.DBNAME])
	default:
		return v, errors.New("disk status not available")
	}
	return v, nil
}
