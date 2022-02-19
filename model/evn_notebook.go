package model

import (
	"time"

	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	Title    string `json:"title" gorm:"comment:笔记本标题"`
	CreateBy string `json:"title" gorm:"column:create_by;comment:创建人"`
}
