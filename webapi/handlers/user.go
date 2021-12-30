package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/goodaye/fakeeyes/pkg/ginhandler"
	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/goodaye/fakeeyes/service/user"
)

// User
type UserHandler struct {
	ginhandler.BaseHandler
}

//Router
func (h UserHandler) Router(rg *gin.RouterGroup) {
	user := rg.Group("/User")
	{

		user.POST("/Login", h.Login)
		user.POST("/SignIn", h.SignIn)
		user.POST("/SignUp", h.SignUp)
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
		h.SendFailure(c, "LoginError", err)
		return
	}
	h.SendSuccess(c, u)
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

}
