package usersvc

import (
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
)

// New returns a new user service
func New(db *database.DB, m ...Middleware) Service {
	store := NewStore(db, database.Users)
	model := NewModel(store)
	service := NewService(model)
	service = Decorate(service, m...)
	return service
}
