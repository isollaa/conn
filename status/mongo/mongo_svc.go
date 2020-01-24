package mongo

import (
	"github.com/globalsign/mgo/bson"
)

func (m *Mongo) Info(info string) error {
	result := bson.M{}
	err := m.Session.DB("admin").Run(bson.D{{info, 1}}, &result)
	if err != nil {
		return err
	}
	m.Result = result
	return nil
}
