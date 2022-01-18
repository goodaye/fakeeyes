package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goodaye/fakeeyes/pkg/ginhandler"
	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/goodaye/fakeeyes/protos/response"
	"github.com/goodaye/fakeeyes/service/user"
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
		user.POST("/ListDevices", h.SignUp)
		user.POST("/ConnectDevice", h.SignUp)
	}

}

// Login
func (h UserHandler) Login(c *gin.Context) {

	var err error
	var rs = request.UserLogin{}
	err = h.UnmarshalPost(c, &rs)
	if err != nil {
		return
	}
	u, err := user.Login(rs)
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
	u, err := user.UserSignUp(rs)
	if err != nil {
		h.SendFailure(c, HTTPErrorCode.ProcessDataError, err)
		return
	}
	h.SendSuccess(c, u)

}

//登陆状态检查
func (h UserHandler) CheckLoginStatus(c *gin.Context) {

	token := c.Request.Header.Get(HeaderKey.UserToken)
	if token == "" {
		h.SendFailure(c, HTTPErrorCode.RequestForbidben, fmt.Errorf("user need login"))
		return
	}
	user, err := user.LoginByToken(token)
	if err != nil {
		h.SendFailure(c, HTTPErrorCode.InvalidQueryParameter, err)
		return
	}
	// 设置已经登陆用户到Context中
	c.Set(ContextKey.LoginUser, user)
}

// ListDevices
func (h UserHandler) ListDevices(c *gin.Context) {
	user := c.MustGet(ContextKey.LoginUser).(user.User)
	devs, err := user.ListDevices()
	if err != nil {
		h.SendFailure(c, HTTPErrorCode.ProcessDataError, err)
		return
	}

	h.SendSuccess(c, devs)

}
