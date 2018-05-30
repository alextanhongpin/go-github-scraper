package reposvc

import (
	"context"
	"log"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/constant"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
	"go.uber.org/zap"
)

// Model represents the interface for the repo business logic
type model struct {
	store Store
}

// NewModel returns a pointer to the Model
func NewModel(store Store) Service {
	m := model{store: store}
	if err := m.Init(context.TODO()); err != nil {
		log.Fatal(err)
	}
	return &m
}

// BulkUpsert inserts a list of docs if they do not exists, or updates them if they exist and values differs
func (m *model) BulkUpsert(ctx context.Context, repos []github.Repo) (err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "BulkUpsert"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error upserting repos", zap.Error(err))
		} else {
			zlog.Info("inserted repos", zap.Int("count", len(repos)))
		}
	}(time.Now())
	return m.store.BulkUpsert(repos)
}

// Count returns the total count of the repos
func (m *model) Count(ctx context.Context) (c int, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "BulkUpsert"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error returning repo count", zap.Error(err))
		} else {
			zlog.Info("got repo count", zap.Int("count", c))
		}
	}(time.Now())
	return m.store.Count()
}

// Drop drops the collection
func (m *model) Drop(ctx context.Context) (err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx)
		zlog.With(zap.String("method", "Drop"),
			zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error dropping repo collection", zap.Error(err))
		} else {
			zlog.Info("drop repo collection")
		}
	}(time.Now())
	return m.store.Drop()
}

// Perform initialization of the service, such as setting up
// tables for the storage or indexes
func (m *model) Init(ctx context.Context) (err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "Init"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error init reposvc", zap.Error(err))
		} else {
			zlog.Info("init reposvc")
		}
	}(time.Now())
	return m.store.Init()
}

// FindLastCreatedByUser returns the last created datetime in the format YYYY-MM-DD for a particular user
func (m *model) FindLastCreatedByUser(ctx context.Context, login string) (date string, ok bool) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx)
		zlog.Info("find last created by user",
			zap.String("method", "FindLastCreatedByUser"),
			zap.Duration("took", time.Since(start)),
			zap.String("date", date),
			zap.Bool("default", ok))
	}(time.Now())

	repo, err := m.store.FindLastCreatedByUser(login)
	if err != nil || repo == nil {
		// Github's creation date
		return constant.GithubCreatedAt, false
	}
	t, err := time.Parse(time.RFC3339, repo.CreatedAt)
	if err != nil {
		return constant.GithubCreatedAt, false
	}
	return t.Format("2006-01-02"), true
}

// LanguageCountByUser returns the top languages for a particular user
func (m *model) LanguageCountByUser(ctx context.Context, login string, limit int) (res []schema.LanguageCount, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "LanguageCountByUser"),
				zap.Duration("took", time.Since(start)),
				zap.String("for", login),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error getting language count by user", zap.Error(err))
		} else {
			zlog.Info("got language count")
		}
	}(time.Now())
	return m.store.AggregateLanguageByUser(login, limit)
}

// MostPopularLanguage returns the most frequent language based on repo count in descending order
func (m *model) MostPopularLanguage(ctx context.Context, limit int) (res []schema.LanguageCount, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "MostPopularLanguage"),
				zap.Duration("took", time.Since(start)),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error getting most popular language", zap.Error(err))
		} else {
			zlog.Info("got most popular language")
		}
	}(time.Now())
	return m.store.AggregateLanguages(limit)
}

// MostRecent returns a limited results of repo that are recently updated
func (m *model) MostRecent(ctx context.Context, limit int) (res []schema.Repo, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "MostRecent"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error getting most recent repo", zap.Error(err))
		} else {
			zlog.Info("got most recent repos", zap.Int("count", len(res)))
		}
	}(time.Now())
	return m.store.FindAll(limit, []string{"-updatedAt"})
}

// MostRecentReposByLanguage returns the most recent repos that are updated for a given language
func (m *model) MostRecentReposByLanguage(ctx context.Context, language string, limit int) (res []schema.Repo, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
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
	return m.store.AggregateMostRecentReposByLanguage(language, limit)
}

// MostStars returns a limited results of repos with the most stars
func (m *model) MostStars(ctx context.Context, limit int) (res []schema.Repo, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "MostStars"),
				zap.Duration("took", time.Since(start)),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error getting repos with most stars", zap.Error(err))
		} else {
			zlog.Info("got repos with most stars", zap.Int("count", len(res)))
		}
	}(time.Now())
	return m.store.FindAll(limit, []string{"-stargazers"})
}

// RepoCountByUser returns the users with most repos sorted in descending order
func (m *model) RepoCountByUser(ctx context.Context, limit int) (res []schema.UserCount, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "RepoCountByUser"),
				zap.Duration("took", time.Since(start)),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error getting repo count by user", zap.Error(err))
		} else {
			zlog.Info("got repo count by user", zap.Int("count", len(res)))
		}
	}(time.Now())
	return m.store.AggregateReposByUser(limit)
}

// ReposByLanguage returns the users with the most repo in the particular language
func (m *model) ReposByLanguage(ctx context.Context, language string, limit int) (res []schema.UserCount, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
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
	return m.store.AggregateReposByLanguage(language, limit)
}

func (m *model) WatchersFor(ctx context.Context, login string) (count int64, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "WatchersFor"),
				zap.Duration("took", time.Since(start)),
				zap.String("for", login))
		if err != nil {
			zlog.Warn("error getting watchers for", zap.Error(err))
		} else {
			zlog.Info("got watchers for", zap.Int64("count", count))
		}
	}(time.Now())
	return m.store.WatchersFor(login)
}

func (m *model) StargazersFor(ctx context.Context, login string) (count int64, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "StargazersFor"),
				zap.Duration("took", time.Since(start)),
				zap.String("login", login))
		if err != nil {
			zlog.Warn("error getting stargazers for", zap.Error(err))
		} else {
			zlog.Info("got stargazers for", zap.Int64("count", count))
		}
	}(time.Now())
	return m.store.StargazersFor(login)
}

func (m *model) ForksFor(ctx context.Context, login string) (count int64, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "ForksFor"),
				zap.Duration("took", time.Since(start)),
				zap.String("login", login))
		if err != nil {
			zlog.Warn("error getting forks for", zap.Error(err))
		} else {
			zlog.Info("got forks for", zap.Int64("count", count))
		}
	}(time.Now())
	return m.store.ForksFor(login)
}

func (m *model) KeywordsFor(ctx context.Context, login string, limit int) (res []schema.Keyword, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "KeywordsFor"),
				zap.Duration("took", time.Since(start)),
				zap.String("login", login),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error getting keywords for", zap.Error(err))
		} else {
			zlog.Info("got keywords for", zap.Int("count", len(res)))
		}
	}(time.Now())
	return m.store.KeywordsFor(login, limit)
}

func (m *model) DistinctLogin(ctx context.Context) (res []string, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "DistinctLogin"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error getting distinct login", zap.Error(err))
		} else {
			zlog.Info("got distinct login", zap.Int("count", len(res)))
		}
	}(time.Now())
	return m.store.DistinctLogin()
}

func (m *model) GetProfile(ctx context.Context, login string) (p schema.Profile) {
	var watchers, stargazers, forks int64
	var keywords []schema.Keyword
	var languages []schema.LanguageCount
	var err error
	zlog := logger.RequestIDFromContext(ctx)
	defer func(start time.Time) {
		zlog.Info("got profile", zap.String("method", "GetProfile"),
			zap.Duration("took", time.Since(start)),
			zap.String("login", login),
			zap.Int64("watchers", watchers),
			zap.Int64("stargazers", stargazers),
			zap.Int64("forks", forks),
			zap.Int("keywords", len(keywords)),
			zap.Error(err))
	}(time.Now())

	watchers, err = m.store.WatchersFor(login)
	if err != nil {
		zlog.Warn("error getting watcher count", zap.Error(err))
	}
	stargazers, err = m.store.StargazersFor(login)
	if err != nil {
		zlog.Warn("error getting stargazer count", zap.Error(err))
	}
	forks, err = m.store.ForksFor(login)
	if err != nil {
		zlog.Warn("error getting fork count", zap.Error(err))
	}
	keywords, err = m.store.KeywordsFor(login, 20)
	if err != nil {
		zlog.Warn("error getting keyword count", zap.Error(err))
	}

	languages, err = m.store.AggregateLanguageByUser(login, 20)
	if err != nil {
		zlog.Warn("error fetching language count repos", zap.Error(err))
	}

	return schema.Profile{
		Login:      login,
		Watchers:   watchers,
		Stargazers: stargazers,
		Forks:      forks,
		Keywords:   keywords,
		Languages:  languages,
	}
}
