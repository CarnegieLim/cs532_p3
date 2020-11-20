package db

import (

	"log"
	"context"
	"time"

	
	"go.mongodb.org/mongo-driver/mongo"
   	"go.mongodb.org/mongo-driver/mongo/options"
   	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var ctx context.Context
var cancel_func context.CancelFunc
var db *mongo.Database

func init() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel_func = context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database("imdb")

}

func GetDB() (database *mongo.Database, context context.Context) {
	database = db
	context = ctx
	return
}

func Close() {
	client.Disconnect(ctx)
}