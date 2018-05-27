package analyticsvc

import (
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/internal/schema"
)

// Service represents the analytic service
type Service interface {
	Init() error

	GetUserCount() (*schema.UserCount, error)
	PostUserCount(count int) error

	GetRepoCount() (*RepoCount, error)
	PostRepoCount(count int) error

	GetReposMostRecent() (*ReposMostRecent, error)
	PostReposMostRecent(data []schema.Repo) error

	GetRepoCountByUser() (*RepoCountByUser, error)
	PostRepoCountByUser(users []schema.UserCount) error

	GetReposMostStars() (*ReposMostStars, error)
	PostReposMostStars(repos []schema.Repo) error

	GetMostPopularLanguage() (*MostPopularLanguage, error)
	PostMostPopularLanguage(languages []schema.LanguageCount) error

	GetLanguageCountByUser() (*LanguageCountByUser, error)
	PostLanguageCountByUser(languages []schema.LanguageCount) error

	GetMostRecentReposByLanguage() (*MostRecentReposByLanguage, error)
	PostMostRecentReposByLanguage(repos []schema.RepoLanguage) error

	GetReposByLanguage() (*ReposByLanguage, error)
	PostReposByLanguage(users []schema.UserCountByLanguage) error
}

// New returns a new analytic service model
func New(db *database.DB) Service {
	return NewModel(NewStore(db, database.Analytics))
}