package utils

import (
	"go-gateway/global"
	"github.com/gin-gonic/gin"
	"go-gateway/model/request"
	"go-gateway/pb/common"
	"go.uber.org/zap"
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