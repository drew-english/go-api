package main

import (
	"go-api/internal/http"
	"go-api/internal/models"
	"go-api/internal/pkg/db"
	"go-api/internal/pkg/env"
)

func main() {
	// init DB connection
	mongoDB := db.New(&db.Config{
		DBName: env.GetEnvDefault("DB_DATABASE", "main"),
		Host: env.GetEnvDefault("DB_HOST", "mongo"),
		Port: env.GetEnvDefault("DB_PORT", "27017"),
	})

	// init model services
	modelServices := models.NewService(&models.Config{
		DB: mongoDB,
	})

	// init server
	app, err := http.New(&http.Config{
		Port:   env.GetEnvDefault("APP_PORT", "3000"),
		ModelServices: modelServices,
	})

	if err == nil {
		app.Run()
	}
}
