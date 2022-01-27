package handlers

import (
	"github.com/gin-gonic/gin"
)

//升级Connection
func WSUpgrade(c *gin.Context) {
	conn, err := WSUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		Logger.Error(err)
		return
	}
	c.Set(ContextKey.WSConnection, conn)
}
