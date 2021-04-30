package main

import (
	"fakeeyes/helloworld"
	"fakeeyes/videogate"

	"github.com/liangdas/mqant"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/nats-io/nats.go"
)

func main() {
	// rs := consul.NewRegistry(func(options *registry.Options) {
	// 	options.Addrs = []string{"127.0.0.1:8500"}
	// })

	nc, err := nats.Connect("nats://127.0.0.1:4222", nats.MaxReconnects(10000))
	if err != nil {
		log.Error("nats error %v", err)
		return
	}
	app := mqant.CreateApp(
		module.Debug(true), //只有是在调试模式下才会在控制台打印日志, 非调试模式下只在日志文件中输出日志
		module.Nats(nc),    //指定nats rpc
		// module.Registry(rs), //指定服务发现
	)
	err = app.Run( //模块都需要加到入口列表中传入框架
		helloworld.Module(),
		videogate.Module(),
	)
	if err != nil {
		log.Error(err.Error())
	}
}
