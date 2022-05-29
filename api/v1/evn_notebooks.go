package v1

import (
	"encoding/json"
	"evernote-client/global"
	"evernote-client/model"
	"evernote-client/model/request"
	"evernote-client/model/response"
	"evernote-client/service"
	"evernote-client/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type Notebooks struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type RespondData struct {
	Data []Notebooks `json:"data"`
}

type Note struct {
	//Length    int       `json:"length"`
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type NoteRespondData struct {
	Data []Note `json:"data"`
}

// @Summary 用户删除笔记本
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /notebook/del/:id [DELETE]
func DeleteNotebook(c *gin.Context) {
	oid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	id := uint(oid)

	if err := service.DeleteNotebook(id, getUserID(c)); err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败！"+err.Error(), c)
	} else {
		// 删除Redis缓存
		_ = service.DelRedis("notebook:" + strconv.FormatUint(uint64(getUserID(c)), 10))
		// 返回数据
		response.OkWithMessage("删除成功", c)
	}
}

// @Summary 用户修改笔记本
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /notebook/:id [PATCH]
func UpdateNotebook(c *gin.Context) {
	oid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	id := uint(oid)

	var notebook model.EvnNotebook
	_ = c.ShouldBindJSON(&notebook)
	if err := utils.Verify(notebook, utils.NoteBookVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateNotebook(notebook, id, getUserID(c)); err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		// 删除Redis缓存
		_ = service.DelRedis("notebook:" + strconv.FormatUint(uint64(getUserID(c)), 10))
		// 返回数据
		response.OkWithMessage("更新成功", c)
	}
}

// @Summary 用户新建笔记本
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /notebook/add [post]
func CreateNotebook(c *gin.Context) {
	var notebook model.EvnNotebook
	_ = c.ShouldBindJSON(&notebook)
	if err := utils.Verify(notebook, utils.NoteBookVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateNotebook(notebook, getUserID(c)); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败! "+err.Error(), c)
	} else {
		// 删除Redis缓存
		_ = service.DelRedis("notebook:" + strconv.FormatUint(uint64(getUserID(c)), 10))
		// 返回数据
		response.OkWithMessage("创建成功", c)
	}
}

// @Summary 用户获取笔记本
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notebook/get [get]
func GetNotebooks(c *gin.Context) {
	userID := getUserID(c)
	// 这里做一个Redis缓存
	if err, redisData := service.GetRedis("notebook:" + strconv.FormatUint(uint64(userID), 10)); err == redis.Nil {
		if err, list, total := service.GetNotebooks(userID); err != nil {
			global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
			response.FailWithMessage("获取失败", c)
		} else {
			// 缓存到Redis
			resJson, err := json.Marshal(response.PageResult{
				List:  list,
				Total: total,
			})
			if err != nil {
				global.SYS_LOG.Error("生成json字符串错误", zap.Any("err", err))
				return
			}
			_ = service.SetRedis("notebook:"+strconv.FormatUint(uint64(userID), 10), string(resJson), 3600)
			// 返回结果
			response.OkWithDetailed(response.PageResult{
				List:  list,
				Total: total,
			}, "获取成功", c)
		}
	} else {
		// Redis里有缓存直接返回
		var result response.PageResult
		if err := json.Unmarshal([]byte(redisData), &result); err != nil {
			global.SYS_LOG.Error("解析失败!", zap.Any("err", err))
		} else {
			response.OkWithDetailed(result, "获取成功", c)
		}
	}

}

// 从Gin的Context中获取从jwt解析出来的用户ID
func getUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		global.SYS_LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return 0
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.ID
	}
}
