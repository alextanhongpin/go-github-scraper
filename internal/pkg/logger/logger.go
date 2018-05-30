package logger

import (
	"context"
	"log"
	"os"

	"github.com/rs/xid"

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
	reqID := xid.New()
	return context.WithValue(ctx, RequestID, reqID.String())
}

// RequestIDFromContext returns a global zap logger with the injected fields
func RequestIDFromContext(ctx context.Context) *zap.Logger {
	if v := ctx.Value(RequestID); v != nil {
		return zap.L().With(zap.String("requestId", v.(string)))
	}
	return zap.L()
}

type Logger = zap.Logger

// New returns a new logger
func New() *Logger {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	l = l.Named("main")

	// Inject hostname for discoverability when scaling services in docker
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("error getting hostname")
	} else {
		l = l.With(zap.String("hostname", hostname))
	}
	zap.ReplaceGlobals(l)
	return l
}
