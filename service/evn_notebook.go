package service

import (
	"errors"
	"evernote-client/global"
	"evernote-client/model"

	"gorm.io/gorm"
)

func CreateNotebook(n model.EvnNotebook, id uint) (err error) {
	var notebook model.EvnNotebook
	if !errors.Is(global.SYS_DB.Where("title = ? AND create_by = ?", n.Title, id).First(&notebook).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("笔记本已存在")
	}
	n.CreateBy = id
	err = global.SYS_DB.Create(&n).Error
	return err
}

//@function: GetUserInfoList
//@description: 获取用户数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func GetNotebooks(id uint) (err error, list interface{}, total int64) {
	var notebookList []model.EvnNotebook
	db := global.SYS_DB.Model(&model.EvnNotebook{})
	err = db.Count(&total).Error
	err = db.Where("create_by = ?", id).Find(&notebookList).Error
	return err, notebookList, total
}
