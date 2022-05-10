package main

import (
	"go-api/internal/pkg/logger"
	"go-api/internal/http"
)

func main() {
	appLogger := logger.New("go-api", "v1.0.0", 1)

	app, ok := http.New(&http.Config{
		Logger: appLogger,
		Port: "3000",
	})

	if ok != nil {
		app.Run()
	}
}
