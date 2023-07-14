package services

import (
	"MallSystem/database"
	"MallSystem/model"
	"MallSystem/model/response"
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
	dict := make(map[string]int)
	for _, id := range o.CommoditiesID {
		if _, ok := dict[id.String()]; ok {
			return nil, errors.New("不允许多次操作同一个商品")
		} else if ci, err := database.QueryOneCommodity(&bson.M{"_id": id}); err != nil {
			return nil, err
		} else if ci.Status == model.Sold {
			return nil, errors.New("商品已经卖出")
		} else if ci.UserID != o.SellerID {
			return nil, errors.New("订单卖家信息无效")
		} else {
			totalPrice += ci.Price
			dict[id.String()] = 1
		}
	}
	if totalPrice != o.Price {
		return nil, errors.New("订单价格无效")
	}
	return o, nil
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
		session.AbortTransaction(context.Background())
		return err
	}
	id, err := database.InsertOneOrder(o)
	if err != nil {
		session.AbortTransaction(context.Background())
		return err
	}
	for _, id := range o.CommoditiesID {
		if err := database.SetOneCommodityStatus(&bson.M{"_id": id}, model.Sold); err != nil {
			return nil
		}
	}
	// 操作数据库，将商品状态进行修改，然后处理事务结果
	// 在处理订单时，已经进入了事务，此时查询到未卖出则一定未卖出，可以修改数据库
	log.Println(*id)

	err = session.CommitTransaction(context.Background())
	if err != nil {
		session.AbortTransaction(context.Background())
		return err
	}

	// 在这里开始支付，调用第三方支付平台的接口，传入相关信息，获取支付凭证或重定向
	payInfo := "模拟付款"
	c.JSON(http.StatusOK, response.MakeSucceedResponse(payInfo))
	return nil
}
