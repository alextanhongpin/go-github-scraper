package reposvc

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/app/usersvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/bow"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/constant"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"

	"go.uber.org/zap"
)

// Model represents the interface for the repo business logic
type model struct {
	store  Store
	logger *logger.Logger
}

// NewModel returns a pointer to the Model
func NewModel(store Store, l *logger.Logger) Service {
	m := model{store: store, logger: l}
	if err := m.Init(context.Background()); err != nil {
		log.Fatal(err)
	}
	return &m
}

// BulkUpsert inserts a list of docs if they do not exists, or updates them if they exist and values differs
func (m *model) BulkUpsert(ctx context.Context, repos []github.Repo) (err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
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
		zlog := logger.Wrap(ctx, m.logger).
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
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "Drop"),
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
		zlog := logger.Wrap(ctx, m.logger).
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
		zlog := logger.Wrap(ctx, m.logger)
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
	// Deduct a day
	t = t.AddDate(0, -1, 0)
	return t.Format("2006-01-02"), true
}

// LanguageCountByUser returns the top languages for a particular user
func (m *model) LanguageCountByUser(ctx context.Context, login string, limit int) (res []schema.LanguageCount, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
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
		zlog := logger.Wrap(ctx, m.logger).
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
		zlog := logger.Wrap(ctx, m.logger).
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
		zlog := logger.Wrap(ctx, m.logger).
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
		zlog := logger.Wrap(ctx, m.logger).
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

// MostForks returns a limited results of repos with the most forks
func (m *model) MostForks(ctx context.Context, limit int) (res []schema.Repo, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "MostForks"),
				zap.Duration("took", time.Since(start)),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error getting repos with most forks", zap.Error(err))
		} else {
			zlog.Info("got repos with most forks", zap.Int("count", len(res)))
		}
	}(time.Now())
	return m.store.FindAll(limit, []string{"-forks"})
}

// RepoCountByUser returns the users with most repos sorted in descending order
func (m *model) RepoCountByUser(ctx context.Context, limit int) (res []schema.UserCount, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
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
		zlog := logger.Wrap(ctx, m.logger).
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

func (m *model) DistinctLogin(ctx context.Context) (res []string, err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
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

func (m *model) GetProfile(ctx context.Context, login string) (p usersvc.User) {
	var watchers, stargazers, forks int64
	var languageList []string
	var descriptions []string

	var keywords []schema.Keyword
	var languages []schema.LanguageCount
	var err error
	zlog := logger.Wrap(ctx, m.logger)
	defer func(start time.Time) {
		zlog.Info("got profile", zap.String("method", "GetProfile"),
			zap.Duration("took", time.Since(start)),
			zap.String("login", login),
			zap.Int64("watchers", watchers),
			zap.Int64("stargazers", stargazers),
			zap.Int64("forks", forks),
			zap.Int("keywords", len(keywords)),
			zap.Int("languages", len(languages)),
			zap.Error(err))
	}(time.Now())

	repos, err := m.store.FindAllFor(login)

	for i := 0; i < len(repos); i++ {
		repo := repos[i]

		stargazers += repo.Stargazers
		watchers += repo.Watchers
		forks += repo.Forks
		languageList = append(languageList, repo.Languages...)
		descriptions = append(descriptions, repo.Description)
	}

	topKeywords := bow.Top(bow.Parse(descriptions...), 20)
	for _, k := range topKeywords {
		keywords = append(keywords, schema.Keyword{ID: k.Key, Value: k.Value})
	}
	sort.SliceStable(keywords, func(i, j int) bool {
		return keywords[i].Value > keywords[j].Value
	})

	topLanguages := bow.Top(languageList, 20)
	for _, k := range topLanguages {
		languages = append(languages, schema.LanguageCount{Name: k.Key, Count: k.Value})
	}
	sort.SliceStable(languages, func(i, j int) bool {
		return languages[i].Count > languages[j].Count
	})

	return usersvc.User{
		Login: login,
		Profile: schema.Profile{
			Watchers:   watchers,
			Stargazers: stargazers,
			Forks:      forks,
			Keywords:   keywords,
			Languages:  languages,
		},
	}
}
