package util

import (
	"context"

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

func NewRequestIDWithContext(ctx context.Context, zlog *zap.Logger) (context.Context, *zap.Logger) {
	ctx, reqID := ContextWithRequestID(ctx)
	zlog = zlog.WithOptions(zap.Fields(
		zap.String("requestId", reqID)))
	return ctx, zlog
}
