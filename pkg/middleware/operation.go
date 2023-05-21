package middleware

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	pbAny "github.com/xiaohubai/go-grpc-layout/api/any/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/pkg/email"
	"github.com/xiaohubai/go-grpc-layout/pkg/jwt"
	"github.com/xiaohubai/go-grpc-layout/pkg/kafka"
	"github.com/xiaohubai/go-grpc-layout/pkg/metric"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
)

// Operation 记录请求流水
func Operation() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		claims, _ := c.Get("claims")
		userInfo := claims.(*jwt.Claims)
		uid := userInfo.UID
		reqBody, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer

		c.Next()

		var g errgroup.Group
		g.Go(func() error {
			record := pbAny.OperationRecord{
				Uid:      uid,
				DateTime: time.Now().Local().Format(time.DateTime),
				Ip:       c.ClientIP(),
				Method:   c.Request.Method,
				Path:     c.Request.RequestURI,
				Agent:    c.Request.UserAgent(),
				Status:   int32(c.Writer.Status()),
				Latency:  time.Since(start).String(),
				ReqBody:  string(reqBody),
				RespBody: writer.body.String(),
				TraceID:  tracing.TraceID(c.Request.Context()),
			}
			producer, err := kafka.NewProducer(consts.KafkaTopicOperationRecord)
			if err != nil {
				return err
			}
			return producer.Send(record.String())
		})
		if err := g.Wait(); err != nil {
			email.SendWarn(c.Request.Context(), consts.EmailTitleKafkaProducer, err.Error())
			metric.Count.With(fmt.Sprintf("producer_%s_error", consts.KafkaTopicOperationRecord)).Inc()
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
