package webapi

import "github.com/goodaye/fakeeyes/webapi/handlers"

// var app gin

func InitHandler() {

	// api
	app.GET("/ping", handlers.Handler{}.Pong)
	app.GET("/api/version", handlers.Handler{}.Version)

	// API for system call
	api := app.Group("/api/v1")
	{
		handlers.UserHandler{}.Router(api)
		handlers.DevcieHandler{}.Router(api)
		handlers.AdminHandler{}.Router(api)
	}
}
