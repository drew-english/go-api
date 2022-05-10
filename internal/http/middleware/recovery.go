package middleware

import (
	"github.com/gin-gonic/gin"
	
	"net/http"
	"runtime/debug"
)


func RecoverFromPanic(httpl *HTTPLogger) gin.RecoveryFunc {
	return func(c *gin.Context, recovered interface{}) {
		extraTags := map[string]interface{} {
			"stackTrace": string(debug.Stack()),
		}
		
		if err, ok := recovered.(string); ok {
			extraTags["error"] = err
			c.String(http.StatusInternalServerError, "Internal server error")
		}

		httpl.Logger.ErrorWithExtra(extraTags, "Request panic recovered")
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
