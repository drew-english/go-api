package db

import (
	ctx "context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDX_ALREADY_EXISTS_ERROR struct {
	error
}

type Config struct {
	DBName string
	Host string
	Port string
}

func (conf *Config) uri() string {
	return fmt.Sprintf("mongodb://%s:%s", conf.Host, conf.Port)
}

func New(conf *Config) *mongo.Database {
	client, err := mongo.Connect(ctx.TODO(), options.Client().ApplyURI(conf.uri()))
	if err != nil {
		panic(err)
	}

	if client.Ping(ctx.TODO(), nil) != nil {
		panic("Error connecting to DB")
	}

	return client.Database(conf.DBName)
}


func CreateIndex(coll *mongo.Collection, column string) (string, error) {
	if idxName := fmt.Sprintf("%s_1", column); findIndex(coll, idxName) {
		return idxName, IDX_ALREADY_EXISTS_ERROR{}
	}

	return coll.Indexes().CreateOne(
		ctx.TODO(),
		mongo.IndexModel{
			Keys: bson.D{{ Key: column, Value: 1 }},
			Options: options.Index().SetUnique(true),
		},
	)
}

func findIndex(coll *mongo.Collection, name string) bool {
	cursor, _ := coll.Indexes().List(ctx.TODO())
	var idxs []bson.M
	cursor.All(ctx.TODO(), &idxs)

	for _, idx := range idxs {
		if idx["name"] == name {
			return true
		}
	}

	return false
}
