package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/goodaye/fakeeyes/pkg/ginhandler"
	"github.com/gorilla/websocket"
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
func (h Handler) WSAbort(c *gin.Context, aborterr error) {

	conn := c.MustGet(ContextKey.WSConnection).(*websocket.Conn)
	conn.WriteMessage(websocket.TextMessage, []byte(aborterr.Error()))
	conn.Close()
	c.Abort()
}
