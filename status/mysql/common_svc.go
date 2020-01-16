package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/hokaccha/go-prettyjson"
	"github.com/isollaa/conn/status"
)

type mysql struct {
	// status.Attribute
	DBName     string
	Collection string
	Session    *sql.DB
}

func (m *mysql) Connect(c map[string]string) error {
	// func (m *mysql) Connect(c *status.Config) error {
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s", c["username"], c["password"], c["host"], c["dbName"])
	session, err := sql.Open("mysql", source)
	if err != nil {
		return err
	}
	m.DBName = c["username"]
	m.Collection = c["collection"]
	// m.Attribute = *c.Attribute
	m.Session = session
	return nil
}

func (m *mysql) Close() {
	defer m.Session.Close()
}

func (m *mysql) Ping() error {
	err := m.Session.Ping()
	if err != nil {
		return err
	}
	log.Print("- Mysql server is ok.")

	return nil
}

func (m *mysql) ListDB() error {
	rows, err := m.Session.Query("SHOW DATABASES")
	if err != nil {
		return err
	}
	dbNames := []string{}
	defer rows.Close()
	for rows.Next() {
		dbName := ""
		err = rows.Scan(&dbName)
		if err != nil {
			return err
		}
		dbNames = append(dbNames, dbName)
	}

	v, _ := prettyjson.Marshal(dbNames)
	fmt.Println(string(v))
	return nil
}

func (m *mysql) ListColl() error {
	res, err := m.Session.Query("SHOW TABLES")
	if err != nil {
		return err
	}
	tables := []string{}
	for res.Next() {
		table := ""
		res.Scan(&table)
		tables = append(tables, table)
	}

	v, _ := prettyjson.Marshal(tables)
	fmt.Println("Tables: ", string(v))
	return nil
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
	return &mysql{}
}
