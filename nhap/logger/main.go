package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/TheZeroSlave/zapsentry"
	"go.elastic.co/apm/module/apmzap/v2"
)

var errorLogger *zap.SugaredLogger
var logger *zap.Logger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func GetLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func init() {
	level := GetLoggerLevel("debug")

	// Cấu hình encoder
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	// Khởi tạo write syncer cho console và file
	consoleSyncer := zapcore.Lock(os.Stdout)
	fileSyncer := getFileWriterSyncer("/var/log/test.log") // Thay đổi đường dẫn tới file log của bạn ở đây

	// Kết hợp các write syncer lại với nhau
	multiWriteSyncer := zapcore.NewMultiWriteSyncer(consoleSyncer, fileSyncer)

	// Tạo core của logger
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		multiWriteSyncer,
		level,
	)

	// Cấu hình Sentry
	cfgSentry := zapsentry.Configuration{
		Level: zapcore.ErrorLevel, // when to send message to sentry
		Tags: map[string]string{
			"component": "system",
		},
	}
	sentryCore, err := zapsentry.NewCore(cfgSentry, zapsentry.NewSentryClientFromDSN("https://650d55e7a43c408ca7bbb0fd5a061d5c@sentry.paas.vn/323"))
	if err != nil {
		fmt.Errorf("failed to init zap", zap.Error(err))
	}

	// Tạo logger
	logger = zap.New(
		core,
		zap.WrapCore((&apmzap.Core{}).WrapCore),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)

	// Đính kèm Sentry core vào logger
	zapsentry.AttachCoreToLogger(sentryCore, logger)

	// Khởi tạo SugaredLogger
	errorLogger = logger.Sugar()
}

// getFileWriterSyncer trả về một WriteSyncer để ghi log vào file
func getFileWriterSyncer(filePath string) zapcore.WriteSyncer {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		return nil
	}
	return zapcore.Lock(file)
}

// Các hàm ghi log thông thường
func WithTrace(message string, fields ...interface{}) {
	errorLogger.With(fields...).Debug(message)
}

func Debug(args ...interface{}) {
	errorLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	errorLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	errorLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	errorLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	errorLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	errorLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	errorLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	errorLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	errorLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	errorLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	errorLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	errorLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	errorLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	errorLogger.Fatalf(template, args...)
}
