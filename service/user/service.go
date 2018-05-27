package usersvc

import (
	"github.com/alextanhongpin/go-github-scraper/api/github"
	"github.com/alextanhongpin/go-github-scraper/internal/database"
)

// Service represents the model of the user
type Service interface {
	Init() error
	MostRecent(limit int) ([]User, error)
	BulkUpsert(users []github.User) error
	Drop() error
	FindLastCreated() (string, bool)
	FindLastFetched(limit int) ([]User, error)
	Count() (int, error)
	UpdateOne(login string) error
}

// New returns a new user service
func New(db *database.DB) Service {
	return NewModel(NewStore(db, database.Users))
}
