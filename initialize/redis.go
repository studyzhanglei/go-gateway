package initialize

import (
	"go-gateway/global"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

func Redis() {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
		PoolSize: 20,
		PoolTimeout:  30 * time.Second,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MinIdleConns: 5, //最小连接数量
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.REDIS = client
	}
}
