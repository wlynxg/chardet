package log

import (
	"go.uber.org/zap"
)

func New(name string) *zap.SugaredLogger {
	development, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return development.Sugar().Named(name)
}
