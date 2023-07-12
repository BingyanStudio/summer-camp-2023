package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
 *	购物车结构体
 *
 */
type Cart struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty"`
	UserID  primitive.ObjectID   `bson:"userid"`
	ItemIDs []primitive.ObjectID `bson:"itemids"`
}

type CartItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CartID      primitive.ObjectID `bson:"cartid"`
	CommodityID primitive.ObjectID `bson:"commodityid"`
	Count       int                `bson:"count"`
}
