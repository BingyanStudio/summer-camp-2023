package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus int

/*
 *	0 待支付
 *	1 已支付
 *	2 已发货
 *	3 已完成
 *	4 已退款
 */
const (
	Pending OrderStatus = iota
	Paid
	Shipped
	Completed
	Refunded
	Canceled
)

/*
 *	订单ID
 *	买家ID
 *	卖家ID
 *	订单总价
 *	订单创建时间
 *	订单交付时间
 *	订单状态
 *	商品ID
 *	送货地址ID
 *	备注
 */
type OrderInfo struct {
	ID            primitive.ObjectID   `json:"id" bson:"_id,omitempty" form:"-" binding:"-"`
	UserID        primitive.ObjectID   `json:"userid" bson:"userid" form:"-" binding:"-"`
	SellerID      primitive.ObjectID   `json:"sellerid" bson:"sellerid" form:"sellerid" binding:"required"`
	Price         float64              `json:"price" bson:"price" form:"price" binding:"required"`
	CreateTime    time.Time            `json:"createTime" bson:"createTime" form:"-" binding:"-"`
	DealTime      time.Time            `json:"dealTime" bson:"dealTime" form:"-" binding:"-"`
	Status        OrderStatus          `json:"status" bson:"status" form:"-" binding:"-"`
	CommoditiesID []primitive.ObjectID `json:"commoditiesid" bson:"commoditiesid" form:"commoditiesid" binding:"required"`
	AddressID     primitive.ObjectID   `json:"addressid" bson:"addressid" form:"addressid" binding:"required"`
	Remark        string               `json:"remark" bson:"remark" form:"remark" binding:"-"`
}

/****************************************/
