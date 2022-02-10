package service

import (
	"github.com/gorilla/websocket"
)

// 房间链接池
var RoomPool map[string]*Room

// 房间连接池大小
var poolsize = 100

// 设备链接池管理，
// key :device UUID
// value: device connection
var DeviceConns = map[string]*websocket.Conn{}

func init() {

	// var err error
}
