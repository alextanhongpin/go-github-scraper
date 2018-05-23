package main

import (
	"flag"
	"log"

	"github.com/alextanhongpin/go-github-scraper/api/github"
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/model"
)

func main() {
	githubToken := flag.String("github_token", "", "The Github's access token used to make calls to the GraphQL endpoint")
	dbName := flag.String("db_name", "scraper", "The name of the database")
	dbHost := flag.String("db_host", "mongodb://myuser:mypass@localhost:27017", "The hostname of the database")

	flag.Parse()

	db := database.New(*dbHost, *dbName)
	defer db.Close()

	// Create a new collection with the session
	// sess, users := db.Collection("users")
	// defer sess.Close()
	// users.Find()

	gapi := github.New(*githubToken, "https://api.github.com/graphql", "Malaysia")
	cursor := ""
	hasNextPage := true
	var users []model.User
	for hasNextPage {
		res, err := gapi.FetchUsers("2018-01-01", "2018-02-01", cursor)
		if err != nil {
			break
		}
		log.Println("got user count:", res.Data.Search.UserCount)
		hasNextPage = res.Data.Search.PageInfo.HasNextPage
		cursor = res.Data.Search.PageInfo.EndCursor
		for _, edge := range res.Data.Search.Edges {
			users = append(users, edge.Node)
		}
	}
	log.Println(len(users))
}
