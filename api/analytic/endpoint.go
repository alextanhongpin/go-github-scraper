package analytic

import (
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/util"

	"github.com/julienschmidt/httprouter"
)

// Endpoints represents the services exposed as http routes
type Endpoints interface {
	GetUserCount() httprouter.Handle
}

type endpoints struct {
	api API
}

// NewEndpoints returns a new route for the analytic service
func NewEndpoints(api API) Endpoints {
	return &endpoints{
		api: api,
	}
}

func (e *endpoints) GetUserCount() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		count, err := e.api.GetUserCount()
		util.ResponseJSON(w, count, err)
	}
}
