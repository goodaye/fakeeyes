package webapi

import (
	"github.com/gin-gonic/gin"
	"github.com/goodaye/fakeeyes/config"
	"github.com/goodaye/fakeeyes/webapi/handlers"
	"github.com/goodaye/wire"
)

var (
	app *gin.Engine
)

type SVC struct {
	wire.BaseService
}

func (s SVC) Init() error {

	handlers.Logger = config.Loggers.WebLogger
	return nil
}
func (s SVC) Start() error {
	return StartWeb()
}
