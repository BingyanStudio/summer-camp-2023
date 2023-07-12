package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	e := gin.Default()

	// 注册各路由
	registerUserRoutes(e)
	registerCommodityRoutes(e)

	return e
}
