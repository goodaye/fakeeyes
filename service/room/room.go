package room

import (
	"fmt"

	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/goodaye/fakeeyes/service/device"
	"github.com/goodaye/fakeeyes/service/user"
	"github.com/gorilla/websocket"
)

// 房间信息
type Room struct {
	// 房间名
	Name string
	// 链接的用户
	User *user.User
	// 链接的设备
	Device *device.Device
	// 客户端链接
	ClientConn *websocket.Conn
	// 设备链接
	DeviceConn *websocket.Conn
	DeviceIn   chan []byte
	DeviceOut  chan []byte
	ClientIn   chan []byte
	ClientOut  chan []byte
}

func CreateRoom(user *user.User, client_conn *websocket.Conn, device_sn string) (room *Room, err error) {

	devcie, err := device.DescribeDevice(request.DescribeDevice{SN: device_sn})
	if err != nil {
		return
	}

	device_conn, ok := DeviceConnection[device_sn]
	if !ok {
		return nil, fmt.Errorf("device lost connection")
	}

	room = &Room{
		Name:       fmt.Sprintf("%s's Room", user.Name),
		User:       user,
		Device:     devcie,
		ClientConn: client_conn,
		DeviceConn: device_conn,
		DeviceIn:   make(chan []byte),
		DeviceOut:  make(chan []byte),
		ClientIn:   make(chan []byte),
		ClientOut:  make(chan []byte),
	}
	return
}

func (r *Room) Run() {
	go r.client_in()
	go r.client_out()
	go r.device_out()
	go r.StartCrossMatrix()
}

func (r *Room) Close() {

}

func (r *Room) client_in() {
	defer func() {
		r.ClientConn.Close()
	}()

	for {
		_, message, err := r.ClientConn.ReadMessage()
		if err != nil {
			r.ClientConn.Close()
			break
		}
		r.ClientIn <- message
	}
}
func (r *Room) client_out() {
	defer func() {
		r.ClientConn.Close()
	}()
	for {
		message, ok := <-r.ClientOut
		if !ok {
			r.ClientConn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		r.ClientConn.WriteMessage(websocket.TextMessage, message)
	}
}

func (r *Room) device_out() {
	defer func() {
		r.DeviceConn.Close()
	}()
	for {
		message, ok := <-r.DeviceOut
		if !ok {
			r.DeviceConn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		r.DeviceConn.WriteMessage(websocket.TextMessage, message)
	}
}

func (r *Room) StartCrossMatrix() {

	var message []byte
	for {
		select {
		case message = <-r.ClientIn:
			r.DeviceOut <- message
		case message = <-r.DeviceIn:
			r.ClientOut <- message
		}
	}

}
