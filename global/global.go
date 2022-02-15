package initialize

import (
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	SYS_CONFIG  config.Server
	SYS_VIP     *viper.Viper
	SYS_ZAP_LOG *zap.Logger
)
