package core

import (
	"fmt"
	"go-gateway/global"
	"go-gateway/initialize"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()


	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)

	fmt.Println(address, 9999)
	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))

	fmt.Println(88888888)

	global.LOG.Error(s.ListenAndServe().Error())
}
