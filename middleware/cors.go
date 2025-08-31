package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//origin := c.Request.Header.Get("Origin")
		// 为了兼容通配符的跨域场景，这里不携带凭据（Cookie/Authorization 自动注入），前端改为显式设置头即可
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id,authorization")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT,PATCH")
		// 暴露后端续签头，便于前端读取并替换 token
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, new-token, new-expires-at")
		// 当 Allow-Origin 为 * 时，Allow-Credentials 不能为 true，否则浏览器会拦截
		c.Header("Access-Control-Allow-Credentials", "false")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
