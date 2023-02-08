package service

import (
	"errors"
	"evernote-client/global"
	"evernote-client/model"
	"evernote-client/model/request"
	"evernote-client/utils"
	"evernote-client/utils/upload"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"mime/multipart"
	"strings"
)

// @function: UpdateNickName
// @description: 用户更新昵称
// @param: uid uint,
// @return: err error
func UpdateNickName(uid uint, nickName string) (err error) {
	db := global.DB.Model(&model.EvnUser{})
	err = db.Where("id=?", uid).Update("nickname", nickName).Error
	return err
}

// @function: ChangePassword
// @description: 用户修改密码
// @param: u *model.EvnUser, newPassword string
// @return: err error, userInter *model.EvnUser
func ChangePassword(uid uint, passWord request.ChangePassword) (err error) {
	db := global.DB.Model(&model.EvnUser{})
	var user model.EvnUser
	err = db.Where("id = ? AND password = ?", uid, utils.MD5V([]byte(passWord.OldPass))).First(&user).Update("password", utils.MD5V([]byte(passWord.NewPass))).Error
	return err
}

// @function: UpdateAvatar
// @description: 更新用户头像
// @param: file model.EvnUpload
// @return: error
func UpdateAvatar(uid uint, file model.EvnUpload) (err error) {
	db := global.DB.Model(&model.EvnUser{})
	var user model.EvnUser
	err = db.Where("id = ?", uid).First(&user).Update("headerimg", file.Url).Error
	return err
}

// @function: UploadAvatar
// @description: 根据配置文件判断是文件上传到本地或者七牛云
// @param: header *multipart.FileHeader, noSave string
// @return: err error, file model.ExaFileUploadAndDownload
func UploadAvatar(uid uint, header *multipart.FileHeader, noSave string) (err error, file model.EvnUpload) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := model.EvnUpload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return UpdateAvatar(uid, f), f
	}
	return
}

// @function: UpdateEmail
// @description: 用户修改邮箱
// @param: u *model.EvnUser, newPassword string
// @return: err error, userInter *model.EvnUser
func UpdateEmail(uid uint, user request.ChangeEmail) (err error) {
	db := global.DB.Model(&model.EvnUser{})
	var SysUser model.EvnUser
	err = db.Where("id = ? AND password = ?", uid, utils.MD5V([]byte(user.PassWord))).First(&SysUser).Update("email", user.Email).Error
	return err
}

// @function: Register
// @description: 用户注册
// @param: u model.EvnUser
// @return: err error, userInter model.EvnUser
func Register(u request.Register) (err error) {
	var user model.EvnUser
	if !errors.Is(global.DB.Model(&model.EvnUser{}).Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已存在")
	}
	if !errors.Is(global.DB.Model(&model.EvnUser{}).Where("email = ?", u.Email).First(&user).Error, gorm.ErrRecordNotFound) { // 判断邮箱是否注册
		return errors.New("邮箱已存在")
	}
	l := model.EvnUser{Username: u.Username, Email: u.Email}
	l.UUID = uuid.NewV4()
	// 否则 附加uuid 密码md5简单加密 注册
	l.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Model(&model.EvnUser{}).Create(&l).Error
	return err
}

// @function: Login
// @description: 用户登录
// @param: u *model.EvnUser
// @return: err error, userInter *model.EvnUser
func Login(u *model.EvnUser) (err error, userInter *model.EvnUser) {
	var user model.EvnUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

// @function: Login
// @description: 用户退出登录
// @param: u *model.EvnUser
// @return: err error, userInter *model.EvnUser
func Logout(jwtList model.EvnJwtBlacklist) (err error) {
	err = global.DB.Create(&jwtList).Error
	return
}

// @function: SetUserInfo
// @description: 设置用户信息
// @param: reqUser model.EvnUser
// @return: err error, user model.EvnUser
func SetUserInfo(reqUser model.EvnUser) (err error, user model.EvnUser) {
	err = global.DB.Updates(&reqUser).Error
	return err, reqUser
}

// @function: FindUserById
// @description: 通过id获取用户信息
// @param: id int
// @return: err error, user *model.EvnUser
func FindUserById(id int) (err error, user *model.EvnUser) {
	var u model.EvnUser
	err = global.DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

// @author: [SliverHorn](https://github.com/SliverHorn)
// @function: FindUserByUuid
// @description: 通过uuid获取用户信息
// @param: uuid string
// @return: err error, user *model.EvnUser
func FindUserByUuid(uuid string) (err error, user *model.EvnUser) {
	var u model.EvnUser
	if err = global.DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
