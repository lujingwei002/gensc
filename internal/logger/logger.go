package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

func (l *Logger) Infow(msg string, kvs ...interface{}) {
	l.sugar.Infow(msg, kvs...)
}

func (l *Logger) Errorw(msg string, kvs ...interface{}) {
	l.sugar.Errorw(msg, kvs...)
}

func (l *Logger) Debugw(msg string, kvs ...interface{}) {
	l.sugar.Debugw(msg, kvs...)
}

func (l *Logger) Fatalw(msg string, kvs ...interface{}) {
	l.sugar.Fatalw(msg, kvs...)
}

func (l *Logger) Warnw(msg string, kvs ...interface{}) {
	l.sugar.Warnw(msg, kvs...)
}

func (l *Logger) Info(args ...interface{}) {
	l.sugar.Info(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.sugar.Error(args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.sugar.Debug(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.sugar.Fatal(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.sugar.Warn(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.sugar.Infof(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.sugar.Errorf(format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.sugar.Debugf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.sugar.Fatalf(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.sugar.Warnf(format, args...)
}

const (
	EncoderKey_time       = "time"
	EncoderKey_level      = "level"
	EncoderKey_name       = "logger"
	EncoderKey_caller     = "caller"
	EncoderKey_message    = "msg"
	EncoderKey_stacktrace = "stacktrace"
)

func NewDefaultLogger() *Logger {
	logger, _ := zap.NewDevelopment()
	logger = logger.WithOptions(zap.AddCallerSkip(2))
	sugar := logger.Sugar()
	l := &Logger{
		logger: logger,
		sugar:  sugar,
	}
	return l
}
