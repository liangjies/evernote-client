package v1

import (
	"evernote-client/global"
	"evernote-client/model/response"
	"evernote-client/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 获取笔记历史版本
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /history/:id [get]
func GetHistories(c *gin.Context) {
	oid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	nid := uint(oid)

	if err, list, total := service.GetHistories(nid, getUserID(c)); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}
