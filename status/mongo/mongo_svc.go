package mongo

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/hokaccha/go-prettyjson"
)

func (m *mongo) Info(info string) error {
	result := bson.M{}
	err := m.Session.DB("admin").Run(bson.D{{info, 1}}, &result)
	if err != nil {
		return err
	}
	println(info, ":")
	v, _ := prettyjson.Marshal(result)
	fmt.Println(string(v))

	return nil
}

func (m *mongo) CollStats() error {
	result := bson.M{}
	err := m.Session.DB(m.DBName).Run(&bson.D{bson.DocElem{"collstats", m.Collection}}, &result)
	if err != nil {
		return err
	}
	println("collstats :")
	v, _ := prettyjson.Marshal(result)
	fmt.Println(string(v))

	return nil
}

func (m *mongo) DbStats() error {
	result := bson.M{}
	err := m.Session.DB(m.DBName).Run("dbstats", &result)
	if err != nil {
		return err
	}

	println("dbstats :")
	v, _ := prettyjson.Marshal(result)
	fmt.Println(string(v))

	return nil
}
