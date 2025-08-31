package router

import (
	v1 "evernote-client/api/v1"

	"github.com/gin-gonic/gin"
)

func InitLikeRouter(Router *gin.RouterGroup) {
	LikeRouter := Router.Group("like")
	{
		// 点赞笔记
		LikeRouter.POST("/note/:noteId", v1.LikeNote)

		// 取消点赞
		LikeRouter.DELETE("/note/:noteId", v1.UnlikeNote)

		// 检查用户是否点赞
		LikeRouter.GET("/check/:noteId", v1.CheckUserLike)

		// 获取笔记点赞数量
		LikeRouter.GET("/count/:noteId", v1.GetNoteLikeCount)

		// 获取笔记点赞用户列表
		LikeRouter.GET("/users/:noteId", v1.GetNoteLikeUsers)

		// 获取点赞排行榜
		LikeRouter.GET("/ranking", v1.GetLikeRanking)

		// 获取用户点赞的笔记列表
		LikeRouter.GET("/user/notes", v1.GetUserLikedNotes)

		// 获取热门笔记
		LikeRouter.GET("/hot", v1.GetHotNotes)

		// 同步点赞数据（管理员功能）
		LikeRouter.POST("/sync", v1.SyncLikeData)
	}
}
