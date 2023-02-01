package global

import (
	"evernote-client/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG config.Server
	VIP    *viper.Viper
	LOG    *zap.Logger
	REDIS  *redis.Client
)
