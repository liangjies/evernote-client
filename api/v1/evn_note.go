package v1

import (
	"evernote-client/global"
	"evernote-client/model"
	"evernote-client/model/request"
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
// @Router /notes/:id [DELETE]
func DeleteNote(c *gin.Context) {
	oid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	nid := uint(oid)

	if err := service.DeleteNote(nid, getUserID(c)); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败！"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Summary 用户修改笔记
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /notes/update [POST]
func UpdateNote(c *gin.Context) {
	var note model.EvnNote
	_ = c.ShouldBindJSON(&note)
	if err := utils.Verify(note, utils.NoteBookVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateNote(note, getUserID(c)); err != nil {
		global.LOG.Error("保存失败!", zap.Any("err", err))
		response.FailWithMessage("保存失败", c)
	} else {
		response.OkWithMessage("保存成功", c)
	}
}

// @Summary 用户新建笔记
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /notes/add [post]
func CreateNote(c *gin.Context) {
	var note model.EvnNote

	_ = c.ShouldBindJSON(&note)
	if note.Title == "" && note.Content == "" {
		response.FailWithMessage("请输入一些内容😊", c)
		return
	}
	if note.Title == "" {
		note.Title = "未命名标题"
	}

	if id, err := service.CreateNote(note, getUserID(c)); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		// 更新笔记本缓存
		_ = UpdateNotebookRedis(getUserID(c))
		response.OkWithDetailed(response.AddResult{ID: id}, "创建成功", c)
	}
}

// @Summary 获取笔记本笔记列表
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

	if err, list, total, title := service.GetNotes(nid, getUserID(c)); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, title, c)
	}
}

// @Summary 用户根据id获取笔记详情
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notes/:id [get]
func GetNoteById(c *gin.Context) {
	oid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	nid := uint(oid)

	if err, list := service.GetNoteById(nid, getUserID(c)); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.NoteResult{
			List: list,
		}, "获取成功", c)
	}
}

// @Summary 获取所有笔记
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notes/all [get]
func GetAllNotes(c *gin.Context) {
	if err, list, total := service.GetAllNotes(getUserID(c)); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}

// @Summary 搜索笔记
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notes/search [post]
func SearchNote(c *gin.Context) {
	var SearchParams request.SearchNoteParams
	err := c.ShouldBindJSON(&SearchParams)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err, list, total := service.SearchNote(SearchParams.SearchKey, SearchParams.NotebookId, getUserID(c)); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysNoteSearchResponse{
			List:      list,
			SearchKey: SearchParams.SearchKey,
			Total:     total,
		}, "获取成功", c)
	}
}
