package service

import (
	"evernote-client/global"
	"evernote-client/model"
)

//@function: RevertNote
//@description: 用户从废纸篓恢复笔记
//@param: nid uint, uid uint
//@return: err error
func RevertNote(nid uint, uid uint) (err error) {
	tx := global.SYS_DB.Begin()
	var note model.EvnNote

	// 恢复笔记判断原笔记本是否存在
	var count uint
	err = tx.Raw("SELECT count(1) FROM `evn_notebooks` a  LEFT JOIN evn_notes b ON a.id=b.notebook_id WHERE b.id=?", nid).Scan(&count).Error

	var notebookId uint
	if count == 0 {
		err = tx.Model(&model.EvnNotebook{}).Select("id").Where("create_by=? AND main=1", uid).First(&notebookId).Error
	}

	err = tx.Where("id = ? AND create_by = ? AND del_flag=1", nid, uid).First(&note).Updates(map[string]interface{}{"del_flag": 0, "notebook_id": notebookId}).Error
	if err == nil {
		err = tx.Exec("UPDATE evn_notebooks SET note_counts=(SELECT COUNT(1) FROM evn_notes WHERE notebook_id=(SELECT notebook_id FROM evn_notes WHERE id=? AND del_flag=0)) WHERE id=(SELECT notebook_id FROM evn_notes WHERE id=?)", nid, nid).Error
	}
	//回滚
	if err != nil {
		tx.Rollback()
	}
	return tx.Commit().Error
}

//@function: DeleteTrash
//@description: 用户彻底删除笔记
//@param: nid uint, uid uint
//@return: err error
func DeleteTrash(nid uint, uid uint) (err error) {
	var note model.EvnNote
	err = global.SYS_DB.Where("id = ? AND create_by = ?", nid, uid).Delete(&note).Error
	return err
}

//@function: GetTrashs
//@description: 获取所有废纸篓笔记
//@param: uid uint
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
