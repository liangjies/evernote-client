package v1

import (
	"evernote-client/global"
	"evernote-client/model"
	"evernote-client/model/request"
	"evernote-client/model/response"
	"evernote-client/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 用户更新昵称
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/nickName [PATCH]
func UpdateNickName(c *gin.Context) {
	var user model.SysUser
	err := c.ShouldBindJSON(&user)
	if err != nil || user.NickName == "" {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := service.UpdateNickName(getUserID(c), user.NickName); err != nil {
		global.SYS_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败！"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Summary 用户修改密码
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [PATCH]
func ChangePassword(c *gin.Context) {
	var passWord request.ChangePassword
	err := c.ShouldBindJSON(&passWord)
	if err != nil || passWord.OldPass == "" || passWord.NewPass == "" {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := service.ChangePassword(getUserID(c), passWord); err != nil {
		global.SYS_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags User
// @Summary 用户上传头像
// @Produce  application/json
// @Param file File
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadAvatar [post]
func UploadAvatar(c *gin.Context) {
	var file model.FileUpload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.SYS_LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	err, file = service.UploadAvatar(getUserID(c), header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.SYS_LOG.Error("修改数据库链接失败!", zap.Any("err", err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(response.FileUploadResponse{File: file}, "上传成功", c)
}

// @Summary 用户修改邮箱
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/email [PATCH]
func UpdateEmail(c *gin.Context) {
	var user request.ChangeEmail
	err := c.ShouldBindJSON(&user)
	if err != nil || user.Email == "" || user.PassWord == "" {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := service.UpdateEmail(getUserID(c), user); err != nil {
		global.SYS_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败，密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}
