package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// 默认的upgrade
var WSUpgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
