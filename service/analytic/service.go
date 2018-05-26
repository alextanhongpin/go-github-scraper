package analytic

import (
	"github.com/alextanhongpin/go-github-scraper/internal/database"
)

// NewService returns a new analytic service model
func NewService(db *database.DB) Model {
	return NewModel(NewStore(db, database.Analytics))
}
