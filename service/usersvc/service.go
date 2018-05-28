package usersvc

import (
	"github.com/alextanhongpin/go-github-scraper/api/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
)

// Service represents the model of the user
type Service interface {
	BulkUpsert(users []github.User) error
	Count() (int, error)
	Drop() error
	Init() error
	FindOne(login string) (*User, error)
	FindLastCreated() (string, bool)
	FindLastFetched(limit int) ([]User, error)
	MostRecent(limit int) ([]User, error)
	PickLogin() ([]string, error)
	UpdateOne(login string) error
}

// New returns a new user service
func New(db *database.DB) Service {
	return NewModel(NewStore(db, database.Users))
}
