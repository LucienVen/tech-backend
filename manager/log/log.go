package log

import (
	"fmt"
	"github.com/LucienVen/tech-backend/bootstrap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger(env *bootstrap.Env) {
	//logger, _ := zap.NewProduction()
	wirterSyncer := getLogWriter(env.AppName)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, wirterSyncer, zap.DebugLevel)

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

// 指定日志目录
func getLogWriter(appName string) zapcore.WriteSyncer {
	fileName := fmt.Sprintf("../logs/%s.log", appName)
	file, _ := os.Create(fileName)
	return zapcore.AddSync(file)
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
