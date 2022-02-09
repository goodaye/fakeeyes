package service

import (
	"fmt"

	"github.com/goodaye/fakeeyes/protos/command"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

// 房间模式
type ModeType int

// 模式列表
var Modes = struct {
	Normal ModeType
	Echo   ModeType
}{
	Normal: 1,
	Echo:   2,
}

// 房间信息
type Room struct {
	// 房间名
	Name string
	// 链接的用户
	User *User
	// 链接的设备
	Device *Device
	// 客户端链接
	ClientConn *websocket.Conn
	// 设备链接
	DeviceConn *websocket.Conn
	DeviceIn   chan []byte
	DeviceOut  chan []byte
	ClientIn   chan []byte
	ClientOut  chan []byte
	Mode       ModeType
}

func CreateRoom(user *User, client_conn *websocket.Conn, device_sn string) (room *Room, err error) {

	// devcie, err := DescribeDevice(request.DescribeDevice{SN: device_sn})
	// if err != nil {
	// 	return
	// }

	// device_conn, ok := DeviceConnection[device_sn]
	// if !ok {
	// 	return nil, fmt.Errorf("device lost connection")
	// }

	room = &Room{
		Name:       fmt.Sprintf("%s's Room", user.Name),
		User:       user,
		Device:     nil,
		ClientConn: client_conn,
		DeviceConn: nil,
		DeviceIn:   make(chan []byte),
		DeviceOut:  make(chan []byte),
		ClientIn:   make(chan []byte),
		ClientOut:  make(chan []byte),
		Mode:       Modes.Echo,
	}
	return
}

func (r *Room) Run() {
	r.StreamON()
	go r.StartCrossMatrix()
}

func (r *Room) Close() {

}

func (r *Room) StreamON() {

	// clientin
	go func() {
		defer func() {
			r.ClientConn.Close()
		}()
		for {
			mt, message, err := r.ClientConn.ReadMessage()

			if err != nil {
				r.ClientConn.Close()
				break
			}
			if mt == websocket.TextMessage {
				fmt.Println(string(message))
			} else if mt == websocket.BinaryMessage {
				r.ClientIn <- message
			}

		}

	}()
	// client out
	go func() {
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
	}()

	// device out
	go func() {
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
	}()

	// device in
	go func() {
		if r.DeviceConn == nil {
			return
		}
		defer func() {
			r.DeviceConn.Close()
		}()

		for {

			_, message, err := r.DeviceConn.ReadMessage()
			if err != nil {
				r.DeviceConn.Close()
				break
			}
			r.DeviceIn <- message
		}
	}()

}

func (r *Room) StartCrossMatrix() {

	var message []byte
	for {
		select {
		case message = <-r.ClientIn:
			var err error
			cmd := command.Command{}
			err = proto.Unmarshal(message, &cmd)
			if err != nil {
				return
			}
			// 如果是设备命令，转发到设备出口
			if cmd.Type != command.Command_Device {
				break
			}
			switch r.Mode {
			case Modes.Normal:
				r.DeviceOut <- message
			case Modes.Echo:
				r.ClientOut <- message
			}
		case message = <-r.DeviceIn:
			r.ClientOut <- message
		}
	}
}
