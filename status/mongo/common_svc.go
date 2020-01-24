package mongo

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	s "github.com/isollaa/conn/status"
)

type Mongo struct {
	DBName     string
	Collection string
	Result     interface{}
	Session    *mgo.Session
}

func (m *Mongo) Connect(c s.Config) error {
	source := fmt.Sprintf("%s:%d", c.GetString(s.HOST), c.GetInt(s.PORT))
	session, err := mgo.Dial(source)
	if err != nil {
		return err
	}
	m.DBName = c.GetString(s.DBNAME)
	m.Collection = c.GetString(s.COLLECTION)
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
