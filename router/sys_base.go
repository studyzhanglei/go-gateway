package router

import (
	"github.com/gin-gonic/gin"
	"go-gateway/api/v1"
)


func InitCommonRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("")
	{
		BaseRouter.GET("/", v1.Index)
	}
	return BaseRouter
}

func InitGrpcRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("")
	{
		BaseRouter.GET("grpc2", v1.Grpc2)
		BaseRouter.GET("grpc3", v1.Grpc3)
	}
	return BaseRouter
}



func InitLoginRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("").Use()
	{
		BaseRouter.POST("login", v1.Login)
	}
	return BaseRouter
}

