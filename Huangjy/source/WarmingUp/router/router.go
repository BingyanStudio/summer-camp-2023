package router

import (
	"log"
	"myserver/controller"
	"myserver/controller/middleware"
	"myserver/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	log.Println("running...")
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	router.POST("/login", controller.LoginHandler)
	router.POST("/register", controller.RegisterHandler)
	user := router.Group("/me")
	user.Use(middleware.Authorizate(model.User))
	{
		user.GET("", func(c *gin.Context) {
			c.String(200, "Hello!")
		})
		user.POST("info", controller.UserUpdateInfoHandler)
	}
	admin := router.Group("/admin")
	admin.Use(middleware.Authorizate(model.Admin))
	{
		admin.GET("", func(c *gin.Context) {
			c.String(200, "hello admin!")
		})
		admin.DELETE("/user/:user_id", controller.AdminDeleteHandler)
		admin.GET("/user/:user_id", controller.AdminGetOneUser)
		admin.GET("/user/all", controller.AdminGetAllUsers)
	}
	router.Run(":8080")
}
