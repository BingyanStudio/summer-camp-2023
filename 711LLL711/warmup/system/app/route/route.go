package route

import (
	"net/http"
	"system/app/controller"
	"system/app/route/middleware"
	"system/app/shared/server"
	mysession "system/app/shared/session"
	myuser "system/app/shared/user"

	"github.com/gin-gonic/gin"
)

// 设置路由
func Routes() {

	// 用户注册
	server.Ginserver.GET("/register", func(context *gin.Context) {
		context.HTML(http.StatusOK, "register.html", gin.H{
			"title": "注册",
		})
	})

	server.Ginserver.POST("/register", func(context *gin.Context) {
		var user myuser.User
		user.Username = context.PostForm("username")
		user.Password = context.PostForm("password")
		user.Id = context.PostForm("id")
		user.Phone = context.PostForm("phone")
		user.Email = context.PostForm("email")
		user.Role = context.PostForm("role")
		controller.RegisterUser(user, context)
	})

	// 用户登录
	server.Ginserver.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{
			"title": "登录",
		})
	})
	server.Ginserver.POST("/login", func(context *gin.Context) {
		var user myuser.User
		user.Id = context.PostForm("id")
		user.Password = context.PostForm("password")
		controller.LoginUser(user, context)
	})

	// 普通用户更改个人信息
	server.Ginserver.GET("/update", middleware.AuthMiddleware(), func(context *gin.Context) {
		session, _ := mysession.GetSession(context.Request)

		//id := mysession.Getid(context)
		context.HTML(http.StatusOK, "update.html", gin.H{
			"title": "更新",
			"id":    session.Id,
		})
	})

	server.Ginserver.POST("/update", middleware.AuthMiddleware(), func(context *gin.Context) {
		var user myuser.User
		user.Username = context.PostForm("username")
		user.Password = context.PostForm("password")
		user.Phone = context.PostForm("phone")
		user.Email = context.PostForm("email")
		user.Role = context.PostForm("role")
		session, _ := mysession.GetSession(context.Request)
		user.Id = session.Id
		//user.Id = mysession.Getid(context)
		controller.Updateuser(user, context)
	})

	// 仅管理员可用的路由组

	adminGroup := server.Ginserver.Group("/admin")
	adminGroup.Use(middleware.Auth_admin_Middleware())

	// 删除普通用户--删除请求中的参数中的id记录
	adminGroup.DELETE("/del", func(context *gin.Context) {
		controller.DeleteUser(context)
	})

	// 获取一个成员
	adminGroup.GET("/getuser", func(context *gin.Context) {
		controller.GetUser(context)
	})

	// 获取所有成员信息
	adminGroup.GET("/getalluser", func(context *gin.Context) {
		controller.GetAllUsers(context)
	})

}
