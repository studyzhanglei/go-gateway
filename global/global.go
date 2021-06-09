package global

import (
	"gin-vue-admin/config"
	"github.com/go-redis/redis"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	grpcpool "gin-vue-admin/utils/grpc"
)


var (
	DB     *gorm.DB
	REDISPOOL  *redigo.Pool
	REDIS *redis.Client
	CONFIG config.Server
	VIPER     *viper.Viper
	LOG   *zap.Logger
	GRPC_POOL *grpcpool.Pool
)
