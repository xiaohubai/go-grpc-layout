package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-grpc-layout/pkg/metric"
)

// Metrics returns a gin.HandlerFunc for exporting some Web metrics
func Metrics(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		status := fmt.Sprintf("%d", c.Writer.Status())
		path := c.Request.URL.Path
		method := c.Request.Method

		labels := []string{serviceName, status, path, method}

		// no response content will return -1
		respSize := c.Writer.Size()
		if respSize < 0 {
			respSize = 0
		}

		metric.ReqCount.With(labels...).Inc()
		metric.CurReqCount.With(labels...).Inc()
		metric.ReqDuration.With(labels...).Observe((time.Since(start).Seconds()))

	}
}
