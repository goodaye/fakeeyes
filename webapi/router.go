package webapi

import "github.com/goodaye/fakeeyes/webapi/handlers"

// var app gin

func InitHandler() {

	app.GET("/user", handlers.User{}.Login)
	// app.GET("/user", handlers.WsPage)

}
