package models

import (
	"reflect"
	"regexp"
	"strings"

	"go-api/internal/models/app_setting"
	"go-api/internal/pkg/db"
	"go-api/internal/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	AppSettings *app_setting.AppSettings
}

type Config struct {
	DB *mongo.Database
}

func NewService(conf *Config) *Services {
	appSettingService := app_setting.NewService(conf.DB)
	initSchema(app_setting.AppSetting{}, appSettingService.Collection)
	
	return &Services{AppSettings: appSettingService}
}

func initSchema(model interface{}, coll *mongo.Collection) {
	t := reflect.TypeOf(model)
	schemaUniqueTag := regexp.MustCompile(`(^|,)unique($|,)`)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if schemaTag := field.Tag.Get("schema"); schemaTag != "" {
			if schemaUniqueTag.MatchString(schemaTag) {
				name, err := db.CreateIndex(coll, strings.Split(field.Tag.Get("bson"), ",")[0])

				switch err {
				case nil:
					logger.Info("Created index name=" + name)
				case db.IDX_ALREADY_EXISTS_ERROR{}:
					logger.Info("Skipped creating index name=" + name)
				default:
					logger.Error("Unable to create index", err)
				}
			}
		}
	}
}
