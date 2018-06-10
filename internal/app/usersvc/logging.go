package usersvc

import (
	"context"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
	"go.uber.org/zap"
)

// Logging adds logging capabilities to the service
func Logging(l *logger.Logger) Middleware {
	return func(s Service) Service {
		return &loggingMiddleware{
			service: s,
			logger:  l,
		}
	}
}

type loggingMiddleware struct {
	logger  *logger.Logger
	service Service
}

func (m *loggingMiddleware) FindLastCreated(ctx context.Context) (lastCreated string, ok bool) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("FindLastCreated"),
			logger.Duration(start))

		L.Info("get last created user",
			zap.String("date", lastCreated),
			zap.Bool("default", ok))
	}(time.Now())

	return m.service.FindLastCreated(ctx)
}

func (m *loggingMiddleware) BulkUpsert(ctx context.Context, users []github.User) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("BulkUpsert"),
			logger.Duration(start),
			zap.Int("count", len(users)))

		logger.Maybe(L, "bulk upsert users", err)
	}(time.Now())

	return m.service.BulkUpsert(ctx, users)
}

func (m *loggingMiddleware) FindLastFetched(ctx context.Context, limit int) (res []User, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("FindLastFetched"),
			logger.Duration(start),
			zap.Int("limit", limit))

		logger.Maybe(L, "get last fetched user", err)
	}(time.Now())

	return m.service.FindLastFetched(ctx, limit)
}

func (m *loggingMiddleware) UpdateOne(ctx context.Context, login string) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateOne"),
			logger.Duration(start),
			zap.String("login", login))

		logger.Maybe(L, "update one user", err)
	}(time.Now())

	return m.service.UpdateOne(ctx, login)
}

func (m *loggingMiddleware) Count(ctx context.Context) (count int, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("Count"),
			logger.Duration(start))

		logger.Maybe(L, "get user count", err)
	}(time.Now())

	return m.service.Count(ctx)
}

func (m *loggingMiddleware) BulkUpdate(ctx context.Context, users []User) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("BulkUpdate"),
			logger.Duration(start),
			zap.Int("count", len(users)))

		logger.Maybe(L, "bulk update user", err)
	}(time.Now())

	return m.service.BulkUpdate(ctx, users)
}

func (m *loggingMiddleware) WithRepos(ctx context.Context, count int) (users []User, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("WithRepos"),
			logger.Duration(start),
			zap.Int("greaterThan", count))

		logger.Maybe(L, "get users with repos greater than", err)
	}(time.Now())

	return m.service.WithRepos(ctx, count)
}

func (m *loggingMiddleware) DistinctCompany(ctx context.Context) (res []string, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("DistinctCompany"),
			logger.Duration(start),
			zap.Int("count", len(res)))

		logger.Maybe(L, "get companies with certain counts", err)
	}(time.Now())

	return m.service.DistinctCompany(ctx)
}

func (m *loggingMiddleware) FindByCompany(ctx context.Context, company string) (res []schema.User, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("FindByCompany"),
			logger.Duration(start),
			zap.String("company", company))

		logger.Maybe(L, "get users by company", err)
	}(time.Now())

	return m.service.FindByCompany(ctx, company)
}

func (m *loggingMiddleware) AggregateCompany(ctx context.Context, min, max int) (res []schema.Company, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("AggregateCompany"),
			logger.Duration(start),
			zap.Int("min", min),
			zap.Int("max", max))

		logger.Maybe(L, "get companies", err)
	}(time.Now())

	return m.service.AggregateCompany(ctx, min, max)
}

func (m *loggingMiddleware) FindOne(ctx context.Context, login string) (res *User, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("FindOne"),
			logger.Duration(start),
			zap.String("login", login))

		logger.Maybe(L, "find one", err)
	}(time.Now())

	return m.service.FindOne(ctx, login)
}
