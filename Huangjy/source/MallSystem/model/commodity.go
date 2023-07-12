package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryIndex int

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
 *	商品信息结构体
 *	商品ID
 *	发布者ID
 *	。。。
 */
type CommodityInfo struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID       primitive.ObjectID `json:"userid" bson:"userid"`
	Title        string             `json:"title" bson:"title"`
	Desc         string             `json:"desc" bson:"desc"`
	Price        float64            `json:"price" bson:"price"`
	Picture      string             `json:"picture" bson:"picture"`
	ViewCount    int                `json:"viewCount" bson:"viewCount"`
	CollectCount int                `json:"collectCount" bson:"collectCount"`
}

/*
 *	商品类别信息
 */
type CommodityCategory struct {
	CommodityID primitive.ObjectID
	Category    CategoryIndex
}
