package service

import (
	"evernote-client/global"
	"evernote-client/model"
)

//@function: RevertNote
//@description: 用户从废纸篓恢复笔记
//@param: id uint, uid uint
//@return: err error
func RevertNote(nid uint, uid uint) (err error) {
	var note model.EvnNote
	err = global.SYS_DB.Where("id = ? AND create_by = ? AND del_flag=1", nid, uid).First(&note).Update("del_flag", 0).Error
	return err
}

//@function: DeleteNotebook
//@description: 用户彻底删除笔记
//@param: id uint, uid uint
//@return: err error
func DeleteTrash(nid uint, uid uint) (err error) {
	var note model.EvnNote
	err = global.SYS_DB.Where("id = ? AND create_by = ?", nid, uid).Delete(&note).Error
	return err
}

//@function: GetNotebooks
//@description: 用户获取废纸篓
//@param: nid uint, uid uint
//@return: err error, list interface{}, total int64
func GetTrashs(uid uint) (err error, list interface{}, total int64) {
	var noteList []model.EvnNote
	db := global.SYS_DB.Model(&model.EvnNote{})
	err = db.Where("create_by = ? AND del_flag=1", uid).Find(&noteList).Error
	err = db.Count(&total).Error
	return err, noteList, total
}

//@function: GetTrashById
//@description: 用户根据id获取废纸篓笔记详情
//@param: nid uint, uid uint
//@return: err error, list interface{}, total int64
func GetTrashById(nid uint, uid uint) (err error, list interface{}) {
	var noteList []model.EvnNote
	db := global.SYS_DB.Model(&model.EvnNote{})
	err = db.Where("id = ? AND create_by = ? AND del_flag=1", nid, uid).Find(&noteList).Error
	return err, noteList
}
