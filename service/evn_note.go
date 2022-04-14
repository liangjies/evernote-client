package service

import (
	"evernote-client/global"
	"evernote-client/model"
	"evernote-client/utils"
)

//@function: DeleteNote
//@description: 用户删除笔记
//@param: nid uint, uid uint
//@return: err error
func DeleteNote(nid uint, uid uint) (err error) {
	tx := global.SYS_DB.Begin()
	var note model.EvnNote
	err = tx.Where("id = ? AND create_by = ? AND del_flag=0", nid, uid).First(&note).Update("del_flag", 1).Error
	if err == nil {
		err = tx.Exec("UPDATE evn_notebooks SET note_counts=(SELECT COUNT(1) FROM evn_notes WHERE notebook_id=(SELECT notebook_id FROM evn_notes WHERE id=?)) WHERE id=(SELECT notebook_id FROM evn_notes WHERE id=?)", nid, nid).Error
	}
	//回滚
	if err != nil {
		tx.Rollback()
	}
	return tx.Commit().Error
}

//@function: UpdateNote
//@description: 用户修改笔记
//@param: n model.EvnNote, uid uint
//@return: err error
func UpdateNote(n model.EvnNote, uid uint) (err error) {
	db := global.SYS_DB.Model(&model.EvnNote{})
	var note model.EvnNote

	// 生成笔记片段
	rs := []rune(utils.StripTags(n.Content))
	i := 40
	if len(rs) < 40 {
		i = len(rs)
	}
	n.Snippet = string(rs[:i])

	// 查询是否没有修改
	var count uint
	err = db.Raw("SELECT count(id) FROM `evn_notes` WHERE id=? AND content=?", n.ID, n.Content).Scan(&count).Error

	if n.NotebookId != 0 {
		err = db.Select("title", "content", "notebook_id", "snippet").Where("id = ? AND create_by = ? AND del_flag=0", n.ID, uid).First(&note).Updates(&n).Error
	} else {
		err = db.Select("content").Where("id = ? AND create_by = ? AND del_flag=0", n.ID, uid).First(&note).Updates(&n).Error
	}
	// 保存历史记录
	if count == 0 {
		var version uint
		var versionMin uint
		var versionCount uint
		row := db.Raw("SELECT IFNULL(max(version),0),count(1),IFNULL(min(version),0) FROM `evn_histories`  WHERE note_id=?", n.ID).Row()
		row.Scan(&version, &versionCount, &versionMin)

		history := model.EvnHistory{NoteId: n.ID, Content: n.Content, Version: version + 1}
		err = global.SYS_DB.Model(&model.EvnHistory{}).Create(&history).Error

		// 历史记录保存限制
		var VersionMax uint = 20
		if global.SYS_CONFIG.System.VersionMax != 0 {
			VersionMax = global.SYS_CONFIG.System.VersionMax
		}
		if versionCount > VersionMax {
			err = global.SYS_DB.Model(&model.EvnHistory{}).Where("note_id=? AND version=?", n.ID, versionMin).Delete(&model.EvnHistory{}).Error
		}
	}

	return err
}

//@function: CreateNote
//@description: 用户新建笔记
//@param: n model.EvnNote, uid uint
//@return: id uint, err error
func CreateNote(n model.EvnNote, uid uint) (id uint, err error) {
	tx := global.SYS_DB.Begin()
	n.CreateBy = uid

	// 生成笔记片段
	rs := []rune(utils.StripTags(n.Content))
	i := 40
	if len(rs) < 40 {
		i = len(rs)
	}
	n.Snippet = string(rs[:i])

	// 如果NotebookId为0，默认笔记本
	if n.NotebookId == 0 {
		err = tx.Model(&model.EvnNotebook{}).Select("id").Where("create_by=? AND main=1", uid).First(&n.NotebookId).Error
	}

	err = tx.Create(&n).Error
	if err == nil {
		err = tx.Exec("UPDATE evn_notebooks SET note_counts=(SELECT COUNT(1) FROM evn_notes WHERE notebook_id=? AND del_flag=0) WHERE id=?", n.NotebookId, n.NotebookId).Error
	}
	//回滚
	if err != nil {
		tx.Rollback()
	}

	return n.ID, tx.Commit().Error
}

//@function: GetNotes
//@description: 获取笔记本笔记列表
//@param: nid uint, uid uint
//@return: err error, list interface{}, total int64
func GetNotes(nid uint, uid uint) (err error, list interface{}, total int64, title string) {
	var noteList []model.EvnNote
	db := global.SYS_DB.Model(&model.EvnNote{})
	err = db.Select("CreatedAt", "UpdatedAt", "ID", "Title", "NotebookId").Where("notebook_id = ? AND create_by = ? AND del_flag=0", nid, uid).Order("updated_at desc").Find(&noteList).Error
	err = db.Count(&total).Error
	err = global.SYS_DB.Model(&model.EvnNotebook{}).Select("Title").Where("id = ? AND create_by = ?", nid, uid).First(&title).Error
	return err, noteList, total, title
}

//@function: GetNoteById
//@description: 用户根据id获取笔记详情
//@param: nid uint, uid uint
//@return: err error, list interface{}
func GetNoteById(nid uint, uid uint) (err error, list interface{}) {
	var noteList []model.EvnNote
	db := global.SYS_DB.Model(&model.EvnNote{})
	err = db.Where("id = ? AND create_by = ? AND del_flag=0", nid, uid).Find(&noteList).Error
	return err, noteList
}

//@function: GetAllNotes
//@description: 获取所有笔记
//@param: uid uint
//@return: err error, list interface{}, total int64
func GetAllNotes(uid uint) (err error, list interface{}, total int64) {
	var noteList []model.EvnNote
	db := global.SYS_DB.Model(&model.EvnNote{})
	err = db.Select("CreatedAt", "UpdatedAt", "ID", "Title", "Snippet").Where("create_by = ? AND del_flag=0", uid).Order("updated_at desc").Find(&noteList).Error
	err = db.Count(&total).Error
	return err, noteList, total
}

//@function: SearchNote
//@description: 搜索笔记
//@param: SearchKey string, NotebookId uint, uid uint
//@return: err error, list interface{}, total int64
func SearchNote(SearchKey string, NotebookId uint, uid uint) (err error, list interface{}, total int64) {
	var noteList []model.EvnNote
	db := global.SYS_DB.Model(&model.EvnNote{})
	db.Where("del_flag = 0")
	if NotebookId != 0 {
		db.Where("notebook_id = ?", NotebookId)
	}

	if SearchKey != "" {
		db.Where("title like ? OR content like ?", "%"+SearchKey+"%", "%"+SearchKey+"%")
	}

	err = db.Select("CreatedAt", "UpdatedAt", "ID", "Title", "Snippet").Order("updated_at desc").Find(&noteList).Error
	err = db.Count(&total).Error
	return err, noteList, total
}
