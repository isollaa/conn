package register

import (
	"github.com/isollaa/conn/status"
	"github.com/isollaa/conn/status/mongo"
	"github.com/isollaa/conn/status/sql"
)

func init() {
	status.Register(mongo.New)
	status.Register(sql.New)
}
