package v1

import (
	"time"

	"github.com/gin-gonic/gin"
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

func GetNotebooks(c *gin.Context) {
	var notebooks []Notebooks
	var notebook Notebooks
	var respondData RespondData
	notebook.Id = 1
	notebook.Title = "测试标题"
	notebook.UpdatedAt = time.Now()
	notebook.CreatedAt = time.Now()
	notebooks = append(notebooks, notebook)
	respondData.Data = notebooks
	c.JSON(200, respondData)
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
