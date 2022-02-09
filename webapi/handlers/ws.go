package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 默认的upgrade
var WSUpgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

//升级Connection
func WSUpgrade(c *gin.Context) {
	conn, err := WSUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		Logger.Error(err)
		c.Abort()
		return
	}
	conn.WriteMessage(websocket.TextMessage, []byte("connecting to fakeeyes"))
	c.Set(ContextKey.WSConnection, conn)
}
