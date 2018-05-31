package statsvc

import (
	"context"
	"log"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
	"go.uber.org/zap"
)

// Model represents the interface for the analytic business logic
type (
	model struct {
		store  Store
		logger *logger.Logger
	}
)

// NewModel returns a new analytic model
func NewModel(s Store, l *logger.Logger) Service {
	m := model{
		store:  s,
		logger: l,
	}
	if err := m.Init(context.Background()); err != nil {
		log.Fatal(err)
	}
	return &m
}

func (m *model) Init(ctx context.Context) (err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "Init"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error initializing statsvc", zap.Error(err))
		} else {
			zlog.Info("initialize statsvc")
		}
	}(time.Now())
	return m.store.Init()
}

func (m *model) GetUserCount(ctx context.Context) (res *UserCount, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "GetUserCount"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error getting user count", zap.Error(err))
		} else {
			zlog.Info("get user count",
				zap.Int("count", res.Count))
		}
	}(time.Now())
	return m.store.GetUserCount()
}

func (m *model) PostUserCount(ctx context.Context, count int) (err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "PostUserCount"),
				zap.Duration("took", time.Since(start)),
				zap.Int("count", count))
		if err != nil {
			zlog.Error("error posting user count", zap.Error(err))
		} else {
			zlog.Info("post user count")
		}
	}(time.Now())
	return m.store.PostUserCount(count)
}

func (m *model) GetRepoCount(ctx context.Context) (res *RepoCount, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "GetRepoCount"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error getting repo count", zap.Error(err))
		} else {
			zlog.Info("get user count", zap.Int("count", res.Count))
		}
	}(time.Now())
	return m.store.GetRepoCount()
}

func (m *model) PostRepoCount(ctx context.Context, count int) (err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "PostRepoCount"),
				zap.Duration("took", time.Since(start)),
				zap.Int("count", count))
		if err != nil {
			zlog.Error("error posting repo count", zap.Error(err))
		} else {
			zlog.Info("post user count")
		}
	}(time.Now())
	return m.store.PostRepoCount(count)
}

func (m *model) GetReposMostRecent(ctx context.Context) (res *ReposMostRecent, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "GetReposMostRecent"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error getting repos most count", zap.Error(err))
		} else {
			zlog.Info("get repos most recent")
		}
	}(time.Now())
	return m.store.GetReposMostRecent()
}

func (m *model) PostReposMostRecent(ctx context.Context, repos []schema.Repo) (err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "PostReposMostRecent"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error posting repos most recent", zap.Error(err))
		} else {
			zlog.Info("post repos most recent", zap.Int("count", len(repos)))
		}
	}(time.Now())
	return m.store.PostReposMostRecent(repos)
}

func (m *model) GetRepoCountByUser(ctx context.Context) (res *RepoCountByUser, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "GetRepoCountByUser"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error getting repo count by user", zap.Error(err))
		} else {
			zlog.Info("get repos count by user")
		}
	}(time.Now())
	return m.store.GetRepoCountByUser()
}

func (m *model) PostRepoCountByUser(ctx context.Context, users []schema.UserCount) (err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "PostRepoCountByUser"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error posting repo count by user", zap.Error(err))
		} else {
			zlog.Info("post repo count by user", zap.Int("count", len(users)))
		}
	}(time.Now())
	return m.store.PostRepoCountByUser(users)
}

func (m *model) GetReposMostStars(ctx context.Context) (res *ReposMostStars, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "GetReposMostStars"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error getting repos most stars", zap.Error(err))
		} else {
			zlog.Info("get repos most stars")
		}
	}(time.Now())
	return m.store.GetReposMostStars()
}

func (m *model) PostReposMostStars(ctx context.Context, repos []schema.Repo) (err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "PostReposMostStars"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error posting repos most stars", zap.Error(err))
		} else {
			zlog.Info("post repos most stars")
		}
	}(time.Now())
	return m.store.PostReposMostStars(repos)
}

func (m *model) GetMostPopularLanguage(ctx context.Context) (res *MostPopularLanguage, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "GetMostPopularLanguage"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error getting most popular language", zap.Error(err))
		} else {
			zlog.Info("get most poplular language")
		}
	}(time.Now())
	return m.store.GetMostPopularLanguage()
}

func (m *model) PostMostPopularLanguage(ctx context.Context, languages []schema.LanguageCount) (err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "PostMostPopularLanguage"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error posting most popular language", zap.Error(err))
		} else {
			zlog.Info("post most popular language")
		}
	}(time.Now())
	return m.store.PostMostPopularLanguage(languages)
}

func (m *model) GetLanguageCountByUser(ctx context.Context) (res *LanguageCountByUser, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "GetLanguageCountByUser"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error getting language count by users", zap.Error(err))
		} else {
			zlog.Info("get language count by user")
		}
	}(time.Now())
	return m.store.GetLanguageCountByUser()
}

func (m *model) PostLanguageCountByUser(ctx context.Context, languages []schema.LanguageCount) (err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "PostLanguageCountByUser"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error posting language count by user", zap.Error(err))
		} else {
			zlog.Info("post language count by user")
		}
	}(time.Now())
	return m.store.PostLanguageCountByUser(languages)
}

func (m *model) GetMostRecentReposByLanguage(ctx context.Context) (res *MostRecentReposByLanguage, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "GetMostRecentReposByLanguage"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error getting most recent repos by language", zap.Error(err))
		} else {
			zlog.Info("get most recent repos by language")
		}
	}(time.Now())
	return m.store.GetMostRecentReposByLanguage()
}

func (m *model) PostMostRecentReposByLanguage(ctx context.Context, repos []schema.RepoLanguage) (err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "PostMostRecentReposByLanguage"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error posting most recent repos by language", zap.Error(err))
		} else {
			zlog.Info("post most recent repos by language")
		}
	}(time.Now())
	return m.store.PostMostRecentReposByLanguage(repos)
}

func (m *model) GetReposByLanguage(ctx context.Context) (res *ReposByLanguage, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "GetReposByLanguage"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error getting repos by language", zap.Error(err))
		} else {
			zlog.Info("get repos by language")
		}
	}(time.Now())
	return m.store.GetReposByLanguage()
}

func (m *model) PostReposByLanguage(ctx context.Context, users []schema.UserCountByLanguage) (err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "PostReposByLanguage"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Error("error posting repos by language", zap.Error(err))
		} else {
			zlog.Info("post repos by language")
		}
	}(time.Now())
	return m.store.PostReposByLanguage(users)
}
