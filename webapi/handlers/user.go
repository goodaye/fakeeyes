package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goodaye/fakeeyes/pkg/copy"
	"github.com/goodaye/fakeeyes/pkg/ginhandler"
	"github.com/goodaye/fakeeyes/protos"
	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/goodaye/fakeeyes/protos/response"
	"github.com/goodaye/fakeeyes/service"

	"github.com/gorilla/websocket"
)

// User
type UserHandler struct {
	ginhandler.BaseHandler
}

//Router
func (h UserHandler) Router(rg *gin.RouterGroup) {
	// 用户操作
	rg.POST("/UserLogin", h.Login)
	rg.POST("/UserSignIn", h.SignIn)
	rg.POST("/UserSignUp", h.SignUp)
	user := rg.Group("/User", h.CheckLoginStatus)
	{
		user.POST("/ListDevices", h.ListDevices)
	}
	// ws := rg.Group("/User", h.CheckLoginStatus, WSUpgrade)
	ws := rg.Group("/User", WSUpgrade, h.WSCheckLoginStatus)
	{
		// WebSocket
		ws.GET("/ConnectDevice", h.ConnectDevice)
	}

}

func (h UserHandler) WSAbort(c *gin.Context, aborterr error) {

	conn := c.MustGet(ContextKey.WSConnection).(*websocket.Conn)
	conn.WriteMessage(websocket.TextMessage, []byte(aborterr.Error()))
	conn.Close()
	c.Abort()
}

// Login
func (h UserHandler) Login(c *gin.Context) {

	var err error
	var rs = request.UserLogin{}
	err = h.UnmarshalPost(c, &rs)
	if err != nil {
		return
	}
	u, err := service.Login(rs)
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

// 登陆
func (h UserHandler) SignIn(c *gin.Context) {
	h.Login(c)
}

// 登出
func (h UserHandler) SignOut(c *gin.Context) {

}

// 注册
func (h UserHandler) SignUp(c *gin.Context) {
	var err error
	var rs = request.UserSignUp{}
	err = h.UnmarshalPost(c, &rs)
	if err != nil {
		return
	}
	u, err := service.UserSignUp(rs)
	if err != nil {
		h.SendFailure(c, HTTPErrorCode.ProcessDataError, err)
		return
	}
	h.SendSuccess(c, u)

}

//登陆状态检查
func (h UserHandler) CheckLoginStatus(c *gin.Context) {

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

//WebSocket登陆状态检查
func (h UserHandler) WSCheckLoginStatus(c *gin.Context) {

	conn := c.MustGet(ContextKey.WSConnection).(*websocket.Conn)
	token := c.Request.Header.Get(protos.HeaderKey.UserToken)
	if token == "" {
		h.WSAbort(c, fmt.Errorf("user need login"))
		// conn.WriteMessage(websocket.CloseMessage, []byte("user need login"))
		// c.Abort()
		return
	}
	user, err := service.LoginByToken(token)
	if err != nil {
		// h.SendFailure(c, HTTPErrorCode.InvalidQueryParameter, err)
		conn.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
		c.Abort()
		return
	}
	// 设置已经登陆用户到Context中
	c.Set(ContextKey.LoginUser, user)
}

// ListDevices
func (h UserHandler) ListDevices(c *gin.Context) {
	user := c.MustGet(ContextKey.LoginUser).(*service.User)
	devs, err := user.ListDevices()
	if err != nil {
		h.SendFailure(c, HTTPErrorCode.ProcessDataError, err)
		return
	}
	var resp = []response.DeviceInfo{}
	copy.StructSliceCopy(devs, &resp)
	h.SendSuccess(c, resp)
}

// ConnectDevice connect to device
func (h UserHandler) ConnectDevice(c *gin.Context) {

	var err error
	u := c.MustGet(ContextKey.LoginUser).(*service.User)
	conn := c.MustGet(ContextKey.WSConnection).(*websocket.Conn)
	var req request.ConnectDevice
	err = c.Bind(&req)
	if err != nil {
		h.WSAbort(c, err)
	}
	if req.DeviceUUID == "" {
		h.WSAbort(c, fmt.Errorf("need argument device_uuid"))
		return
	}

	// _, err = room.CreateRoom(user, conn, req.DeviceUUID)
	// if err != nil {
	// 	h.WSAbort(c, err)
	// }
	_, err = u.ConnectDevice(req, conn)
	if err != nil {
		h.WSAbort(c, err)
	}
	conn.WriteMessage(websocket.TextMessage, []byte("connect to server success "))
}
