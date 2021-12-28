package rdb

import (
	"fmt"

	"github.com/go-xorm/xorm"
	"github.com/goodaye/fakeeyes/config"
	"github.com/goodaye/wire"
)

var defaultModel *Model

func init() {
	wire.Append(SVC{})
}

type SVC struct {
	wire.BaseService
}

// Init
func (s SVC) Init() error {
	return Initialize(config.GlobalConfig.MetaDB)
}

// Initialize package db variable
func Initialize(config config.DBConfig) error {

	model, err := NewModel(config)
	if err != nil {
		return err
	}
	defaultModel = model
	return nil
}

// GetModel return default Model
func GetModel() *Model {
	if defaultModel == nil {
		panic(fmt.Errorf("default model Not Initialize"))
	}
	err := defaultModel.engine.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	return defaultModel
}

// NewSession
func NewSession() *xorm.Session {
	return GetModel().NewSession()
}

// Engine
func Engine() *xorm.Engine {
	return GetModel().GetEngine()
}

func SyncDB() error {
	return defaultModel.SyncDB()
}
func DropDB() error {
	return defaultModel.DropDB()
}
