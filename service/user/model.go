package usersvc

import (
	"log"
	"time"

	"github.com/alextanhongpin/go-github-scraper/api/github"
)

type model struct {
	store Store
}

// NewModel returns a new model with the store
func NewModel(store Store) Service {
	m := model{store: store}
	if err := m.Init(); err != nil {
		log.Fatal(err)
	}
	return &m
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

func (m *model) FindLastCreated() (string, bool) {
	user, err := m.store.FindLastCreated()
	if err != nil || user == nil {
		// Github's creation date
		return "2008-04-01", false
	}
	t, err := time.Parse(time.RFC3339, user.CreatedAt)
	if err != nil {
		return "2008-04-01", false
	}
	return t.Format("2006-01-02"), true
}

func (m *model) FindLastFetched(limit int) ([]User, error) {
	return m.store.FindAll(limit, []string{"-fetchedAt"})
}

func (m *model) Count() (int, error) {
	return m.store.Count()
}

func (m *model) UpdateOne(login string) error {
	return m.store.UpdateOne(login)
}
