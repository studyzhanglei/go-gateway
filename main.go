package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
)


func main() {
	global.VIPER = core.Viper()      // 初始化Viper
	global.LOG = core.Zap()       	 // 初始化zap日志库
	global.DB = initialize.Gorm() 	 // gorm连接数据库
	global.GRPC_POOL   = core.GetGrpcClient("1234") // 初始化GRPC连接

	core.RunWindowsServer()
}
