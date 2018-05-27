package util

import (
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// LoggerWithRequestID returns a logger with requestId field set
func LoggerWithRequestID(zlog *zap.Logger) *zap.Logger {
	requestID, err := uuid.NewV4()
	if err != nil {
		zlog.Warn("error generating uuid", zap.Error(err))
		return zlog
	}
	return zlog.WithOptions(zap.Fields(
		zap.String("requestId", requestID.String())))
}
