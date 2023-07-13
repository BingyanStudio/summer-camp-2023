package controller

import (
	"MallSystem/model"
	"MallSystem/model/response"
	"MallSystem/services"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PurchaseHandler(c *gin.Context) {
	var (
		o model.OrderInfo
	)
	if err := c.ShouldBind(&o); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	if err := services.OrderService(&o, c); err != nil {
		if err == context.DeadlineExceeded {
			c.JSON(http.StatusInternalServerError, response.TimeoutError)
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		return
	}
}
