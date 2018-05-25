package main

import (
	"flag"
	"log"

	"github.com/alextanhongpin/go-github-scraper/api/analytic"
	"github.com/alextanhongpin/go-github-scraper/api/repo"
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/internal/util"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := flag.String("port", ":8080", "The tcp port of the application")
	githubToken := flag.String("github_token", "", "The Github's access token used to make calls to the GraphQL endpoint")
	githubURI := flag.String("github_uri", "https://api.github.com/graphql", "The Github's GraphQL endpoint")
	dbName := flag.String("db_name", "scraper", "The name of the database")
	dbHost := flag.String("db_host", "mongodb://myuser:mypass@localhost:27017", "The hostname of the database")
	flag.Parse()

	_ = port
	_ = githubToken
	_ = githubURI

	db := database.New(*dbHost, *dbName)
	defer db.Close()

	// gapi := github.New(*githubToken, *githubURI)
	// users, err := gapi.FetchUsers("malaysia", "2018-01-01", "2018-02-01", "")
	// if err != nil {
	// 	log.Println(err)
	// }
	// repos, err := gapi.FetchReposCursor("alextanhongpin", "2018-01-01", "2018-01-15")
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(len(repos), repos)

	s := repo.New(db, "repos")
	repos, err := s.FindAll(10, []string{"-stargazers"})
	if err != nil {
		panic(err)
	}

	for _, u := range repos {
		// if t, err := time.Parse(u.CreatedAt, time.RFC3339); err != nil {
		// log.Println(err)
		// } else {
		log.Printf("found users: %#v\n", u)
		// }
	}
	r := httprouter.New()

	// Setup services
	analytic.New(db, "analytics", r)

	server := util.NewHTTPServer(*port, r)
	log.Printf("listening to port *%s. press ctrl + c to cancel.\n", *port)
	log.Fatal(server.ListenAndServe())
}
