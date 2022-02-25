package service

import (
	"evernote-client/global"
	"evernote-client/model"
)

//@function: DeleteNotebook
//@description: 用户删除笔记本
//@param: id uint, uid uint
//@return: err error
func DeleteNote(nid uint, uid uint) (err error) {
	var note model.EvnNote
	err = global.SYS_DB.Where("id = ? AND create_by = ? AND del_flag=0", nid, uid).First(&note).Update("del_flag", 1).Error
	return err
}

//@function: UpdateNote
//@description: 用户修改笔记
//@param: n model.EvnNote, nid uint, uid uint
//@return: err error
func UpdateNote(n model.EvnNote, nid uint, uid uint) (err error) {
	var note model.EvnNote
	err = global.SYS_DB.Select("title", "content").Where("id = ? AND create_by = ? AND del_flag=0", nid, uid).First(&note).Updates(&n).Error
	return err
}

//@function: CreateNote
//@description: 用户新建笔记
//@param: n model.EvnNote, nid uint, uid uint
//@return: err error
func CreateNote(n model.EvnNote, nid uint, uid uint) (err error) {
	n.CreateBy = uid
	n.NotebookId = nid
	n.DelFlag = false
	err = global.SYS_DB.Create(&n).Error
	return err
}

//@function: GetNotebooks
//@description: 用户获取笔记列表
//@param: nid uint, uid uint
//@return: err error, list interface{}, total int64
func GetNotes(nid uint, uid uint) (err error, list interface{}, total int64) {
	var noteList []model.EvnNote
	db := global.SYS_DB.Model(&model.EvnNote{})
	err = db.Count(&total).Error
	err = db.Select("CreatedAt", "UpdatedAt", "ID", "Title", "NotebookId").Where("notebook_id = ? AND create_by = ? AND del_flag=0", nid, uid).Find(&noteList).Error
	return err, noteList, total
}

//@function: GetNotebooks
//@description: 用户根据id获取笔记详情
//@param: nid uint, uid uint
//@return: err error, list interface{}, total int64
func GetNoteById(nid uint, uid uint) (err error, list interface{}) {
	var noteList []model.EvnNote
	db := global.SYS_DB.Model(&model.EvnNote{})
	err = db.Where("id = ? AND create_by = ? AND del_flag=0", nid, uid).Find(&noteList).Error
	return err, noteList
}
