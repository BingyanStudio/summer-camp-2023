package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	cartColName string = "carts"
)

var (
	cartCol *mongo.Collection
)

func initCartCollection() {
	cartCol = db.Collection(cartColName)
}

func AddOneToCart(filter *bson.M, id *primitive.ObjectID) error {
	update := &bson.M{
		"$push": bson.M{"itemsid": *id},
	}
	ctx, cancel := makeContext()
	defer cancel()
	_, err := cartCol.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// func QueryOneCart(filter *bson.M) (*model.Cart, error) {
// 	ctx, cancel := makeContext()
// 	defer cancel()
// 	// 下班
// }

// func DeleteOneInCart(filter *bson.M) error {
// 	// 下班
// }
