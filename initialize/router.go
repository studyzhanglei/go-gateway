package initialize

import (
	"github.com/gin-gonic/gin"
	"go-gateway/global"
	"go-gateway/middleware"
	"go-gateway/router"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Base()) //公共中间件
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitGrpcRouter(PublicGroup) // GRPC test
	}


	global.LOG.Info("router register success")
	return Router
}
