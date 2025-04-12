package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetClientMongoDB() (*mongo.Client, error) {
	config := NewConfig()
	clientOptions := options.Client().ApplyURI(config.CreateURIMongoDB())
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil

}
