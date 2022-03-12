package model

import (
	"time"

	"github.com/satori/go.uuid"
	"evernote-client/global"
)

type SysUser struct {
	global.SYS_MODEL
	UUID      uuid.UUID `json:"-" gorm:"column:uuid;comment:用户UUID"`          // 用户UUID
	Username  string    `json:"userName" gorm:"column:username;comment:用户登录名"`   // 用户登录名
	NickName  string    `json:"nickName" gorm:"column:nickname;comment:用户昵称"`   // 用户昵称
	HeaderImg  string    `json:"headerImg" gorm:"column:headerimg;comment:用户昵称"`   // 用户头像
	Password  string    `json:"-"  gorm:"column:password;comment:用户登录密码"` // 用户登录密码
	LoginTime time.Time `json:"-"  gorm:"comment:用户上次登录时间"`                      // 上次登录时间
}
