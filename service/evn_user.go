package service

import (
	"evernote-client/global"
	"evernote-client/model"
)

//@function: UpdateNickName
//@description: 用户更新昵称
//@param: uid uint,
//@return: err error
func UpdateNickName(uid uint, nickName string) (err error) {
	db := global.SYS_DB.Model(&model.SysUser{})
	err = db.Where("id=?", uid).Update("nickname", nickName).Error
	return err
}
