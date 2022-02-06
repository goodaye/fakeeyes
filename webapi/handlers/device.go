package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goodaye/fakeeyes/pkg/ginhandler"
	"github.com/goodaye/fakeeyes/protos"
	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/goodaye/fakeeyes/protos/response"
	"github.com/goodaye/fakeeyes/service"
)

type DevcieHandler struct {
	ginhandler.BaseHandler
}

//Router
func (h DevcieHandler) Router(rg *gin.RouterGroup) {
	rg.POST("/RegisterDevice", h.Register)
	device := rg.Group("/Device", h.CheckLoginStatus)
	{
		device.POST("/SendHeartBeat", h.Register)
	}
	ws := rg.Group("/Device", h.CheckLoginStatus)
	{
		ws.GET("Connect", h.Connect)
	}
}

func (h DevcieHandler) Register(c *gin.Context) {
	var err error
	var rs = request.DeviceInfo{}
	err = h.UnmarshalPost(c, &rs)
	if err != nil {
		return
	}
	u, err := service.RegisterDevice(rs)
	if err != nil {
		h.SendFailure(c, HTTPErrorCode.RequestForbidben, err)
		return
	}
	resp := response.UserLogin{
		Name:       u.Name,
		LastLogin:  u.LastLogin,
		Token:      u.Token,
		ExpireTime: u.ExpireTime,
	}
	h.SendSuccess(c, resp)

}

//登陆状态检查
func (h DevcieHandler) CheckLoginStatus(c *gin.Context) {

	token := c.Request.Header.Get(protos.HeaderKey.UserToken)
	if token == "" {
		h.SendFailure(c, HTTPErrorCode.RequestForbidben, fmt.Errorf("user need login"))
		return
	}
	user, err := service.LoginByToken(token)
	if err != nil {
		h.SendFailure(c, HTTPErrorCode.InvalidQueryParameter, err)
		return
	}
	// 设置已经登陆用户到Context中
	c.Set(ContextKey.LoginUser, user)
}

func (h DevcieHandler) SendHeartBeat(c *gin.Context) {

}

func (h DevcieHandler) Connect(c *gin.Context) {

}
