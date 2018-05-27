package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/alextanhongpin/go-github-scraper/api/github"
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/internal/util"
	"github.com/alextanhongpin/go-github-scraper/service/analytic"
	"github.com/alextanhongpin/go-github-scraper/service/repo"
	"github.com/alextanhongpin/go-github-scraper/service/user"
	"github.com/alextanhongpin/go-github-scraper/worker"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func main() {
	// Setup environment variables
	port := flag.String("port", ":8080", "The tcp port of the application")
	githubToken := flag.String("github_token", "", "The Github's access token used to make calls to the GraphQL endpoint")
	githubURI := flag.String("github_uri", "https://api.github.com/graphql", "The Github's GraphQL endpoint")
	dbName := flag.String("db_name", "scraper", "The name of the database")
	dbHost := flag.String("db_host", "mongodb://myuser:mypass@localhost:27017", "The hostname of the database")
	flag.Parse()

	// Setup Logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	// Setup Database
	db := database.New(*dbHost, *dbName)
	defer db.Close()

	// Setup services
	asvc := analyticsvc.New(db)
	rsvc := reposvc.New(db)
	usvc := usersvc.New(db)
	gsvc := github.New(*githubToken, *githubURI)

	// Setup workers
	w := worker.New(gsvc, asvc, usvc, rsvc, logger)
	_ = w
	// c1 := w.NewFetchUsers("*/20 * * * * *")
	// c1.Start()
	// c2 := w.NewFetchRepos("*/20 * * * * *")
	// c2.Start()
	// c3 := w.NewAnalyticBuilder("*/20 * * * * *")
	// c3.Start()

	// Setup router
	r := httprouter.New()

	// Setup profiling
	r.HandlerFunc(http.MethodGet, "/debug/pprof/", pprof.Index)
	r.HandlerFunc(http.MethodGet, "/debug/pprof/cmdline", pprof.Cmdline)
	r.HandlerFunc(http.MethodGet, "/debug/pprof/profile", pprof.Profile)
	r.HandlerFunc(http.MethodGet, "/debug/pprof/symbol", pprof.Symbol)
	r.HandlerFunc(http.MethodGet, "/debug/pprof/trace", pprof.Trace)
	r.Handler(http.MethodGet, "/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handler(http.MethodGet, "/debug/pprof/heap", pprof.Handler("heap"))
	r.Handler(http.MethodGet, "/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Handler(http.MethodGet, "/debug/pprof/block", pprof.Handler("block"))

	// Setup endpoints
	usersvc.MakeEndpoints(usvc, r)
	analyticsvc.MakeEndpoints(asvc, r)
	reposvc.MakeEndpoints(rsvc, r)

	// Setup server
	server := util.NewHTTPServer(*port, r)
	log.Printf("listening to port *%s. press ctrl + c to cancel.\n", *port)
	log.Fatal(server.ListenAndServe())
}
