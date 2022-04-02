package service

import (
	"evernote-client/global"
	"evernote-client/model"
	"evernote-client/utils"
)

//@function: DeleteNotebook
//@description: 用户删除笔记本
//@param: id uint, uid uint
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
//@param: n model.EvnNote, nid uint, uid uint
//@return: err error
func UpdateNote(n model.EvnNote, uid uint) (err error) {
	var note model.EvnNote
	rs := []rune(utils.StripTags(n.Content))
	i := 40
	if len(rs) < 40 {
		i = len(rs)
	}
	n.Snippet = string(rs[:i])
	err = global.SYS_DB.Select("title", "content", "notebook_id", "snippet").Where("id = ? AND create_by = ? AND del_flag=0", n.ID, uid).First(&note).Updates(&n).Error
	return err
}

//@function: CreateNote
//@description: 用户新建笔记
//@param: n model.EvnNote, nid uint, uid uint
//@return: err error
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

//@function: 获取笔记本笔记列表
//@description: 用户获取笔记列表
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

// @Summary 用户获取笔记列表
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notes [get]
func GetAllNotes(uid uint) (err error, list interface{}, total int64) {
	var noteList []model.EvnNote
	db := global.SYS_DB.Model(&model.EvnNote{})
	err = db.Select("CreatedAt", "UpdatedAt", "ID", "Title", "Snippet").Where("create_by = ? AND del_flag=0", uid).Order("updated_at desc").Find(&noteList).Error
	err = db.Count(&total).Error
	return err, noteList, total
}

// @Summary 搜索笔记
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notes/search [post]
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
