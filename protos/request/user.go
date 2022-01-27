package request

// 用户登陆
type UserLogin struct {
	Name string `json:"name"`
}

// 用户登陆
type UserSignIn struct {
	Name string `json:"name"`
}

// 用户注册
type UserSignUp struct {
	Name string `json:"name"`
}

//用户登陆检查头
type UserLoginHeader struct {
	Token string `json:"token"`
}

// 列出Room列表
type ListRooms struct {
	PageRequest
}

//
type ListDevices struct {
	PageRequest
}

//链接设备
type ConnectDevice struct {
	// UUID 设备的UUID
	DeviceUUID string `json:"device_uuid" form:"device_uuid"`
}
