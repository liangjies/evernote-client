package v1

import (
	"evernote-client/global"
	"evernote-client/model"
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
