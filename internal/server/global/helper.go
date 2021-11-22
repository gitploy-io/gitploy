package global

import (
	"go.uber.org/zap/zapcore"

	"github.com/gitploy-io/gitploy/pkg/e"
)

// GetZapLogLevel return the warning level if the error is managed in the system.
func GetZapLogLevel(err error) zapcore.Level {
	if !e.IsError(err) {
		return zapcore.ErrorLevel
	}

	return zapcore.WarnLevel
}
