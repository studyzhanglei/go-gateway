package global

import (
	"gin-vue-admin/config"
	"gin-vue-admin/utils/timer"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	grpcpool "gin-vue-admin/utils/grpc"
)


var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG   *zap.Logger
	RequestLog   *zap.Logger
	GVA_Timer timer.Timer = timer.NewTimerTask()
	GRPC  *grpc.ClientConn
	GRPC_POOL *grpcpool.Pool
	NUM int8
)
