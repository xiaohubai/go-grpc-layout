package logger

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/xiaohubai/go-grpc-layout/pkg/configs"
)

func New(serviceInfo *configs.ServiceInfo) log.Logger {
	return log.With(
		log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"env", serviceInfo.Env,
		"caller", log.DefaultCaller,
		"service.id", serviceInfo.Id,
		"service.name", serviceInfo.Name,
		"service.version", serviceInfo.Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
}
