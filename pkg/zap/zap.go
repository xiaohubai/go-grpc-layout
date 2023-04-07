package zap

import (
	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/xiaohubai/go-grpc-layout/configs"
	"github.com/xiaohubai/go-grpc-layout/pkg/serviceInfo"

	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewZapLogger(c *configs.Zap, serviceInfo *serviceInfo.ServiceInfo) log.Logger {
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
		panic("load log failed")
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(fileWriter), level)
	lg := zap.New(core).WithOptions()
	logger := kzap.NewLogger(lg)

	return log.With(
		logger,
		"ts", log.DefaultTimestamp,
		"env", serviceInfo.Env,
		"service_id", serviceInfo.Id,
		"service_name", serviceInfo.Name,
		"service_version", serviceInfo.Version,
		"trace_id", tracing.TraceID(),
		"caller", log.DefaultCaller,
	)
}
