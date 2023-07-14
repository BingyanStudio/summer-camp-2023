package controller

import (
	"MallSystem/database"
	"MallSystem/model"
	"MallSystem/model/response"
	"MallSystem/services"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PurchaseHandler(c *gin.Context) {
	var (
		o model.OrderInfo
	)
	sellerid, err := primitive.ObjectIDFromHex(c.Query("sellerid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	commoditiesid := c.QueryArray("commodities")
	if len(commoditiesid) == 0 {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	o.Price, err = strconv.ParseFloat(c.Query("price"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	address, err := primitive.ObjectIDFromHex(c.Query("address"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	o.SellerID = sellerid
	for _, i := range commoditiesid {
		id, err := primitive.ObjectIDFromHex(i)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.InvalidInfoError)
			return
		}
		o.CommoditiesID = append(o.CommoditiesID, id)
	}
	o.AddressID = address

	if err := services.OrderService(&o, c); err != nil {
		if err == context.DeadlineExceeded {
			c.JSON(http.StatusInternalServerError, response.TimeoutError)
		} else {
			c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		}
		return
	}
}

// 第三方的平台通知服务器支付信息，服务端定义一个回调函数AliPayNotify(c *gin.Context)
func AlipayNotify(c *gin.Context) {
	// 如果支付成功，则向支付宝返回成功信息，并修改订单状态为已付款，
	// 订单信息会由alipay通过gin的上下文返回，根据此订单查找数据
	// 如果支付失败，返回失败信息，传入的订单将商品修改为原来的状态，这里默认成功
	id, err := primitive.ObjectIDFromHex(c.Query("orderid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.InvalidInfoError)
		return
	}
	if c.Query("status") == "succeed" {
		if err := database.SetOneOrderStatus(&bson.M{"_id": id}, &bson.M{"$set": bson.M{"status": model.Paid, "dealTime": time.Now()}}); err != nil {
			c.JSON(http.StatusInternalServerError, response.TimeoutError)
			return
		}
	} else {
		if err := database.SetOneOrderStatus(&bson.M{"_id": id}, &bson.M{"$set": bson.M{"status": model.Canceled, "dealTime": time.Now()}}); err != nil {
			c.JSON(http.StatusInternalServerError, response.TimeoutError)
			return
		}
		o, err := database.QueryOneOrder(&bson.M{"_id": id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.TimeoutError)
			return
		}
		for _, ci := range o.CommoditiesID {
			if err := database.SetOneCommodityStatus(&bson.M{"_id": ci}, model.Selling); err != nil {
				c.JSON(http.StatusInternalServerError, response.TimeoutError)
				return
			}
		}
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
}
