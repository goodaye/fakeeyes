package rdb

import (
	"time"
)

// User user
type User struct {
	ID          int64     `json:"id" xorm:"not null  pk autoincr INT"`
	UID         string    `json:"uid" xorm:"not null  unique VARCHAR(255) comment('用户UID')"`
	Name        string    `json:"name" xorm:"not null unique VARCHAR(255) comment('用户名')"`
	LastLogin   time.Time `json:"last_login" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" `
	GmtCreated  time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `
	GmtModified time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
}

type UserSession struct {
	ID          int64     `json:"id" xorm:"not null pk autoincr INT"`
	UserID      int64     `json:"user_id" xorm:"not null BIGINT unique comment('用户ID')"`
	Token       string    `json:"token" xorm:"not null  VARCHAR(255) unique comment('用户token')"`
	ExpireTime  time.Time `json:"expire_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" `
	GmtCreated  time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `
	GmtModified time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
}

type DeviceSession struct {
	ID          int64     `json:"id" xorm:"not null pk autoincr INT"`
	UserID      int64     `json:"user_id" xorm:"not null BIGINT unique comment('用户ID')"`
	Token       string    `json:"token" xorm:"not null  VARCHAR(255) unique comment('用户token')"`
	ExpireTime  time.Time `json:"expire_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" `
	GmtCreated  time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `
	GmtModified time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
}

// Device
type Device struct {
	ID           int64     `json:"id" xorm:"not null pk autoincr INT"`
	UUID         string    `json:"uuid" xorm:" not null unique   VARCHAR(255) comment('设备唯一编码UUID')"`
	SN           string    `json:"sn" xorm:"not null unique VARCHAR(255) comment('设备SN')"`
	Name         string    `json:"name" xorm:"not null  VARCHAR(255) comment('设备名')"`
	ModelName    string    `json:"model_name" xorm:" null  VARCHAR(255) comment('设备名,比如树莓派/MacBook')"`
	ModelID      string    `json:"model_id" xorm:" null  VARCHAR(255) comment('设备名,比如树莓派4B/MacBook16,1')"`
	CPUModel     string    `json:"cpu_model" xorm:" null  VARCHAR(255) comment('处理器型号')"`
	CPUModelID   string    `json:"cpu_model_id" xorm:" null  VARCHAR(255) comment('处理器型号ID')"`
	CPUSpeed     string    `json:"cpu_speed" xorm:" null  VARCHAR(255) comment('处理器速度')"`
	CPUArch      string    `json:"cpu_arch" xorm:" null  VARCHAR(255) comment('处理器架构,x86, arm')"`
	CPUVendor    string    `json:"cpu_vendor" xorm:" null  VARCHAR(255) comment('处理器速度')"`
	CPUCores     int       `json:"cpu_cores" xorm:" null  INT comment('处理器核心数')"`
	CPUSocket    int       `json:"cpu_socket" xorm:" null  INT comment('处理器Socket个数')"`
	Manufacturer string    `json:"manufacturer" xorm:" null  VARCHAR(255) comment('设备制造商,比如,亚博')"`
	OSName       string    `json:"os_name" xorm:" null  VARCHAR(255) comment('OS名字')"`
	OSVersion    string    `json:"os_version" xorm:" null  VARCHAR(255) comment('OS Version')"`
	HardwareUUID string    `json:"hardware_uuid" xorm:" null  VARCHAR(255) comment('硬件UUID')"`
	Status       int       `json:"status" xorm:" null  tinyint comment('设备制状态,比如 在线/离线')"`
	Uptime       int64     `json:"up_time " xorm:" null int comment('设备启动时间(单位秒)')"`
	LastLogin    time.Time `json:"last_login" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP comment('上次登陆时间')" `
	GmtCreated   time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `
	GmtModified  time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
}

// Room 房间
type Room struct {
	ID          int64     `json:"id" xorm:"not null pk autoincr INT"`
	UUID        string    `json:"uuid" xorm:"not null unique  VARCHAR(255) unique comment('房间UUID')"`
	Name        string    `json:"name" xorm:"not null VARCHAR(255) comment('房间名')"`
	Status      int       `json:"status" xorm:" null  VARCHAR(255) comment('设备制状态')"`
	GmtCreated  time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `
	GmtModified time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
}

// //RoomUser 房间中的人
// type RoomUser struct {
// 	ID          int64 `json:"id" xorm:"not null pk autoincr INT"`
// 	RoomID      int64
// 	UserID      int64
// 	Role        int
// 	Status      int
// 	GmtCreated  time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `
// 	GmtModified time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
// }

//房间中的设备
type RoomDevice struct {
	ID          int64 `json:"id" xorm:"not null pk autoincr INT"`
	RoomID      int64
	DeviceID    int64
	GmtCreated  time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `
	GmtModified time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
}
