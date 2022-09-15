package db

import (
	"context"
	"fmt"

	"logger-service/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func ConnectDB(config config.Config) (*mongo.Client, error) {
	credential := options.Credential{
		Username: config.MongoUser,
		Password: config.MongoPass,
	}

	opts := options.Client().ApplyURI(config.DbUrl).SetAuth(credential)

	var err error
	client, err = mongo.Connect(context.TODO(), opts)

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")
	return client, err
}
