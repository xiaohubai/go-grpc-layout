package middleware

import (
	"runtime"

	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/errors/v1"
	"github.com/xiaohubai/go-grpc-layout/pkg/response"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"

	"go.opentelemetry.io/otel/attribute"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx, span := tracing.NewSpan(c.Request.Context(), "panic")
				c.Request = c.Request.WithContext(ctx)
				defer span.End()

				buf := make([]byte, 64<<10)
				buf = buf[:runtime.Stack(buf, false)]

				span.SetAttributes(attribute.Key("painc").String(string(buf)))
				response.Fail(c, v1.Error_Fail, nil)
				c.Abort()
			}
		}()
		c.Next()
	}
}
