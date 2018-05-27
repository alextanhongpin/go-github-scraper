package profilesvc

import (
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/internal/schema"
)

// Service represents the interface for the profile service
type Service interface {
	Init() error
	GetProfile(login string) (*schema.Profile, error)
	UpdateProfile(login string, profile schema.Profile) error
	BulkUpsert(profiles []schema.Profile) error
}

// New returns a new profile service
func New(db *database.DB) Service {
	return NewModel(NewStore(db, database.Profiles))
}
