package routes

import (
	"MallSystem/controller"

	"github.com/gin-gonic/gin"
)

func registerImageRoutes(e *gin.Engine) {
	image := e.Group("/image")

	image.GET("/small/:filename", controller.GetSmallImageHandler)
	image.GET("/:filename", controller.GetImageHandler)

}
