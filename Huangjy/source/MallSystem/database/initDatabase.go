package database

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	hostname string
	port     string
	username string
	password string
	database string
	timeout  int64
)

var (
	uri    string
	client *mongo.Client
	db     *mongo.Database
)

func MakeSession() (mongo.Session, error) {
	return client.StartSession()
}

func makeContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
}

func baseInsertOne(col *mongo.Collection, i interface{}, opts ...*options.InsertOneOptions) (*primitive.ObjectID, error) {
	ctx, cancel := makeContext()
	defer cancel()
	if result, err := col.InsertOne(ctx, i, opts...); err != nil {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		} else {
			return nil, err
		}
	} else {
		id := result.InsertedID.(primitive.ObjectID)
		return &id, nil
	}
}

func InitDatabase() {
	hostname = viper.GetString("MongoDB.HostName")
	port = viper.GetString("MongoDB.Port")
	username = viper.GetString("MongoDB.Username")
	password = viper.GetString("MongoDB.Password")
	database = viper.GetString("MongoDB.Database")
	timeout = viper.GetInt64("MongoDB.Timeout")

	uri = "mongodb://" + username + ":" + password + "@" + hostname + ":" + port

	ctx, cancel := makeContext()
	defer cancel()
	c, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if ctx.Err() != nil {
		log.Fatal(ctx.Err())
	}
	client = c
	db = client.Database(database)

	initUserCollection()
	initCommodityCollection()
	initOrderCollection()

}
