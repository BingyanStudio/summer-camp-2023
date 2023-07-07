package db

import (
	"context"
	"errors"
	"log"
	"myserver/model"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri string = "mongodb://admin:0000@127.0.0.1:27017"
)

var (
	db        *mongo.Database
	users_col *mongo.Collection
	info_col  *mongo.Collection
)

func init() {
	opt := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), opt)

	if err != nil {
		log.Fatal("Failed to connect to db, uri:", uri)
	} else {
		log.Println("successfully connecting to the db, uri:", uri)
	}
	db = client.Database("goserver")
	users_col = db.Collection("userinfo")
	info_col = db.Collection("others")
}

func InsertUser(u *model.UserInfo) error {
	filter := bson.D{{Key: "info", Value: "maxid"}}
	update := bson.D{{
		Key: "$inc", Value: bson.D{{
			Key: "value", Value: 1,
		}},
	}}
	result := info_col.FindOneAndUpdate(
		context.TODO(), filter, update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)
	var d bson.M
	if err := result.Decode(&d); err != nil {
		log.Println(err)
		return err
	}
	u.ID = uint32(d["value"].(int64))
	if _, err := users_col.InsertOne(context.TODO(), u); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func QueryUser(filter *bson.D, opts ...*options.FindOptions) ([]*model.UserInfo, error) {
	info, err := baseQuery(context.TODO(), filter, users_col, reflect.TypeOf(model.UserInfo{}), opts...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	users := make([]*model.UserInfo, len(info))
	for i, u := range info {
		p := (*u).(model.UserInfo)
		users[i] = &p
	}
	return users, nil
}

func UpdateUser(filter *bson.D, u *model.UserUpdate) error {
	update := bson.D{{
		Key: "$set", Value: bson.D{{
			Key: u.Info, Value: u.Value,
		}},
	}}
	result, err := users_col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	} else if result.MatchedCount == 0 {
		errors.New("No document matched the filter")
	}
	return nil
}

func DeleteUser(filter *bson.D) error {
	result, err := users_col.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return err
	} else if result.DeletedCount == 0 {
		return errors.New("No document matched the filter")
	}
	return nil
}

func QueryOtherInfo(info string) interface{} {
	result := info_col.FindOne(context.TODO(), &bson.D{{"info", info}})
	if result == nil {
		return nil
	}
	var result_map map[string]interface{}
	if err := result.Decode(&result_map); err != nil {
		log.Println(err)
		return nil
	}
	return result_map["value"]
}

func baseQuery(ctx context.Context, filter *bson.D, c *mongo.Collection, t reflect.Type, opts ...*options.FindOptions) ([]*interface{}, error) {
	cur, err := c.Find(ctx, filter, opts...)
	infos := []*interface{}{}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for cur.Next(context.TODO()) {
		p := reflect.New(t)
		if err = cur.Decode(p.Interface()); err != nil {
			log.Println(err)
			return infos, err
		}
		c := p.Elem().Interface()
		infos = append(infos, &c)
	}
	return infos, nil
}
