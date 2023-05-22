package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/pkg/email"
	"github.com/xiaohubai/go-grpc-layout/pkg/jwt"
	"github.com/xiaohubai/go-grpc-layout/pkg/kafka"
	"github.com/xiaohubai/go-grpc-layout/pkg/metric"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils"
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
			var respbody map[string]interface{}
			var reqsBody map[string]interface{}
			_ = json.Unmarshal(writer.body.Bytes(), &respbody)
			_ = json.Unmarshal(reqBody, &reqsBody)
			record := map[string]interface{}{
				"uid":      uid,
				"dateTime": time.Now().Local().Format(time.DateTime),
				"ip":       c.ClientIP(),
				"method":   c.Request.Method,
				"path":     c.Request.RequestURI,
				"agent":    c.Request.UserAgent(),
				"status":   int32(c.Writer.Status()),
				"latency":  time.Since(start).String(),
				"reqBody":  reqsBody,
				"respBody": respbody,
				"traceID":  tracing.TraceID(c.Request.Context()),
			}
			producer, err := kafka.NewProducer(consts.KafkaTopicOperationRecord)
			if err != nil {
				return err
			}
			data := utils.JsonToMarshal(record)
			return producer.Send(data)
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
