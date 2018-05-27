package analyticsvc

import "github.com/alextanhongpin/go-github-scraper/internal/schema"

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

func (m *model) GetUserCount() (*UserCount, error) {
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
