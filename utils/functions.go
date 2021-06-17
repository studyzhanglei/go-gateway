package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/studyzhanglei/grpc-proto/pb/common"
	"go-gateway/global"
	"go-gateway/model/request"
	"go-gateway/model/response"
	"go.uber.org/zap"
	"go-gateway/utils/errors"
)

func GetTraceLog(ctx *gin.Context) *zap.Logger {
	l := ctx.Value(REQUET_LOG_KEY)
	if l != nil {
		return l.(*zap.Logger)
	}

	return global.LOG
}

func GetGrpcHeader(ctx *gin.Context) *common.CommonHeader {

	var uid uint = 0
	if claims, exists := ctx.Get("claims"); exists {
		uid = claims.(request.MyClaims).Jti.Uid
	}

	header :=  &common.CommonHeader{
		TraceId: ctx.Value("trace-id").(string),
		OperatorId: uint64(uid),
	}

	return header
}

func ErrorHandle(ctx *gin.Context) {
	if err := recover();err != nil {
		fmt.Println("进入错误捕获处理")
		e, ok := err.(*errors.Error)
		if ok { //是自定义error
			response.Result(int(e.GetCode()), map[string]interface{}{}, e.Error(), ctx)
			return
		}

		if global.CONFIG.System.Env != "prod" {
			response.Result(response.ERROR, map[string]interface{}{}, fmt.Sprintf("%s", err), ctx)
			return
		}

		response.Result(response.ERROR, map[string]interface{}{}, "System error.", ctx)
		return
	}
}