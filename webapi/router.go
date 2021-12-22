package webapi

import "github.com/goodaye/fakeeyes/webapi/handlers"

// var app gin

func InitHandler() {

	app.GET("/ws", handlers.WsPage)

}
