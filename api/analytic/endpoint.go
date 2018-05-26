package analytic

import (
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/util"

	"github.com/julienschmidt/httprouter"
)

// Endpoints represents the services exposed as http routes
type (
	Endpoints interface {
		GetUserCount() httprouter.Handle
	}

	endpoints struct {
		model Model
	}
)

// NewEndpoints returns a new route for the analytic service
func NewEndpoints(model Model) Endpoints {
	return &endpoints{
		model: model,
	}
}

func (e *endpoints) GetUserCount() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		count, err := e.model.GetUserCount()
		util.ResponseJSON(w, count, err)
	}
}
