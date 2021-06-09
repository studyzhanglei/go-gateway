package router

import (
	"go-gateway/api/v1"
	"go-gateway/middleware"
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