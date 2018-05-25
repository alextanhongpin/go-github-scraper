package analytic

import (
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/julienschmidt/httprouter"
)

// New returns a new analytic service model
func New(db *database.DB, collection string, r *httprouter.Router) {
	store := NewStore(db, collection)
	api := NewAPI(store)
	endpoints := NewEndpoints(api)

	r.GET("/analytics", endpoints.GetUserCount())
}
