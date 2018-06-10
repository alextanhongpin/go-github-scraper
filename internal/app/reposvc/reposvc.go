package reposvc

import (
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
)

// At the heart of the reposvc is a builder pattern that is responsible
// building each stages of the service - store, model and service,
// which can be further decorated

// New returns a new service with store
func New(db *database.DB, middlewares ...Middleware) Service {
	store := NewStore(db, database.Repos)
	model := NewModel(store)
	service := NewService(model)

	service = Decorate(service, middlewares...)
	return service
}
