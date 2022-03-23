package v1

import (
	"evernote-client/global"
	"evernote-client/model"
	"evernote-client/model/response"
	"evernote-client/service"
	"evernote-client/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 用户删除笔记
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /notebooks [post]
func DeleteNote(c *gin.Context) {
	oid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	nid := uint(oid)

	if err := service.DeleteNote(nid, getUserID(c)); err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败！"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Summary 用户修改笔记本
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /notebooks [post]
func UpdateNote(c *gin.Context) {
	var note model.EvnNote
	_ = c.ShouldBindJSON(&note)
	if err := utils.Verify(note, utils.NoteBookVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateNote(note, getUserID(c)); err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Summary 用户新建笔记
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /notes/to/:id [post]
func CreateNote(c *gin.Context) {

	var note model.EvnNote
	_ = c.ShouldBindJSON(&note)
	if err := utils.Verify(note, utils.NoteTitleVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if id, err := service.CreateNote(note, getUserID(c)); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(response.AddResult{ID: id}, "创建成功", c)
	}
}

// @Summary 用户获取笔记列表
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notes/from/:id [get]
func GetNotes(c *gin.Context) {
	oid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	nid := uint(oid)

	if err, list, total := service.GetNotes(nid, getUserID(c)); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}

// @Summary 用户根据id获取笔记详情
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notes/from/:id [get]
func GetNoteById(c *gin.Context) {
	oid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	nid := uint(oid)

	if err, list := service.GetNoteById(nid, getUserID(c)); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.NoteResult{
			List: list,
		}, "获取成功", c)
	}
}

// @Summary 用户获取笔记列表
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notes/all [get]
func GetAllNotes(c *gin.Context) {
	if err, list, total := service.GetAllNotes(getUserID(c)); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}
