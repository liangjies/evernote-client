package v1

import (
	"evernote-client/model/response"
	"evernote-client/service"
	"strconv"

	"evernote-client/model/request"

	"github.com/gin-gonic/gin"
)

var likeService = service.NewLikeService()

// 点赞笔记
func LikeNote(c *gin.Context) {
	noteIDStr := c.Param("noteId")
	noteID, err := strconv.ParseUint(noteIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("笔记ID格式错误", c)
		return
	}

	// 从JWT中获取用户信息
	userID := getUserID(c)
	username := getUsername(c)

	err = likeService.LikeNote(uint(noteID), userID, username)
	if err != nil {
		response.FailWithMessage("点赞失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("点赞成功", c)
}

// 取消点赞
func UnlikeNote(c *gin.Context) {
	noteIDStr := c.Param("noteId")
	noteID, err := strconv.ParseUint(noteIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("笔记ID格式错误", c)
		return
	}

	userID := getUserID(c)

	err = likeService.UnlikeNote(uint(noteID), userID)
	if err != nil {
		response.FailWithMessage("取消点赞失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("取消点赞成功", c)
}

// 检查用户是否点赞了笔记
func CheckUserLike(c *gin.Context) {
	noteIDStr := c.Param("noteId")
	noteID, err := strconv.ParseUint(noteIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("笔记ID格式错误", c)
		return
	}

	userID := getUserID(c)

	isLiked, err := likeService.IsUserLiked(uint(noteID), userID)
	if err != nil {
		response.FailWithMessage("检查点赞状态失败", c)
		return
	}

	response.OkWithData(map[string]interface{}{
		"isLiked": isLiked,
	}, c)
}

// 获取笔记点赞数量
func GetNoteLikeCount(c *gin.Context) {
	noteIDStr := c.Param("noteId")
	noteID, err := strconv.ParseUint(noteIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("笔记ID格式错误", c)
		return
	}

	count, err := likeService.GetNoteLikeCount(uint(noteID))
	if err != nil {
		response.FailWithMessage("获取点赞数量失败", c)
		return
	}

	response.OkWithData(map[string]interface{}{
		"likeCount": count,
	}, c)
}

// 获取笔记点赞用户列表
func GetNoteLikeUsers(c *gin.Context) {
	noteIDStr := c.Param("noteId")
	noteID, err := strconv.ParseUint(noteIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("笔记ID格式错误", c)
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	users, err := likeService.GetNoteLikeUsers(uint(noteID), limit)
	if err != nil {
		response.FailWithMessage("获取点赞用户列表失败", c)
		return
	}

	response.OkWithData(map[string]interface{}{
		"users": users,
	}, c)
}

// 获取点赞排行榜
func GetLikeRanking(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	ranking, err := likeService.GetLikeRanking(limit)
	if err != nil {
		response.FailWithMessage("获取排行榜失败", c)
		return
	}

	response.OkWithData(map[string]interface{}{
		"ranking": ranking,
	}, c)
}

// 获取用户点赞的笔记列表
func GetUserLikedNotes(c *gin.Context) {
	userID := getUserID(c)

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	notes, total, err := likeService.GetUserLikedNotes(userID, page, pageSize)
	if err != nil {
		response.FailWithMessage("获取用户点赞笔记失败", c)
		return
	}

	response.OkWithData(map[string]interface{}{
		"notes":    notes,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}, c)
}

// 获取热门笔记
func GetHotNotes(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	notes, err := likeService.GetHotNotes(limit)
	if err != nil {
		response.FailWithMessage("获取热门笔记失败", c)
		return
	}

	response.OkWithData(map[string]interface{}{
		"notes": notes,
	}, c)
}

// 同步点赞数据
func SyncLikeData(c *gin.Context) {
	err := likeService.SyncLikeData()
	if err != nil {
		response.FailWithMessage("同步点赞数据失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("同步点赞数据成功", c)
}

// 辅助函数：从JWT获取用户名
func getUsername(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		// 如果没有JWT claims，尝试从header获取
		username := c.GetHeader("Username")
		if username == "" {
			return "未知用户"
		}
		return username
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.Username
	}
}
