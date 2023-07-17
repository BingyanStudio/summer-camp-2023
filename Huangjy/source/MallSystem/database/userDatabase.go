package database

import (
	"MallSystem/model"
	_ "log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userColName string = "users"
)

var (
	userCol *mongo.Collection
)

func initUserCollection() {
	userCol = db.Collection(userColName)
}

func InsertOneUser(u *model.UserInfo) (*primitive.ObjectID, error) {
	return baseInsertOne(userCol, *u)
}

func IncreaseOneUserBeViewedCount(filter *bson.M) {
	ctx, cancel := makeContext()
	defer cancel()
	update := bson.M{
		"$inc": bson.M{"beViewedCount": 1},
	}
	userCol.UpdateOne(ctx, filter, update)
}

func QueryOneUser(filter *bson.M) (*model.UserInfo, error) {
	var u model.UserInfo
	ctx, cancel := makeContext()
	defer cancel()
	result := userCol.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		} else {
			return nil, err
		}
	}
	if err := result.Decode(&u); err != nil {
		return nil, err
	}
	return &u, nil
}
