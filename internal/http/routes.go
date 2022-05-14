package http

import (
	"go-api/internal/handlers"
	"go-api/internal/models"

	"github.com/gin-gonic/gin"
)

type requestMethod string
const (
	get requestMethod = "GET"
	post requestMethod = "POST"
	patch requestMethod = "PATCH"
	delete requestMethod = "DELETE"
)

type route struct {
	method requestMethod
	path string
	handler gin.HandlerFunc
}

func v1Routes(modelService *models.Services) *[]route {
	appSettingHandler := handlers.AppSettingHandler{AppSettings: modelService.AppSettings}

	return &[]route {
		{ post, "/app_setting", appSettingHandler.Create },
		{ get, "/app_setting/:name", appSettingHandler.Show },
		{ patch, "/app_setting/:name", appSettingHandler.Update },
		{ delete, "/app_setting/:name", appSettingHandler.Delete },
	}
}

func ConfigRoutes(engine *gin.Engine, modelService *models.Services) {
	apiV1 := engine.Group("/api/v1")

	for _, r := range *v1Routes(modelService) {
		apiV1.Handle(string(r.method), r.path, r.handler)
	}
}