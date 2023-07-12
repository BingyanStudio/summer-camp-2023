package routes

import (
	"MallSystem/controller"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(e *gin.Engine) {
	user := e.Group("/user")

	user.POST("", controller.UserRegisterHandler)
	user.POST("/login", controller.UserLoginHandler)

}
