package mongo

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	cc "github.com/isollaa/conn/config"
	s "github.com/isollaa/conn/driver"
)

type Mongo struct {
	DBName     string
	Collection string
	Result     interface{}
	Session    *mgo.Session
}

func (m *Mongo) AutoFill(c cc.Config) {
	if c[cc.PORT] == 0 {
		c[cc.PORT] = 27017
	}
	if c[cc.DBNAME] == "" {
		c[cc.DBNAME] = "xsaas_ctms"
	}
}

func (m *Mongo) Connect(c cc.Config) error {
	source := fmt.Sprintf("%s:%d", c.GetString(cc.HOST), c.GetInt(cc.PORT))
	session, err := mgo.Dial(source)
	if err != nil {
		return err
	}
	m.DBName = c.GetString(cc.DBNAME)
	m.Collection = c.GetString(cc.COLLECTION)
	m.Session = session
	return nil
}

func (m *Mongo) Close() {
	defer m.Session.Close()
}

func (m *Mongo) Ping() error {
	err := m.Session.Ping()
	if err != nil {
		return err
	}
	m.Result = fmt.Sprintf("-- MongoDB server is ok.")
	return nil
}

func (m *Mongo) ListDB() error {
	result, err := m.Session.DatabaseNames()
	if err != nil {
		return err
	}
	m.Result = result
	return nil
}

func (m *Mongo) ListColl() error {
	result, err := m.Session.DB(m.DBName).CollectionNames()
	if err != nil {
		return err
	}
	m.Result = result
	return nil
}

func (m *Mongo) DiskSpace(info string) error {
	var query interface{}
	switch info {
	case "db":
		query = "dbstats"
	case "coll":
		query = &bson.D{bson.DocElem{"collstats", m.Collection}}
	}
	result := bson.M{}
	err := m.Session.DB(m.DBName).Run(query, &result)
	if err != nil {
		return err
	}
	m.Result = result
	return nil
}

func New() s.CommonFeature {
	return &Mongo{}
}
