package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logLevel = zapcore.InfoLevel
)

func New(name string) *zap.SugaredLogger {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level.SetLevel(logLevel)
	development, _ := cfg.Build()
	return development.Sugar().Named(name)
}

func SetLogLevel(level string) {
	parseLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return
	}
	logLevel = parseLevel
}
