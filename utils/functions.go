package utils

import (
	"gin-vue-admin/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetTraceLog(ctx *gin.Context) *zap.Logger {
	l := ctx.Value(REQUET_LOG_KEY)
	if l != nil {
		return l.(*zap.Logger)
	}

	return global.LOG
}

