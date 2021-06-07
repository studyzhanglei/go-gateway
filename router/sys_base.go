package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base").Use(middleware.NeedInit())
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("captcha", v1.Captcha)
	}
	return BaseRouter
}

func InitGrpcRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("").Use(middleware.NeedInit())
	{
		BaseRouter.Any("grpc", v1.Grpc)
		BaseRouter.Any("grpc2", v1.Grpc2)
		BaseRouter.Any("grpc3", v1.Grpc3)
	}
	return BaseRouter
}