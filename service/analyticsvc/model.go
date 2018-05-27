package analyticsvc

import (
	"github.com/alextanhongpin/go-github-scraper/internal/schema"
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

func (m *model) Init() error {
	return m.store.Init()
}

func (m *model) GetUserCount() (*schema.UserCount, error) {
	// Validate...
	// Tracing...
	return m.store.GetUserCount()
}

func (m *model) PostUserCount(count int) error {
	return m.store.PostUserCount(count)
}

func (m *model) GetRepoCount() (*RepoCount, error) {
	return m.store.GetRepoCount()
}
func (m *model) PostRepoCount(count int) error {
	return m.store.PostRepoCount(count)
}

func (m *model) GetReposMostRecent() (*ReposMostRecent, error) {
	return m.store.GetReposMostRecent()
}

func (m *model) PostReposMostRecent(data []schema.Repo) error {
	return m.store.PostReposMostRecent(data)
}

func (m *model) GetRepoCountByUser() (*RepoCountByUser, error) {
	return m.store.GetRepoCountByUser()
}

func (m *model) PostRepoCountByUser(users []schema.UserCount) error {
	return m.store.PostRepoCountByUser(users)
}

func (m *model) GetReposMostStars() (*ReposMostStars, error) {
	return m.store.GetReposMostStars()
}

func (m *model) PostReposMostStars(repos []schema.Repo) error {
	return m.store.PostReposMostStars(repos)
}

func (m *model) GetMostPopularLanguage() (*MostPopularLanguage, error) {
	return m.store.GetMostPopularLanguage()
}

func (m *model) PostMostPopularLanguage(languages []schema.LanguageCount) error {
	return m.store.PostMostPopularLanguage(languages)
}

func (m *model) GetLanguageCountByUser() (*LanguageCountByUser, error) {
	return m.store.GetLanguageCountByUser()
}

func (m *model) PostLanguageCountByUser(languages []schema.LanguageCount) error {
	return m.store.PostLanguageCountByUser(languages)
}

func (m *model) GetMostRecentReposByLanguage() (*MostRecentReposByLanguage, error) {
	return m.store.GetMostRecentReposByLanguage()
}

func (m *model) PostMostRecentReposByLanguage(repos []schema.RepoLanguage) error {
	return m.store.PostMostRecentReposByLanguage(repos)
}

func (m *model) GetReposByLanguage() (*ReposByLanguage, error) {
	return m.store.GetReposByLanguage()
}

func (m *model) PostReposByLanguage(users []schema.UserCountByLanguage) error {
	return m.store.PostReposByLanguage(users)
}
