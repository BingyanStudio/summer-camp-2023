package database

import (
	"MallSystem/model"
	_ "log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	colName string = "users"
)

var (
	col *mongo.Collection
)

func initUserCollection() {
	col = db.Collection(colName)
}

func InsertOneUser(u *model.UserInfo) error {
	ctx, cancel := makeContext()
	defer cancel()
	if _, err := col.InsertOne(ctx, *u); err != nil {
		return err
	}
	return nil
}

func QueryOneUser(filter *bson.M) (*model.UserInfo, error) {
	var u model.UserInfo
	ctx, cancel := makeContext()
	defer cancel()
	result := col.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return nil, err
	}
	if err := result.Decode(&u); err != nil {
		return nil, err
	}
	return &u, nil
}
