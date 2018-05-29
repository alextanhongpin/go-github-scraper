package profilesvc

import (
	"context"
	"log"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
	"go.uber.org/zap"
)

type model struct {
	store Store
}

// NewModel returns a new model that fulfils the Service interface
func NewModel(store Store) Service {
	m := model{store: store}
	if err := m.Init(context.Background()); err != nil {
		log.Fatal(err)
	}
	return &m
}

// Perform initialization of the service, such as setting up
// tables for the storage or indexes
func (m *model) Init(ctx context.Context) error {
	zlog := logger.RequestIDFromContext(ctx)
	zlog.Info("init")
	return m.store.Init()
}

func (m *model) GetProfile(ctx context.Context, login string) (*schema.Profile, error) {
	return m.store.GetProfile(login)
}

func (m *model) UpdateProfile(ctx context.Context, login string, profile schema.Profile) error {
	return m.store.UpdateProfile(login, profile)
}

func (m *model) BulkUpsert(ctx context.Context, profiles []schema.Profile) error {
	zlog := logger.RequestIDFromContext(ctx)
	zlog.Info("bulk upsert profile", zap.Int("count", len(profiles)))
	return m.store.BulkUpsert(profiles)
}
