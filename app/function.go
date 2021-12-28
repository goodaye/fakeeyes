package main

import (
	"github.com/goodaye/fakeeyes/dao/rdb"
	"github.com/goodaye/wire"
)

// SyncDB  syncdb
func SyncDB() error {
	err := wire.Init()
	if err != nil {
		return err
	}
	err = rdb.SyncDB()
	return err
}

// DropDB  dropdb
func DropDB() error {
	err := wire.Init()
	if err != nil {
		return err
	}
	err = rdb.DropDB()
	return err
}
