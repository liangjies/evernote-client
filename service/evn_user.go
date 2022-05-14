package service

import (
	"evernote-client/global"
	"evernote-client/model"
	"evernote-client/model/request"
	"evernote-client/utils"
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

//@function: ChangePassword
//@description: 用户修改密码
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser
func ChangePassword(uid uint, passWord request.ChangePassword) (err error) {
	db := global.SYS_DB.Model(&model.SysUser{})
	var user model.SysUser
	err = db.Where("id = ? AND password = ?", uid, utils.MD5V([]byte(passWord.OldPass))).First(&user).Update("password", utils.MD5V([]byte(passWord.NewPass))).Error
	return err
}
