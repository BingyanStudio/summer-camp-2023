package database

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
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

func makeContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
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
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	client = c
	db = client.Database(database)

	initUserCollection()

}
