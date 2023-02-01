package model

import (
	"evernote-client/global"
)

type EvnJwtBlacklist struct {
	global.SYS_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
