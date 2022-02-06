package handlers

import (
	"github.com/sirupsen/logrus"
)

//
var Logger *logrus.Logger

// 上下文放的key
var ContextKey = struct {
	LoginUser    string
	LoginDevice  string
	WSConnection string
}{
	LoginUser:    "LoginUser",
	LoginDevice:  "LoginDevice",
	WSConnection: "WSConnection",
}
