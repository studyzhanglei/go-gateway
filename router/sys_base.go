package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)


func InitGrpcRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("").Use(middleware.NeedInit())
	{
		BaseRouter.Any("grpc2", v1.Grpc2)
		BaseRouter.Any("grpc3", v1.Grpc3)
	}
	return BaseRouter
}