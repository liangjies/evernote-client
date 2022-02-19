package model

import (
	"time"

	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	Title   string `json:"title" gorm:"comment:笔记标题"`
	Content string `json:"content" gorm:"comment:笔记内容"`
}
