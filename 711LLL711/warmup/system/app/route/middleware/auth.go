package middleware

import (
	"fmt"
	"log"
	"net/http"
	mysession "system/app/shared/session"

	"github.com/gin-gonic/gin"
)

// 检查用户是否已登录的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := mysession.GetSession(c.Request)
		log.Println("session:", session)
		if err != nil {
			fmt.Printf("error in AuthMiddleware: %v\n", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if session == nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		// 用户已登录，继续处理请求
		c.Next()
	}
}

// 检查管理员是否登录的中间件
func Auth_admin_Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查用户是否已登录
		session, err := mysession.GetSession(c.Request)
		if err != nil {
			fmt.Printf("error in AuthMiddleware: %v\n", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if session == nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		// 检查是否是管理员
		if session.Role != "admin" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// 用户已登录且为管理员，继续处理请求
		c.Next()
	}
}
