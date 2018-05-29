package logger

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// ContextKey represents the type of the context key
type ContextKey string

// RequestID represents the request id that is passed in the context
const RequestID = ContextKey("RequestID")

// WrapContextWithRequestID wraps the existing context with the requestID field before returning them
func WrapContextWithRequestID(ctx context.Context) context.Context {
	if v := ctx.Value(RequestID); v != nil {
		return ctx
	}

	reqID, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return context.WithValue(ctx, RequestID, reqID.String())
}

// RequestIDFromContext returns a global zap logger with the injected fields
func RequestIDFromContext(ctx context.Context) *zap.Logger {
	if v := ctx.Value(RequestID); v != nil {
		return zap.L().With(zap.String("requestId", v.(string)))
	}
	return zap.L()
}
