package logger

import (
	"context"
	"log"
	"os"

	"github.com/rs/xid"

	"go.uber.org/zap"
)

// Logger represents a type alias for the logger
type Logger = zap.Logger

// ContextKey represents the type of the context key
type ContextKey string

// RequestID represents the request id that is passed in the context
const RequestID = ContextKey("RequestID")

// WrapContextWithRequestID wraps the existing context with the requestID field before returning them
func WrapContextWithRequestID(ctx context.Context) context.Context {
	if v := ctx.Value(RequestID); v != nil {
		return ctx
	}
	return context.WithValue(ctx, RequestID, xid.New().String())
}

// Wrap takes a context and existing logger and returns the logger with the injected correlation request id
func Wrap(ctx context.Context, l *Logger) *Logger {
	if v := ctx.Value(RequestID); v != nil {
		return l.With(zap.String("requestId", v.(string)))
	}
	return l
}

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
