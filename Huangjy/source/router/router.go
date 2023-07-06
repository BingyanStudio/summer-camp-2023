package router

import (
	"log"
	"myserver/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	log.Println("running...")
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	user := router.Group("/user")
	{
		user.POST("/login", controller.LoginHandler)
		user.POST("/", controller.RegisterHandler)
	}
	router.Run(":8080")
}
