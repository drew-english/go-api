package http

import (
	"fmt"

	"go-api/internal/http/middleware"
	"go-api/internal/models"

	"github.com/gin-gonic/gin"
)

type HTTP struct {
	gin *gin.Engine
	config *Config
}

type Config struct {
	Port string
	ModelServices *models.Services
}

func New(conf *Config) (*HTTP, error) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()

	app.Use(middleware.RequestLog)
	app.Use(gin.CustomRecovery(middleware.RecoverFromPanic()))
	ConfigRoutes(app, conf.ModelServices)

	return &HTTP{app, conf}, nil
}

func (h *HTTP) Run() {
	if h.config.Port != "" {
		h.gin.Run(fmt.Sprintf(":%s", h.config.Port))
	} else {
		h.gin.Run()
	}
}
