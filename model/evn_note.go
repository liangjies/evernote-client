package model

import (
	//"time"

	"gorm.io/gorm"
)

type EvnNote struct {
	gorm.Model
	Title      string `json:"title" gorm:"size:50;column:title;comment:笔记标题"`
	Content    string `json:"content" gorm:"column:content;type:text;comment:笔记内容"`
	NotebookId uint   `json:"notebookId" gorm:"column:notebook_id;comment:笔记所属笔记本"`
	CreateBy   uint   `json:"createBy" gorm:"index;column:create_by;comment:创建人"`
}
