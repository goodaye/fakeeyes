package service

import (
	"github.com/go-xorm/xorm"
	"github.com/goodaye/fakeeyes/dao/rdb"
)

type Entity struct {
	Session *xorm.Session
}

func (e *Entity) WithSession() *xorm.Session {

	if e.Session == nil {
		e.Session = rdb.NewSession()
	}
	return e.Session
}
