package webapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goodaye/fakeeyes/config"
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
	gin.DisableConsoleColor()
	if config.Loggers.AccessLogger != nil {
		gin.DefaultWriter = config.Loggers.AccessLogger.Out
	}
	app = gin.Default()
	app.RemoveExtraSlash = true
	// //检查static目录是否存在
	// StaticFilePath := path.Join(DefaultTemplatePath, "static")
	// fileinfo, err := os.Stat(StaticFilePath)
	// if err != nil {
	// 	return err
	// }
	// if !fileinfo.IsDir() {
	// 	return handlers.ErrorPathIsNotDir
	// }
	// app.Static("/webapi/static", StaticFilePath)
	InitHandler()

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	go func() {
		address := fmt.Sprintf(":%d", config.GlobalConfig.HTTP.Port)
		fmt.Printf("start on address:  %s \n", address)
		err := app.Run(address)
		fmt.Println(err.Error())
	}()
	return nil
}
