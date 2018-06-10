package usersvc

import (
	"context"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
	"go.opencensus.io/trace"
)

// Tracing decorates the service with tracing capabilities
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

func (m *tracingMiddleware) FindLastCreated(ctx context.Context) (string, bool) {
	ctx, span := trace.StartSpan(ctx, "FindLastCreated")
	defer span.End()

	return m.service.FindLastCreated(ctx)
}

func (m *tracingMiddleware) BulkUpsert(ctx context.Context, users []github.User) error {
	ctx, span := trace.StartSpan(ctx, "BulkUpsert")
	defer span.End()

	return m.service.BulkUpsert(ctx, users)
}

func (m *tracingMiddleware) FindLastFetched(ctx context.Context, limit int) ([]User, error) {
	ctx, span := trace.StartSpan(ctx, "FindLastFetched")
	defer span.End()

	return m.service.FindLastFetched(ctx, limit)
}

func (m *tracingMiddleware) UpdateOne(ctx context.Context, login string) error {
	ctx, span := trace.StartSpan(ctx, "UpdateOne")
	defer span.End()

	return m.service.UpdateOne(ctx, login)
}

func (m *tracingMiddleware) Count(ctx context.Context) (int, error) {
	ctx, span := trace.StartSpan(ctx, "ctx")
	defer span.End()

	return m.service.Count(ctx)
}

func (m *tracingMiddleware) BulkUpdate(ctx context.Context, users []User) error {
	ctx, span := trace.StartSpan(ctx, "BulkUpdate")
	defer span.End()

	return m.service.BulkUpdate(ctx, users)
}

func (m *tracingMiddleware) WithRepos(ctx context.Context, count int) ([]User, error) {
	ctx, span := trace.StartSpan(ctx, "WithRepos")
	defer span.End()

	return m.service.WithRepos(ctx, count)
}

func (m *tracingMiddleware) DistinctCompany(ctx context.Context) ([]string, error) {
	ctx, span := trace.StartSpan(ctx, "DistinctCompany")
	defer span.End()

	return m.service.DistinctCompany(ctx)
}

func (m *tracingMiddleware) FindByCompany(ctx context.Context, company string) ([]schema.User, error) {
	ctx, span := trace.StartSpan(ctx, "FindByCompany")
	defer span.End()

	return m.service.FindByCompany(ctx, company)
}

func (m *tracingMiddleware) AggregateCompany(ctx context.Context, min, max int) ([]schema.Company, error) {
	ctx, span := trace.StartSpan(ctx, "AggregateCompany")
	defer span.End()

	return m.service.AggregateCompany(ctx, min, max)
}

func (m *tracingMiddleware) FindOne(ctx context.Context, login string) (*User, error) {
	ctx, span := trace.StartSpan(ctx, "FindOne")
	defer span.End()

	return m.service.FindOne(ctx, login)
}
