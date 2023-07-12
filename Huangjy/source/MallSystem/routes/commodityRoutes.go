package routes

import (
	"MallSystem/controller"
	"MallSystem/utils"

	"github.com/gin-gonic/gin"
)

func registerCommodityRoutes(e *gin.Engine) {
	commodity := e.Group("/commodities")

	commodity.GET("/", controller.GetCommoditiesHandler)
	commodity.GET("/hot", controller.GetHotKeywordHandler)

	authorized := commodity.Use(utils.MiddlewareJWTAuthorize())

	authorized.POST("/", controller.PubCommodityHandler)
}
