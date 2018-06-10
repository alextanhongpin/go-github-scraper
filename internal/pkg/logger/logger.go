package logger

import (
	"context"
	"log"
	"os"
	"time"

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
func Wrap(ctx context.Context, l *Logger, fields ...zap.Field) *Logger {
	if v := ctx.Value(RequestID); v != nil {
		fields = append(fields, zap.String("requestId", v.(string)))
		return l.With(fields...)
	}
	return l
}

// Method is a custom logger field to returns the field Method
func Method(m string) zap.Field {
	return zap.String("method", m)
}

// Duration is a custom logger field to return the field Duration
func Duration(s time.Time) zap.Field {
	return zap.Duration("took", time.Since(s))
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

	// TODO: Inject application version so that it is visible in the logs
	zap.ReplaceGlobals(l)
	return l
}
