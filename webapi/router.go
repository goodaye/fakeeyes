package webapi

import "github.com/goodaye/fakeeyes/webapi/handlers"

// var app gin

func InitHandler() {

	user := app.Group("/User")
	{
		handlers.UserHandler{}.Router(user)
	}
	device := app.Group("/Device")
	{
		handlers.DevcieHandler{}.Router(device)
	}

}
