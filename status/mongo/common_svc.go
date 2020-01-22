package mongo

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/isollaa/conn/status"
)

type Mongo struct {
	// status.Attribute
	DBName     string
	Collection string
	Session    *mgo.Session
}

func (m *Mongo) Connect(c map[string]string) error {
	source := fmt.Sprintf("%s:%s", c["host"], c["port"])
	session, err := mgo.Dial(source)
	if err != nil {
		return err
	}
	m.DBName = c["dbName"]
	m.Collection = c["collection"]
	// m.Attribute = *c.Attribute
	m.Session = session
	return nil
}

func (m *Mongo) Close() {
	defer m.Session.Close()
}

func (m *Mongo) Ping() (interface{}, error) {
	err := m.Session.Ping()
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("-- MongoDB server is ok."), nil
}

func (m *Mongo) ListDB() (interface{}, error) {
	result, err := m.Session.DatabaseNames()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *Mongo) ListColl() (interface{}, error) {
	result, err := m.Session.DB(m.DBName).CollectionNames()
	if err != nil {
		return result, err
	}
	return result, nil
}

func New() status.CommonFeature {
	return &Mongo{}
}
