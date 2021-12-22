package webapi

import (
	"fmt"
	"testing"

	"github.com/goodaye/wire"
)

func init() {
	// err := wire.Init()
	// if err != nil {
	// 	panic(err)
	// }
}
func TestStartWeb(*testing.T) {

	wire.Append(SVC{})
	err := wire.Run()
	if err != nil {
		fmt.Println(err)
	}
}
