package reposvc

import (
	"github.com/alextanhongpin/go-github-scraper/api/github"
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/internal/schema"
)

// Service represents the interface for the repo service
type Service interface {
	BulkUpsert(repos []github.Repo) error
	Count() (int, error)
	Drop() error
	Init() error
	FindLastCreatedByUser(login string) (string, bool)
	LanguageCountByUser(login string, limit int) ([]schema.LanguageCount, error)
	MostPopularLanguage(limit int) ([]schema.LanguageCount, error)
	MostRecent(limit int) ([]schema.Repo, error)
	MostRecentReposByLanguage(language string, limit int) ([]schema.Repo, error)
	MostStars(limit int) ([]schema.Repo, error)
	RepoCountByUser(limit int) ([]schema.UserCount, error)
	ReposByLanguage(language string, limit int) ([]schema.UserCount, error)
	WatchersFor(login string) (int64, error)
	StargazersFor(login string) (int64, error)
	ForksFor(login string) (int64, error)
	WordCount(login string, limit int) ([]WordCount, error)
}

// New returns a new service with store
func New(db *database.DB) Service {
	return NewModel(NewStore(db, database.Repos))
}
