package log

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
)

// Logger 日志器结构
type Logger struct {
	*zap.Logger
	*zap.SugaredLogger
}

// GetLogger 获取日志器
func GetLogger() *Logger {
	return &Logger{
		Logger:        logger,
		SugaredLogger: sugarLogger,
	}
}

// InitLogger 初始化日志系统
func InitLogger() {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "app"
	}
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	writerSyncer := getLogWriter(appName)
	encoder := getEncoder()
	level := getLogLevel(appEnv)
	core := zapcore.NewCore(encoder, writerSyncer, level)

	// zap.AddCaller() 添加调用函数信息
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()
}

// getEncoder 指定编码器（如何写入日志）
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 修改时间编码器
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志文件大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogLevel 根据环境设置日志等级
func getLogLevel(mode string) zapcore.Level {
	switch mode {
	case "prod", "production":
		return zap.InfoLevel
	default:
		return zap.DebugLevel
	}
}

// getLogWriter 获取写入器（文件 + 控制台）
func getLogWriter(appName string) zapcore.WriteSyncer {
	// 使用绝对路径
	execPath, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("cannot get executable path: %v", err))
	}

	logDir := filepath.Join(filepath.Dir(execPath), "logs")
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

// Info 记录信息级别日志
func Info(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Info(msg, fields...)
	}
}

// Warn 记录警告级别日志
func Warn(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Warn(msg, fields...)
	}
}

// Error 记录错误级别日志
func Error(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Error(msg, fields...)
	}
}

// Debug 记录调试级别日志
func Debug(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Debug(msg, fields...)
	}
}

// Infof 支持格式化信息级别日志
func Infof(template string, args ...interface{}) {
	if sugarLogger != nil {
		sugarLogger.Infof(template, args...)
	}
}

// Warnf 支持格式化警告级别日志
func Warnf(template string, args ...interface{}) {
	if sugarLogger != nil {
		sugarLogger.Warnf(template, args...)
	}
}

// Errorf 支持格式化错误级别日志
func Errorf(template string, args ...interface{}) {
	if sugarLogger != nil {
		sugarLogger.Errorf(template, args...)
	}
}

// Debugf 支持格式化调试级别日志
func Debugf(template string, args ...interface{}) {
	if sugarLogger != nil {
		sugarLogger.Debugf(template, args...)
	}
}

// Sync 清理日志缓冲
func Sync() {
	if logger != nil {
		_ = logger.Sync()
	}
}
