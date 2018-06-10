package github

import (
	"context"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"go.uber.org/zap"
)

// Middleware represents the middleware func
type (
	Middleware func(Service) Service

	loggingMiddleware struct {
		service Service
		logger  *logger.Logger
	}
)

// Decorate represents the decorator that adds capabilities to services with middlewares
func Decorate(s Service, ms ...Middleware) Service {
	decorated := s
	for _, m := range ms {
		decorated = m(decorated)
	}
	return decorated
}

// LoggingMiddleware adds logging capabilities to the service
func LoggingMiddleware(l *logger.Logger) Middleware {
	return func(s Service) Service {
		return &loggingMiddleware{
			service: s,
			logger:  l,
		}
	}
}

func (m *loggingMiddleware) FetchUsersCursor(ctx context.Context, location, start, end string, limit int) (users []User, err error) {
	defer func(s time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("FetchUsersCursor"),
			logger.Duration(s),
			zap.String("location", location),
			zap.String("start", start),
			zap.String("end", end),
			zap.Int("limit", limit))

		logger.Maybe(L, "fetch users cursor", err)
	}(time.Now())

	return m.service.FetchUsersCursor(ctx, location, start, end, limit)
}

func (m *loggingMiddleware) FetchReposCursor(ctx context.Context, login, start, end string, limit int) (repos []Repo, err error) {
	defer func(s time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("FetchReposCursor"),
			logger.Duration(s),
			zap.String("login", login),
			zap.String("start", start),
			zap.String("end", end),
			zap.Int("limit", limit))

		logger.Maybe(L, "fetch repos cursor", err)
	}(time.Now())

	return m.service.FetchReposCursor(ctx, login, start, end, limit)
}
