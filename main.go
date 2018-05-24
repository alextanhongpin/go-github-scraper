package main

import (
	"flag"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/database"
)

type User struct {
	Login     string    `json:"login,omitempty" bson:"login,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Count     int64     `json:"count,omitempty"`
}

func main() {
	githubToken := flag.String("github_token", "", "The Github's access token used to make calls to the GraphQL endpoint")
	dbName := flag.String("db_name", "scraper", "The name of the database")
	dbHost := flag.String("db_host", "mongodb://myuser:mypass@localhost:27017", "The hostname of the database")

	_ = githubToken
	flag.Parse()

	db := database.New(*dbHost, *dbName)
	defer db.Close()

	// gapi := github.New(*githubToken, "https://api.github.com/graphql", "Malaysia")
	// cursor := ""
	// hasNextPage := true
	// var users []model.User
	// for hasNextPage {
	// 	res, err := gapi.FetchUsers("2018-01-01", "2018-02-01", cursor)
	// 	if err != nil {
	// 		break
	// 	}
	// 	log.Println("got user count:", res.Data.Search.UserCount)
	// 	hasNextPage = res.Data.Search.PageInfo.HasNextPage
	// 	cursor = res.Data.Search.PageInfo.EndCursor
	// 	for _, edge := range res.Data.Search.Edges {
	// 		users = append(users, edge.Node)
	// 	}
	// }
	// log.Println(len(users))
}
