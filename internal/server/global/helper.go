package global

import (
	"go.uber.org/zap/zapcore"

	"github.com/gitploy-io/gitploy/pkg/e"
)

func GetZapLogLevel(err error) zapcore.Level {
	if !e.IsError(err) {
		return zapcore.ErrorLevel
	}

	if err.(*e.Error).Code == e.ErrorCodeInternalError {
		return zapcore.ErrorLevel
	}

	return zapcore.WarnLevel
}
