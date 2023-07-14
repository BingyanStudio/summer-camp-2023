package database

import (
	"MallSystem/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	orderColName string = "orders"
)

var (
	orderCol *mongo.Collection
)

func initOrderCollection() {
	orderCol = db.Collection(orderColName)
}

func InsertOneOrder(o *model.OrderInfo) (*primitive.ObjectID, error) {
	return baseInsertOne(orderCol, *o)
}

func SetOneOrderStatus(filter *bson.M, update *bson.M) error {
	ctx, cancel := makeContext()
	defer cancel()
	_, err := orderCol.UpdateOne(ctx, filter, update)
	if err != nil {
		if ctx.Err() != nil {
			return ctx.Err()
		} else {
			return err
		}
	}
	return nil
}

func QueryOneOrder(filter *bson.M) (*model.OrderInfo, error) {
	ctx, cancel := makeContext()
	defer cancel()
	var o model.OrderInfo
	result := orderCol.FindOne(ctx, filter)
	if result.Err() != nil {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		} else {
			return nil, result.Err()
		}
	}
	if err := result.Decode(&o); err != nil {
		return nil, err
	}
	return &o, nil
}
