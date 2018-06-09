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
		zlog := logger.Wrap(ctx, l.logger).
			With(zap.String("method", "BulkUpsert"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error upserting repos", zap.Error(err))
		} else {
			zlog.Info("inserted repos", zap.Int("count", len(repos)))
		}
	}(time.Now())
	return l.next.BulkUpsert(ctx, repos)
}

func (l *loggingMiddleware) Count(ctx context.Context) (c int, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, l.logger).
			With(zap.String("method", "Count"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error returning repo count", zap.Error(err))
		} else {
			zlog.Info("got repo count", zap.Int("count", c))
		}
	}(time.Now())
	return l.next.Count(ctx)
}

func (l *loggingMiddleware) LastCreatedBy(ctx context.Context, login string) (date string, ok bool) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, l.logger)
		zlog.Info("find last created by user",
			zap.String("method", "LastCreatedBy"),
			zap.Duration("took", time.Since(start)),
			zap.String("date", date),
			zap.Bool("default", ok))
	}(time.Now())
	return l.next.LastCreatedBy(ctx, login)
}

func (l *loggingMiddleware) MostPopularLanguage(ctx context.Context, limit int) (res []schema.LanguageCount, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, l.logger).
			With(zap.String("method", "MostPopularLanguage"),
				zap.Duration("took", time.Since(start)),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error getting most popular language", zap.Error(err))
		} else {
			zlog.Info("got most popular language")
		}
	}(time.Now())
	return l.next.MostPopularLanguage(ctx, limit)
}

func (l *loggingMiddleware) MostRecent(ctx context.Context, limit int) (res []schema.Repo, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, l.logger).
			With(zap.String("method", "MostRecent"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error getting most recent repo", zap.Error(err))
		} else {
			zlog.Info("got most recent repos", zap.Int("count", len(res)))
		}
	}(time.Now())
	return l.next.MostRecent(ctx, limit)
}

func (l *loggingMiddleware) MostRecentReposByLanguage(ctx context.Context, language string, limit int) (res []schema.Repo, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, l.logger).
			With(zap.String("method", "MostRecentReposByLanguage"),
				zap.Duration("took", time.Since(start)),
				zap.String("language", language),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error getting most recent repos by language", zap.Error(err))
		} else {
			zlog.Info("got most recent repos by language", zap.Int("count", len(res)))
		}
	}(time.Now())
	return l.next.MostRecentReposByLanguage(ctx, language, limit)
}

func (l *loggingMiddleware) MostStars(ctx context.Context, limit int) (res []schema.Repo, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, l.logger).
			With(zap.String("method", "MostStars"),
				zap.Duration("took", time.Since(start)),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error getting repos with most stars", zap.Error(err))
		} else {
			zlog.Info("got repos with most stars", zap.Int("count", len(res)))
		}
	}(time.Now())
	return l.next.MostStars(ctx, limit)
}

func (l *loggingMiddleware) MostForks(ctx context.Context, limit int) (res []schema.Repo, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, l.logger).
			With(zap.String("method", "MostForks"),
				zap.Duration("took", time.Since(start)),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error getting repos with most forks", zap.Error(err))
		} else {
			zlog.Info("got repos with most forks", zap.Int("count", len(res)))
		}
	}(time.Now())
	return l.next.MostForks(ctx, limit)
}

func (l *loggingMiddleware) RepoCountByUser(ctx context.Context, limit int) (res []schema.UserCount, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, l.logger).
			With(zap.String("method", "RepoCountByUser"),
				zap.Duration("took", time.Since(start)),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error getting repo count by user", zap.Error(err))
		} else {
			zlog.Info("got repo count by user", zap.Int("count", len(res)))
		}
	}(time.Now())
	return l.next.RepoCountByUser(ctx, limit)
}

func (l *loggingMiddleware) ReposByLanguage(ctx context.Context, language string, limit int) (res []schema.UserCount, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, l.logger).
			With(zap.String("method", "ReposByLanguage"),
				zap.Duration("took", time.Since(start)),
				zap.String("language", language),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error getting repos by language", zap.Error(err))
		} else {
			zlog.Info("got repos by language", zap.Int("count", len(res)))
		}
	}(time.Now())
	return l.next.ReposByLanguage(ctx, language, limit)
}

func (l *loggingMiddleware) Distinct(ctx context.Context, field string) (res []string, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, l.logger).
			With(zap.String("method", "Distinct"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error getting distinct login", zap.Error(err))
		} else {
			zlog.Info("got distinct login", zap.Int("count", len(res)))
		}
	}(time.Now())
	return l.next.Distinct(ctx, field)
}

func (l *loggingMiddleware) GetProfile(ctx context.Context, login string) (p *usersvc.User, err error) {
	zlog := logger.Wrap(ctx, l.logger)
	defer func(start time.Time) {
		zlog.Info("got profile", zap.String("method", "GetProfile"),
			zap.Duration("took", time.Since(start)),
			zap.String("login", login),
			zap.Int64("watchers", p.Watchers),
			zap.Int64("stargazers", p.Stargazers),
			zap.Int64("forks", p.Forks),
			zap.Int("keywords", len(p.Keywords)),
			zap.Int("languages", len(p.Languages)))
	}(time.Now())
	return l.next.GetProfile(ctx, login)
}
