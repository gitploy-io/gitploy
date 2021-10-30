package global

import (
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/pkg/e"
)

// LogWithError handles the level of log by the error code.
func LogWithError(logger *zap.Logger, message string, err error) {
	ge, ok := err.(*e.Error)
	if !ok {
		logger.Error(message, zap.Error(err))
		return
	}

	if ge.Code == e.ErrorCodeInternalError {
		logger.Error(message, zap.Error(err))
		return
	}

	logger.Warn(message, zap.Error(err))
	return
}
