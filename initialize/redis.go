package initialize

import (
	"context"
	"evernote-client/global"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisConfig := global.CONFIG.Redis
	global.REDIS = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := global.REDIS.Ping(ctx).Result()
	if err != nil {
		global.LOG.Error("Redis连接失败", zap.Error(err))
		panic("Redis连接失败")
	}

	global.LOG.Info("Redis连接成功")
}
