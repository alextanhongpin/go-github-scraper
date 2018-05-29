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
func (m *model) BulkUpsert(ctx context.Context, repos []github.Repo) error {
	zlog := logger.RequestIDFromContext(ctx)
	zlog.Info("bulk upsert repos", zap.Int("count", len(repos)))
	return m.store.BulkUpsert(repos)
}

// Count returns the total count of the repos
func (m *model) Count(ctx context.Context) (int, error) {
	return m.store.Count()
}

// Drop drops the collection
func (m *model) Drop(ctx context.Context) error {
	zlog := logger.RequestIDFromContext(ctx)
	zlog.Warn("drop")
	return m.store.Drop()
}

// Perform initialization of the service, such as setting up
// tables for the storage or indexes
func (m *model) Init(ctx context.Context) error {
	zlog := logger.RequestIDFromContext(ctx)
	zlog.Info("init")
	return m.store.Init()
}

// FindLastCreatedByUser returns the last created datetime in the format YYYY-MM-DD for a particular user
func (m *model) FindLastCreatedByUser(ctx context.Context, login string) (string, bool) {
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
func (m *model) LanguageCountByUser(ctx context.Context, login string, limit int) ([]schema.LanguageCount, error) {
	return m.store.AggregateLanguageByUser(login, limit)
}

// MostPopularLanguage returns the most frequent language based on repo count in descending order
func (m *model) MostPopularLanguage(ctx context.Context, limit int) ([]schema.LanguageCount, error) {
	return m.store.AggregateLanguages(limit)
}

// MostRecent returns a limited results of repo that are recently updated
func (m *model) MostRecent(ctx context.Context, limit int) ([]schema.Repo, error) {
	return m.store.FindAll(limit, []string{"-updatedAt"})
}

// MostRecentReposByLanguage returns the most recent repos that are updated for a given language
func (m *model) MostRecentReposByLanguage(ctx context.Context, language string, limit int) ([]schema.Repo, error) {
	return m.store.AggregateMostRecentReposByLanguage(language, limit)
}

// MostStars returns a limited results of repos with the most stars
func (m *model) MostStars(ctx context.Context, limit int) ([]schema.Repo, error) {
	return m.store.FindAll(limit, []string{"-stargazers"})
}

// RepoCountByUser returns the users with most repos sorted in descending order
func (m *model) RepoCountByUser(ctx context.Context, limit int) ([]schema.UserCount, error) {
	return m.store.AggregateReposByUser(limit)
}

// ReposByLanguage returns the users with the most repo in the particular language
func (m *model) ReposByLanguage(ctx context.Context, language string, limit int) ([]schema.UserCount, error) {
	return m.store.AggregateReposByLanguage(language, limit)
}

func (m *model) WatchersFor(ctx context.Context, login string) (int64, error) {
	return m.store.WatchersFor(login)
}

func (m *model) StargazersFor(ctx context.Context, login string) (int64, error) {
	return m.store.StargazersFor(login)
}

func (m *model) ForksFor(ctx context.Context, login string) (int64, error) {
	return m.store.ForksFor(login)
}

func (m *model) KeywordsFor(ctx context.Context, login string, limit int) ([]schema.Keyword, error) {
	return m.store.KeywordsFor(login, limit)
}

func (m *model) DistinctLogin(ctx context.Context) ([]string, error) {
	return m.store.DistinctLogin()
}

func (m *model) GetProfile(ctx context.Context, login string) schema.Profile {
	zlog := logger.RequestIDFromContext(ctx)
	zlog = zlog.With(zap.String("login", login))

	watchers, err := m.store.WatchersFor(login)
	if err != nil {
		zlog.Warn("error getting watcher count", zap.Error(err))
	}
	stargazers, err := m.store.StargazersFor(login)
	if err != nil {
		zlog.Warn("error getting stargazer count", zap.Error(err))
	}
	forks, err := m.store.ForksFor(login)
	if err != nil {
		zlog.Warn("error getting fork count", zap.Error(err))
	}
	keywords, err := m.store.KeywordsFor(login, 20)
	if err != nil {
		zlog.Warn("error getting keyword count", zap.Error(err))
	}

	languages, err := m.store.AggregateLanguageByUser(login, 20)
	if err != nil {
		zlog.Warn("error fetching language count repos", zap.Error(err))
	}

	zlog.Info("updated profile",
		zap.String("login", login),
		zap.Int64("watchers", watchers),
		zap.Int64("stargazers", stargazers),
		zap.Int64("forks", forks),
		zap.Int("keywords", len(keywords)))

	return schema.Profile{
		Login:      login,
		Watchers:   watchers,
		Stargazers: stargazers,
		Forks:      forks,
		Keywords:   keywords,
		Languages:  languages,
	}
}
