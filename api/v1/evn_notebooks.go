package v1

import (
	"evernote-client/global"
	"evernote-client/model"
	"evernote-client/model/request"
	"evernote-client/model/response"
	"evernote-client/service"
	"evernote-client/utils"
	"time"

	"github.com/gin-gonic/gin"
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

func CreateNotebook(c *gin.Context) {
	var notebook model.EvnNotebook
	_ = c.ShouldBindJSON(&notebook)
	if err := utils.Verify(notebook, utils.NoteBookVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateNotebook(notebook, getUserID(c)); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func GetNotebooks(c *gin.Context) {
	if err, list, total := service.GetNotebooks(getUserID(c)); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}

func GetNote(c *gin.Context) {
	// password := c.Param("password")
	// username := c.Param("username")
	// fmt.Println(username)
	// fmt.Println(password)

	//var notebooks []Notebooks
	var notes []Note
	var note Note
	var noteRespondData NoteRespondData
	//note.Length = 0
	note.Id = 1
	note.Title = "笔记测试标题"
	note.Content = "笔记测试内容"
	note.UpdatedAt = time.Now()
	note.CreatedAt = time.Now()
	notes = append(notes, note)
	noteRespondData.Data = notes
	c.JSON(200, noteRespondData)
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
