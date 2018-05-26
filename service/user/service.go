package user

import "github.com/alextanhongpin/go-github-scraper/internal/database"

// NewService returns a new user service
func NewService(db *database.DB) Model {
	return NewModel(NewStore(db, database.Users))
}
