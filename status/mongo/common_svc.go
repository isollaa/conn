package mongo

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	"github.com/hokaccha/go-prettyjson"
	"github.com/isollaa/conn/status"
)

type mongo struct {
	// status.Attribute
	DBName     string
	Collection string
	Session    *mgo.Session
}

func (m *mongo) Connect(c map[string]string) error {
	// func (m *mongo) Connect(c *status.Config) error {
	session, err := mgo.Dial(c["host"])
	if err != nil {
		return err
	}
	m.DBName = c["username"]
	m.Collection = c["collection"]
	// m.Attribute = *c.Attribute
	m.Session = session
	return nil
}

func (m *mongo) Close() {
	defer m.Session.Close()
}

func (m *mongo) Ping() error {
	err := m.Session.Ping()
	if err != nil {
		return err
	}
	log.Print("- MongoDB server is ok.")

	return nil
}

func (m *mongo) ListDB() error {
	result := []string{}
	result, err := m.Session.DatabaseNames()
	if err != nil {
		return err
	}

	v, _ := prettyjson.Marshal(result)
	fmt.Println(string(v))
	return nil
}

func (m *mongo) ListColl() error {
	result := []string{}
	result, err := m.Session.DB(m.DBName).CollectionNames()
	if err != nil {
		return err
	}

	v, _ := prettyjson.Marshal(result)
	fmt.Println("Collections:", string(v))
	return nil
}

func New() status.CommonFeature {
	return &mongo{}
}
