package service

import (
	"fmt"
	"testing"

	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/goodaye/wire"
)

func init() {
	err := wire.Init()
	if err != nil {
		panic(err)
	}
}

func TestRegisterDevice(t *testing.T) {

	devinfo := request.DeviceInfo{
		SN:        "fakesn-xyz",
		Name:      "mockdevice",
		ModelName: "testmachine",
		ModelID:   "xyz--abc",
		CPUArch:   "x86_64",
	}
	result, err := RegisterDevice(devinfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}
