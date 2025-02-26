package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func TraceLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Process request
		c.Next()

		// Stop timer
		latency := time.Since(start).Milliseconds()
		statusCode := c.Writer.Status()

		data := fmt.Sprintf("route %s:(%s) finished with http-%d. latency: %d", method, path, statusCode, latency)
		Log(c).Info(data)
	}
}
