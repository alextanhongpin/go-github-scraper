package main

import (
	"flag"
	"log"

	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/internal/util"
	"github.com/alextanhongpin/go-github-scraper/service/analytic"
	"github.com/alextanhongpin/go-github-scraper/service/repo"
	"github.com/alextanhongpin/go-github-scraper/service/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := flag.String("port", ":8080", "The tcp port of the application")
	githubToken := flag.String("github_token", "", "The Github's access token used to make calls to the GraphQL endpoint")
	githubURI := flag.String("github_uri", "https://api.github.com/graphql", "The Github's GraphQL endpoint")
	dbName := flag.String("db_name", "scraper", "The name of the database")
	dbHost := flag.String("db_host", "mongodb://myuser:mypass@localhost:27017", "The hostname of the database")
	flag.Parse()

	// _ = port
	_ = githubToken
	_ = githubURI

	db := database.New(*dbHost, *dbName)
	defer db.Close()

	// Setup services
	asvc := analytic.NewService(db)
	rsvc := repo.NewService(db)
	usvc := user.NewService(db)

	if err := rsvc.Init(); err != nil {
		log.Println("error setting repo service", err)
	}
	if err := usvc.Init(); err != nil {
		log.Println("error setting user service", err)
	}

	// if count, err := rsvc.Count(); err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Printf("repos count: %#v", count)
	// }

	// if count, err := usvc.MostRecent(20); err != nil {
	// 	log.Println("error", err)
	// } else {
	// 	log.Printf("user count: %#v", count)
	// }

	// if err := usvc.Drop(); err != nil {
	// 	log.Println("droperror", err)
	// }
	// Setup external api calls
	// gsvc := github.NewAPI(*githubToken, *githubURI)
	// if users, err := gsvc.FetchUsersCursor("Malaysia", "2008-04-01", "2009-01-01", 30); err != nil {
	// 	log.Println("FetchUsersCursorError", err)
	// } else {
	// 	log.Printf("got users %#v\n", users)
	// 	if err := usvc.BulkUpsert(users); err != nil {
	// 		log.Println("BulkUpsert users", err)
	// 	}
	// }
	// repos, err := gsvc.FetchReposCursor("alextanhongpin", "2008-01-01", "2018-06-01", 30)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := rsvc.BulkUpsert(repos); err != nil {
	// 	log.Fatal(err)
	// }

	if count, err := rsvc.MostRecent(5); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("most recent repos: %#v", count)
	}

	// Setup endpoints
	r := httprouter.New()
	analytic.NewEndpoints(asvc, r)

	server := util.NewHTTPServer(*port, r)
	log.Printf("listening to port *%s. press ctrl + c to cancel.\n", *port)
	log.Fatal(server.ListenAndServe())
}
