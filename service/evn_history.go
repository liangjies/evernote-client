package service

import (
	"evernote-client/global"
	"evernote-client/model"
)

//@function: GetHistories
//@description: 获取笔记历史版本
//@param: nid uint, uid uint
//@return: err error, list interface{}, total int64
func GetHistories(nid uint, uid uint) (err error, list interface{}, total int64) {
	var historyList []model.EvnHistory
	db := global.SYS_DB.Model(&model.EvnHistory{})
	// 验证
	var count uint
	err = global.SYS_DB.Model(&model.EvnNote{}).Select("count(1)").Where("id = ? AND create_by = ? AND del_flag=0", nid, uid).Scan(&count).Error

	if count == 0 {
		return err, historyList, total
	}

	err = db.Select("CreatedAt", "ID", "Version", "Content", "NoteId").Where("note_id = ?", nid).Order("Version desc").Find(&historyList).Error
	err = db.Count(&total).Error
	return err, historyList, total
}

//@function: RecoverHistory
//@description: 恢复笔记历史版本
//@param: vnHistory model.EvnHistory, uid uint
//@return: err error, list interface{}, total int64
func RecoverHistory(evnHistory model.EvnHistory, uid uint) (err error) {
	db := global.SYS_DB.Model(&model.EvnHistory{})
	var evnNote model.EvnNote
	evnNote.ID = evnHistory.NoteId
	err = db.Select("Content").Where("id=?", evnHistory.ID).Scan(&evnNote.Content).Error

	err = UpdateNote(evnNote, uid)
	return err
}
