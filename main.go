package main

import (
	"go-api/internal/http"
	"go-api/internal/models"
	"go-api/internal/pkg/db"
	"go-api/internal/pkg/logger"
)

func main() {
	// init logger
	appLogger := logger.New("go-api", "v1.0.0", 1)

	// init DB connection
	mongoDB := db.New(&db.Config{
		DBName: "main",
		Host:   "mongo",
		Port:   "27017",
	})

	// init model services
	modelServices := model_services.New(&model_services.Config{
		DB: mongoDB,
	})

	// init server
	app, err := http.New(&http.Config{
		Logger: appLogger,
		Port:   "3000",
		ModelServices: modelServices,
	})

	if err == nil {
		app.Run()
	}
}
