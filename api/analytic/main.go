package analytic

import (
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/julienschmidt/httprouter"
)

// NewService returns a new analytic service model
func NewService(db *database.DB, collection string, r *httprouter.Router) {
	store := NewStore(db, collection)
	model := NewModel(store)
	endpoints := NewEndpoints(model)

	r.GET("/analytics", endpoints.GetUserCount())
}
