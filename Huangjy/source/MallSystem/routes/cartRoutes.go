package routes

import (
	"MallSystem/controller"
	"MallSystem/utils"

	"github.com/gin-gonic/gin"
)

func registerCartRoutes(e *gin.Engine) {
	cart := e.Group("/cart")
	auth := cart.Use(utils.MiddlewareJWTAuthorize())

	auth.GET("/", controller.GetCartHandler)
	auth.POST("/", controller.AddCartHandler)
	auth.DELETE("/", controller.DeleteCartHandler)
}
