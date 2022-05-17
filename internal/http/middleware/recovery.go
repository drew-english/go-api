package middleware

import (
	"net/http"
	"runtime/debug"

	"go-api/internal/pkg/logger"

	"github.com/gin-gonic/gin"
)

func RecoverFromPanic() gin.RecoveryFunc {
	return func(c *gin.Context, recovered interface{}) {
		extraTags := map[string]interface{} {
			"stackTrace": string(debug.Stack()),
		}
		
		if err, ok := recovered.(string); ok {
			extraTags["error"] = err
			c.String(http.StatusInternalServerError, "Internal server error")
		}

		logger.ErrorWithExtra(extraTags, "Request panic recovered")
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
