package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/goodaye/fakeeyes/service"
)

// Admin
type AdminHandler struct {
	Handler
}

//Router
func (h AdminHandler) Router(rg *gin.RouterGroup) {
	// 用户操作
	user := rg.Group("/Admin")
	{
		user.POST("/ListOnlineDevices", h.ListOnlineDevices)
		user.GET("/ListOnlineDevices", h.ListOnlineDevices)
	}

}

func (h AdminHandler) ListOnlineDevices(c *gin.Context) {

	resp := []string{}
	for key := range service.DeviceConns {
		resp = append(resp, key)
	}
	h.SendSuccess(c, resp)
}
