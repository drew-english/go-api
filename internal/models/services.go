package model_services

import (
	"go-api/internal/models/app_setting"
	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	AppSettings *app_setting.AppSettings
}

type Config struct {
	DB *mongo.Database
}

func New(conf *Config) *Services {
	appSettingService := app_setting.AppSettings{Collection: conf.DB.Collection(app_setting.APP_SETTING_COLLECTION)}
	
	return &Services{AppSettings: &appSettingService}
}
