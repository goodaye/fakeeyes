package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goodaye/fakeeyes/protos"
	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/goodaye/fakeeyes/protos/response"
	"github.com/goodaye/fakeeyes/service"
	"github.com/gorilla/websocket"
)

type DevcieHandler struct {
	Handler
}

//Router
func (h DevcieHandler) Router(rg *gin.RouterGroup) {
	rg.POST("/RegisterDevice", h.Register)
	device := rg.Group("/Device", h.CheckLoginStatus)
	{
		device.POST("/SendHeartBeat", h.SendHeartBeat)
	}
	ws := rg.Group("/Device", WSUpgrade, h.WSCheckLoginStatus)
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
func (h DevcieHandler) _CheckLoginStatus(c *gin.Context) (errcode string, err error) {

	token := c.Request.Header.Get(protos.HeaderKey.DeviceToken)
	if token == "" {
		return HTTPErrorCode.RequestForbidben, fmt.Errorf("device need login")
	}
	dev, err := service.DeviceLoginByToken(token)
	if err != nil {
		return HTTPErrorCode.InvalidQueryParameter, err
	}
	// 设置已经登陆用户到Context中
	c.Set(ContextKey.LoginDevice, dev)
	return
}

func (h DevcieHandler) CheckLoginStatus(c *gin.Context) {

	code, err := h._CheckLoginStatus(c)
	if err != nil {
		h.SendFailure(c, code, err)
		return
	}
}

//WebSocket登陆状态检查
func (h DevcieHandler) WSCheckLoginStatus(c *gin.Context) {

	_, err := h._CheckLoginStatus(c)
	if err != nil {
		h.WSAbort(c, err)
		return
	}
}

// SendHeartBeat 不会修改token，只是保持心跳
func (h DevcieHandler) SendHeartBeat(c *gin.Context) {
	var err error
	dev := c.MustGet(ContextKey.LoginDevice).(*service.Device)
	var rs = request.DeviceInfo{}
	err = h.UnmarshalPost(c, &rs)
	if err != nil {
		return
	}
	err = dev.SendHeartBeat(rs)
	if err != nil {
		h.SendFailure(c, HTTPErrorCode.ProcessDataError, err)
		return
	}
}
func (h DevcieHandler) Connect(c *gin.Context) {
	var err error
	dev := c.MustGet(ContextKey.LoginDevice).(*service.Device)
	conn := c.MustGet(ContextKey.WSConnection).(*websocket.Conn)

	err = dev.Connect(conn)
	if err != nil {
		h.WSAbort(c, err)
	}
	err = conn.WriteMessage(websocket.TextMessage, []byte("connect to server success "))
	if err != nil {
		fmt.Println(err.Error())
	}
}
