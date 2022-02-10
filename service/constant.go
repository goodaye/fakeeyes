package service

import "time"

// 用户token 过期时间
var UserTokenExpireDuration = 1 * time.Hour

var DeviceState = struct {
	Online  int
	Offline int
}{
	Online:  1,
	Offline: 2,
}
