package protos

// 设备平台
var DevicePlatform = struct {
	// Mac
	Darwin string
	// Linux
	Linux string
}{
	Darwin: "darwin",
	Linux:  "linux",
}

var APIName = struct {
	RegisterDevice string
}{
	RegisterDevice: "RegisterDevice",
}

// Httpheader key
var HeaderKey = struct {
	UserToken   string
	DeviceToken string
}{
	UserToken:   "UserToken",
	DeviceToken: "DeviceToken",
}
