package routes

import (
	"MallSystem/controller"
	"MallSystem/utils"

	"github.com/gin-gonic/gin"
)

func registerPurchaseRoutes(e *gin.Engine) {
	authorized := e.Group("/purchase").Use(utils.MiddlewareJWTAuthorize())

	authorized.POST("/direct", controller.PurchaseHandler)
}
