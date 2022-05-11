package middleware

import (
	"fmt"
	"time"

	"go-api/internal/pkg/logger"
	"github.com/gin-gonic/gin"	
)

type HTTPLogger struct {
	Logger *logger.LogHandler
}

func (httpl *HTTPLogger)RequestLog(c *gin.Context) {
	startResponse := time.Now()
	c.Next()
	responseDuration := time.Since(startResponse)

	httpl.Logger.InfoWithExtra(map[string]interface{} {
		"duration": fmt.Sprintf("%.2fms", float64(responseDuration.Nanoseconds() / 1e6)),
		"method": c.Request.Method,
		"path": c.Request.RequestURI,
		"status": c.Writer.Status(),
	})
}
