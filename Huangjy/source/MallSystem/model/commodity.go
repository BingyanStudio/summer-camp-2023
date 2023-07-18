package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryIndex int
type CommodityStatus int

/*
 * 1 电子设备
 * 2 书籍资料
 * 3 宿舍百货
 * 4 美妆护肤
 * 5 女装
 * 6 男装
 * 7 鞋帽配饰
 * 8 门票卡券
 * 9 其它
 */
const (
	All CategoryIndex = iota
	Electronics
	Books
	DormGoods
	Beauty
	WomenClothing
	MenClothing
	ShoesHatsAccessories
	TicketsCoupons
	Others
)

/*
 *	0 已卖出
 *	1 售卖中
 */
const (
	Sold CommodityStatus = iota
	Selling
)

/*
 *	商品信息结构体
 *	商品ID
 *	发布者ID
 *	上架时间
 *	商品名
 *	商品简介
 *	商品价格
 *	商品类别
 *	浏览数
 *	收藏数
 *	商品状态
 */
type CommodityInfo struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty" form:"-" binding:"-"`
	UserID       primitive.ObjectID `json:"userid" bson:"userid" form:"-" binding:"-"`
	LaunchTime   time.Time          `json:"launchTime" bson:"launchTime" form:"-" binding:"-"`
	Title        string             `json:"title" bson:"title" form:"title" binding:"required"`
	Desc         string             `json:"desc" bson:"desc" form:"desc" binding:"required"`
	Price        float64            `json:"price" bson:"price" form:"price" binding:"required"`
	Category     CategoryIndex      `json:"category" bson:"category" form:"category" binding:"required"`
	Picture      string             `json:"picture" bson:"picture" form:"-" binding:"-"`
	ViewCount    int                `json:"viewCount" bson:"viewCount" form:"-" binding:"-"`
	CollectCount int                `json:"collectCount" bson:"collectCount" form:"-" binding:"-"`
	Status       CommodityStatus    `json:"status" bson:"status" form:"-" binding:"-"`
}

/****************************************/

type CommoditySearchResult struct {
	ID       primitive.ObjectID `json:"id"`
	Title    string             `json:"title"`
	Price    float64            `json:"price"`
	Category CategoryIndex      `json:"category"`
	Picture  string             `json:"picture"`
}
