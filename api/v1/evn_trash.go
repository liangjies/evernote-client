package v1

import (
	"evernote-client/global"
	"evernote-client/model/response"
	"evernote-client/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 用户从废纸篓恢复笔记
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /notebooks [post]
func RevertNote(c *gin.Context) {
	oid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	nid := uint(oid)

	if err := service.RevertNote(nid, getUserID(c)); err != nil {
		global.SYS_LOG.Error("恢复失败!", zap.Any("err", err))
		response.FailWithMessage("恢复失败", c)
	} else {
		response.OkWithMessage("恢复成功", c)
	}
}

// @Summary 用户彻底删除笔记
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /notebooks [post]
func DeleteTrash(c *gin.Context) {
	oid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	nid := uint(oid)

	if err := service.DeleteTrash(nid, getUserID(c)); err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败！"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Summary 用户获取废纸篓
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notes/from/:id [get]
func GetTrashs(c *gin.Context) {
	if err, list, total := service.GetTrashs(getUserID(c)); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}
