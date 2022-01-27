package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// 默认的upgrade
var WSUpgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

//
var Logger *logrus.Logger

// Httpheader key
var HeaderKey = struct {
	UserToken   string
	DeviceToken string
}{
	UserToken:   "UserToken",
	DeviceToken: "DeviceToken",
}

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
