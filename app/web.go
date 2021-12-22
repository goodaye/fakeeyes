package main

import (
	"github.com/goodaye/fakeeyes/webapi"
	"github.com/goodaye/wire"
)

// StartWorker start worker
func StartWeb() error {
	wire.Append(webapi.SVC{})
	return wire.Run()
}
