package rdb

import (
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	"github.com/goodaye/fakeeyes/config"
	"github.com/goodaye/fakeeyes/protos/db"

	//使用 mysql 驱动
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"
)

var tableList []interface{}

func init() {
	tableList = []interface{}{
		new(db.User),
		new(db.Device),
		new(db.Room),
	}
}

// Model  DB engine
type Model struct {
	engine     *xorm.Engine
	dbName     string
	config     config.DBConfig
	dataSource string
}

// GetEngine  return engine
func (model *Model) GetEngine() *xorm.Engine {
	return model.engine
}

// NewModel Create DBEngine instance
func NewModel(conf config.DBConfig) (*Model, error) {
	engine, err := xorm.NewEngine("mysql", conf.Datasource)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	engine.SetMapper(core.GonicMapper{})
	engine.ShowSQL(conf.SQLLog)

	engine.SetMaxOpenConns(conf.MaxOpenConn)
	engine.SetMaxIdleConns(conf.MaxIdleConn)
	var dbengine = Model{
		engine:     engine,
		dbName:     "mysql",
		config:     conf,
		dataSource: conf.Datasource,
	}
	return &dbengine, nil
}

// Reconnect  重连
func (model *Model) Reconnect() error {

	if model.engine != nil {
		model.engine.Close()
	}
	newengine, err := xorm.NewEngine(model.dbName, model.dataSource)

	if err != nil {
		return err
	}
	newengine.SetMapper(core.GonicMapper{})
	newengine.ShowSQL(model.config.SQLLog)

	newengine.SetMaxOpenConns(model.config.MaxOpenConn)
	newengine.SetMaxIdleConns(model.config.MaxIdleConn)

	if model.engine != nil {
		model.engine.Close()
	}
	model.engine = newengine
	return nil
}

//SyncDB sync table defined in  struct.go
func (model *Model) SyncDB() error {

	err := model.engine.Sync2(tableList...)
	return err
}

// DropDB drop table from db
func (model *Model) DropDB() error {

	err := model.engine.DropTables(tableList...)
	return err
}

// CleanDB clean table from db
func (model *Model) CleanDB() error {

	var err error
	for _, table := range tableList {
		tablename := model.engine.TableName(table)
		fmt.Println(tablename)
		sql := fmt.Sprintf("truncate %s", tablename)
		_, err = model.engine.Exec(sql)
		if err != nil {
			return err
		}
	}
	return err
}

// NewSession Create NewSession From DB Engine
func (model *Model) NewSession() *xorm.Session {
	session := model.engine.NewSession()
	session.Begin()
	return session
}
