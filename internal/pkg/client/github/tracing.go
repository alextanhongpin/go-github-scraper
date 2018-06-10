package github

import (
	"context"

	"go.opencensus.io/trace"
)

// Tracing adds tracing capability to the service
func Tracing() Middleware {
	return func(s Service) Service {
		return &tracingMiddleware{
			service: s,
		}
	}
}

type tracingMiddleware struct {
	service Service
}

func (m *tracingMiddleware) FetchUsersCursor(ctx context.Context, location, start, end string, limit int) (users []User, err error) {
	ctx, span := trace.StartSpan(ctx, "FetchUsersCursor")
	defer span.End()

	span.AddAttributes(
		trace.StringAttribute("start", start),
		trace.StringAttribute("end", end),
		trace.Int64Attribute("limit", int64(limit)))

	return m.service.FetchUsersCursor(ctx, location, start, end, limit)
}
func (m *tracingMiddleware) FetchReposCursor(ctx context.Context, login, start, end string, limit int) (repos []Repo, err error) {
	ctx, span := trace.StartSpan(ctx, "FetchReposCursor")
	defer span.End()

	span.AddAttributes(
		trace.StringAttribute("start", start),
		trace.StringAttribute("end", end),
		trace.Int64Attribute("limit", int64(limit)))

	return m.service.FetchReposCursor(ctx, login, start, end, limit)
}
