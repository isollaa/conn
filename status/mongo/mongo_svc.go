package mongo

import (
	"github.com/globalsign/mgo/bson"
)

func (m *Mongo) Info(info string) (interface{}, error) {
	result := bson.M{}
	err := m.Session.DB("admin").Run(bson.D{{info, 1}}, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *Mongo) CollStats() (interface{}, error) {
	result := bson.M{}
	err := m.Session.DB(m.DBName).Run(&bson.D{bson.DocElem{"collstats", m.Collection}}, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *Mongo) DbStats() (interface{}, error) {
	result := bson.M{}
	err := m.Session.DB(m.DBName).Run("dbstats", &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
