package model

import (
	"evernote-client/global"
)

type EvnHistory struct {
	global.SYS_MODEL
	Content string `json:"content" gorm:"column:content;type:text;comment:笔记内容"`
	NoteId  uint   `json:"noteId" gorm:"index;column:note_id;comment:笔记ID"`
	Version uint   `json:"version" gorm:"column:version;comment:版本"`
}
