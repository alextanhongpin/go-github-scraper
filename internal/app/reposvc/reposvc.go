package reposvc

import (
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
)

// New returns a new service with store
func New(db *database.DB, middlewares ...Middleware) Service {
	store := NewStore(db, database.Repos)
	model := NewModel(store)
	service := NewService(model)

	// Apply decorator
	service = Decorate(service, middlewares...)
	return service
}
