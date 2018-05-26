package user

import "github.com/alextanhongpin/go-github-scraper/api/github"

type (
	// Model represents the model of the user
	Model interface {
		Init() error
		MostRecent(limit int) ([]User, error)
		BulkUpsert(users []github.User) error
		Drop() error
	}

	model struct {
		store Store
	}
)

// NewModel returns a new model with the store
func NewModel(store Store) Model {
	return &model{store: store}
}

func (m *model) Init() error {
	return m.store.Init()
}

func (m *model) MostRecent(limit int) ([]User, error) {
	return m.store.FindAll(limit, []string{"-updatedAt"})
}

func (m *model) BulkUpsert(users []github.User) error {
	return m.store.BulkUpsert(users)
}

func (m *model) Drop() error {
	return m.store.Drop()
}
