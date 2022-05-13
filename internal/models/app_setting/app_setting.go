package app_setting

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const APP_SETTING_COLLECTION = "appSettings"

type NOT_FOUND_ERROR struct {
	error
}

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
		if err == mongo.ErrNoDocuments {
			err = NOT_FOUND_ERROR{}
		}

		return nil, err
	}

	return &foundSetting, nil
}
