package db

import "time"

// User user
type User struct {
	ID          int64
	Name        string
	LastLogin   time.Time
	GmtCreated  time.Time
	GmtModified time.Time
}

// Device
type Device struct {
	ID           int64
	Name         string
	SN           string
	Type         string
	Class        string
	Manufacturer string
	Status       int
	LastLogin    time.Time
	GmtCreated   time.Time
	GmtModified  time.Time
}

// Room 房间
type Room struct {
	ID          int64
	UUID        string
	Name        string
	Status      int
	GmtCreated  time.Time
	GmtModified time.Time
}

//RoomUser 房间中的人
type RoomUser struct {
	ID          int64
	RoomID      int64
	UserID      int64
	Role        int
	Status      int
	GmtCreated  time.Time
	GmtModified time.Time
}

//房间中的设备
type RoomDevice struct {
	ID          int64
	RoomID      int64
	DeviceID    int64
	GmtCreated  time.Time
	GmtModified time.Time
}
