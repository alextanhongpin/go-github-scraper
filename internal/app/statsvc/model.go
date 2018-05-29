package statsvc

import (
	"context"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
)

// Model represents the interface for the analytic business logic
type (
	model struct {
		store Store
	}
)

// NewModel returns a new analytic model
func NewModel(s Store) Service {
	return &model{
		store: s,
	}
}

func (m *model) Init(ctx context.Context) error {
	zlog := logger.RequestIDFromContext(ctx)
	zlog.Info("init")
	return m.store.Init()
}

func (m *model) GetUserCount(ctx context.Context) (*schema.UserCount, error) {
	// Validate...
	// Tracing...
	return m.store.GetUserCount()
}

func (m *model) PostUserCount(ctx context.Context, count int) error {
	return m.store.PostUserCount(count)
}

func (m *model) GetRepoCount(ctx context.Context) (*RepoCount, error) {
	return m.store.GetRepoCount()
}
func (m *model) PostRepoCount(ctx context.Context, count int) error {
	return m.store.PostRepoCount(count)
}

func (m *model) GetReposMostRecent(ctx context.Context) (*ReposMostRecent, error) {
	return m.store.GetReposMostRecent()
}

func (m *model) PostReposMostRecent(ctx context.Context, data []schema.Repo) error {
	return m.store.PostReposMostRecent(data)
}

func (m *model) GetRepoCountByUser(ctx context.Context) (*RepoCountByUser, error) {
	return m.store.GetRepoCountByUser()
}

func (m *model) PostRepoCountByUser(ctx context.Context, users []schema.UserCount) error {
	return m.store.PostRepoCountByUser(users)
}

func (m *model) GetReposMostStars(ctx context.Context) (*ReposMostStars, error) {
	return m.store.GetReposMostStars()
}

func (m *model) PostReposMostStars(ctx context.Context, repos []schema.Repo) error {
	return m.store.PostReposMostStars(repos)
}

func (m *model) GetMostPopularLanguage(ctx context.Context) (*MostPopularLanguage, error) {
	return m.store.GetMostPopularLanguage()
}

func (m *model) PostMostPopularLanguage(ctx context.Context, languages []schema.LanguageCount) error {
	return m.store.PostMostPopularLanguage(languages)
}

func (m *model) GetLanguageCountByUser(ctx context.Context) (*LanguageCountByUser, error) {
	return m.store.GetLanguageCountByUser()
}

func (m *model) PostLanguageCountByUser(ctx context.Context, languages []schema.LanguageCount) error {
	return m.store.PostLanguageCountByUser(languages)
}

func (m *model) GetMostRecentReposByLanguage(ctx context.Context) (*MostRecentReposByLanguage, error) {
	return m.store.GetMostRecentReposByLanguage()
}

func (m *model) PostMostRecentReposByLanguage(ctx context.Context, repos []schema.RepoLanguage) error {
	return m.store.PostMostRecentReposByLanguage(repos)
}

func (m *model) GetReposByLanguage(ctx context.Context) (*ReposByLanguage, error) {
	return m.store.GetReposByLanguage()
}

func (m *model) PostReposByLanguage(ctx context.Context, users []schema.UserCountByLanguage) error {
	return m.store.PostReposByLanguage(users)
}
