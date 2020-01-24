package sql

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/isollaa/conn/status"
	s "github.com/isollaa/conn/status"
)

type SQL struct {
	Driver     string
	DBName     string
	Collection string
	Result     interface{}
	Session    *sql.DB
}

var cfg *SQLConf

func (m *SQL) Connect(c s.Config) error {
	var err error
	cfg = &SQLConf{
		s.DRIVER:     c.GetString(s.DRIVER),
		s.HOST:       c.GetString(s.HOST),
		s.PORT:       c.GetInt(s.PORT),
		s.USERNAME:   c.GetString(s.USERNAME),
		s.PASSWORD:   c.GetString(s.PASSWORD),
		s.DBNAME:     c.GetString(s.DBNAME),
		s.COLLECTION: c.GetString(s.COLLECTION),
	}
	m.Session, err = sql.Open(c.GetString(s.DRIVER), cfg.GetSource())
	if err != nil {
		return err
	}
	m.Driver = c.GetString(s.DRIVER)
	m.DBName = c.GetString(s.DBNAME)
	m.Collection = c.GetString(s.COLLECTION)
	return nil
}

func (m *SQL) Close() {
	defer m.Session.Close()
}

func (m *SQL) Ping() error {
	err := m.Session.Ping()
	if err != nil {
		return err
	}
	m.Result = fmt.Sprintf("-- %s server is ok.", m.Driver)
	return nil
}

func (m *SQL) ListDB() error {
	rows, err := m.Session.Query(cfg.GetQueryDB())
	if err != nil {
		return err
	}
	defer rows.Close()
	dbNames := []string{}
	for rows.Next() {
		dbName := ""
		err = rows.Scan(&dbName)
		if err != nil {
			return err
		}
		dbNames = append(dbNames, dbName)
	}
	return nil
}

func (m *SQL) ListColl() error {
	tables := []string{}
	rows, err := m.Session.Query(cfg.GetQueryTable())
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		table := ""
		rows.Scan(&table)
		tables = append(tables, table)
	}
	m.Result = tables
	return nil
}

func (m *SQL) DiskSpace(info string) error {
	v, err := cfg.GetDiskSpace(info)
	if err != nil {
		return err
	}
	row, err := m.Session.Query(v["query"])
	if err != nil {
		return err
	}
	defer row.Close()
	table := ""
	for row.Next() {
		row.Scan(&table)
	}
	if table == "" {
		return errors.New("data not found")
	}
	if m.Driver == "mysql" {
		table = table + " kB"
	}
	m.Result = fmt.Sprintf("%s, Disk Size: %s", v["title"], table)
	return nil
}

func New() status.CommonFeature {
	return &SQL{}
}
