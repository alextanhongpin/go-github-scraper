package profilesvc

import (
	"context"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
)

// Service represents the interface for the profile service
type Service interface {
	Init(ctx context.Context) error
	GetProfile(ctx context.Context, login string) (*schema.Profile, error)
	UpdateProfile(ctx context.Context, login string, profile schema.Profile) error
	BulkUpsert(ctx context.Context, profiles []schema.Profile) error
}

// New returns a new profile service
func New(db *database.DB, l *logger.Logger) Service {
	return NewModel(NewStore(db, database.Profiles), l)
}
