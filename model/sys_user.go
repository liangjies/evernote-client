package model

import (
	"gin-vue-admin/global"

	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	Username  string    `json:"userName" gorm:"comment:用户登录名"` // 用户登录名
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`      // 用户登录密码
	LoginTime time.Time // 上次登录时间
}
