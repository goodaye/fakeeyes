package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/goodaye/fakeeyes/pkg/ginhandler"
)

type DevcieHandler struct {
	ginhandler.BaseHandler
}

//Router
func (h DevcieHandler) Router(rg *gin.RouterGroup) {
	device := rg.Group("/Device")
	{
		device.POST("/Login", h.Login)
	}
}

func (d DevcieHandler) Login(c *gin.Context) {

}

func (d DevcieHandler) HeartBeat() {

}

func (d DevcieHandler) Connect() {

}
