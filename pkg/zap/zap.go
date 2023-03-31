package zap

import (
	"os"

	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/xiaohubai/go-grpc-layout/configs"
	"github.com/xiaohubai/go-grpc-layout/pkg/serviceInfo"

	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func New(serviceInfo *serviceInfo.ServiceInfo) log.Logger {
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

func NewZapLogger(c *configs.Zap) log.Logger {
	encoderConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		LineEnding:     zapcore.DefaultLineEnding,     //默认换行
		EncodeLevel:    zapcore.LowercaseLevelEncoder, //小写
		EncodeTime:     zapcore.ISO8601TimeEncoder,    //输出时间
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, //记录调用路径
	}

	fileWriter := &lumberjack.Logger{
		Filename:   c.Filename,
		MaxSize:    int(c.MaxSize),
		MaxBackups: int(c.MaxBackups),
		MaxAge:     int(c.MaxAge),
	}
	zapcore.AddSync(fileWriter)

	var level zapcore.Level
	if err := level.UnmarshalText([]byte(c.Level)); err != nil {
		return nil
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(fileWriter), level)
	lg := zap.New(core).WithOptions()
	logger := kzap.NewLogger(lg)

	return log.With(logger, "ts", log.DefaultTimestamp, "trace_id", tracing.TraceID(), "caller", log.DefaultCaller)
}
