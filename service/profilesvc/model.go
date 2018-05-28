package profilesvc

import (
	"errors"
	"log"

	"github.com/alextanhongpin/go-github-scraper/internal/schema"
)

type model struct {
	store Store
}

// NewModel returns a new model that fulfils the Service interface
func NewModel(store Store) Service {
	m := model{store: store}
	if err := m.Init(); err != nil {
		log.Fatal(err)
	}
	return &m
}

// Perform initialization of the service, such as setting up
// tables for the storage or indexes
func (m *model) Init() error {
	return m.store.Init()
}

func (m *model) GetProfile(login string) (*schema.Profile, error) {
	return m.store.GetProfile(login)
}

func (m *model) UpdateProfile(login string, profile schema.Profile) error {
	return m.store.UpdateProfile(login, profile)
}

func (m *model) BulkUpsert(profiles []schema.Profile) error {
	if len(profiles) > 1000 {
		return errors.New("more than 1000 items")
	}
	return m.store.BulkUpsert(profiles)
}
