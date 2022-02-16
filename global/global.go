package global

import (
	"evernote-client/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	SYS_DB                  *gorm.DB
	SYS_CONFIG              config.Server
	SYS_VIP                 *viper.Viper
	SYS_LOG                 *zap.Logger
	SYS_REDIS               *redis.Client
	SYS_Concurrency_Control = &singleflight.Group{}
)
