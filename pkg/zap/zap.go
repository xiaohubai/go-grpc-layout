package zap

import (
	"fmt"
	"os"
	"path"
	"time"

	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var level zapcore.Level

// Zap 日志组件
func Init() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.Cfg.Zap.Director); !ok { // 判断是否有日志文件夹
		fmt.Printf("create %v directory\n", global.Cfg.Zap.Director)
		_ = os.Mkdir(global.Cfg.Zap.Director, os.ModePerm)
	}
	switch global.Cfg.Zap.Level {
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	if level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if global.Cfg.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		panic(fmt.Errorf("Get Write Syncer Failed err:%v", err.Error()))
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.Cfg.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (cfg zapcore.EncoderConfig) {
	cfg = zapcore.EncoderConfig{
		MessageKey:     "trace_id",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.Cfg.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,     //默认换行
		EncodeLevel:    zapcore.LowercaseLevelEncoder, //小写
		EncodeTime:     CustomTimeEncoder,             //输出时间
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, //记录调用路径
	}
	switch {
	case global.Cfg.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		cfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.Cfg.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		cfg.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.Cfg.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.Cfg.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		cfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return cfg
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.Cfg.Zap.Prefix + "2006-01-02 15:04:05.000"))
}

// GetWriteSyncer 日志分割
func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(global.Cfg.Zap.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(global.Cfg.Zap.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),     //日志清除时间
		zaprotatelogs.WithRotationTime(24*time.Hour), //日志文件创建时间
	)
	if global.Cfg.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
