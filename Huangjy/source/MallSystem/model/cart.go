package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
 *	购物车结构体
 *	购物车ID
 *	用户ID
 *	所有的商品ID
 */
type Cart struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty"`
	UserID  primitive.ObjectID   `bson:"userid"`
	ItemIDs []primitive.ObjectID `bson:"itemids"`
}
