package database

import (
	"MallSystem/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	commodityColName string = "commodities"
)

var (
	commodityCol *mongo.Collection
)

func initCommodityCollection() {
	commodityCol = db.Collection(commodityColName)
}

func InsertOneCommodity(c *model.CommodityInfo) (*primitive.ObjectID, error) {
	return baseInsertOne(commodityCol, *c)
}

func QueryOneCommodity(filter *bson.M) (*model.CommodityInfo, error) {
	ctx, cancel := makeContext()
	defer cancel()
	var c model.CommodityInfo
	result := commodityCol.FindOne(ctx, filter)
	if result.Err() != nil {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		} else {
			return nil, result.Err()
		}
	}
	if err := result.Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func SetOneCommodityStatus(filter *bson.M, status model.CommodityStatus) error {
	ctx, cancel := makeContext()
	defer cancel()
	update := bson.M{
		"$set": bson.M{"status": status},
	}
	_, err := commodityCol.UpdateOne(ctx, filter, update)
	if err != nil {
		if ctx.Err() != nil {
			return ctx.Err()
		} else {
			return err
		}
	}
	return nil
}

func SetOneCommodityImage(filter *bson.M, src string) error {
	ctx, cancel := makeContext()
	defer cancel()
	update := bson.M{
		"$set": bson.M{"picture": src},
	}
	_, err := commodityCol.UpdateOne(ctx, filter, update)
	if err != nil {
		if ctx.Err() != nil {
			return ctx.Err()
		} else {
			return err
		}
	}
	return nil
}

func IncreaseOneCommodityViewCount(filter *bson.M) {
	ctx, cancel := makeContext()
	defer cancel()
	update := bson.M{
		"$inc": bson.M{"viewCount": 1},
	}
	commodityCol.UpdateOne(ctx, filter, update)
}

func QueryCommodities(filter *bson.M, opts ...*options.FindOptions) ([]*model.CommodityInfo, error) {
	ctx, cancel := makeContext()
	defer cancel()
	cur, err := commodityCol.Find(ctx, filter, opts...)
	if err != nil {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		} else {
			return nil, err
		}
	}
	slice := make([]*model.CommodityInfo, 0)
	for cur.Next(context.Background()) {
		c := model.CommodityInfo{}
		cur.Decode(&c)
		slice = append(slice, &c)
	}
	return slice, nil
}
