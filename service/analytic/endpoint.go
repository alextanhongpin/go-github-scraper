package analyticsvc

import (
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/util"

	"github.com/julienschmidt/httprouter"
)

// Endpoints represents the services exposed as http routes
type (
	Endpoints interface {
		GetAnalytics() httprouter.Handle
	}

	endpoints struct {
		svc Service
	}
)

// MakeEndpoints creates a new set of endpoints based on the service provided and router
func MakeEndpoints(svc Service, r *httprouter.Router) {
	e := &endpoints{svc: svc}

	r.GET("/analytics", e.GetAnalytics())
}

func (e *endpoints) GetAnalytics() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		switch r.URL.Query().Get("type") {
		case EnumUserCount:
			count, err := e.svc.GetUserCount()
			util.ResponseJSON(w, count, err)
		case EnumRepoCount:
			count, err := e.svc.GetRepoCount()
			util.ResponseJSON(w, count, err)
		case EnumReposMostRecent:
			count, err := e.svc.GetReposMostRecent()
			util.ResponseJSON(w, count, err)
		case EnumRepoCountByUser:
			count, err := e.svc.GetRepoCountByUser()
			util.ResponseJSON(w, count, err)
		case EnumReposMostStars:
		case EnumMostPopularLanguage:
		case EnumLanguageCountByUser:
		case EnumMostRecentReposByLanguage:
		case EnumReposByLanguage:
		default:
			http.Error(w, "query type is missing", http.StatusBadRequest)
			return
		}

	}
}
