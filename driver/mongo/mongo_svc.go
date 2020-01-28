package mongo

import (
	"github.com/globalsign/mgo/bson"
)

func (m *Mongo) Info(info string) error {

	query := &bson.D{}
	if info == "serverInfo" {
		query = &bson.D{{"serverStatus", 1}, {"repl", 0}, {"metrics", 0}, {"locks", 0}}
	} else {
		query = &bson.D{{info, 1}}
	}
	result := bson.M{}
	err := m.Session.DB("admin").Run(query, &result)
	if err != nil {
		return err
	}
	m.Result = result
	return nil
}
