package reposvc

import (
	"context"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/app/usersvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"

	"go.uber.org/zap"
)

// Middleware utilises the decorator pattern to add new functionality to the reposvc
type Middleware func(Service) Service

// Decorate will decorate the service with the given middlewares and returns the decorated Service
func Decorate(s Service, ms ...Middleware) Service {
	decorated := s
	for _, decorator := range ms {
		decorated = decorator(decorated)
	}
	return decorated
}

// LoggingMiddleware is a middleware that decorates the service with the logging middleware
func LoggingMiddleware(l *logger.Logger) Middleware {
	return func(s Service) Service {
		return &loggingMiddleware{
			next:   s,
			logger: l,
		}
	}
}

type loggingMiddleware struct {
	logger *logger.Logger
	next   Service
}

func (l *loggingMiddleware) BulkUpsert(ctx context.Context, repos []github.Repo) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, l.logger,
			logger.Method("BulkUpsert"),
			logger.Duration(start))

		logger.Maybe(L, "bulk upsert repos", err)
	}(time.Now())
	return l.next.BulkUpsert(ctx, repos)
}

func (l *loggingMiddleware) Count(ctx context.Context) (c int, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, l.logger,
			logger.Method("Count"),
			logger.Duration(start))

		logger.Maybe(L, "get repo count", err)
	}(time.Now())
	return l.next.Count(ctx)
}

func (l *loggingMiddleware) LastCreatedBy(ctx context.Context, login string) (date string, ok bool) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, l.logger)
		L.Info("find last created by user",
			zap.String("method", "LastCreatedBy"),
			zap.Duration("took", time.Since(start)),
			zap.String("date", date),
			zap.Bool("default", ok))
	}(time.Now())
	return l.next.LastCreatedBy(ctx, login)
}

func (l *loggingMiddleware) MostPopularLanguage(ctx context.Context, limit int) (res []schema.LanguageCount, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, l.logger,
			logger.Method("MostPopularLanguage"),
			logger.Duration(start),
			zap.Int("limit", limit))

		logger.Maybe(L, "get most popular language", err)
	}(time.Now())
	return l.next.MostPopularLanguage(ctx, limit)
}

func (l *loggingMiddleware) MostRecent(ctx context.Context, limit int) (res []schema.Repo, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, l.logger,
			logger.Method("MostRecent"),
			logger.Duration(start))

		logger.Maybe(L, "get most recent repos", err)
	}(time.Now())
	return l.next.MostRecent(ctx, limit)
}

func (l *loggingMiddleware) MostRecentReposByLanguage(ctx context.Context, language string, limit int) (res []schema.Repo, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, l.logger,
			logger.Method("MostRecentReposByLanguage"),
			logger.Duration(start),
			zap.String("language", language),
			zap.Int("limit", limit))

		logger.Maybe(L, "get most recent repos by language", err)
	}(time.Now())
	return l.next.MostRecentReposByLanguage(ctx, language, limit)
}

func (l *loggingMiddleware) MostStars(ctx context.Context, limit int) (res []schema.Repo, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, l.logger,
			logger.Method("MostStars"),
			logger.Duration(start),
			zap.Int("limit", limit))

		logger.Maybe(L, "get repos with most stars", err)
	}(time.Now())
	return l.next.MostStars(ctx, limit)
}

func (l *loggingMiddleware) MostForks(ctx context.Context, limit int) (res []schema.Repo, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, l.logger,
			logger.Method("MostForks"),
			logger.Duration(start),
			zap.Int("limit", limit))

		logger.Maybe(L, "get repos with most forks", err)
	}(time.Now())
	return l.next.MostForks(ctx, limit)
}

func (l *loggingMiddleware) RepoCountByUser(ctx context.Context, limit int) (res []schema.UserCount, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, l.logger,
			logger.Method("RepoCountByUser"),
			logger.Duration(start),
			zap.Int("limit", limit))

		logger.Maybe(L, "get repo count by user", err)
	}(time.Now())
	return l.next.RepoCountByUser(ctx, limit)
}

func (l *loggingMiddleware) ReposByLanguage(ctx context.Context, language string, limit int) (res []schema.UserCount, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, l.logger,
			logger.Method("ReposByLanguage"),
			logger.Duration(start),
			zap.String("language", language),
			zap.Int("limit", limit))

		logger.Maybe(L, "get repos by language", err)
	}(time.Now())
	return l.next.ReposByLanguage(ctx, language, limit)
}

func (l *loggingMiddleware) Distinct(ctx context.Context, field string) (res []string, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, l.logger,
			logger.Method("Distinct"),
			logger.Duration(start),
			zap.String("field", field))

		logger.Maybe(L, "get distinct login", err)
	}(time.Now())
	return l.next.Distinct(ctx, field)
}

func (l *loggingMiddleware) GetProfile(ctx context.Context, login string) (p *usersvc.User, err error) {
	L := logger.Wrap(ctx, l.logger)
	defer func(start time.Time) {
		L.Info("get profile",
			logger.Method("GetProfile"),
			logger.Duration(start),
			zap.String("login", login),
			zap.Int64("watchers", p.Watchers),
			zap.Int64("stargazers", p.Stargazers),
			zap.Int64("forks", p.Forks),
			zap.Int("keywords", len(p.Keywords)),
			zap.Int("languages", len(p.Languages)))
	}(time.Now())
	return l.next.GetProfile(ctx, login)
}
