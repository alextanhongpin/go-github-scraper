package repo

import "github.com/alextanhongpin/go-github-scraper/internal/database"

// NewService returns a new service with store
func NewService(db *database.DB) Model {
	return NewModel(NewStore(db, database.Repos))
}
