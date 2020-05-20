package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
)

var logger *zap.Logger
var sl *zap.SugaredLogger

func init() {
	logger = initLogger()
	sl = logger.Sugar()
}


func initLogger() *zap.Logger {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filename := path.Join(wd, "pg-exporter.log")

	lf := &lumberjack.Logger{
		Filename:   filename,
		MaxBackups: 10,
		MaxSize:    100,
		LocalTime:  true,
		Compress:   true,
	}

	fileLevel := zap.NewAtomicLevel()
	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool{
		return level > zapcore.WarnLevel
	})

	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	ws := zapcore.AddSync(lf)

	cores := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), ws, fileLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), os.Stderr, highPriority),
	)
	return zap.New(cores)
}

func Info(msg string, fields ...zap.Field){
	logger.Info(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	sl.Infof(template, args)
}

func Warn(msg string, fields ...zap.Field){
	logger.Warn(msg, fields...)
}

func Warnf(template string, args ...interface{}) {
	sl.Warnf(template, args)
}

func Error(msg string, fields ...zap.Field){
	logger.Error(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	sl.Errorf(template, args)
}
