package http

import (
	"go-api/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	
	"time"
	"fmt"
)

type HTTPLogger struct {
	logger *logger.LogHandler
}

func (l *HTTPLogger) New() gin.HandlerFunc {
	return l.requestLog
}

func (l *HTTPLogger)requestLog(c *gin.Context) {
	startResponse := time.Now()
	c.Next()
	responseDuration := time.Since(startResponse)

	l.logger.InfoWithExtra(map[string]interface{} {
		"duration": fmt.Sprintf("%.2fms", float64(responseDuration.Nanoseconds() / 1e6)),
		"method": c.Request.Method,
		"path": c.Request.RequestURI,
		"status": c.Writer.Status(),
	})
}
