package service

import (
	"errors"
	"evernote-client/global"
	"evernote-client/model"

	"gorm.io/gorm"
)

//@function: DeleteNotebook
//@description: 用户删除笔记本
//@param: id uint, uid uint
//@return: err error
func DeleteNotebook(id uint, uid uint) (err error) {
	err = global.SYS_DB.Where("notebook_id = ? AND del_flag=0", id).First(&model.EvnNote{}).Error
	if err != nil {
		var notebook model.EvnNotebook
		err = global.SYS_DB.Where("id = ? AND create_by= ?", id, uid).Delete(&notebook).Error
		return err
	} else {
		return errors.New("此笔记本下存在笔记不可删除")
	}
}

//@function: UpdateNotebook
//@description: 用户修改笔记本
//@param: n model.EvnNotebook, id uint, uid uint
//@return: err error
func UpdateNotebook(n model.EvnNotebook, id uint, uid uint) (err error) {
	var notebook model.EvnNotebook
	if !errors.Is(global.SYS_DB.Where("title = ? AND id != ? AND create_by = ?", n.Title, id, uid).First(&notebook).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("笔记本已存在")
	}
	err = global.SYS_DB.Where("id = ? AND create_by = ?", id, uid).First(&notebook).Update("title", n.Title).Error
	return err
}

//@function: CreateNotebook
//@description: 用户新建笔记本
//@param: n model.EvnNotebook, id uint
//@return: err error
func CreateNotebook(n model.EvnNotebook, id uint) (err error) {
	var notebook model.EvnNotebook
	if !errors.Is(global.SYS_DB.Where("title = ? AND create_by = ?", n.Title, id).First(&notebook).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("笔记本已存在")
	}
	n.CreateBy = id
	err = global.SYS_DB.Create(&n).Error
	return err
}

//@function: GetNotebooks
//@description: 用户获取笔记本
//@param: id uint
//@return: err error, list interface{}, total int64
func GetNotebooks(id uint) (err error, list interface{}, total int64) {
	var notebookList []model.EvnNotebook
	db := global.SYS_DB.Model(&model.EvnNotebook{})
	err = db.Count(&total).Error
	err = db.Where("create_by = ?", id).Find(&notebookList).Error
	return err, notebookList, total
}
