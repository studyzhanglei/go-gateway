package global

import (
	"go-gateway/config"
	"github.com/go-redis/redis/v8"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	grpcpool "go-gateway/utils/grpc"
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
