package http

import (
	"go-api/internal/pkg/logger"
	"go-api/internal/http/middleware"
	"github.com/gin-gonic/gin"

	"fmt"
)

type HTTP struct {
	gin *gin.Engine
	config *Config
}

type Config struct {
	Logger *logger.LogHandler
	Port string
}

func New(conf *Config) (*HTTP, error) {
	app := gin.New()
	httpl := middleware.HTTPLogger{Logger: conf.Logger}

	app.Use(httpl.RequestLog)
	app.Use(gin.CustomRecovery(middleware.RecoverFromPanic(&httpl)))

	return &HTTP{app, conf}, nil
}

func (h *HTTP) Run() {
	if h.config.Port != "" {
		h.gin.Run(fmt.Sprintf(":%s", h.config.Port))
	} else {
		h.gin.Run()
	}
}
