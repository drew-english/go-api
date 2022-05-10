package http

import (
	"go-api/internal/pkg/logger"
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

func New(c *Config) (*HTTP, error) {
	app := gin.New()

	app.Use((&HTTPLogger{c.Logger}).New())

	return &HTTP{app, c}, nil
}

func (h *HTTP) Run() {
	if h.config.Port != "" {
		h.gin.Run(fmt.Sprintf(":%s", h.config.Port))
	} else {
		h.gin.Run()
	}
}
