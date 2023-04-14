package zap

import (
	"fmt"
	"time"

	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/xiaohubai/go-grpc-layout/configs"

	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func New(c *configs.Zap, g *configs.Global) (log.Logger, error) {
	encoderConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "ts",
		LineEnding:     zapcore.DefaultLineEnding,                              //默认换行
		EncodeLevel:    zapcore.LowercaseLevelEncoder,                          //小写
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"), //输出时间
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, //记录调用路径
	}

	fileWriter := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", c.Filename, time.Now().Format(time.DateOnly)), //文件名
		MaxSize:    int(c.MaxSize),                                                         //M
		MaxBackups: int(c.MaxBackups),                                                      //副本
		MaxAge:     int(c.MaxAge),                                                          //天
		Compress:   c.Compress,                                                             //压缩
	}

	var level zapcore.Level
	if err := level.UnmarshalText([]byte(c.Level)); err != nil {
		return nil, err
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(fileWriter), level)
	lg := zap.New(core).WithOptions()
	logger := kzap.NewLogger(lg)

	return log.With(
		logger,
		"env", g.Env,
		"service_id", g.Id,
		"service_name", g.AppName,
		"service_version", g.Version,
		"trace_id", tracing.TraceID(),
		"caller", log.DefaultCaller,
	), nil
}
