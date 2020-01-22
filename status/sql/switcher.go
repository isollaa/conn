package sql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Config struct {
	Driver     string
	Host       string
	Port       string
	Username   string
	Password   string
	DBName     string
	Collection string
}

func (c *Config) GetSource() string {
	source := ""
	switch c.Driver {
	case "mysql":
		source = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Username, c.Password, c.Host, c.Port, c.DBName)
	case "postgres":
		source = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.Username, c.Password, c.DBName)
	}
	return source
}

func (c *Config) GetQueryDB() string {
	query := ""
	switch c.Driver {
	case "mysql":
		query = "SHOW DATABASES"
	case "postgres":
		query = "SELECT datname FROM pg_database WHERE datistemplate = false"
	}
	return query
}

func (c *Config) GetQueryTable() string {
	query := ""
	switch c.Driver {
	case "mysql":
		query = "SHOW TABLES"
	case "postgres":
		query = "SELECT table_schema,table_name FROM information_schema.tables ORDER BY table_schema,table_name"
	}
	return query
}

func (c *Config) GetDiskSpace(info string) (string, string) {
	msg := ""
	switch info {
	case "db":
		info = "pg_database_size"
		msg = "DB - " + c.DBName
	case "coll":
		info = "pg_total_relation_size"
		c.DBName = c.Collection
		msg = "Table - " + c.Collection
	}
	query := fmt.Sprintf("SELECT pg_size_pretty(%s('%s'))", info, c.DBName)
	return msg, query
}
