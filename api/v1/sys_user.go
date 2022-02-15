package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string `json:"username"`
}

type Auth struct {
	Data    User `json:"data"`
	IsLogin bool `json:"isLogin"`
}

type AuthLogin struct {
	Data User   `json:"data"`
	Msg  string `json:"msg"`
}

func GetAuth(c *gin.Context) {
	var auth Auth
	auth.IsLogin = true
	auth.Data = User{"liangjies"}
	c.JSON(200, auth)
}
func Login(c *gin.Context) {
	password := c.Param("password")
	username := c.Param("username")
	fmt.Println(username)
	fmt.Println(password)

	var user User
	user.UserName = "liangjies"
	var auth AuthLogin
	auth.Data = user
	c.JSON(200, auth)
}
