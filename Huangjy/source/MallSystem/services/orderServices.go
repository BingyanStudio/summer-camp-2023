package services

import (
	"MallSystem/database"
	"MallSystem/model"
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeOrder(o *model.OrderInfo, c *gin.Context) (*model.OrderInfo, error) {
	var (
		totalPrice float64 = 0
	)
	o.UserID, _ = primitive.ObjectIDFromHex(c.GetString("userid")[10:34])
	o.CreateTime = time.Now()
	o.Status = model.Pending

	for _, id := range o.CommoditiesID {
		if ci, err := database.QueryOneCommodity(&bson.M{"_id": id}); err != nil {
			return nil, err
		} else if ci.Status == model.Sold {
			return nil, errors.New("商品已经卖出")
		} else if ci.UserID != o.SellerID {
			return nil, errors.New("订单卖家信息无效")
		} else {
			totalPrice += ci.Price
			// 在这里更新数据库信息，如果某个地方出现错误则事务会立即结束，而不会被提交
			if err2 := database.SetOneCommodityStatus(&bson.M{"_id": id}, model.Sold); err2 != nil {
				return nil, err2
			}
		}
	}
	if totalPrice != o.Price {
		return nil, errors.New("订单价格无效")
	}
	return o, nil
}

// 第三方的平台通知服务器支付信息，服务端定义一个回调函数AliPayNotify(c *gin.Context)
func AlipayNotify(c *gin.Context) {
	// 如果支付成功，则向支付宝返回成功信息，并修改订单状态为已付款，
	// 订单信息会由alipay通过gin的上下文返回，根据此订单查找数据
	// 如果支付失败，返回失败信息，传入的订单将商品修改为原来的状态，这里默认成功
}

func OrderService(o *model.OrderInfo, c *gin.Context) error {
	session, err := database.MakeSession()
	if err != nil {
		return err
	}
	session.StartTransaction()
	defer session.EndSession(context.Background())

	o, err = makeOrder(o, c)
	if err != nil {
		return err
	}

	// 操作数据库，将商品状态进行修改，然后处理事务结果
	// 在处理订单时，已经进入了事务，此时查询到未卖出则一定未卖出，可以修改数据库
	log.Println(o.CreateTime)

	err = session.CommitTransaction(context.Background())
	if err != nil {
		session.AbortTransaction(context.Background())
	}

	// 在这里开始支付，调用第三方支付平台的接口，传入相关信息，获取支付凭证或重定向
	payInfo := "模拟付款"
	c.JSON(http.StatusOK, payInfo)

	/*** 模拟支付成功 ***/

	/******************/
	return nil
}
