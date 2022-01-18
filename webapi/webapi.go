package webapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goodaye/fakeeyes/config"
)

func StartWeb() error {
	gin.DisableConsoleColor()
	if config.Loggers.AccessLogger != nil {
		gin.DefaultWriter = config.Loggers.AccessLogger.Out
	}
	app = gin.Default()
	app.RemoveExtraSlash = true

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
