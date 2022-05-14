package http

import (
	"fmt"

	"go-api/internal/http/middleware"
	"go-api/internal/models"
	"go-api/internal/pkg/logger"
	"github.com/gin-gonic/gin"
)

type HTTP struct {
	gin *gin.Engine
	config *Config
}

type Config struct {
	Logger *logger.LogHandler
	Port string
	ModelServices *models.Services
}

func New(conf *Config) (*HTTP, error) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	httpl := middleware.HTTPLogger{Logger: conf.Logger}

	app.Use(httpl.RequestLog)
	app.Use(gin.CustomRecovery(middleware.RecoverFromPanic(&httpl)))
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
