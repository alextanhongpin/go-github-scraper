package usersvc

import (
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/util"

	"github.com/julienschmidt/httprouter"
)

// Endpoints represents the services exposed as http routes
type (
	Endpoints interface {
		GetUser() httprouter.Handle
		GetUserCount() httprouter.Handle
	}

	endpoints struct {
		svc Service
	}
)

// NewEndpoints creates a new endpoint
func NewEndpoints(svc Service) Endpoints {
	return &endpoints{svc: svc}
}

// MakeEndpoints creates a new endpoint based on the service provided and router
func MakeEndpoints(svc Service, r *httprouter.Router) {
	e := NewEndpoints(svc)

	r.GET("/users/:login", e.GetUser())
	r.GET("/users", e.GetUserCount())
}

func (e *endpoints) GetUserCount() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		count, err := e.svc.Count()
		util.ResponseJSON(w, GetUserCountResponse{count}, err)
	}
}

func (e *endpoints) GetUser() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		login := ps.ByName("login")
		user, err := e.svc.FindOne(login)
		util.ResponseJSON(w, GetUserResponse{user}, err)
	}
}

// Stars: 84Watchers: 84Forks: 18
// Keywords: simple × 43sample × 38example × 19go × 18api × 16nodejs × 15grpc × 15node × 14using × 14golang × 13
// Languages: JavaScript 41%Go 22%Jupyter Notebook 7%Python 6%Rust 4%HTML 3%C# 2%Makefile 2%Scala 2%HCL 2%
