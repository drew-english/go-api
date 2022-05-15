package http

import (
	"go-api/docs"
	"go-api/internal/handlers"
	"go-api/internal/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type requestMethod string
const (
	get requestMethod = "GET"
	put requestMethod = "PUT"
	post requestMethod = "POST"
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
		{ put, "/app_setting", appSettingHandler.Update },
		{ get, "/app_setting/:name", appSettingHandler.Show },
		{ delete, "/app_setting/:name", appSettingHandler.Delete },
	}
}

func ConfigRoutes(engine *gin.Engine, modelService *models.Services) {
	docs.SwaggerInfo.BasePath = "/api"
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiV1 := engine.Group("/api/v1")

	for _, r := range *v1Routes(modelService) {
		apiV1.Handle(string(r.method), r.path, r.handler)
	}
}