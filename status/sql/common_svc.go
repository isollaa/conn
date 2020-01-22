package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/isollaa/conn/status"
)

type SQL struct {
	Driver     string
	DBName     string
	Collection string
	Session    *sql.DB
}

var cfg *Config

func (m *SQL) Connect(c map[string]string) error {
	var err error
	cfg = &Config{
		Driver:     c["driver"],
		Host:       c["host"],
		Port:       c["port"],
		Username:   c["username"],
		Password:   c["password"],
		DBName:     c["dbName"],
		Collection: c["collection"],
	}
	m.Session, err = sql.Open(cfg.Driver, cfg.GetSource())
	if err != nil {
		return err
	}
	m.Driver = cfg.Driver
	m.DBName = cfg.DBName
	m.Collection = cfg.Collection
	return nil
}

func (m *SQL) Close() {
	defer m.Session.Close()
}

func (m *SQL) Ping() (interface{}, error) {
	err := m.Session.Ping()
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("-- %s server is ok.", m.Driver), nil
}

func (m *SQL) ListDB() (interface{}, error) {
	dbNames := []string{}
	rows, err := m.Session.Query(cfg.GetQueryDB())
	if err != nil {
		return dbNames, err
	}
	defer rows.Close()
	for rows.Next() {
		dbName := ""
		err = rows.Scan(&dbName)
		if err != nil {
			return dbNames, err
		}
		dbNames = append(dbNames, dbName)
	}
	return dbNames, nil
}

func (m *SQL) ListColl() (interface{}, error) {
	tables := []string{}
	rows, err := m.Session.Query(cfg.GetQueryTable())
	if err != nil {
		return tables, err
	}
	defer rows.Close()
	for rows.Next() {
		table := ""
		rows.Scan(&table)
		tables = append(tables, table)
	}
	return tables, nil
}

// func (m *mysql) CollData() error {
// 	var (
// 		result  interface{}
// 		results []interface{}
// 	)
// 	rows, err := m.Session.Query("SELECT * FROM " + m.Table)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		err = rows.Scan(&result)
// 		if err != nil {
// 			return err
// 		}
// 		results = append(results, result)
// 	}

// 	v, _ := prettyjson.Marshal(results)
// 	fmt.Println("Results All: ", string(v))
// 	return nil
// }

func New() status.CommonFeature {
	return &SQL{}
}
