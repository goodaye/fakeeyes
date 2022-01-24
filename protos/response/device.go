package response

import (
	"github.com/goodaye/fakeeyes/protos/request"
)

type DeviceInfo struct {
	request.DeviceInfo
	Status int
	Uptime int64
}
