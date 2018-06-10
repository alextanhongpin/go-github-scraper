package mediatorsvc

import (
	"context"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"

	"go.uber.org/zap"
)

// Middleware represents a decorator pattern
type Middleware func(Service) Service

// Decorate decorates a service with a list of provided middlewares
func Decorate(s Service, ms ...Middleware) Service {
	decorated := s
	for _, m := range ms {
		decorated = m(decorated)
	}
	return decorated
}

// LoggingMiddleware represents the decorator pattern to add logging to services
func LoggingMiddleware(l *logger.Logger) Middleware {
	return func(s Service) Service {
		return &loggingMiddleware{
			service: s,
			logger:  l,
		}
	}
}

type loggingMiddleware struct {
	service Service
	logger  *logger.Logger
}

func (m *loggingMiddleware) FetchUsers(ctx context.Context, location string, months int, perPage int) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("FetchUsers"),
			logger.Duration(start),
			zap.String("location", location),
			zap.Int("months", months),
			zap.Int("perPage", perPage))

		logger.Maybe(L, "fetch users", err)
	}(time.Now())

	return m.service.FetchUsers(ctx, location, months, perPage)
}

func (m *loggingMiddleware) FetchRepos(ctx context.Context, userPerPage, repoPerPage int, reset bool) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("FetchRepos"),
			logger.Duration(start),
			zap.Int("userPerPage", userPerPage),
			zap.Int("repoPerPage", repoPerPage))

		logger.Maybe(L, "fetch repos", err)
	}(time.Now())

	return m.service.FetchRepos(ctx, userPerPage, repoPerPage, reset)
}

func (m *loggingMiddleware) UpdateUserCount(ctx context.Context) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateUserCount"),
			logger.Duration(start))

		logger.Maybe(L, "update user count", err)
	}(time.Now())

	return m.service.UpdateUserCount(ctx)
}

func (m *loggingMiddleware) UpdateRepoCount(ctx context.Context) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateRepoCount"),
			logger.Duration(start))

		logger.Maybe(L, "update repo count", err)
	}(time.Now())

	return m.service.UpdateRepoCount(ctx)
}

func (m *loggingMiddleware) UpdateReposMostRecent(ctx context.Context, perPage int) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateReposMostRecent"),
			logger.Duration(start))

		logger.Maybe(L, "update repos most recent", err)
	}(time.Now())

	return m.service.UpdateReposMostRecent(ctx, perPage)
}

func (m *loggingMiddleware) UpdateRepoCountByUser(ctx context.Context, perPage int) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateRepoCountByUser"),
			logger.Duration(start),
			zap.Int("perPage", perPage))

		logger.Maybe(L, "update repo count by user", err)
	}(time.Now())

	return m.service.UpdateRepoCountByUser(ctx, perPage)
}

func (m *loggingMiddleware) UpdateReposMostStars(ctx context.Context, perPage int) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateReposMostStars"),
			logger.Duration(start),
			zap.Int("perPage", perPage))

		logger.Maybe(L, "update repos most stars", err)
	}(time.Now())

	return m.service.UpdateReposMostStars(ctx, perPage)
}

func (m *loggingMiddleware) UpdateReposMostForks(ctx context.Context, perPage int) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateReposMostForks"),
			logger.Duration(start),
			zap.Int("perPage", perPage))

		logger.Maybe(L, "update repos most forks", err)
	}(time.Now())

	return m.service.UpdateReposMostForks(ctx, perPage)
}

func (m *loggingMiddleware) UpdateLanguagesMostPopular(ctx context.Context, perPage int) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateLanguagesMostPopular"),
			logger.Duration(start),
			zap.Int("perPage", perPage))

		logger.Maybe(L, "update languages most popular", err)
	}(time.Now())

	return m.service.UpdateLanguagesMostPopular(ctx, perPage)
}

func (m *loggingMiddleware) UpdateMostRecentReposByLanguage(ctx context.Context, perPage int) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateMostRecentReposByLanguage"),
			logger.Duration(start),
			zap.Int("perPage", perPage))

		logger.Maybe(L, "update most recent repos by language", err)
	}(time.Now())

	return m.service.UpdateMostRecentReposByLanguage(ctx, perPage)
}

func (m *loggingMiddleware) UpdateReposByLanguage(ctx context.Context, perPage int) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateReposByLanguage"),
			logger.Duration(start),
			zap.Int("perPage", perPage))

		logger.Maybe(L, "update repos by language", err)
	}(time.Now())

	return m.service.UpdateReposByLanguage(ctx, perPage)
}

func (m *loggingMiddleware) UpdateProfile(ctx context.Context, numWorkers int) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateProfile"),
			logger.Duration(start),
			zap.Int("numWorkers", numWorkers))

		logger.Maybe(L, "update profile", err)
	}(time.Now())

	return m.service.UpdateProfile(ctx, numWorkers)
}

func (m *loggingMiddleware) UpdateMatches(ctx context.Context) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateMatches"),
			logger.Duration(start))

		logger.Maybe(L, "update matches", err)
	}(time.Now())

	return m.service.UpdateMatches(ctx)
}

func (m *loggingMiddleware) UpdateUsersByCompany(ctx context.Context, min, max int) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateUsersByCompany"),
			logger.Duration(start))

		logger.Maybe(L, "update users by company", err)
	}(time.Now())

	return m.service.UpdateUsersByCompany(ctx, min, max)
}

func (m *loggingMiddleware) UpdateCompanyCount(ctx context.Context) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("UpdateCompanyCount"),
			logger.Duration(start))

		logger.Maybe(L, "update company count", err)
	}(time.Now())

	return m.service.UpdateCompanyCount(ctx)
}
