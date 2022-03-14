package model

import (
	"evernote-client/global"
)

type EvnNotebook struct {
	global.SYS_MODEL
	Title      string `json:"title" gorm:"size:50;column:title;comment:笔记本标题"`
	CreateBy   uint   `json:"_" gorm:"index;column:create_by;comment:创建人"`
	NoteCounts uint   `json:"noteCounts" gorm:"column:note_counts;comment:笔记数量"`
}
