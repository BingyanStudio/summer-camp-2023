package database

import (
	"MallSystem/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

func InsertOneCommodity(c *model.CommodityInfo) error {
	ctx, cancel := makeContext()
	defer cancel()
	if _, err := commodityCol.InsertOne(ctx, *c); err != nil {
		return err
	}
	return nil
}

func QueryCommodities(filter *bson.M, opts ...*options.FindOptions) ([]*model.CommodityInfo, error) {
	ctx, cancel := makeContext()
	defer cancel()
	cur, err := commodityCol.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	slice := make([]*model.CommodityInfo, 0)
	for cur.Next(context.Background()) {
		c := model.CommodityInfo{}
		cur.Decode(&c)
		slice = append(slice, &c)
	}
	return slice, nil
}
