package main

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
)

type CtxKey string

func NewRequestID(ctx context.Context, key CtxKey) context.Context {
	if v := ctx.Value(key); v != nil {
		log.Println("found vlaue", v)
		return ctx
	}

	reqID, _ := uuid.NewV4()
	ctx = context.WithValue(ctx, CtxKey("RequestID"), reqID)
	return ctx
}

func main() {

	ctx := context.Background()
	// ctx = context.WithValue(ctx, CtxKey("RequestID"), "123456")

	ctx = NewRequestID(ctx, CtxKey("RequestID"))
	log.Println(ctx.Value("RequestID"))
}
