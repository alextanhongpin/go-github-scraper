package util

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

// ContextKey refers to the context key
type ContextKey string

func ContextWithRequestID(ctx context.Context) (context.Context, string) {
	key := ContextKey("RequestID")
	if v := ctx.Value(key); v != nil {
		return ctx, v.(string)
	}

	reqID, _ := uuid.NewV4()
	return context.WithValue(ctx, key, reqID.String()), reqID.String()
}
