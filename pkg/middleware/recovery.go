package middleware

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"

	"go.opentelemetry.io/otel/attribute"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx, span := tracing.NewSpan(c.Request.Context(), "recover")
				c.Request = c.Request.WithContext(ctx)
				defer span.End()

				buf := make([]byte, 64<<10)
				buf = buf[:runtime.Stack(buf, false)]
				bufs := string(buf)

				log.Errorw("traceId", tracing.TraceID(c.Request.Context()), "msg", "recover", "painc", bufs)

				span.SetAttributes(attribute.Key("err").String(fmt.Sprintf("%s", err)))
				span.SetAttributes(attribute.Key("painc").String(bufs))
				response.Fail(c, errors.Failed, nil)
				c.Abort()
			}
		}()
		c.Next()
	}
}
