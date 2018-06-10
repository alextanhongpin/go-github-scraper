package statsvc

import "github.com/alextanhongpin/go-github-scraper/internal/pkg/database"

// New returns a new stat service
func New(db *database.DB, ms ...Middleware) Service {
	store := NewStore(db, database.Stats)
	model := NewModel(store)
	service := NewService(model)
	service = Decorate(service, ms...)
	return service
}
