package model

import (
	"time"

	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`    // 用户UUID
	Username  string    `json:"userName" gorm:"comment:用户登录名"` // 用户登录名
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`      // 用户登录密码
	LoginTime time.Time // 上次登录时间
}
