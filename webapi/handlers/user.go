package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/goodaye/fakeeyes/pkg/ginhandler"
)

// User
type User struct {
	ginhandler.BaseHandler
}

// Login
func (u User) Login(c *gin.Context) {

}
