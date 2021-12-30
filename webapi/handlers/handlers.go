package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/goodaye/fakeeyes/pkg/ginhandler"
)

type Handler struct {
	ginhandler.BaseHandler
}

func (h Handler) Pong(c *gin.Context) {

	h.SendSuccess(c, "Pong")
}

func (h Handler) Version(c *gin.Context) {

	h.SendSuccess(c, "v1")
}
