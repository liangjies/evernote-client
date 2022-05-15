package service

import (
	"evernote-client/global"
	"evernote-client/model"
	"evernote-client/model/request"
	"evernote-client/utils"
	"evernote-client/utils/upload"
	"mime/multipart"
	"strings"
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

//@function: UpdateAvatar
//@description: 更新用户头像
//@param: file model.FileUpload
//@return: error
func UpdateAvatar(uid uint, file model.FileUpload) (err error) {
	db := global.SYS_DB.Model(&model.SysUser{})
	var user model.SysUser
	err = db.Where("id = ?", uid).First(&user).Update("headerimg", file.Url).Error
	return err
}

//@function: UploadAvatar
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload
func UploadAvatar(uid uint, header *multipart.FileHeader, noSave string) (err error, file model.FileUpload) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := model.FileUpload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return UpdateAvatar(uid, f), f
	}
	return
}

//@function: UpdateEmail
//@description: 用户修改邮箱
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser
func UpdateEmail(uid uint, user request.ChangeEmail) (err error) {
	db := global.SYS_DB.Model(&model.SysUser{})
	var SysUser model.SysUser
	err = db.Where("id = ? AND password = ?", uid, utils.MD5V([]byte(user.PassWord))).First(&SysUser).Update("email", user.Email).Error
	return err
}
