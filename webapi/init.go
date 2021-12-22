package webapi

import (
	"github.com/gin-gonic/gin"
	"github.com/goodaye/wire"
)

var (
	app *gin.Engine
)

type SVC struct {
	wire.BaseService
}

func (s SVC) Init() error {

	return nil
}
func (s SVC) Start() error {
	return StartWeb()
}
