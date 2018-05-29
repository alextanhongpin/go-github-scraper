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
		ctx := r.Context()
		var res interface{}
		var err error
		switch r.URL.Query().Get("type") {
		case EnumUserCount:
			res, err = e.svc.GetUserCount(ctx)
		case EnumRepoCount:
			res, err = e.svc.GetRepoCount(ctx)
		case EnumReposMostRecent:
			res, err = e.svc.GetReposMostRecent(ctx)
		case EnumRepoCountByUser:
			res, err = e.svc.GetRepoCountByUser(ctx)
		case EnumReposMostStars:
			res, err = e.svc.GetReposMostStars(ctx)
		case EnumMostPopularLanguage:
			res, err = e.svc.GetMostPopularLanguage(ctx)
		case EnumLanguageCountByUser:
			res, err = e.svc.GetLanguageCountByUser(ctx)
		case EnumMostRecentReposByLanguage:
			res, err = e.svc.GetMostRecentReposByLanguage(ctx)
		case EnumReposByLanguage:
			res, err = e.svc.GetReposByLanguage(ctx)
		default:
			http.Error(w, "query type is missing", http.StatusBadRequest)
			return
		}
		util.ResponseJSON(w, res, err)
	}
}
