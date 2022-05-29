package initialize

import (
	"evernote-client/global"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.SYS_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.SYS_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.SYS_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.SYS_REDIS = client
	}
}
