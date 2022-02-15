// service
package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RespondData struct {
	Status    int     `json:"status"`
	U         float64 `json:"u"`
	I         float64 `json:"i"`
	Process   string  `json:"process"`
	Pointcode string  `json:"pointcode"`
	Mutiple   float64 `json:"mutiple"`
	Ia        float64 `json:"ia"`
	E         float64 `json:"e"`
	Msg       string  `json:"msg"`
}

func Login(c *gin.Context) {
	password := c.Param("password")
	username := c.Param("username")
	fmt.Println(username)
	fmt.Println(password)

	var respondData RespondData
	c.JSON(200, respondData)
}
