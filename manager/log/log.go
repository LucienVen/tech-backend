package log

import (
	"fmt"
	"github.com/LucienVen/tech-backend/bootstrap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger(env *bootstrap.Env) {
	//logger, _ = zap.NewProduction()
	writerSyncer := getLogWriter(env.AppName)
	encoder := getEncoder()

	level := getLogLevel(env.AppEnv)
	core := zapcore.NewCore(encoder, writerSyncer, level)

	// zap.AddCaller() 添加调用函数信息
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()
}

// 指定编码器（如何写入日志）
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 修改时间编码器
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志文件大学字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 根据环境设置日志等级
func getLogLevel(mode string) zapcore.Level {
	switch mode {
	case "prod", "production":
		return zap.InfoLevel
	default:
		return zap.DebugLevel
	}
}

// 获取写入器（文件 + 控制台）
func getLogWriter(appName string) zapcore.WriteSyncer {
	logDir := "../logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(fmt.Sprintf("cannot create log directory: %v", err))
	}

	fileName := filepath.Join(logDir, fmt.Sprintf("%s.log", appName))
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("cannot open log file: %v", err))
	}

	fileSyncer := zapcore.AddSync(file)
	consoleSyncer := zapcore.AddSync(os.Stdout)
	return zapcore.NewMultiWriteSyncer(fileSyncer, consoleSyncer)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// 支持格式化日志（sugar）
func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}
func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}
func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}
func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}

// 清理缓冲
func Sync() {
	_ = logger.Sync()
}
