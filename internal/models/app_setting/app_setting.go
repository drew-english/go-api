package app_setting

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

const APP_SETTING_COLLECTION = "appSettings"

type AppSetting struct {
	Name string	`json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
}

type AppSettings struct {
	Collection *mongo.Collection
}

func NewService(db *mongo.Database) *AppSettings {
	return &AppSettings{db.Collection(APP_SETTING_COLLECTION)}
}

func (as *AppSettings) Create(newSetting *AppSetting) (err error) {
	_, err = as.Collection.InsertOne(context.TODO(), newSetting)
	return
}

func (as *AppSettings) Find(name string) (*AppSetting, error) {
	var foundSetting AppSetting

	filter := bson.D{{ Key: "name", Value: name }}
	err := as.Collection.FindOne(context.TODO(), filter).Decode(&foundSetting)
	if err != nil {
		return nil, err
	}

	return &foundSetting, nil
}
