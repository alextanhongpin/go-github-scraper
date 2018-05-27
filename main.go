package main

import (
	"flag"
	"log"

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

	// users, err := usvc.FindLastFetched(10)
	// if err != nil {
	// 	log.Println("fail to fetch users", err)
	// }
	// log.Printf("%#v\n", users)

	// if err := usvc.Drop(); err != nil {
	// 	log.Fatal("fail to drop users", err)
	// }
	// if err := rsvc.Drop(); err != nil {
	// 	log.Fatal("fail to drop repos", err)
	// }
	if count, err := usvc.Count(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("user count: %#v", count)
	}

	if count, err := rsvc.ReposByLanguage("JavaScript", 20); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("repos count: %#v", count)
	}

	// MostRecent(limit int) ([]Repo, error)
	// MostStars(limit int) ([]Repo, error)
	// Count() (int, error)
	// MostPopularLanguage(limit int) ([]LanguageCount, error)
	// RepoCountByUser(limit int) ([]UserCount, error)
	// LanguageCountByUser(login string, limit int) ([]LanguageCount, error)
	// MostRecentReposByLanguage(language string, limit int) ([]Repo, error)
	// ReposByLanguage(language string, limit int) ([]Repo, error)

	// if repos, err := rsvc.MostRecent(20); err != nil {
	// 	log.Println("error", err)
	// } else {
	// 	log.Printf("repos: %#v", repos)
	// }
	// if count, err := usvc.MostRecent(20); err != nil {
	// 	log.Println("error", err)
	// } else {
	// 	log.Printf("user count: %#v", count)
	// }

	// Setup workers
	w := worker.New(gsvc, asvc, usvc, rsvc, logger)
	// c1 := w.NewFetchUsers("*/20 * * * * *")
	// c1.Start()
	// c2 := w.NewFetchRepos("*/20 * * * * *")
	// c2.Start()
	c3 := w.NewAnalyticBuilder("*/20 * * * * *")
	c3.Start()

	// Setup endpoints
	r := httprouter.New()
	usersvc.MakeEndpoints(usvc, r)
	analyticsvc.MakeEndpoints(asvc, r)
	reposvc.MakeEndpoints(rsvc, r)

	server := util.NewHTTPServer(*port, r)
	log.Printf("listening to port *%s. press ctrl + c to cancel.\n", *port)
	log.Fatal(server.ListenAndServe())
}
