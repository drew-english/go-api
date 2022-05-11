package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	DBName string

	Host string
	Port string
}

func (conf *Config) uri() string {
	return fmt.Sprintf("mongodb://%s:%s", conf.Host, conf.Port)
}

func New(conf *Config) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.uri()))
	if err != nil {
		panic(err)
	}

	return client.Database(conf.DBName)
}
