package database

import (
	"MallSystem/model"
	_ "log"

	"go.mongodb.org/mongo-driver/bson"
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

func InsertOneUser(u *model.UserInfo) error {
	ctx, cancel := makeContext()
	defer cancel()
	if _, err := userCol.InsertOne(ctx, *u); err != nil {
		if ctx.Err() != nil {
			return ctx.Err()
		} else {
			return err
		}
	}
	return nil
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
