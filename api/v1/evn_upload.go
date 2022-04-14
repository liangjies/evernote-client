package v1

import (
	"evernote-client/global"
	"evernote-client/model"
	"evernote-client/model/response"
	"evernote-client/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /upload [post]
func UploadFile(c *gin.Context) {
	var file model.FileUpload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.SYS_LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	err, file = service.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.SYS_LOG.Error("修改数据库链接失败!", zap.Any("err", err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(response.FileUploadResponse{File: file}, "上传成功", c)
}
