package routes

import (
	"MallSystem/controller"
	"MallSystem/utils"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(e *gin.Engine) {
	user := e.Group("/user")

	user.POST("", controller.UserRegisterHandler)
	user.POST("/login", controller.UserLoginHandler)

	user.GET("/:userid", controller.UserInfoHandler)

	me := user.Use(utils.MiddlewareJWTAuthorize())

	me.GET("/me", controller.SelfInfoHandler)

}
