package main

import (
	"flag"

	"github.com/alextanhongpin/go-github-scraper/internal/database"
)

func main() {
	githubToken := flag.String("github_token", "", "The Github's access token used to make calls to the GraphQL endpoint")
	githubURI := flag.String("github_uri", "https://api.github.com/graphql", "The Github's GraphQL endpoint")
	dbName := flag.String("db_name", "scraper", "The name of the database")
	dbHost := flag.String("db_host", "mongodb://myuser:mypass@localhost:27017", "The hostname of the database")
	flag.Parse()

	db := database.New(*dbHost, *dbName)
	defer db.Close()

	_ = githubToken
	_ = githubURI
	// gapi := github.New(*githubToken, *githubURI)
	// repos, err := gapi.FetchReposCursor("alextanhongpin", "2018-01-01", "2018-01-15")
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(len(repos), repos)
}
